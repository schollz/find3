+++
date = "2017-04-10T16:42:43+01:00"
draft = false
weight = 180
description = "Navigation, menus and content sliding from outside the page"
title = "Offcanvas"
bref="Offcanvas makes it look like a menu is sliding from the outside of the page. This may be useful in many different cases, one of them being when you need to save space on screen and don't have to display sidebar at all times"
toc= true
+++

<h3 class="section-head" id="h-demo"><a href="#h-demo">Demo</a></h3>
<p>This feature is very easy to set up. It is based on <code>data-component</code> set to <var>offcanvas</var>, <code>data-target</code> set to the menu layer (<var>#offcanvas-right</var> in example below) and a feature-specific <code>data-direction</code> which is required for right menu to be set to <var>right</var>.</p>
<p><a class="button outline" data-component="offcanvas" data-target="#offcanvas-left" href="#"><i class="kube-menu"></i> Open Left</a> <a class="button outline" data-component="offcanvas" data-direction="right" data-target="#offcanvas-right" href="#">Open Right <i class="kube-menu"></i></a></p>
<div class="hide" id="offcanvas-left">
  <nav>
    <ul>
      <li>
        <a href="">Home</a>
      </li>
      <li>
        <a href="">About</a>
      </li>
      <li>
        <a href="">Showcase</a>
      </li>
      <li>
        <a href="">Help</a>
      </li>
      <li>
        <a href="">Contact</a>
      </li>
    </ul>
  </nav>
</div>
<div class="hide" id="offcanvas-right">
  <a class="close" href="#"></a>
  <nav>
    <ul>
      <li>
        <a href="">Home</a>
      </li>
      <li>
        <a href="">About</a>
      </li>
      <li>
        <a href="">Showcase</a>
      </li>
      <li>
        <a href="">Help</a>
      </li>
      <li>
        <a href="">Contact</a>
      </li>
    </ul>
  </nav>
</div>
<pre class="code skip">// Left
<span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#"</span> <span class="hljs-attr">data-component</span>=<span class="hljs-string">"offcanvas"</span> <span class="hljs-attr">data-target</span>=<span class="hljs-string">"#offcanvas-left"</span>&gt;</span>Open Left<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span>

<span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"offcanvas-left"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"hide"</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">nav</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">ul</span>&gt;</span>
            <span class="hljs-tag">&lt;<span class="hljs-name">li</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
        <span class="hljs-tag">&lt;/<span class="hljs-name">ul</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">nav</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>

// Right + Close
<span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#"</span> <span class="hljs-attr">data-component</span>=<span class="hljs-string">"offcanvas"</span> <span class="hljs-attr">data-target</span>=<span class="hljs-string">"#offcanvas-right"</span> <span class="hljs-attr">data-direction</span>=<span class="hljs-string">"right"</span>&gt;</span>Open Right<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span>

<span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"offcanvas-right"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"hide"</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"close"</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">nav</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">ul</span>&gt;</span>
            <span class="hljs-tag">&lt;<span class="hljs-name">li</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
        <span class="hljs-tag">&lt;/<span class="hljs-name">ul</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">nav</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
</pre>
<h3 class="section-head" id="h-settings"><a href="#h-settings">Settings</a></h3>
<h5>target</h5>
<ul>
  <li>Type: <var>string</var></li>
  <li>Default: <var>null</var></li>
</ul>
<p>Sets ID selector of an element that will slide from the side.</p>
<h5>push</h5>
<ul>
  <li>Type: <var>boolean</var></li>
  <li>Default: <var>true</var></li>
</ul>
<h5>direction</h5>
<ul>
  <li>Type: <var>string</var></li>
  <li>Default: <var>'left'</var></li>
</ul>
<p>The direction in which page will shift to give way for the sidebar menu. Default is <var>left</var> and is not required for left-side navigation, however, <var>right</var> value must be set for the right-side menu to work.</p>
<h5>clickOutside</h5>
<ul>
  <li>Type: <var>boolean</var></li>
  <li>Default: <var>true</var></li>
</ul>
<p>Unless set to <var>false</var>, clicking anywhere on a page will make side menu to close.</p>
<h5>width</h5>
<ul>
  <li>Type: <var>string</var></li>
  <li>Default: <var>'250px'</var></li>
</ul>
<p>Sidebar width in pixels.</p>
<h5>animation</h5>
<ul>
  <li>Type: <var>boolean</var></li>
  <li>Default: <var>true</var></li>
</ul>
<p>Setting this to <var>false</var> turns off opening and closing animation.</p>
<h3 class="section-head" id="h-callbacks"><a href="#h-callbacks">Callbacks</a></h3>
<h5>open</h5>
<pre class="code skip">$(<span class="hljs-string">'#my-offcanvas'</span>).on(<span class="hljs-string">'open.offcanvas'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});
</pre>
<h5>opened</h5>
<pre class="code skip">$(<span class="hljs-string">'#my-offcanvas'</span>).on(<span class="hljs-string">'opened.offcanvas'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});
</pre>
<h5>close</h5>
<pre class="code skip">$(<span class="hljs-string">'#my-offcanvas'</span>).on(<span class="hljs-string">'close.offcanvas'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});
</pre>
<h5>closed</h5>
<pre class="code skip">$(<span class="hljs-string">'#my-offcanvas'</span>).on(<span class="hljs-string">'closed.offcanvas'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});
</pre>