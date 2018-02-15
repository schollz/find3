+++
title = "Modal"
draft = false
weight = 150
date = "2017-04-10T16:42:18+01:00"
description = "Clean, functional, and extensible modal window dialogs"
bref="Modal windows are used for various reasons and purposes. Kube makes the way you design and operate modal windows very straightforward. First, you set a <code>div</code> which will represent the content of the modal window, then you set up a way to open and close this window, and finally you add a way to call the modal from the page"
toc = true
+++

<h4 class="demo-head" id="h-calling">Calling</h4>
<p>Here you add an actual call to the modal window. Clicking the "Open" button below will launch a <var>modal</var> with content of a <var>#my-modal</var> div. You can use buttons or links to open modals in Kube. Clicking anywhere outside of the modal or hitting <kbd>ESC</kbd> will close the window.</p>
<div class="example">
  <p><button class="button" data-component="modal" data-loaded="true" data-target="#my-modal" data-width="760px">Open</button></p>
  <pre class="code skip">// <span class="hljs-keyword">Call</span>
&lt;button <span class="hljs-keyword">data</span>-component=<span class="hljs-string">"modal"</span> <span class="hljs-keyword">data</span>-target=<span class="hljs-string">"#my-modal"</span>&gt;<span class="hljs-keyword">Open</span>&lt;/button&gt;

// Markup
&lt;<span class="hljs-keyword">div</span> <span class="hljs-keyword">id</span>=<span class="hljs-string">"my-modal"</span> <span class="hljs-keyword">class</span>=<span class="hljs-string">"modal-box hide"</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-keyword">class</span>=<span class="hljs-string">"modal"</span>&gt;
        &lt;span <span class="hljs-keyword">class</span>=<span class="hljs-string">"close"</span>&gt;&lt;/span&gt;
        &lt;<span class="hljs-keyword">div</span> <span class="hljs-keyword">class</span>=<span class="hljs-string">"modal-header"</span>&gt;Modal Header&lt;/<span class="hljs-keyword">div</span>&gt;
        &lt;<span class="hljs-keyword">div</span> <span class="hljs-keyword">class</span>=<span class="hljs-string">"modal-body"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;/<span class="hljs-keyword">div</span>&gt;
&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
</div>
<p>Open from url</p>
<div class="example">
  <p><button class="button" data-component="modal" data-loaded="true" data-target="#ui-modal" data-url="/tests/modal.html">Open from url</button></p>
  <pre class="code skip">// <span class="hljs-keyword">Call</span>
&lt;button <span class="hljs-keyword">data</span>-component=<span class="hljs-string">"modal"</span> <span class="hljs-keyword">data</span>-target=<span class="hljs-string">"#ui-modal"</span> <span class="hljs-keyword">data</span>-<span class="hljs-keyword">url</span>=<span class="hljs-string">"modal.html"</span>&gt;<span class="hljs-keyword">Open</span>&lt;/button&gt;

// Markup
&lt;<span class="hljs-keyword">div</span> <span class="hljs-keyword">id</span>=<span class="hljs-string">"ui-modal"</span> <span class="hljs-keyword">class</span>=<span class="hljs-string">"modal-box hide"</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-keyword">class</span>=<span class="hljs-string">"modal"</span>&gt;
        &lt;span <span class="hljs-keyword">class</span>=<span class="hljs-string">"close"</span>&gt;&lt;/span&gt;
        &lt;<span class="hljs-keyword">div</span> <span class="hljs-keyword">class</span>=<span class="hljs-string">"modal-header"</span>&gt;UI Modal&lt;/<span class="hljs-keyword">div</span>&gt;
        &lt;<span class="hljs-keyword">div</span> <span class="hljs-keyword">class</span>=<span class="hljs-string">"modal-body"</span>&gt;... <span class="hljs-keyword">content</span> <span class="hljs-keyword">from</span> modal.html ...&lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;/<span class="hljs-keyword">div</span>&gt;
&lt;/<span class="hljs-keyword">div</span>&gt;

// modal.html
&lt;p&gt;...&lt;/p&gt;
&lt;a href=<span class="hljs-string">"#"</span> <span class="hljs-keyword">data</span>-<span class="hljs-keyword">action</span>=<span class="hljs-string">"modal-close"</span>&gt;<span class="hljs-keyword">Close</span>&lt;/a&gt;
</pre>
</div>
<p>Direct open</p>
<div class="example">
  <p><button class="button" onclick="$.modalwindow({ target: '#ui-modal', url: '/tests/modal.html' });">Direct Open</button></p>
  <pre class="code skip"><span class="hljs-comment">// Call</span>
&lt;<span class="hljs-keyword">button</span> <span class="hljs-keyword">onclick</span>=<span class="hljs-string">"$.modalwindow({ target: '#ui-modal', url: 'modal.html' });"</span>&gt;Open&lt;/<span class="hljs-keyword">button</span>&gt;
</pre>
</div>
<div class="modal-box hide" id="ui-modal">
  <div class="modal">
    <span class="close"></span>
    <div class="modal-header">
      UI Modal
    </div>
    <div class="modal-body"></div>
  </div>
