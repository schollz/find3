+++
date = "2017-04-10T16:42:12+01:00"
draft = false
weight = 140
description = "Popular and adjustable menu and navigation option"
title = "Dropdown"
bref= "Dropdown menus in Kube are simple and intuitive. You've got a link with <code>data-component='dropdown'</code> and a <code>div</code> with some kind of HTML content. Clicking outside of the dropdown or hitting <kbd>Esc</kbd> closes dropdown"
toc = true
+++

<h3 class="section-head" id="h-demo"><a href="#h-demo">Demo</a></h3>
<div class="example">
  <p><a data-component="dropdown" data-loaded="true" data-target="#dropdown1" href="#">Show Dropdown <span class="caret down"></span></a></p>
  <div class="dropdown hide" id="dropdown1">
    <a class="close show-sm" href=""></a>
    <ul>
      <li>
        <a href="">Item 1</a>
      </li>
      <li>
        <a href="">Item 2</a>
      </li>
      <li class="active">
        <a href="">Item 3</a>
      </li>
      <li>
        <a href="">Item 4</a>
      </li>
      <li>
        <a href=""><span class="label primary">Item 5</span></a>
      </li>
      <li>
        <a href=""><span class="label error">Item 6</span></a>
      </li>
    </ul>
  </div>
  <pre class="code skip">// Toggle
<span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#"</span> <span class="hljs-attr">data-component</span>=<span class="hljs-string">"dropdown"</span> <span class="hljs-attr">data-target</span>=<span class="hljs-string">"#my-dropdown"</span>&gt;</span>Show <span class="hljs-tag">&lt;<span class="hljs-name">span</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"caret down"</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">span</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span>

// Dropdown
<span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"dropdown hide"</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"my-dropdown"</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">""</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"close show-sm"</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">ul</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">li</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">""</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">ul</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
</pre>
</div>
<div class="example">
  <p><button class="button outline" data-component="dropdown" data-loaded="true" data-target="#dropdown2">Show Dropdown <span class="caret down"></span></button></p>
  <div class="dropdown hide" id="dropdown2">
    <div style="padding: 24px;">
      <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p><a data-action="dropdown-close" href="#">Close</a>
    </div>
  </div>
  <pre class="code skip">// Toggle
<span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">data-component</span>=<span class="hljs-string">"dropdown"</span> <span class="hljs-attr">data-target</span>=<span class="hljs-string">"#my-dropdown"</span>&gt;</span>Show Dropdown <span class="hljs-tag">&lt;<span class="hljs-name">span</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"caret down"</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">span</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span>

// Dropdown
<span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"dropdown hide"</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"my-dropdown"</span>&gt;</span>

    <span class="hljs-tag">&lt;<span class="hljs-name">p</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">p</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#"</span> <span class="hljs-attr">data-action</span>=<span class="hljs-string">"dropdown-close"</span>&gt;</span>Close<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span>

<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
</pre>
</div>
<h3 class="section-head">Navigation Fixed</h3>
<p>Scroll down a bit to make a fixed navigation bar appear at the top of the page, and then try out "Account" dropdown.</p>
<div class="example">
  <div data-component="sticky" data-loaded="true" id="navbar-demo">
    <div id="navbar-brand">
      <a href=""><img alt="Brand" src="/img/kube/brand.png"></a>
    </div>
    <nav id="navbar-main">
      <ul>
        <li>
          <a href="#">Shop</a>
        </li>
        <li>
          <a href="#">News</a>
        </li>
        <li>
          <a href="#">Contact</a>
        </li>
        <li>
          <a href="#">Blog</a>
        </li>
        <li>
          <a data-component="dropdown" data-loaded="true" data-target="#dropdown-fixed" href="">Account <span class="caret down"></span></a>
        </li>
      </ul>
    </nav>
  </div>
  <div class="dropdown hide" id="dropdown-fixed">
    <ul>
      <li>
        <a href="">Billing</a>
      </li>
      <li>
        <a href="">Log Out</a>
      </li>
    </ul>
  </div>
  <pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"navbar-demo"</span> <span class="hljs-attr">data-component</span>=<span class="hljs-string">"sticky"</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">nav</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">ul</span>&gt;</span>
            <span class="hljs-tag">&lt;<span class="hljs-name">li</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">""</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
            <span class="hljs-tag">&lt;<span class="hljs-name">li</span>&gt;</span>
                <span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">""</span> <span class="hljs-attr">data-component</span>=<span class="hljs-string">"dropdown"</span> <span class="hljs-attr">data-target</span>=<span class="hljs-string">"#dropdown-fixed"</span>&gt;</span>
                    Account
                    <span class="hljs-tag">&lt;<span class="hljs-name">span</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"caret down"</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">span</span>&gt;</span>
                <span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span>
            <span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
        <span class="hljs-tag">&lt;/<span class="hljs-name">ul</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">nav</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>

// Dropdown
<span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"dropdown hide"</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"dropdown-fixed"</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">ul</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">li</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">""</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">ul</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
</pre>
</div>
<h3 class="section-head" id="h-settings"><a href="#h-settings">Settings</a></h3>
<h5>target</h5>
<ul>
  <li>Type: <var>string</var></li>
  <li>Default: <var>null</var></li>
</ul>
<p>Sets an ID of a target dropdown layer.</p>
<h5>height</h5>
<ul>
  <li>Type: <var>int</var></li>
  <li>Default: <var>false</var></li>
</ul>
<p>Sets dropdown height.</p>
<h5>width</h5>
<ul>
  <li>Type: <var>int</var></li>
  <li>Default: <var>false</var></li>
</ul>
<p>Sets dropdown width.</p>
<h5>animation</h5>
<ul>
  <li>Type: <var>boolean</var></li>
  <li>Default: <var>true</var></li>
</ul>
<p>Turns opening and closing animation on and off.</p>
<h3 class="section-head" id="h-callbacks"><a href="#h-callbacks">Callbacks</a></h3>
<h5>open</h5>
<pre class="code">$(<span class="hljs-string">'#my-dropdown'</span>).on(<span class="hljs-string">'open.dropdown'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});</pre>
<h5>opened</h5>
<pre class="code">$(<span class="hljs-string">'#my-dropdown'</span>).on(<span class="hljs-string">'opened.dropdown'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});</pre>
<h5>close</h5>
<pre class="code">$(<span class="hljs-string">'#my-dropdown'</span>).on(<span class="hljs-string">'close.dropdown'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});</pre>
<h5>closed</h5>
<pre class="code">$(<span class="hljs-string">'#my-dropdown'</span>).on(<span class="hljs-string">'closed.dropdown'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});</pre>