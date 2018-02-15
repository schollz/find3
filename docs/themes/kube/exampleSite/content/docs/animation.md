+++
description = "Add some motion, shaking, pulsing, sliding and more"
title = "Animation"
date = "2017-04-10T16:43:08+01:00"
draft = false
weight = 200
bref="Although it is quite easy to overuse animation effects, Kube makes it pretty easy to create meaningful, smooth and overall nice looking animation. Feel free to click every button below to see what it does, and then only use those effects that are essential to your project"
toc = true
script = 'animation'
+++

<h3 class="section-head" id="h-slide"><a href="#h-slide">Slide</a></h3>
<div class="example">
  <div class="demo-animation-wrap">
    <div class="demo-animation-box">
      <div id="animation-box-slide-up">
        slideUp
      </div>
    </div>
    <div class="demo-animation-box">
      <div class="hide" id="animation-box-slide-down">
        slideDown
      </div>
    </div>
  </div><a class="demo-animation-btn" data-back="slideDown" data-el="#animation-box-slide-up" data-to="slideUp" href="#" id="slide-up-btn">Slide Up</a> <a class="demo-animation-btn" data-back="slideUp" data-el="#animation-box-slide-down" data-to="slideDown" href="#" id="slide-down-btn">Slide Down</a>
  <pre class="code"><span class="hljs-variable">$(</span><span class="hljs-string">'#element'</span>).animation(<span class="hljs-string">'slideUp'</span>);
<span class="hljs-variable">$(</span><span class="hljs-string">'#element'</span>).animation(<span class="hljs-string">'slideDown'</span>);</pre>
</div>
<h3 class="section-head" id="h-fade"><a href="#h-fade">Fade</a></h3>
<div class="example">
  <div class="demo-animation-wrap">
    <div class="demo-animation-box">
      <div class="hide" id="animation-box-fade-in">
        fadeIn
      </div>
    </div>
    <div class="demo-animation-box">
      <div id="animation-box-fade-out">
        fadeOut
      </div>
    </div>
  </div><a class="demo-animation-btn" data-back="fadeOut" data-el="#animation-box-fade-in" data-to="fadeIn" href="#" id="fade-in-btn">Fade In</a> <a class="demo-animation-btn" data-back="fadeIn" data-el="#animation-box-fade-out" data-to="fadeOut" href="#" id="fade-out-btn">Fade Out</a>
  <pre class="code"><span class="hljs-variable">$(</span><span class="hljs-string">'#element'</span>).animation(<span class="hljs-string">'fadeIn'</span>);
<span class="hljs-variable">$(</span><span class="hljs-string">'#element'</span>).animation(<span class="hljs-string">'fadeOut'</span>);</pre>
</div>
<h3 class="section-head" id="h-flip"><a href="#h-flip">Flip</a></h3>
<div class="example">
  <div class="demo-animation-wrap">
    <div class="demo-animation-box">
      <div class="hide" id="animation-box-flip-in">
        flipIn
      </div>
    </div>
    <div class="demo-animation-box">
      <div id="animation-box-flip-out">
        flipOut
      </div>
    </div>
  </div><a class="demo-animation-btn" data-back="flipOut" data-el="#animation-box-flip-in" data-to="flipIn" href="#" id="flip-in-btn">Flip In</a> <a class="demo-animation-btn" data-back="flipIn" data-el="#animation-box-flip-out" data-to="flipOut" href="#" id="flip-out-btn">Flip Out</a>
  <pre class="code"><span class="hljs-variable">$(</span><span class="hljs-string">'#element'</span>).animation(<span class="hljs-string">'flipIn'</span>);
<span class="hljs-variable">$(</span><span class="hljs-string">'#element'</span>).animation(<span class="hljs-string">'flipOut'</span>);</pre>
</div>
<h3 class="section-head" id="h-zoom"><a href="#h-zoom">Zoom</a></h3>
<div class="example">
  <div class="demo-animation-wrap">
    <div class="demo-animation-box">
      <div class="hide" id="animation-box-zoom-in">
        zoomIn
      </div>
    </div>
    <div class="demo-animation-box">
      <div id="animation-box-zoom-out">
        zoomOut
      </div>
    </div>
  </div><a class="demo-animation-btn" data-back="zoomOut" data-el="#animation-box-zoom-in" data-to="zoomIn" href="#" id="zoom-in-btn">Zoom In</a> <a class="demo-animation-btn" data-back="zoomIn" data-el="#animation-box-zoom-out" data-to="zoomOut" href="#" id="zoom-out-btn">Zoom Out</a>
  <pre class="code"><span class="hljs-variable">$(</span><span class="hljs-string">'#element'</span>).animation(<span class="hljs-string">'zoomIn'</span>);
