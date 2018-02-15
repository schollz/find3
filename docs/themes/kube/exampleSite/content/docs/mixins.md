+++
weight = 230
description = "Mixins save countless hours and bring results faster"
date = "2017-04-10T16:43:36+01:00"
title = "Mixins"
draft = false
bref="Mixins are a great way to produce things and effects way faster with Sass then with pure CSS. Kube has a lot to offer in this respect; feel free to use any of these mixins with any Kube components."
toc = true
+++

<h3 class="section-head" id="h-get-started"><a href="#h-get-started">Get Started</a></h3>
<p>Kube has been designed to help you with web development, that's why it's so easy to use Kube when building websites. To move forward quickly and efficiently, just link <code>kube.scss</code> from Kube package: this file contains variables, mixins and everything you need to simplify daily routine tasks.</p>
<p>For example, import <code>kube.scss</code> into your <code>master.scss</code> styles file, which you will later compile into <code>master.css</code></p>
<pre class="code skip"><span class="hljs-comment">// master.scss</span>
@<span class="hljs-keyword">import</span> <span class="hljs-string">"dist/scss/kube.scss"</span>;
</pre>
<p>Now all Kube's variables and mixins are readily available in <code>master.scss</code>, and you can use them whenever needed.</p>
<pre class="code skip"><span class="hljs-comment">// master.scss</span>
@<span class="hljs-keyword">import</span> <span class="hljs-string">"dist/scss/kube.scss"</span>;

<span class="hljs-comment">// use mixins</span>
<span class="hljs-selector-id">#my-layout</span> {
    @<span class="hljs-keyword">include</span> flex;
}

<span class="hljs-comment">// use variables</span>
<span class="hljs-selector-id">#my-layout</span> {
    <span class="hljs-attribute">padding</span>: <span class="hljs-variable">$base-line</span>;
}
</pre>
<h3 class="section-head" id="h-fonts"><a href="#h-fonts">Fonts</a></h3>
<p>Generates a font-family declarations for text, headings, buttons or inputs.</p>
<pre class="code skip"><span class="hljs-comment">// import Kube</span>
@<span class="hljs-keyword">import</span> <span class="hljs-string">"dist/scss/kube.scss"</span>;

<span class="hljs-comment">// use mixins</span>
@<span class="hljs-keyword">include</span> text-font(<span class="hljs-string">"Lato, 'Helvetica Neue', sans-serif"</span>);
@<span class="hljs-keyword">include</span> headings-font(<span class="hljs-string">"Lato, 'Helvetica Neue', sans-serif"</span>);
@<span class="hljs-keyword">include</span> buttons-font(<span class="hljs-string">"Lato, 'Helvetica Neue', sans-serif"</span>);
@<span class="hljs-keyword">include</span> inputs-font(<span class="hljs-string">"Lato, 'Helvetica Neue', sans-serif"</span>);
</pre>
<p>CSS Output</p>
<pre class="code skip"><span class="hljs-comment">// Text</span>
<span class="hljs-selector-tag">body</span> {
    <span class="hljs-attribute">font-family</span>: Lato, <span class="hljs-string">'Helvetica Neue'</span>, sans-serif;
}

<span class="hljs-comment">// Headings</span>
<span class="hljs-selector-tag">h1</span><span class="hljs-selector-class">.title</span>, <span class="hljs-selector-tag">h1</span>, <span class="hljs-selector-tag">h2</span>, <span class="hljs-selector-tag">h3</span>, <span class="hljs-selector-tag">h4</span>, <span class="hljs-selector-tag">h5</span>, <span class="hljs-selector-tag">h6</span>, <span class="hljs-selector-class">.h1</span>, <span class="hljs-selector-class">.h2</span>, <span class="hljs-selector-class">.h3</span>, <span class="hljs-selector-class">.h4</span>, <span class="hljs-selector-class">.h5</span>, <span class="hljs-selector-class">.h6</span> {
    <span class="hljs-attribute">font-family</span>: Lato, <span class="hljs-string">'Helvetica Neue'</span>, sans-serif;
}

<span class="hljs-comment">// Buttons</span>
<span class="hljs-selector-tag">button</span>, <span class="hljs-selector-class">.button</span> {
    <span class="hljs-attribute">font-family</span>: Lato, <span class="hljs-string">'Helvetica Neue'</span>, sans-serif;
}

