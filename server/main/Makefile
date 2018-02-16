.PHONY: scratch, install, basicbuild, server, server1, server2, server3, dev1, dev2, dev3


TAG=$(shell git tag)
HASH=$(shell git log --pretty=format:"%h" -n 1)
LDFLAGS=-ldflags "-s -w -X main.Version=${TAG}-${HASH}"


basicbuild:
	rerun -p "**/*.{go,tmpl,css,js}" --ignore 'bindata.go' make server1

update:
	dep ensure -v
	dep ensure -v -update
	# remove things that import "testing" so that the flags are not included
	rm -rf vendor/golang.org/x/text/internal/testtext/
	rm -rf vendor/golang.org/x/net/nettest
	go build -v -a
	
release:
	docker pull karalabe/xgo-latest
	go get github.com/karalabe/xgo
	mkdir -p bin 
	xgo -go "1.9.2" -dest bin ${LDFLAGS} -targets linux/amd64,linux/arm-6,darwin/amd64,windows/amd64 github.com/schollz/kiki
	# cd bin && upx --brute kiki-linux-amd64


server1:
	go build -v ${LDFLAGS}
	./main -debug

docs:
	cd doc && make serve
