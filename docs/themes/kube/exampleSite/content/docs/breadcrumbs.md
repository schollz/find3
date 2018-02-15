+++
date = "2017-04-10T16:41:54+01:00"
weight = 110
description = "Traces for the users so they won't get lost"
title = "Breadcrumbs"
draft = false
bref= "Breadcrumbs in Kube are formed as lists with default separator. This separator can be customized with ease by simply changing a single CSS line"
toc = true
+++

<h3 class="section-head" id="h-base"><a href="#h-base">Base</a></h3>
<div class="example">
  <nav class="breadcrumbs">
    <ul>
      <li>
        <a href="#">Home</a>
      </li>
      <li>
        <a href="#">Shop</a>
      </li>
      <li>
        <a href="#">Catalog</a>
      </li>
      <li>
        <a href="#">T-Shirts</a>
      </li>
      <li><span>Brand</span></li>
    </ul>
  </nav>
  <pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">nav</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"breadcrumbs"</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">ul</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">li</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">""</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">li</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">span</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">span</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">ul</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">nav</span>&gt;</span>
</pre>
</div>
<h3 class="section-head" id="h-centered"><a href="#h-centered">Centered</a></h3>
<div class="example">
  <nav class="breadcrumbs push-center">
    <ul>
      <li>
        <a href="#">Home</a>
      </li>
      <li>
        <a href="#">Shop</a>
      </li>
      <li>
        <a href="#">Catalog</a>
      </li>
      <li>
        <a href="#">T-Shirts</a>
      </li>
      <li class="active">
        <a href="">Brand</a>
      </li>
    </ul>
  </nav>
  <pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">nav</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"breadcrumbs push-center"</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">ul</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">li</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">""</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">li</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"active"</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">""</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">ul</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">nav</span>&gt;</span>
</pre>
</div>
<h3 class="section-head" id="h-custom-separator"><a href="#h-custom-separator">Custom Separator</a></h3>
<div class="example">
  <nav class="breadcrumbs" id="breadcrumbs-custom-separator">
    <ul>
      <li>
        <a href="#">Home</a>
      </li>
      <li>
        <a href="#">Shop</a>
      </li>
      <li>
        <a href="#">Catalog</a>
      </li>
      <li>
        <a href="#">T-Shirts</a>
      </li>
      <li><span>Brand</span></li>
    </ul>
  </nav>
  <pre class="code skip"><span class="hljs-comment">// css</span>
<span class="hljs-meta">#breadcrumbs-custom-separator li:after {</span>
<span class="hljs-symbol">    content:</span> <span class="hljs-string">'&gt;'</span>;
}

<span class="hljs-comment">// html</span>
<span class="hljs-params">&lt;nav id="breadcrumbs-custom-separator" class="breadcrumbs"&gt;</span>
    <span class="hljs-params">&lt;ul&gt;</span>
        <span class="hljs-params">&lt;li&gt;</span><span class="hljs-params">&lt;a href=""&gt;</span>...<span class="hljs-params">&lt;/a&gt;</span><span class="hljs-params">&lt;/li&gt;</span>
        <span class="hljs-params">&lt;li&gt;</span><span class="hljs-params">&lt;span&gt;</span>...<span class="hljs-params">&lt;/span&gt;</span><span class="hljs-params">&lt;/li&gt;</span>
    <span class="hljs-params">&lt;/ul&gt;</span>
<span class="hljs-params">&lt;/nav&gt;</span>
</pre>
</div>