+++
date = "2017-04-10T16:42:50+01:00"
draft = false
weight = 180
description = "Smoothly and reliably collapse elements for extra convenience"
title = "Collapse"
bref="Collapsable elements are horizontally aligned tabs, in a way. Jokes aside, collapsable elements are useful and easy to setup, both for direct purpose of switching between content while collapsing everything else, and for navigation use cases"
toc = true
+++

<h3 class="section-head" id="h-base"><a href="#h-base">Base</a></h3>
<div class="example">
  <div class="my-collapse" data-component="collapse" id="my-collapse">
    <h4><a class="collapse-toggle" href="#collapse-box-1">Item 1</a></h4>
    <div class="collapse-box hide" id="collapse-box-1">
      <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
    </div>
    <h4><a class="collapse-toggle" href="#collapse-box-2">Item 2</a></h4>
    <div class="collapse-box" id="collapse-box-2">
      <p>I look active on load.</p>
      <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
    </div>
    <h4><a class="collapse-toggle" href="#collapse-box-3">Item 3</a></h4>
    <div class="collapse-box hide" id="collapse-box-3">
      <h6>Important Heading</h6>
      <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
      <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
    </div>
  </div>

<pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"my-collapse"</span> <span class="hljs-attr">data-component</span>=<span class="hljs-string">"collapse"</span>&gt;</span>

    <span class="hljs-tag">&lt;<span class="hljs-name">h4</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#box-1"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"collapse-toggle"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">h4</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"collapse-box hide"</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"box-1"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>

    <span class="hljs-tag">&lt;<span class="hljs-name">h4</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#box-2"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"collapse-toggle"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">h4</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"collapse-box"</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"box-2"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>

    <span class="hljs-tag">&lt;<span class="hljs-name">h4</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#box-3"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"collapse-toggle"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">h4</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"collapse-box hide"</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"box-3"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>

<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
</pre>
</div>
<h3 class="section-head" id="h-navigation-example"><a href="#h-navigation-example">Navigation Example</a></h3>
<p>Here you can see a list of items forming a navigation menu, with a nested list which acts as a collapsable element</p>
<div class="example">
  <ul id="demo-nav-collapse">
    <li>
      <a href="#">Installation</a>
    </li>
    <li>
      <a href="#">Configuration</a>
    </li>
    <li>
      <a href="#">Adding Content</a>
    </li>
    <li>
      <a href="#">Templates</a>
    </li>
    <li data-component="collapse">
      <a class="collapse-toggle" href="#languages-box-2">Languages <span class="caret down"></span></a>
      <ul class="collapse-box" id="languages-box-2">
        <li>
          <a href="#">Setup</a>
        </li>
        <li>
          <a href="#">Translating content</a>
        </li>
        <li>
          <a href="#">Language variables</a>
        </li>
        <li>
          <a href="#">Supporting RTL</a>
        </li>
      </ul>
    </li>
    <li>
      <a href="#">Settings</a>
    </li>
    <li>
      <a href="#">Callbacks</a>
    </li>
    <li>
      <a href="#">API</a>
    </li>
  </ul>
  <pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">ul</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">li</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">li</span> <span class="hljs-attr">data-component</span>=<span class="hljs-string">"collapse"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#toggle-box"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"collapse-toggle"</span>&gt;</span>
            Toggle
            <span class="hljs-tag">&lt;<span class="hljs-name">span</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"caret down"</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">span</span>&gt;</span>
        <span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">ul</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"toggle-box"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"collapse-box"</span>&gt;</span>
            <span class="hljs-tag">&lt;<span class="hljs-name">li</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
        <span class="hljs-tag">&lt;/<span class="hljs-name">ul</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">li</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">ul</span>&gt;</span>
</pre>
</div>
<h3 class="section-head" id="h-settings"><a href="#h-settings">Settings</a></h3>
<p>Settings are pretty straightforward: you can toggle, denote active element, turn animation on and off, and more.</p>
<h5>toggle</h5>
<ul>
  <li>Type: <var>boolean</var></li>
  <li>Default: <var>true</var></li>
</ul>
<h5>active</h5>
<ul>
  <li>Type: <var>string or boolean</var></li>
  <li>Default: <var>false</var></li>
</ul>
<h5>toggleClass</h5>
<ul>
  <li>Type: <var>string</var></li>
  <li>Default: <var>'collapse-toggle'</var></li>
</ul>
<p>Sets a class of a collapsable object.</p>
<h5>boxClass</h5>
<ul>
  <li>Type: <var>string</var></li>
  <li>Default: <var>'collapse-box'</var></li>