<span class="hljs-variable">$(</span><span class="hljs-string">'#element'</span>).animation(<span class="hljs-string">'zoomOut'</span>);</pre>
</div>
<h3 class="section-head" id="h-rotate"><a href="#h-rotate">Rotate</a></h3>
<div class="example">
  <div class="demo-animation-wrap">
    <div class="demo-animation-box">
      <div id="animation-box-rotate">
        rotate
      </div>
    </div>
  </div><a class="demo-animation-btn" href="#" id="rotate-btn">Rotate</a>
  <pre class="code"><span class="hljs-variable">$(</span><span class="hljs-string">'#element'</span>).animation(<span class="hljs-string">'rotate'</span>);</pre>
</div>
<h3 class="section-head" id="h-shake"><a href="#h-shake">Shake</a></h3>
<div class="example">
  <div class="demo-animation-wrap">
    <div class="demo-animation-box">
      <div id="animation-box-shake">
        shake
      </div>
    </div>
  </div><a class="demo-animation-btn" href="#" id="shake-btn">Shake</a>
  <pre class="code"><span class="hljs-variable">$(</span><span class="hljs-string">'#element'</span>).animation(<span class="hljs-string">'shake'</span>);</pre>
</div>
<h3 class="section-head" id="h-pulse"><a href="#h-pulse">Pulse</a></h3>
<div class="example">
  <div class="demo-animation-wrap">
    <div class="demo-animation-box">
      <div id="animation-box-pulse">
        pulse
      </div>
    </div>
  </div><a class="demo-animation-btn" href="#" id="pulse-btn">Pulse</a>
  <pre class="code"><span class="hljs-variable">$(</span><span class="hljs-string">'#element'</span>).animation(<span class="hljs-string">'pulse'</span>);</pre>
</div>
<h3 class="section-head" id="h-slide-in"><a href="#h-slide-in">Slide In</a></h3>
<div class="example">
  <div class="demo-animation-wrap">
    <div class="demo-animation-box">
      <div class="hide" id="animation-box-slide-in-right">
        slideInRight
      </div>
    </div>
    <div class="demo-animation-box">
      <div class="hide" id="animation-box-slide-in-left">
        slideInLeft
      </div>
    </div>
    <div class="demo-animation-box">
      <div class="hide" id="animation-box-slide-in-down">
        slideInDown
      </div>
    </div>
  </div><a class="demo-animation-btn" data-back="slideOutRight" data-el="#animation-box-slide-in-right" data-to="slideInRight" href="#" id="slide-in-right-btn">Slide In Right</a> <a class="demo-animation-btn" data-back="slideOutLeft" data-el="#animation-box-slide-in-left" data-to="slideInLeft" href="#" id="slide-in-left-btn">Slide In Left</a> <a class="demo-animation-btn" data-back="slideOutUp" data-el="#animation-box-slide-in-down" data-to="slideInDown" href="#" id="slide-in-down-btn">Slide In Down</a>
  <pre class="code"><span class="hljs-variable">$(</span><span class="hljs-string">'#element'</span>).animation(<span class="hljs-string">'slideInRight'</span>);
<span class="hljs-variable">$(</span><span class="hljs-string">'#element'</span>).animation(<span class="hljs-string">'slideInLeft'</span>);
<span class="hljs-variable">$(</span><span class="hljs-string">'#element'</span>).animation(<span class="hljs-string">'slideInDown'</span>);</pre>
</div>
<h3 class="section-head" id="h-slide-out"><a href="#h-slide-out">Slide Out</a></h3>
<div class="example">
  <div class="demo-animation-wrap">
    <div class="demo-animation-box">
      <div id="animation-box-slide-out-right">
        slideOutRight
      </div>
    </div>
    <div class="demo-animation-box">
      <div id="animation-box-slide-out-left">
        slideOutLeft
      </div>
    </div>
    <div class="demo-animation-box">
      <div id="animation-box-slide-out-up">
        slideOutUp
      </div>
    </div>
  </div><a class="demo-animation-btn" data-back="slideInRight" data-el="#animation-box-slide-out-right" data-to="slideOutRight" href="#" id="slide-out-right-btn">Slide Out Right</a> <a class="demo-animation-btn" data-back="slideInLeft" data-el="#animation-box-slide-out-left" data-to="slideOutLeft" href="#" id="slide-out-left-btn">Slide Out Left</a> <a class="demo-animation-btn" data-back="slideInDown" data-el="#animation-box-slide-out-up" data-to="slideOutUp" href="#" id="slide-out-up-btn">Slide Out Up</a>
  <pre class="code"><span class="hljs-variable">$(</span><span class="hljs-string">'#element'</span>).animation(<span class="hljs-string">'slideOutRight'</span>);
<span class="hljs-variable">$(</span><span class="hljs-string">'#element'</span>).animation(<span class="hljs-string">'slideOutLeft'</span>);
<span class="hljs-variable">$(</span><span class="hljs-string">'#element'</span>).animation(<span class="hljs-string">'slideOutUp'</span>);</pre>
</div>
<h3 class="section-head" id="h-callback"><a href="#h-callback">Callback</a></h3>
<pre class="code">$(<span class="hljs-string">'#element'</span>).animation(<span class="hljs-string">'fadeIn'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// code ...</span>
});</pre>