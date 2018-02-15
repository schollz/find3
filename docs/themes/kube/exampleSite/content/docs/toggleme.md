+++
title = "Toggleme"
date = "2017-04-10T16:42:59+01:00"
draft = false
weight = 190
description = "Display or hide elements with simple toggle"
bref= "Toggleme is a great way to add a binary option to either display some content or hide it. It works like a charm for menus, disclaimers and so much more!"
toc = true
+++

<h3 class="section-head" id="h-demo"><a href="#h-demo">Demo</a></h3>
<p>Toggleme works on mobile devices as well as on desktops. To see Toggleme in action, just resize this window or open this page on a mobile device.</p>
<div class="example">
  <div class="show-sm">
    <a data-component="toggleme" data-target="#navbar" href="#"><b>Toggle</b></a>
  </div>
  <div class="hide-sm" id="navbar">
    <div id="navbar-demo">
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
            <a href="#">Account</a>
          </li>
        </ul>
      </nav>
    </div>
  </div>
  <pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"show-sm"</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#"</span> <span class="hljs-attr">data-component</span>=<span class="hljs-string">"toggleme"</span> <span class="hljs-attr">data-target</span>=<span class="hljs-string">"#navbar"</span>&gt;</span>Toggle<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>

<span class="hljs-tag">&lt;<span class="hljs-name">nav</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"navbar"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"hide-sm"</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">ul</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">li</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">ul</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">nav</span>&gt;</span>
</pre>
</div>
<h3 class="section-head" id="h-usage"><a href="#h-usage">Usage</a></h3>
<div class="example">
  <p><button class="button outline" data-component="toggleme" data-target="#togglebox-target-1" data-text="Hide Me">Show Me</button></p>
  <div class="togglebox-box hide" id="togglebox-target-1">
    <h3>Ok, I'm opened. Now hide me</h3>
    <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
  </div>
  <pre class="code skip">&lt;button data-component=<span class="hljs-string">"toggleme"</span> data-target=<span class="hljs-string">"#togglebox-target"</span> data-<span class="hljs-built_in">text</span>=<span class="hljs-string">"Hide Me"</span>&gt;Show Me&lt;/button&gt;

&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">id</span>=<span class="hljs-string">"togglebox-target"</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"hide"</span>&gt;
    ...
&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
</div>
<h3 class="section-head" id="h-multiple-targets"><a href="#h-multiple-targets">Multiple targets</a></h3>
<div class="example">
  <p><button class="button primary outline" data-component="toggleme" data-target="#togglebox-target-3, #togglebox-target-4" data-text="Hide Me">Show Me</button></p>
  <div class="togglebox-box hide" id="togglebox-target-3">
    <h3>Ok, I'm opened. Now hide me</h3>
    <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
  </div>
  <div class="togglebox-box hide" id="togglebox-target-4">
    <h3>... and, I'm opened too. Now hide me</h3>
    <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
  </div>
  <pre class="code skip">&lt;button data-component=<span class="hljs-string">"toggleme"</span> data-target=<span class="hljs-string">"#togglebox-target-3, #togglebox-target-4"</span> data-<span class="hljs-built_in">text</span>=<span class="hljs-string">"Hide Me"</span>&gt;Show Me&lt;/button&gt;

&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">id</span>=<span class="hljs-string">"togglebox-target-1"</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"hide"</span>&gt;
    ...
&lt;/<span class="hljs-keyword">div</span>&gt;

&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">id</span>=<span class="hljs-string">"togglebox-target-2"</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"hide"</span>&gt;
    ...
&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
</div>
<h3 class="section-head" id="h-settings"><a href="#h-settings">Settings</a></h3>
<h5>target</h5>
<ul>
  <li>Type: <var>string</var></li>
  <li>Default: <var>null</var></li>
  <li>Possible values: <var>selector/selectors</var></li>