<span class="hljs-comment">// Inputs</span>
<span class="hljs-selector-tag">input</span>, <span class="hljs-selector-tag">textarea</span>, select {
    <span class="hljs-attribute">font-family</span>: Lato, <span class="hljs-string">'Helvetica Neue'</span>, sans-serif;
}
</pre>
<h3 class="section-head" id="h-breakpoints"><a href="#h-breakpoints">Breakpoints</a></h3>
<p>Breakpoint for small screens (max-width 768px by default).</p>
<pre class="code skip">@<span class="hljs-keyword">include</span> breakpoint(sm) {
    <span class="hljs-selector-class">.span</span> {
        <span class="hljs-attribute">display</span>: none;
    }
}
</pre>
<p>Breakpoint for medium screens (min-width 1024px by default).</p>
<pre class="code skip">@<span class="hljs-keyword">include</span> breakpoint(md) {
    <span class="hljs-selector-class">.span</span> {
        <span class="hljs-attribute">display</span>: none;
    }
}
</pre>
<p>Breakpoint for large screens (min-width 1200px by default).</p>
<pre class="code skip">@<span class="hljs-keyword">include</span> breakpoint(lg) {
    <span class="hljs-selector-class">.span</span> {
        <span class="hljs-attribute">display</span>: none;
    }
}
</pre>
<p>Custom breakpoints:</p>
<pre class="code skip"><span class="hljs-comment">// min-width 768px;</span>
<span class="hljs-variable">@include</span> breakpoint(<span class="hljs-number">768px</span>) {}

<span class="hljs-comment">// min-width 768px and max-width 1024px;</span>
<span class="hljs-variable">@include</span> breakpoint(<span class="hljs-number">768px</span>, <span class="hljs-number">1024px</span>) {}

<span class="hljs-comment">// max-width 1024px;</span>
<span class="hljs-variable">@include</span> breakpoint(<span class="hljs-number">0</span>, <span class="hljs-number">1024px</span>) {}
</pre>
<h3 class="section-head" id="h-grid"><a href="#h-grid">Grid</a></h3>
<h5 id="h-grid-row">Row</h5>
<p>Generates a grid row.</p>
<pre class="code skip"><span class="hljs-selector-class">.my-row</span> {
    <span class="hljs-variable">@include</span> grid-row;
}
</pre>
<p>CSS Output</p>
<pre class="code skip"><span class="hljs-selector-class">.my-row</span> {
    <span class="hljs-attribute">display</span>: -ms-flexbox;
    <span class="hljs-attribute">display</span>: -webkit-flex;
    <span class="hljs-attribute">display</span>: flex;

    <span class="hljs-attribute">-ms-flex-direction</span>: row;
    <span class="hljs-attribute">-webkit-flex-direction</span>: row;
    <span class="hljs-attribute">flex-direction</span>: row;

    <span class="hljs-attribute">-ms-flex-wrap</span>: wrap;
    <span class="hljs-attribute">-webkit-flex-wrap</span>: wrap;
    <span class="hljs-attribute">flex-wrap</span>: wrap;
}
</pre>
<h5 id="h-grid-media">Media Grid</h5>
<p>Generates a media grid. See <a href="/kube/docs/grid/#h-media-grid">live example</a>.</p>
<pre class="code skip"><span class="hljs-selector-class">.my-media-grid</span> {
    @include grid-media-<span class="hljs-attribute">columns</span>(<span class="hljs-number">3</span>);
}
</pre>
<p>CSS Output</p>
<pre class="code skip"><span class="hljs-selector-class">.my-media-grid</span> {
    -webkit-<span class="hljs-attribute">column-count</span>: <span class="hljs-number">3</span>;
    -moz-<span class="hljs-attribute">column-count</span>: <span class="hljs-number">3</span>;
    <span class="hljs-attribute">column-count</span>: <span class="hljs-number">3</span>;

    <span class="hljs-comment">// column gap is specified</span>
    <span class="hljs-comment">// in the grid settings (variables.scss) as $grid-gutter</span>
    -webkit-<span class="hljs-attribute">column-gap</span>: <span class="hljs-number">2%</span>;
    -moz-<span class="hljs-attribute">column-gap</span>: <span class="hljs-number">2%</span>;
    <span class="hljs-attribute">column-gap</span>: <span class="hljs-number">2%</span>;
}

<span class="hljs-selector-class">.my-media-grid</span> &gt; <span class="hljs-selector-tag">div</span> {
    <span class="hljs-attribute">display</span>: inline-block;
    <span class="hljs-attribute">width</span>: <span class="hljs-number">100%</span>;
}

