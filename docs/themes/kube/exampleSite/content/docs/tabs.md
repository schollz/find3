+++
title = "Tabs"
date = "2017-04-10T16:41:40+01:00"
draft = false
weight = 100
description = "User-friendly navigation for your content or app"
bref= "Tabs in Kube are crafted the way one would expect from the world's best CSS framework. With versatile API, wide range of settings an options, with callbacks and live examples. Take a look!"
toc= true
+++

<h3 class="section-head" id="h-base"><a href="#h-base">Base</a></h3>
<p>Here's an example of basic tabs setup. Tabs bar is an unordered list, and each tab in tabs bar is a list item. For each tab there's a corresponding div, which contains the body of the tab (it can be any kind of HTML).</p>
<div class="example">
  <nav class="tabs" data-component="tabs">
    <ul>
      <li class="active">
        <a href="#tab1">Home</a>
      </li>
      <li>
        <a href="#tab2">Shop</a>
      </li>
      <li>
        <a href="#tab3">Catalog</a>
      </li>
      <li>
        <a href="#tab4">T-Shirts</a>
      </li>
      <li>
        <a href="#tab5">Brand</a>
      </li>
    </ul>
  </nav>
  <div id="tab1">
    <h5>Home</h5>
    <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
  </div>
  <div class="hide" id="tab2">
    <h5>Shop</h5>
    <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
    <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
  </div>
  <div class="hide" id="tab3">
    <h5>Catalog</h5>
    <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
  </div>
  <div class="hide" id="tab4">
    <h5>T-Shirts</h5>
    <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
  </div>
  <div class="hide" id="tab5">
    <h5>Brand</h5>
    <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
  </div>
  <pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">nav</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"tabs"</span> <span class="hljs-attr">data-component</span>=<span class="hljs-string">"tabs"</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">ul</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">li</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"active"</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#tab1"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">li</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#tab2"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">ul</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">nav</span>&gt;</span>

<span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"tab1"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"tab2"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
    </pre>
</div>
<h3 class="section-head" id="h-javaScript-behavior"><a href="#h-javaScript-behavior">JavaScript behavior</a></h3>
<p>Tabs are interactive elements, and can be opened, closed, switched, destroyed and selected programmatically via JavaScript. Try this example, poke around and then have a look at the self-explanatory code example below.</p>
<div class="example">
  <p><button class="button outline" onclick="$('#tabs-demo').tabs('destroy');">destroy</button> <button class="button outline" onclick="$('#tabs-demo').tabs('prev');">prev</button> <button class="button outline" onclick="$('#tabs-demo').tabs('next');">next</button> <button class="button outline" onclick="$('#tabs-demo').tabs('open', 4);">open</button> <button class="button outline" onclick="$('#tabs-demo').tabs('close', '#tab-demo9');">close</button> <button class="button outline" onclick="$('#tabs-demo').tabs('closeAll');">closeAll</button></p><br>
  <nav class="tabs" data-component="tabs" id="tabs-demo">
    <ul>
      <li>
        <a href="#tab-demo6">Home</a>
      </li>
      <li class="active">
        <a href="#tab-demo7">Shop</a>
      </li>
      <li>
        <a href="#tab-demo8">Catalog</a>
      </li>
      <li>
        <a href="#tab-demo9">T-Shirts</a>
      </li>
      <li>
        <a href="#tab-demo10">Brand</a>
      </li>
    </ul>
  </nav>
  <div id="tab-demo6">
    <h5>Home</h5>
    <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
  </div>
  <div id="tab-demo7">
    <h5>Shop</h5>
    <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
  </div>
  <div id="tab-demo8">
    <h5>Catalog</h5>
    <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
  </div>
  <div id="tab-demo9">
    <h5>T-Shirts</h5>
    <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
  </div>
  <div id="tab-demo10">
    <h5>Brand</h5>
    <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
  </div>
  <pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">onclick</span>=<span class="hljs-string">"$('#tabs').tabs('destroy');"</span>&gt;</span>destroy<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">onclick</span>=<span class="hljs-string">"$('#tabs').tabs('prev');"</span>&gt;</span>prev<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">onclick</span>=<span class="hljs-string">"$('#tabs').tabs('next');"</span>&gt;</span>next<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">onclick</span>=<span class="hljs-string">"$('#tabs').tabs('open', 4);"</span>&gt;</span>open<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">onclick</span>=<span class="hljs-string">"$('#tabs').tabs('close', '#tab4');"</span>&gt;</span>close<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">onclick</span>=<span class="hljs-string">"$('#tabs').tabs('closeAll');"</span>&gt;</span>closeAll<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span>

