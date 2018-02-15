+++
date = "2017-04-10T16:41:16+01:00"
weight = 70
description = "Full variety of pressible, clickable and pushable buttons"
title = "Buttons"
draft = false
bref =  "Buttons in Kube are minimalistic, designed for instant and convenient customization, and are ready to be pushed, pressed, clicked and manipulated in whichever ways. A lot of interactive components use buttons, and you should too!"
toc = true
+++

<h3 class="section-head" id="h-primary"><a href="#h-primary">Primary</a></h3>
<div class="example">
  <button class="button">Button</button> <a class="button" href="#">Button</a>
  <pre class="code"><span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button"</span>&gt;</span>Button<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button"</span>&gt;</span>Button<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span></pre>
</div>
<h3 class="section-head" id="h-secondary"><a href="#h-secondary">Secondary</a></h3>
<div class="example">
  <button class="button secondary">Button</button> <a class="button secondary" href="#">Button</a>
  <pre class="code"><span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button secondary"</span>&gt;</span>Button<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button secondary"</span>&gt;</span>Button<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span></pre>
</div>
<h3 class="section-head" id="h-outline"><a href="#h-outline">Outline</a></h3>
<p>Outline class does exactly what it is supposed to do: styles a button with an outline with no fill color.</p>
<div class="example">
  <button class="button outline">Button</button> <button class="button secondary outline">Button</button>
  <pre class="code">&lt;<span class="hljs-keyword">button</span> class=<span class="hljs-string">"button outline"</span>&gt;<span class="hljs-keyword">Button</span>&lt;/<span class="hljs-keyword">button</span>&gt;
&lt;<span class="hljs-keyword">button</span> class=<span class="hljs-string">"button secondary outline"</span>&gt;<span class="hljs-keyword">Button</span>&lt;/<span class="hljs-keyword">button</span>&gt;</pre>
</div>
<h3 class="section-head" id="h-disabled"><a href="#h-disabled">Disabled</a></h3>
<p>Disabled buttons are automatically styles with muted colors, inactive and not clickable. As you can see, there's no need to set disabled class for the button element, and disabled argument is enough. Having this said, both a link and a button input require disabled class.</p>
<div class="example">
  <a class="button disabled" href="#" role="button">Link</a> <button class="button secondary" disabled>Button</button> <button class="button outline" disabled>Button</button> <button class="button secondary outline" disabled>Button</button>
</div>
<pre class="code"><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button disabled"</span> <span class="hljs-attr">role</span>=<span class="hljs-string">"button"</span>&gt;</span>Link<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button secondary"</span> <span class="hljs-attr">disabled</span>&gt;</span>Button<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span></pre>
<h3 class="section-head" id="h-small"><a href="#h-small">Small</a></h3>
<p>Buttons come in all sizes and shapes.</p>
<div class="example">
  <a class="button small" href="#" role="button">Link</a> <button class="button secondary small">Button</button>
</div>
<pre class="code"><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button small"</span> <span class="hljs-attr">role</span>=<span class="hljs-string">"button"</span>&gt;</span>Link<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button secondary small"</span>&gt;</span>Button<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span></pre>
<h3 class="section-head" id="h-big"><a href="#h-big">Big</a></h3>
<div class="example">
  <a class="button big" href="#" role="button">Link</a> <button class="button secondary big">Button</button>
</div>
<pre class="code"><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button big"</span> <span class="hljs-attr">role</span>=<span class="hljs-string">"button"</span>&gt;</span>Link<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button secondary big"</span>&gt;</span>Button<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span></pre>
<h3 class="section-head" id="h-large"><a href="#h-large">Large</a></h3>
<div class="example">
  <a class="button large" href="#" role="button">Link</a> <button class="button secondary large">Button</button>
</div>
<pre class="code"><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button large"</span> <span class="hljs-attr">role</span>=<span class="hljs-string">"button"</span>&gt;</span>Link<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button secondary large"</span>&gt;</span>Button<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span></pre>
<h3 class="section-head" id="h-upper"><a href="#h-upper">Upper</a></h3>
<div class="example">
  <a class="button upper" href="#" role="button">Link</a> <button class="button secondary upper">Button</button> <a class="button upper outline" href="#" role="button">Link</a> <button class="button secondary upper outline">Button</button>