@<span class="hljs-keyword">media</span> (max-width: 768px) {
    <span class="hljs-selector-class">.my-media-grid</span> {
        -webkit-<span class="hljs-attribute">column-count</span>: <span class="hljs-number">1</span>;
        -moz-<span class="hljs-attribute">column-count</span>: <span class="hljs-number">1</span>;
        <span class="hljs-attribute">column-count</span>: <span class="hljs-number">1</span>;
    }
}
</pre>
<h3 class="section-head" id="h-flex"><a href="#h-flex">Flex</a></h3>
<pre class="code skip"><span class="hljs-variable">@include</span> flex;
<span class="hljs-variable">@include</span> flex-basis($basis);

<span class="hljs-comment">// items</span>
<span class="hljs-variable">@include</span> flex-items-wrap;
<span class="hljs-variable">@include</span> flex-items-nowrap;
<span class="hljs-variable">@include</span> flex-items-row
<span class="hljs-variable">@include</span> flex-items-column;
<span class="hljs-variable">@include</span> flex-items-left;
<span class="hljs-variable">@include</span> flex-items-right;
<span class="hljs-variable">@include</span> flex-items-center;
<span class="hljs-variable">@include</span> flex-items-space-between;
<span class="hljs-variable">@include</span> flex-items-space-around;
<span class="hljs-variable">@include</span> flex-items-top;
<span class="hljs-variable">@include</span> flex-items-middle;
<span class="hljs-variable">@include</span> flex-items-bottom;

<span class="hljs-comment">// item</span>
<span class="hljs-variable">@include</span> flex-item-grow($grow);
<span class="hljs-variable">@include</span> flex-item-auto;
<span class="hljs-variable">@include</span> flex-item-one;
<span class="hljs-variable">@include</span> flex-item-shrink($shrink);
<span class="hljs-variable">@include</span> flex-item-width($width);
</pre>
<h3 class="section-head" id="h-gradients"><a href="#h-gradients">Gradients</a></h3>
<h5 id="h-gradients-vertical">Vertical</h5>
<pre class="code skip">@<span class="hljs-keyword">include</span> gradient-vertical(<span class="hljs-variable">$startColor</span>, <span class="hljs-variable">$endColor</span>);</pre>
<div class="demo-gradient demo-gradient-vertical"></div>
<pre class="code skip"><span class="hljs-symbol">@include</span> gradient-vertical-<span class="hljs-keyword">to</span>-opacity($startColor, $opacity)<span class="hljs-comment">;</span></pre>
<div class="demo-gradient demo-gradient-vertical-to-opacity"></div>
<h5 id="h-gradients-horizontal">Horizontal</h5>
<pre class="code skip">@<span class="hljs-keyword">include</span> gradient-horizontal(<span class="hljs-variable">$startColor</span>, <span class="hljs-variable">$endColor</span>);</pre>
<div class="demo-gradient demo-gradient-horizontal"></div>
<pre class="code skip"><span class="hljs-symbol">@include</span> gradient-horizontal-<span class="hljs-keyword">to</span>-opacity($startColor, $opacity)<span class="hljs-comment">;</span></pre>
<div class="demo-gradient demo-gradient-horizontal-to-opacity"></div>
<h5 id="h-gradients-radial">Radial</h5>
<pre class="code skip">@include gradient-<span class="hljs-keyword">radial</span>($innerColor, $outerColor);</pre>
<div class="demo-gradient demo-gradient-radial"></div>
<h3 class="section-head" id="h-utils"><a href="#h-utils">Utils</a></h3>
<h5 id="h-utils-clearfix">Clearfix</h5>
<p>Provides an easy way to include a clearfix for containing floats.</p>
<pre class="code skip"><span class="hljs-selector-class">.layout</span> {
    <span class="hljs-variable">@include</span> clearfix;
}
</pre>
<p>CSS Output</p>
<pre class="code skip"><span class="hljs-selector-class">.layout</span><span class="hljs-selector-pseudo">:after</span> {
    <span class="hljs-attribute">content</span>: <span class="hljs-string">''</span>;
    <span class="hljs-attribute">display</span>: table;
    <span class="hljs-attribute">clear</span>: both;
}
</pre>
<h5 id="h-utils-transition">Transition</h5>
<p>This mixin provides a shorthand syntax for transitions.</p>
<pre class="code skip"><span class="hljs-comment">// by default 'all linear .2s'</span>
<span class="hljs-meta">@include</span> transition;