</div>
<div class="modal-box hide" id="my-modal">
  <div class="modal">
    <span class="close"></span>
    <div class="modal-header">
      My Modal
    </div>
    <div class="modal-body">
      <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
      <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
      <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
      <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
      <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
      <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
      <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
      <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
      <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
    </div>
  </div>
</div>
<h4 class="section-head" id="h-actions"><a href="#h-actions">Actions</a></h4>
<p>Using <var>modal-close</var> action you now introducing a way to close you window, using a link or a button.</p>
<pre class="code skip">&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">id</span>=<span class="hljs-string">"my-modal"</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"modal-box hide"</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"modal"</span>&gt;
        &lt;span <span class="hljs-built_in">class</span>=<span class="hljs-string">"close"</span>&gt;&lt;/span&gt;
        &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"modal-header"</span>&gt;Modal Header&lt;/<span class="hljs-keyword">div</span>&gt;
        &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"modal-body"</span>&gt;
            ...
            &lt;a href=<span class="hljs-string">"#"</span> data-action=<span class="hljs-string">"modal-close"</span>&gt;Close&lt;/a&gt;
        &lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;/<span class="hljs-keyword">div</span>&gt;
&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
<h3 class="section-head" id="h-settings"><a href="#h-settings">Settings</a></h3>
<h5>target</h5>
<ul>
  <li>Type: <var>string</var></li>
  <li>Default: <var>null</var></li>
</ul>
<p>Defines a content layer for the modal window</p>
<h5>url</h5>
<ul>
  <li>Type: <var>string</var></li>
  <li>Default: <var>false</var></li>
</ul>
<p>Defines a URL in case your modal is opening via a URL.</p>
<h5>header</h5>
<ul>
  <li>Type: <var>string</var></li>
  <li>Default: <var>false</var></li>
</ul>
<p>Sets the header for the modal window. Optional, and is <var>false</var> by default.</p>
<h5>width</h5>
<ul>
  <li>Type: <var>string</var></li>
  <li>Default: <var>'600px'</var></li>
</ul>
<h5>height</h5>
<ul>
  <li>Type: <var>string</var></li>
  <li>Default: <var>false</var></li>
</ul>
<h5>maxHeight</h5>
<ul>
  <li>Type: <var>string</var></li>
  <li>Default: <var>false</var></li>
</ul>
<p>This setting defines the maximum height of the window. A scrollbar will be introduced in case there's more text than <var>maxHeight</var> can accommodate.</p>
<h5>position</h5>
<ul>
  <li>Type: <var>string</var></li>
  <li>Default: <var>'center'</var></li>
</ul>
<p>This is where your modal appears when opened.</p>
<h5>overlay</h5>
<ul>
  <li>Type: <var>boolean</var></li>
  <li>Default: <var>true</var></li>
</ul>
<p>When this is set to <var>false</var>, you modal window will just appear on top ow your page, without an overlay effect. By default, your page will be "covered" with a semi-transparent layer when you open a modal.</p>
<h5>animation</h5>
<ul>
  <li>Type: <var>boolean</var></li>
  <li>Default: <var>true</var></li>
</ul>
<p>Turns opening and closing animation on and off.</p>
<h3 class="section-head" id="h-callbacks"><a href="#h-callbacks">Callbacks</a></h3>
<h5>open</h5>
<pre class="code skip">$(<span class="hljs-string">'#my-modal'</span>).on(<span class="hljs-string">'open.modal'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});
</pre>
<h5>opened</h5>
<pre class="code skip">$(<span class="hljs-string">'#my-modal'</span>).on(<span class="hljs-string">'opened.modal'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});
</pre>
<h5>close</h5>
<pre class="code skip">$(<span class="hljs-string">'#my-modal'</span>).on(<span class="hljs-string">'close.modal'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});
</pre>
<h5>closed</h5>
<pre class="code skip">$(<span class="hljs-string">'#my-modal'</span>).on(<span class="hljs-string">'closed.modal'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});
</pre>
<h3 class="section-head" id="h-api"><a href="#h-api">API</a></h3>
<p>You can use these API methods to programmatically operate and modify modal windows.</p>
<h5>close</h5>
<pre class="code skip"><span class="hljs-variable">$(</span><span class="hljs-string">'#my-modal'</span>).modal(<span class="hljs-string">'close'</span>);
</pre>
<h5>setHeader</h5>
<pre class="code skip"><span class="hljs-variable">$(</span><span class="hljs-string">'#my-modal'</span>).modal(<span class="hljs-string">'setHeader'</span>, <span class="hljs-string">'My Header'</span>);
</pre>
<p>This is another way to set a header for the modal on the fly without introducing a <code>div</code> with a <var>modal-header</var> class.</p>
<h5>setContent</h5>
<pre class="code skip"><span class="hljs-variable">$(</span><span class="hljs-string">'#my-modal'</span>).modal(<span class="hljs-string">'setContent'</span>, <span class="hljs-string">'My Content'</span>);
</pre>
<p>Content of the modal window can be set up on the fly as well. No need for a <code>div</code> with <var>modal</var> class.</p>
<h5>setWidth</h5>
<pre class="code skip"><span class="hljs-variable">$(</span><span class="hljs-string">'#my-modal'</span>).modal(<span class="hljs-string">'setWidth'</span>, <span class="hljs-string">'800px'</span>);
</pre>