</ul>
<p>Defines which layer or layers should be displayed upon clicking the "open" button.</p>
<pre class="code skip">&lt;<span class="hljs-keyword">button </span><span class="hljs-meta">data</span>-component=<span class="hljs-string">"toggleme"</span> <span class="hljs-meta">data</span>-target=<span class="hljs-string">"#togglebox-target"</span>&gt;Show Me&lt;/<span class="hljs-keyword">button&gt;
</span>
&lt;<span class="hljs-keyword">button </span><span class="hljs-meta">data</span>-component=<span class="hljs-string">"toggleme"</span> <span class="hljs-meta">data</span>-target=<span class="hljs-string">"#togglebox-target-1, #togglebox-target-2"</span>&gt;Show Me&lt;/<span class="hljs-keyword">button&gt;
</span>
&lt;<span class="hljs-keyword">button </span><span class="hljs-meta">data</span>-component=<span class="hljs-string">"toggleme"</span> <span class="hljs-meta">data</span>-target=<span class="hljs-string">".togglebox-target"</span>&gt;Show Me&lt;/<span class="hljs-keyword">button&gt;
</span>
</pre>
<h5>text</h5>
<ul>
  <li>Type: <var>string</var></li>
  <li>Default: <var>''</var></li>
</ul>
<p>Defines what text should a button have after the layer has been opened. If not set, button text will not change.</p>
<pre class="code skip">&lt;<span class="hljs-keyword">button </span><span class="hljs-meta">data</span>-component=<span class="hljs-string">"toggleme"</span> <span class="hljs-meta">data</span>-target=<span class="hljs-string">"#togglebox-target"</span> <span class="hljs-meta">data</span>-text=<span class="hljs-string">"Hide Me"</span>&gt;Show Me&lt;/<span class="hljs-keyword">button&gt;
</span>
</pre>
<h3 class="section-head" id="h-callbacks"><a href="#h-callbacks">Callbacks</a></h3>
<h5>open</h5>
<pre class="code skip">$(<span class="hljs-string">'#togglebox-target'</span>).on(<span class="hljs-string">'open.toggleme'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});
</pre>
<h5>opened</h5>
<pre class="code skip">$(<span class="hljs-string">'#togglebox-target'</span>).on(<span class="hljs-string">'opened.toggleme'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});
</pre>
<h5>close</h5>
<pre class="code skip">$(<span class="hljs-string">'#togglebox-target'</span>).on(<span class="hljs-string">'close.toggleme'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});
</pre>
<h5>closed</h5>
<pre class="code skip">$(<span class="hljs-string">'#togglebox-target'</span>).on(<span class="hljs-string">'closed.toggleme'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});
</pre>
<h3 class="section-head" id="h-api"><a href="#h-api">API</a></h3>
<h5>toggle</h5>
<pre class="code skip"><span class="hljs-variable">$(</span><span class="hljs-string">'#togglebox-target'</span>).toggleme(<span class="hljs-string">'toggle'</span>);
</pre>
<h5>open</h5>
<pre class="code skip"><span class="hljs-variable">$(</span><span class="hljs-string">'#togglebox-target'</span>).toggleme(<span class="hljs-string">'open'</span>);
</pre>
<h5>close</h5>
<pre class="code skip"><span class="hljs-variable">$(</span><span class="hljs-string">'#togglebox-target'</span>).toggleme(<span class="hljs-string">'close'</span>);
</pre>
<h5>isOpened</h5>
<pre class="code skip"><span class="hljs-keyword">var</span> isOpened = $(<span class="hljs-string">'#togglebox-target'</span>).toggleme(<span class="hljs-string">'isOpened'</span>);
</pre>
<h5>isClosed</h5>
<pre class="code skip"><span class="hljs-keyword">var</span> isClosed = $(<span class="hljs-string">'#togglebox-target'</span>).toggleme(<span class="hljs-string">'isClosed'</span>);
</pre>
<h5>destroy</h5>
<pre class="code skip"><span class="hljs-variable">$(</span><span class="hljs-string">'#togglebox-target'</span>).toggleme(<span class="hljs-string">'destroy'</span>);
</pre>