<span class="hljs-comment">// custom transitions</span>
<span class="hljs-meta">@include</span> transition(all .<span class="hljs-number">2</span>s ease-<span class="hljs-keyword">in</span>-<span class="hljs-keyword">out</span>);
<span class="hljs-meta">@include</span> transition(opacity <span class="hljs-number">1</span>s ease-<span class="hljs-keyword">in</span>, width .<span class="hljs-number">2</span>s ease-<span class="hljs-keyword">in</span>);
</pre>
<h5 id="h-utils-transform">Transform</h5>
<p>Provides a shorthand syntax for transforms.</p>
<pre class="code skip"><span class="hljs-selector-class">.span</span> {
    @include <span class="hljs-attribute">transform</span>(rotate(<span class="hljs-number">90deg</span>));
}
</pre>
<p>CSS Output</p>
<pre class="code skip"><span class="hljs-selector-class">.span</span> {
    <span class="hljs-attribute">-webkit-transform</span>: <span class="hljs-built_in">rotate</span>(90deg);
    <span class="hljs-attribute">-moz-transform</span>: <span class="hljs-built_in">rotate</span>(90deg);
    <span class="hljs-attribute">-ms-transform</span>: <span class="hljs-built_in">rotate</span>(90deg);
    <span class="hljs-attribute">transform</span>: <span class="hljs-built_in">rotate</span>(90deg);
}
</pre>
<h5 id="h-utils-blur">Blur</h5>
<p>Provides a shorthand syntax for blur filter.</p>
<pre class="code skip"><span class="hljs-selector-class">.span</span> {
    <span class="hljs-variable">@include</span> blur(<span class="hljs-number">5px</span>);
}
</pre>
<p>CSS Output</p>
<pre class="code skip"><span class="hljs-selector-class">.span</span> {
    <span class="hljs-attribute">-webkit-filter</span>: <span class="hljs-built_in">blur</span>(5px);
    <span class="hljs-attribute">-moz-filter</span>: <span class="hljs-built_in">blur</span>(5px);
    <span class="hljs-attribute">-ms-filter</span>: <span class="hljs-built_in">blur</span>(5px);
    <span class="hljs-attribute">filter</span>: <span class="hljs-built_in">blur</span>(5px);
}
</pre>
<h5 id="h-utils-retina-image">Retina Image</h5>
<p>Retina image must have a suffix <code>-2x</code>, for example: <code>image-2x.jpg</code></p>
<pre class="code skip">@include retina-<span class="hljs-built_in">background</span>-<span class="hljs-built_in">image</span>($<span class="hljs-built_in">image</span>-url, $<span class="hljs-built_in">image</span>-type, $<span class="hljs-built_in">image</span>-<span class="hljs-built_in">width</span>, $<span class="hljs-built_in">image</span>-<span class="hljs-built_in">height</span>);

// $<span class="hljs-built_in">image</span>-type - jpg, png, gif
// $<span class="hljs-built_in">image</span>-<span class="hljs-built_in">height</span> - optional
</pre>
<p>Example:</p>
<pre class="code skip"><span class="hljs-selector-class">.brand</span> {
    @include retina-<span class="hljs-attribute">background-image</span>(<span class="hljs-string">'../logo'</span>, <span class="hljs-string">'png'</span>, <span class="hljs-number">100px</span>);
}
</pre>
<p>CSS Output</p>
<pre class="code skip"><span class="hljs-selector-class">.brand</span> {
    <span class="hljs-attribute">background-repeat</span>: no-repeat;
    <span class="hljs-attribute">background-image</span>: <span class="hljs-built_in">url</span>(<span class="hljs-string">"../logo.png"</span>);
}
@<span class="hljs-keyword">media</span> only screen and (-webkit-min-device-pixel-ratio: <span class="hljs-number">2</span>),
       only screen and (min--moz-device-pixel-ratio: <span class="hljs-number">2</span>),
       only screen and (-o-min-device-pixel-ratio: <span class="hljs-number">2</span> / <span class="hljs-number">1</span>),
       only screen and (min-device-pixel-ratio: <span class="hljs-number">2</span>),
       only screen and (min-resolution: <span class="hljs-number">192dpi</span>),
       only screen and (min-resolution: <span class="hljs-number">2dppx</span>)
       {
            <span class="hljs-selector-class">.brand</span> {
                <span class="hljs-attribute">background-image</span>: <span class="hljs-built_in">url</span>(<span class="hljs-string">"../logo-2x.png"</span>);
                <span class="hljs-attribute">background-size</span>: <span class="hljs-number">100px</span> auto;
            }
       }
</pre>