</ul>
<p>Sets a class for collapsable object's content</p>
<h5>animation</h5>
<ul>
  <li>Type: <var>boolean</var></li>
  <li>Default: <var>true</var></li>
</ul>
<p>Turns animation on and off.</p>
<h3 class="section-head" id="h-callbacks"><a href="#h-callbacks">Callbacks</a></h3>
<h5>open</h5>
<pre class="code skip">$(<span class="hljs-string">'#my-collapse'</span>).on(<span class="hljs-string">'open.collapse'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});
</pre>
<h5>opened</h5>
<pre class="code skip">$(<span class="hljs-string">'#my-collapse'</span>).on(<span class="hljs-string">'opened.collapse'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});
</pre>
<h5>close</h5>
<pre class="code skip">$(<span class="hljs-string">'#my-collapse'</span>).on(<span class="hljs-string">'close.collapse'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});
</pre>
<h5>closed</h5>
<pre class="code skip">$(<span class="hljs-string">'#my-collapse'</span>).on(<span class="hljs-string">'closed.collapse'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});
</pre>
<h3 class="section-head" id="h-api"><a href="#h-api">API</a></h3>
<p>We love APIs, and for interactive elements such as collapsable elements we offer ways to programmatically open and close individual elements, or all at once.</p>
<div class="example">
  <p><button class="button outline" onclick="$('#my-collapse-api').collapse('open', '#box-2');">Open</button> <button class="button outline" onclick="$('#my-collapse-api').collapse('close', '#box-2');">Close</button> <button class="button outline" onclick="$('#my-collapse-api').collapse('openAll');">Open All</button> <button class="button outline" onclick="$('#my-collapse-api').collapse('closeAll');">Close All</button></p><br>
  <div class="my-collapse" data-component="collapse" id="my-collapse-api">
    <h4><a class="collapse-toggle" href="#box-1">Item 1</a></h4>
    <div class="collapse-box hide" id="box-1">
      <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
    </div>
    <h4><a class="collapse-toggle" href="#box-2">Item 2</a></h4>
    <div class="collapse-box hide" id="box-2">
      <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
    </div>
    <h4><a class="collapse-toggle" href="#box-3">Item 3</a></h4>
    <div class="collapse-box hide" id="box-3">
      <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
      <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
    </div>
  </div>
  <pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">onclick</span>=<span class="hljs-string">"$('#my-collapse-api').collapse('open', '#box-2');"</span>&gt;</span>Open<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">onclick</span>=<span class="hljs-string">"$('#my-collapse-api').collapse('close', '#box-2');"</span>&gt;</span>Close<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">onclick</span>=<span class="hljs-string">"$('#my-collapse-api').collapse('openAll');"</span>&gt;</span>Open All<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">onclick</span>=<span class="hljs-string">"$('#my-collapse-api').collapse('closeAll');"</span>&gt;</span>Close All<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span>

<span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"my-collapse-api"</span> <span class="hljs-attr">data-component</span>=<span class="hljs-string">"collapse"</span>&gt;</span>

    <span class="hljs-tag">&lt;<span class="hljs-name">h4</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#box-1"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"collapse-toggle"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">h4</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"collapse-box hide"</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"box-1"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>

    <span class="hljs-tag">&lt;<span class="hljs-name">h4</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#box-2"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"collapse-toggle"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">h4</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"collapse-box hide"</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"box-2"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>

    <span class="hljs-tag">&lt;<span class="hljs-name">h4</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#box-3"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"collapse-toggle"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">h4</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"collapse-box hide"</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"box-3"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>

<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
</pre>
</div>
<h5>open</h5>
<pre class="code skip"><span class="hljs-variable">$(</span><span class="hljs-string">'#my-collapse'</span>).collapse(<span class="hljs-string">'open'</span>);
</pre>
<h5>close</h5>
<pre class="code skip"><span class="hljs-variable">$(</span><span class="hljs-string">'#my-collapse'</span>).collapse(<span class="hljs-string">'close'</span>);
</pre>
<h5>openAll</h5>
<pre class="code skip"><span class="hljs-variable">$(</span><span class="hljs-string">'#my-collapse'</span>).collapse(<span class="hljs-string">'openAll'</span>);
</pre>
<h5>closeAll</h5>
<pre class="code skip"><span class="hljs-variable">$(</span><span class="hljs-string">'#my-collapse'</span>).collapse(<span class="hljs-string">'closeAll'</span>);
</pre>