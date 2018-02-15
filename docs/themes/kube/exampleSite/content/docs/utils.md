+++
date = "2017-04-10T16:43:30+01:00"
draft = false
weight = 280
description = "A dozen of useful utilities that come with Kube"
title = "Utils"
bref= "Here you can find over a dozen of examples of little tiny utilities, that can make developer's life that much easier, and your project progress that much faster."
toc = true
+++

<h3 class="section-head" id="h-group"><a href="#h-group">Group</a></h3>
<p>Combines float elements to group with a clearfix.</p>
<div class="example">
  <div class="group">
    <div class="float-left">
      This text is visibly floating left
    </div>
    <div class="float-right">
      This text looks like a case of right float
    </div>
  </div>
  <pre class="code skip">&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"group"</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"float-left"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"float-right"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
</div>
<h3 class="section-head" id="h-visibility"><a href="#h-visibility">Visibility</a></h3>
<p>Below is an invisible <code>div</code> with class <var>invisible</var>. You can't see it, because it is invisible. You can't see invisible things.</p>
<div class="example">
  <div class="invisible">
    invisible
  </div>
  <div class="visible">
    But you can see this <code>div</code> because it has <var>visible</var> class.
  </div>
  <pre class="code skip">&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"invisible"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"visible"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
</div>
<h3 class="section-head" id="h-display"><a href="#h-display">Display</a></h3>
<p>This is a very useful little bit. Whenever you need to hide some text or an element on a small screen, just throw in <var>hide-sm</var> class. Or, if you need to specifically show something only on small screens, <var>show-sm</var> class is yours for the job. You can also just plain hide stuff with <var>hide</var> class. Can you see red words "I'm hidden" below? Exactly.</p>
<div class="example">
  <div class="hide red">
    I'm hidden
  </div>
  <div class="hide-sm">
    This text will not show up on a small screen.
  </div>
  <p>Resize your window and trust us on this <span class="hide-sm">↑</span>&nbsp;<span class="show-sm">↓</span></p>
  <div class="show-sm">
    This text will only show up on a small screen
  </div>
  <pre class="code skip">&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"hide"</span>&gt;I'm hidden&lt;/<span class="hljs-keyword">div</span>&gt;
&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"hide-sm"</span>&gt;This <span class="hljs-built_in">text</span> will <span class="hljs-keyword">not</span> show up <span class="hljs-keyword">on</span> a small screen.&lt;/<span class="hljs-keyword">div</span>&gt;
&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"show-sm"</span>&gt;This will only show up <span class="hljs-keyword">on</span> a small screen&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
</div>
<h3 class="section-head" id="h-print"><a href="#h-print">Print</a></h3>
<p>One more neat feature of Kube. It helps you produce better ready-to-print pages by simply hiding irrelevant things.</p>
<div class="example">
  <div class="hide-print">
    This will be hidden on print, because it is some sort of web-specific thing.
  </div>
  <div class="show-print">
    This will be printed, because this text is somehow more relevant on paper than on screen.
  </div>
  <pre class="code skip">&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"hide-print"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"show-print"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
</div>
<h3 class="section-head" id="h-video-container"><a href="#h-video-container">Video Container</a></h3>
<p>Helps to serve responsive video to various devices.</p>
<div class="example">
  <div class="video-container">
    <iframe allowfullscreen frameborder="0" height="315" src="https://www.youtube.com/embed/nywsA8wCCfY" width="560"></iframe>
  </div>
  <pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"video-container"</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">iframe</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">iframe</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
</pre>
</div>
<h3 class="section-head" id="h-close"><a href="#h-close">Close</a></h3>
<p>Closing icon to close anything you want.</p>
<div class="example">
  <span class="close small"></span>
  <pre class="code skip">&lt;span <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"close small"</span>&gt;<span class="xml"><span class="hljs-tag">&lt;/<span class="hljs-name">span</span>&gt;</span></span></pre>
</div>
<div class="example">
  <span class="close"></span>
  <pre class="code skip">&lt;span <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"close"</span>&gt;<span class="xml"><span class="hljs-tag">&lt;/<span class="hljs-name">span</span>&gt;</span></span></pre>
</div>
<div class="example">
  <span class="close big"></span>
  <pre class="code skip">&lt;span <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"close big"</span>&gt;<span class="xml"><span class="hljs-tag">&lt;/<span class="hljs-name">span</span>&gt;</span></span></pre>
</div>
<h3 class="section-head" id="h-caret"><a href="#h-caret">Caret</a></h3>
<p>Kube has already built-in four-direction caret.</p>
<div class="example">
  <span class="caret down"></span> <span class="caret up"></span> <span class="caret left"></span> <span class="caret right"></span>
  <pre class="code skip">&lt;span <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"caret down"</span>&gt;<span class="xml"><span class="hljs-tag">&lt;/<span class="hljs-name">span</span>&gt;</span></span>
&lt;span <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"caret up"</span>&gt;<span class="xml"><span class="hljs-tag">&lt;/<span class="hljs-name">span</span>&gt;</span></span>
&lt;span <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"caret left"</span>&gt;<span class="xml"><span class="hljs-tag">&lt;/<span class="hljs-name">span</span>&gt;</span></span>
&lt;span <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"caret right"</span>&gt;<span class="xml"><span class="hljs-tag">&lt;/<span class="hljs-name">span</span>&gt;</span></span>
</pre>
</div>
<p>Example of usage:</p>
<div class="example">
  <button class="button">Button <span class="caret down white"></span></button> &nbsp;&nbsp;&nbsp; <a href="#">Link <span class="caret up"></span></a> &nbsp;&nbsp;&nbsp; Text <span class="caret down"></span> &nbsp;&nbsp;&nbsp; <button class="button secondary outline">Button <span class="caret down"></span></button>
</div>
<h3 class="section-head" id="h-icons"><a href="#h-icons">Icons</a></h3>
<p>Some useful icons are already built-in to Kube.</p>
<div class="example">
  <i class="kube-search"></i>
  <pre class="code skip">&lt;i <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"kube-search"</span>&gt;<span class="xml"><span class="hljs-tag">&lt;/<span class="hljs-name">i</span>&gt;</span></span></pre>
</div>
<div class="example">
  <i class="kube-menu"></i>
  <pre class="code skip">&lt;i <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"kube-menu"</span>&gt;<span class="xml"><span class="hljs-tag">&lt;/<span class="hljs-name">i</span>&gt;</span></span></pre>
</div>
<div class="example">
  <i class="kube-calendar"></i>
  <pre class="code skip">&lt;i <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"kube-calendar"</span>&gt;<span class="xml"><span class="hljs-tag">&lt;/<span class="hljs-name">i</span>&gt;</span></span></pre>
</div>