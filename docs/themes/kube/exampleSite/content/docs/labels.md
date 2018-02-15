+++
date = "2017-04-10T16:40:43+01:00"
title = "Labels"
draft = false
weight = 80
description = "Label things with all sorts of labels"
bref = "Labels have all standard Kube styling options available: outline, states (success, error, warning), inverted color, and more. You can also customize labels to your requirements, and place them inside other elements, such as buttons, for example"
toc = true
+++

<h3 class="section-head" id="h-base"><a href="#h-base">Base</a></h3>
<div class="example">
  <span class="label">Default</span> &nbsp; <span class="label success">Success</span> &nbsp; <span class="label error">Error</span> &nbsp; <span class="label warning">Warning</span> &nbsp; <span class="label focus">Focus</span> &nbsp; <span class="label black">Black</span> &nbsp; <span class="example-inverted-box"><span class="label inverted">Inverted</span></span>
</div>
<h3 class="section-head" id="h-outline"><a href="#h-outline">Outline</a></h3>
<div class="example">
  <span class="label outline">Default</span> &nbsp; <span class="label success outline">Success</span> &nbsp; <span class="label error outline">Error</span> &nbsp; <span class="label warning outline">Warning</span> &nbsp; <span class="label focus outline">Focus</span> &nbsp; <span class="label black outline">Black</span> &nbsp; <span class="example-inverted-box"><span class="label inverted outline">Inverted</span></span>
</div>
<h3 class="section-head" id="h-upper"><a href="#h-upper">Upper</a></h3>
<div class="example">
  <span class="label upper">Default</span> &nbsp; <span class="label success upper">Success</span> &nbsp; <span class="label error upper">Error</span> &nbsp; <span class="label warning upper">Warning</span> &nbsp; <span class="label focus upper">Focus</span> &nbsp; <span class="label black upper">Black</span> &nbsp; <span class="example-inverted-box"><span class="label inverted upper">Inverted</span></span>
</div>
<h3 class="section-head" id="h-tag"><a href="#h-tag">Tag</a></h3>
<div class="example">
  <span class="label tag"><a href="#">Default</a></span> &nbsp; <span class="label tag success">Success</span> &nbsp; <span class="label tag error">Error</span> &nbsp; <span class="label tag warning">Warning</span> &nbsp; <span class="label tag focus"><a href="#">Focus</a></span> &nbsp; <span class="label tag black">Black</span> &nbsp; <span class="example-inverted-box"><span class="label tag inverted">Inverted</span></span>
</div>
<h3 class="section-head" id="h-badges"><a href="#h-badges">Badges</a></h3>
<p>Labels are ideal for use as badges with badge class.</p>
<div class="example">
  <span class="label badge"><a href="#">1</a></span> <span class="label badge error">2</span> <span class="label badge success">3</span> <span class="label badge warning">4</span> <span class="label badge focus"><a href="#">5</a></span> <span class="label badge black">6</span> <span class="example-inverted-box"><span class="label badge inverted">7</span></span>
</div>
<h3 class="section-head" id="h-outline-badges"><a href="#h-outline-badges">Outline Badges</a></h3>
<div class="example">
  <span class="label badge outline"><a href="#">1</a></span> <span class="label badge error outline">2</span> <span class="label badge success outline">3</span> <span class="label badge warning outline">4</span> <span class="label badge focus outline"><a href="#">5</a></span> <span class="label badge black outline">6</span> <span class="example-inverted-box"><span class="label badge inverted outline">7</span></span>
</div>
<h3 class="section-head" id="h-custom"><a href="#h-custom">Custom</a></h3>
<p>You can use mixins to customize your labels and badges.</p>
<div class="example">
  <span class="label custom">Label</span> &nbsp; <span class="label badge custom">1</span> &nbsp; <span class="label tag custom">Tag</span>
  <pre class="code skip"><span class="hljs-comment">&lt;!-- scss --&gt;</span>
.label.custom {
    // $text-color, $back-color
    @include label(#fff, #ea48a7);
}

<span class="hljs-comment">&lt;!-- html --&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">span</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"label custom"</span>&gt;</span>Label<span class="hljs-tag">&lt;/<span class="hljs-name">span</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">span</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"label badge custom"</span>&gt;</span>1<span class="hljs-tag">&lt;/<span class="hljs-name">span</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">span</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"label tag custom"</span>&gt;</span>Tag<span class="hljs-tag">&lt;/<span class="hljs-name">span</span>&gt;</span>
</pre>
</div>