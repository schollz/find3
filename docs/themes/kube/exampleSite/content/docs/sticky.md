+++
date = "2017-04-10T16:42:26+01:00"
title = "Sticky"
draft = false
weight = 160
description = "Make navigation menus stick to the top on scroll"
bref="Sticky navigation makes it easy to affix menus to the top of the page whenever a user scrolls the page down. This could be very helpful and useful on long pages. Sticky is disabled on mobile devices to avoid unwanted content overlaps and to preserve valuable screen real estate"
toc = true
+++

<h3 class="section-head" id="h-demo"><a href="#h-demo">Demo</a></h3>
<p>Scroll down to fix the navigation.</p>
<div data-component="sticky" id="navbar-demo">
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
<pre class="code skip">&lt;<span class="hljs-keyword">div</span> data-component=<span class="hljs-string">"sticky"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
<h3 class="section-head" id="h-settings"><a href="#h-settings">Settings</a></h3>
<h5>offset</h5>
<ul>
  <li>Type: <var>int</var></li>
  <li>Default: <var>0</var></li>
</ul>
<p>Sets top offset in pixels when navigation is fixed.</p>
<h3 class="section-head" id="h-callbacks"><a href="#h-callbacks">Callbacks</a></h3>
<h5>fixed</h5>
<p>Using this callback, you can act upon menu becoming fixed at the top of the page.</p>
<pre class="code skip">$(<span class="hljs-string">'#my-nav'</span>).on(<span class="hljs-string">'fixed.sticky'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});
</pre>
<h5>unfixed</h5>
<p>Whenever fixed position of the menu is released, you can do something with this callback.</p>
<pre class="code skip">$(<span class="hljs-string">'#my-nav'</span>).on(<span class="hljs-string">'unfixed.sticky'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});
</pre><br>
<br>
<p><em>Some scrollable material</em></p>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>