<span class="hljs-tag">&lt;<span class="hljs-name">nav</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"tabs"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"tabs"</span> <span class="hljs-attr">data-component</span>=<span class="hljs-string">"tabs"</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">ul</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">li</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#tab1"</span>&gt;</span>Home<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">li</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"active"</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#tab2"</span>&gt;</span>Shop<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">li</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#tab3"</span>&gt;</span>Catalog<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">li</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#tab4"</span>&gt;</span>T-Shirts<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">li</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#tab5"</span>&gt;</span>Brand<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">ul</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">nav</span>&gt;</span>

<span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"tab1"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"tab2"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"tab3"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"tab4"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"tab4"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
</pre>
</div>
<h3 class="section-head" id="h-equals"><a href="#h-equals">Equals</a></h3>
<p>Often it is important to create tabs that are equal in width regardless of content. <var>data-equals</var> is here to help you with this task.</p>
<div class="example">
  <nav class="tabs" data-component="tabs" data-equals="true">
    <ul>
      <li class="active">
        <a href="#tab11">Tab 1</a>
      </li>
      <li>
        <a href="#tab12">Tab 2</a>
      </li>
    </ul>
  </nav>
  <div id="tab11">
    <h5>Tab 1</h5>
    <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
  </div>
  <div id="tab12">
    <h5>Tab 2</h5>
    <p>...</p>
  </div>
  <pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">nav</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"tabs"</span> <span class="hljs-attr">data-component</span>=<span class="hljs-string">"tabs"</span> <span class="hljs-attr">data-equals</span>=<span class="hljs-string">"true"</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">ul</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">li</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"active"</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#tab1"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">li</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#tab12"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">ul</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">nav</span>&gt;</span>

<span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"tab1"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"tab2"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
</pre>
</div>
<h3 class="section-head" id="h-livetabs"><a href="#h-livetabs">Livetabs</a></h3>
<p>Livetabs provide a very seamless and smooth experience by blending content and tabs in this kind of live manner. Set up <var>data-live</var> class, and make sure your tabs have this same class. It's that simple.</p>
<div class="example">
  <nav data-component="tabs" data-live=".tab-live" id="livetabs"></nav>
  <div class="tab-live" data-title="General" id="tab-general">
    <h5>General</h5>
    <p>This is very General Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
  </div>
  <div class="tab-live" data-title="Additional" id="tab-additional">
    <h5>Additional</h5>
    <p>Quite an additional is displayed here: Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
  </div>
  <pre class="code skip">&lt;nav <span class="hljs-built_in">id</span>=<span class="hljs-string">"livetabs"</span> data-component=<span class="hljs-string">"tabs"</span> data-live=<span class="hljs-string">".tab-live"</span>&gt;&lt;/nav&gt;

&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">id</span>=<span class="hljs-string">"tab-general"</span> data-title=<span class="hljs-string">"General"</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"tab-live"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">id</span>=<span class="hljs-string">"tab-additional"</span> data-title=<span class="hljs-string">"Additional"</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"tab-live"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
</div>
<h3 class="section-head" id="h-active"><a href="#h-active">Active</a></h3>
<p>To denote active tab, and to let users know where they are, use class <var>active</var>.</p>
<pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">nav</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"tabs"</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">ul</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">li</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">""</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">li</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"active"</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">""</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">li</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">""</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">ul</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">nav</span>&gt;</span>
</pre>
<h3 class="section-head" id="h-settings"><a href="#h-settings">Settings</a></h3>
<h5>equals</h5>
<ul>
  <li>Type: <var>boolean</var></li>
  <li>Default: <var>false</var></li>