</div>
<pre class="code"><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button upper"</span> <span class="hljs-attr">role</span>=<span class="hljs-string">"button"</span>&gt;</span>Link<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button secondary upper"</span>&gt;</span>Button<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button upper outline"</span> <span class="hljs-attr">role</span>=<span class="hljs-string">"button"</span>&gt;</span>Link<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button secondary upper outline"</span>&gt;</span>Button<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span></pre>
<h3 class="section-head" id="h-round"><a href="#h-round">Round</a></h3>
<div class="example">
  <a class="button round" href="#" role="button">Link</a> <button class="button secondary round">Button</button> <a class="button round outline" href="#" role="button">Link</a> <button class="button secondary round outline">Button</button>
</div>
<pre class="code"><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button round"</span> <span class="hljs-attr">role</span>=<span class="hljs-string">"button"</span>&gt;</span>Link<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button secondary round"</span>&gt;</span>Button<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button round outline"</span> <span class="hljs-attr">role</span>=<span class="hljs-string">"button"</span>&gt;</span>Link<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button secondary round outline"</span>&gt;</span>Button<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span></pre>
<h3 class="section-head" id="h-inverted"><a href="#h-inverted">Inverted</a></h3>
<p>For use on darker backgrounds, you can just introduce inverted class to have your button flip its color to the opposite one.</p>
<div class="example bg-darkgray">
  <a class="button inverted" href="#" role="button">Link</a> &nbsp; <button class="button inverted outline">Button</button> &nbsp;
  <pre class="code"><span class="hljs-tag">&lt;<span class="hljs-name">a</span> <span class="hljs-attr">href</span>=<span class="hljs-string">"#"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button inverted"</span> <span class="hljs-attr">role</span>=<span class="hljs-string">"button"</span>&gt;</span>Link<span class="hljs-tag">&lt;/<span class="hljs-name">a</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button inverted outline"</span>&gt;</span>Button<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span></pre>
</div>
<h3 class="section-head" id="h-width"><a href="#h-width">Width</a></h3>
<div class="example">
  <p><button class="button w100">100%</button></p>
  <p><button class="button secondary w50">50%</button></p>
  <pre class="code"><span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button w100"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button secondary w50"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span></pre>
</div>
<h3 class="section-head" id="h-icons"><a href="#h-icons">Icons</a></h3>
<div class="example">
  <button class="button"><i class="kube-calendar"></i></button> <button class="button secondary"><i class="kube-calendar"></i> Change</button> <button class="button outline"><i class="kube-search"></i></button> <button class="button secondary outline"><i class="kube-search"></i> Search</button>
  <pre class="code"><span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button"</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">i</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"kube-calendar"</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">i</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button secondary"</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">i</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"kube-calendar"</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">i</span>&gt;</span> Change<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span>

<span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button outline"</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">i</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"kube-search"</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">i</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button secondary outline"</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">i</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"kube-search"</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">i</span>&gt;</span> Search<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span></pre>
</div>
<h3 class="section-head" id="h-custom"><a href="#h-custom">Custom</a></h3>
<p>With some Sass magic, you can customize your buttons in a snap. Just include a color class declaration and set the color itself, and you're done.</p>
<div class="example">
  <button class="button red">Button</button> <button class="button red outline">Button</button>
  <pre class="code"><span class="hljs-comment">// scss</span>
.button.red {
    <span class="hljs-comment">// $text-color, $back-color</span>
    @include <span class="hljs-keyword">button</span>(<span class="hljs-meta">#fff, #ff3366);</span>
}

<span class="hljs-comment">// html</span>
&lt;<span class="hljs-keyword">button</span> class=<span class="hljs-string">"button red"</span>&gt;<span class="hljs-keyword">Button</span>&lt;/<span class="hljs-keyword">button</span>&gt;
&lt;<span class="hljs-keyword">button</span> class=<span class="hljs-string">"button red outline"</span>&gt;<span class="hljs-keyword">Button</span>&lt;/<span class="hljs-keyword">button</span>&gt;</pre>
</div>