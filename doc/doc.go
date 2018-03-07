// Copyright 2016 The ProjectX Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package frontend provides a web server that serves documentation and meta
// tags to instruct "go get" where to find the ProjectX source repository.
package doc

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/NYTimes/gziphandler"
	"github.com/russross/blackfriday"
)

const (
	extMarkdown  = ".md"
	docHostname  = "projectx.io"      // redirect doc requests to this host
	testHostname = "test.projectx.io" // don't redirect requests to this host
)

type server struct {
	handlers http.Handler // stack of wrapped http.Handlers
	docHTML  map[string][]byte
	docTitle map[string]string
	tmpl     struct {
		doc, download *template.Template
	}
}

// NewServer initializes and returns a new HTTP server serving documentation
// from the given directory. A nil Config prevents the server from building and
// serving binaries for download.
func NewServer(docs string) (http.Handler, error) {
	s := &server{}

	if err := s.parseTemplates(filepath.Join(docs, "templates")); err != nil {
		return nil, fmt.Errorf("parsing templates: %v", err)
	}
	if err := s.parseDocs(docs); err != nil {
		return nil, fmt.Errorf("parsing docs: %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(s.handleDoc))
	mux.Handle("/images/", http.FileServer(http.Dir(docs)))
	mux.Handle("/issue/", redirectHandler("/issue/", "https://github.com/githubuser/projectx/issues/"))
	s.handlers = goGetHandler{gziphandler.GzipHandler(canonicalHostHandler{mux})}

	return s, nil
}

// redirectHandler redirects requests from prefix/something to target/something.
// It rejects non-empty values of "something" that don't parse as int.
func redirectHandler(prefix, target string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		suffix := strings.TrimPrefix(r.URL.Path, prefix)
		if suffix != "" {
			if _, err := strconv.Atoi(suffix); err != nil {
				http.NotFound(w, r)
				return
			}
		}
		http.Redirect(w, r, target+suffix, http.StatusFound)
	})
}

type pageData struct {
	Title    string
	Content  interface{}
	FileName string
}

func (s *server) handleDoc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	switch r.URL.Path {
	case "/":
		s.renderDoc(w, "index.md")
	case "/doc/":
		s.renderDoc(w, "doc.md")
	case "/doc":
		http.Redirect(w, r, "/doc/", http.StatusFound)
	default:
		if !strings.HasPrefix(r.URL.Path, "/doc/") {
			http.NotFound(w, r)
			return
		}
		base := filepath.Base(r.URL.Path)
		switch base {
		case "index.md":
			http.Redirect(w, r, "/", http.StatusFound)
		case "doc.md":
			http.Redirect(w, r, "/doc/", http.StatusFound)
		default:
			s.renderDoc(w, base)
		}
	}
}

func (s *server) renderDoc(w http.ResponseWriter, fn string) {
	b, ok := s.docHTML[fn]
	if !ok {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if err := s.tmpl.doc.Execute(w, pageData{
		Title:    s.docTitle[fn] + " · ProjectX",
		Content:  template.HTML(b),
		FileName: fn,
	}); err != nil {
		log.Printf("Error executing doc content template: %s", err)
		return
	}
}

// ServeHTTP satisfies the http.Handler interface for a server.
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.TLS != nil {
		w.Header().Set("Strict-Transport-Security", "max-age=86400; includeSubDomains")
	}
	s.handlers.ServeHTTP(w, r)
}

func (s *server) parseTemplates(dir string) (err error) {
	s.tmpl.doc, err = template.ParseFiles(filepath.Join(dir, "base.tmpl"), filepath.Join(dir, "doc.tmpl"))
	if err != nil {
		return err
	}
	s.tmpl.download, err = template.ParseFiles(filepath.Join(dir, "base.tmpl"), filepath.Join(dir, "download.tmpl"))
	return err
}

func (s *server) parseDocs(dir string) error {
	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	var (
		html  = map[string][]byte{}
		title = map[string]string{}
	)
	for _, fi := range fis {
		fn := fi.Name()
		if filepath.Ext(fn) != extMarkdown {
			continue
		}
		b, err := ioutil.ReadFile(filepath.Join(dir, fn))
		if err != nil {
			return err
		}
		html[fn] = blackfriday.MarkdownCommon(b)
		title[fn] = docTitle(b)
	}
	s.docHTML = html
	s.docTitle = title
	return nil
}

// docTitle extracts the first Markdown header in the given document body.
// It expects the first line to be of the form
// 	# Title string
// If not, it will return "Untitled".
func docTitle(b []byte) string {
	if len(b) > 2 && b[0] == '#' {
		if i := bytes.IndexByte(b, '\n'); i != -1 {
			// On Windows we need to strip out the '\r' as well
			if b[i-1] == '\r' {
				i--
			}
			return string(b[2:i])
		}
	}
	return "Untitled"
}

type goGetHandler struct {
	http.Handler
}

func (h goGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Handler.ServeHTTP(w, r)
}

type canonicalHostHandler struct {
	http.Handler
}

func (h canonicalHostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Redirect requests to foo.projectx.io to projectx.io.
	if r.Host != docHostname && r.Host != testHostname && strings.HasSuffix(r.Host, "."+docHostname) {
		u := *r.URL
		u.Host = docHostname
		http.Redirect(w, r, u.String(), http.StatusFound)
		return
	}
	h.Handler.ServeHTTP(w, r)
}

// reloadHandler is a http.Handler wrapper that watches the contents of dir for
// modifications and reloads the underlying http.Handler when changes are
// made. It is intended to be used in local mode to facilitate rapid editing.
type reloadHandler struct {
	dir  string
	load func() (http.Handler, error)

	mu           sync.Mutex
	handler      http.Handler
	lastScan     time.Time
	lastModified time.Time
}

func (h *reloadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// scanInterval is time to wait between scanning for new changes.
	// It should be set long enough to avoid scanning on every request,
	// but not so long as to make editing tedious.
	const scanInterval = 5 * time.Second

	h.mu.Lock()
	if h.handler != nil && time.Since(h.lastScan) > scanInterval {
		err := filepath.Walk(h.dir, func(_ string, fi os.FileInfo, _ error) error {
			if t := fi.ModTime(); t.After(h.lastModified) {
				h.lastModified = t
				h.handler = nil
			}
			return nil
		})
		if err != nil {
			h.mu.Unlock()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if h.handler == nil {
		var err error
		h.handler, err = h.load()
		if err != nil {
			h.mu.Unlock()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	handler := h.handler
	h.mu.Unlock()

	handler.ServeHTTP(w, r)
}