</ul>
<p>Making all tabs in a set equal width</p>
<h5>active</h5>
<ul>
  <li>Type: <var>string</var></li>
  <li>Default: <var>false</var></li>
</ul>
<p>Responsible for denoting active tab.</p>
<h5>live</h5>
<ul>
  <li>Type: <var>string</var></li>
  <li>Default: <var>false</var></li>
</ul>
<p>Responsible for live tabs (see <a href="#h-livetabs">Livetabs</a> example)</p>
<h5>hash</h5>
<ul>
  <li>Type: <var>boolean</var></li>
  <li>Default: <var>true</var></li>
</ul>
<p>This will add a hash and an ID to the URL.</p>
<h3 class="section-head" id="h-callbacks"><a href="#h-callbacks">Callbacks</a></h3>
<h5>init</h5>
<pre class="code skip">$(<span class="hljs-string">'#tabs'</span>).on(<span class="hljs-string">'init.tabs'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});
</pre>
<h5>next</h5>
<pre class="code skip">$(<span class="hljs-string">'#tabs'</span>).on(<span class="hljs-string">'next.tabs'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params">$item</span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});
</pre>
<h5>prev</h5>
<pre class="code skip">$(<span class="hljs-string">'#tabs'</span>).on(<span class="hljs-string">'prev.tabs'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params">$item</span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});
</pre>
<h5>open</h5>
<pre class="code skip">$(<span class="hljs-string">'#tabs'</span>).on(<span class="hljs-string">'open.tabs'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});
</pre>
<h5>opened</h5>
<pre class="code skip">$(<span class="hljs-string">'#tabs'</span>).on(<span class="hljs-string">'opened.tabs'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});
</pre>
<h5>close</h5>
<pre class="code skip">$(<span class="hljs-string">'#tabs'</span>).on(<span class="hljs-string">'close.tabs'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});
</pre>
<h5>closed</h5>
<pre class="code skip">$(<span class="hljs-string">'#tabs'</span>).on(<span class="hljs-string">'closed.tabs'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});
</pre>
<h3 class="section-head" id="h-api"><a href="#h-api">API</a></h3>
<h5>open</h5>
<pre class="code skip"><span class="hljs-variable">$(</span><span class="hljs-string">'#tabs'</span>).tabs(<span class="hljs-string">'open'</span>, <span class="hljs-number">1</span>);

<span class="hljs-regexp">//</span> <span class="hljs-keyword">or</span>

<span class="hljs-variable">$(</span><span class="hljs-string">'#tabs'</span>).tabs(<span class="hljs-string">'open'</span>, <span class="hljs-string">'#tab1'</span>);
</pre>
<h5>close</h5>
<pre class="code skip"><span class="hljs-variable">$(</span><span class="hljs-string">'#tabs'</span>).tabs(<span class="hljs-string">'close'</span>, <span class="hljs-number">1</span>);

<span class="hljs-regexp">//</span> <span class="hljs-keyword">or</span>

<span class="hljs-variable">$(</span><span class="hljs-string">'#tabs'</span>).tabs(<span class="hljs-string">'close'</span>, <span class="hljs-string">'#tab1'</span>);
</pre>
<h5>closeAll</h5>
<pre class="code skip"><span class="hljs-variable">$(</span><span class="hljs-string">'#tabs'</span>).tabs(<span class="hljs-string">'closeAll'</span>);
</pre>
<h5>next</h5>
<pre class="code skip"><span class="hljs-variable">$(</span><span class="hljs-string">'#tabs'</span>).tabs(<span class="hljs-string">'next'</span>);
</pre>
<h5>prev</h5>
<pre class="code skip"><span class="hljs-variable">$(</span><span class="hljs-string">'#tabs'</span>).tabs(<span class="hljs-string">'prev'</span>);
</pre>
<h5>destroy</h5>
<pre class="code skip"><span class="hljs-variable">$(</span><span class="hljs-string">'#tabs'</span>).tabs(<span class="hljs-string">'destroy'</span>);
</pre>