+++
draft = false
weight = 30
description = "Flexible, fully responsive and ready to use"
title = "Grid"
bref = "Kube uses fully responsive, flexbox-enabled 12-column grid. You can combine or divide columns, nest them, center them horizontally or vertically, and do all sorts of things. Here are a few examples of what you can do with Kube's grid."
toc = true
+++

<div id="main">
  <h3 class="section-head" id="h-columns"><a href="#h-columns">Columns</a></h3>
  <p>Columns are building blocks for many websites. Here's how columns are formed in Kube, and here's how you can use them right away.</p>
  <div class="example demo-grid">
    <div class="row">
      <div class="col col-1">
        1
      </div>
      <div class="col col-11">
        11
      </div>
    </div>
    <div class="row">
      <div class="col col-2">
        2
      </div>
      <div class="col col-10">
        10
      </div>
    </div>
    <div class="row">
      <div class="col col-3">
        3
      </div>
      <div class="col col-9">
        9
      </div>
    </div>
    <div class="row">
      <div class="col col-4">
        4
      </div>
      <div class="col col-8">
        8
      </div>
    </div>
    <div class="row">
      <div class="col col-5">
        5
      </div>
      <div class="col col-7">
        7
      </div>
    </div>
    <div class="row">
      <div class="col col-6">
        6
      </div>
      <div class="col col-6">
        6
      </div>
    </div>
    <div class="row">
      <div class="col col-7">
        7
      </div>
      <div class="col col-5">
        5
      </div>
    </div>
    <div class="row">
      <div class="col col-8">
        8
      </div>
      <div class="col col-4">
        4
      </div>
    </div>
    <div class="row">
      <div class="col col-9">
        9
      </div>
      <div class="col col-3">
        3
      </div>
    </div>
    <div class="row">
      <div class="col col-10">
        10
      </div>
      <div class="col col-2">
        2
      </div>
    </div>
    <div class="row">
      <div class="col col-11">
        11
      </div>
      <div class="col col-1">
        1
      </div>
    </div>
    <div class="row">
      <div class="col col-12">
        12
      </div>
    </div>
    <pre class="code skip">&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"row"</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-4"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-8"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
  </div>
  <h3 class="section-head" id="h-gutters"><a href="#h-gutters">Gutters</a></h3>
  <div class="example demo-grid">
    <div class="row gutters">
      <div class="col col-3">
        3
      </div>
      <div class="col col-9">
        9
      </div>
    </div>
    <pre class="code skip">&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"row gutters"</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-3"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-9"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
  </div>
  <h3 class="section-head" id="h-auto-width"><a href="#h-auto-width">Auto Width</a></h3>
  <div class="example demo-grid">
    <div class="row auto">
      <div class="col">
        auto
      </div>
      <div class="col">
        auto
      </div>
      <div class="col">
        auto
      </div>
    </div>
    <pre class="code skip">&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"row auto"</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
  </div>
  <h5>Auto + Gutters</h5>
  <div class="example demo-grid">
    <div class="row gutters auto">
      <div class="col">
        auto
      </div>
      <div class="col">
        auto
      </div>
      <div class="col">
        auto
      </div>
      <div class="col">
        auto
      </div>
    </div>
    <pre class="code skip">&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"row gutters auto"</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
  </div>
  <h3 class="section-head" id="h-bricks"><a href="#h-bricks">Bricks</a></h3>
  <div class="example demo-grid">
    <div class="row">
      <div class="col col-6">
        50%
      </div>
      <div class="col col-6">
        50%
      </div>
      <div class="col col-6">
        50%
      </div>
      <div class="col col-6">
        50%
      </div>
    </div>
    <pre class="code skip">&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"row"</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-6"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-6"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-6"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-6"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
  </div>
  <h5>Bricks + Gutters</h5>
  <div class="example demo-grid">
    <div class="row gutters">
      <div class="col col-4">
        33%
      </div>
      <div class="col col-4">
        33%
      </div>
      <div class="col col-4">
        33%
      </div>
      <div class="col col-4">
        33%
      </div>
      <div class="col col-4">
        33%
      </div>
      <div class="col col-4">
        33%
      </div>
    </div>
    <pre class="code skip">&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"row gutters"</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-4"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-4"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-4"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-4"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
  </div>
  <h3 class="section-head" id="h-offset"><a href="#h-offset">Offset</a></h3>
  <p>Offsetting columns is very simple in Kube. Just use column's class as usual, for example, <code>col col-2</code> and then add the offset value <code>offset-4</code>. This will offset this column and all following columns by 4.</p>
  <div class="example demo-grid">
    <div class="row">
      <div class="col col-2">
        2
      </div>
      <div class="col col-6 offset-4">
        6
      </div>
    </div>
    <pre class="code skip">&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"row"</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-2"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-6 offset-4"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
  </div>
  <h3 class="section-head" id="h-nested"><a href="#h-nested">Nested</a></h3>
  <p>Here's an example of nesting columns within columns. In this example, we have a single row, with two columns .col col-6, a div nested within the second column, with .row end class to denote where to inject nested columns, and the three nested columns.</p>
  <div class="example demo-grid">
    <div class="row">
      <div class="col col-6">
        6
      </div>
      <div class="col col-6 demo-col-nested">
        <div class="row">
          <div class="col col-4">
            4
          </div>
          <div class="col col-4">
            4
          </div>
          <div class="col col-4">
            4
          </div>
        </div>
      </div>
    </div>
    <pre class="code skip">&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"row"</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-6"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-6"</span>&gt;
        &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"row"</span>&gt;
            &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-4"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
            &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-4"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
            &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-4"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
        &lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;/<span class="hljs-keyword">div</span>&gt;
&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
  </div>
  <h3 class="section-head" id="h-equal-height"><a href="#h-equal-height">Equal Height Columns</a></h3>
  <p>Columns are equal height by default in Kube</p>
  <div class="example demo-grid">
    <div class="row">
      <div class="col col-6">
        1
      </div>
      <div class="col col-6">
        1<br>
        2<br>
        3
      </div>
    </div>
    <pre class="code skip">&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"row"</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-6"</span>&gt;
        ...
    &lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-6"</span>&gt;
        ...
        ...
        ...
    &lt;/<span class="hljs-keyword">div</span>&gt;
&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
  </div>
  <h3 class="section-head" id="h-reordering"><a href="#h-reordering">Reordering</a></h3>
  <h5>First</h5>
  <div class="example demo-grid">
    <div class="row">
      <div class="col col-4">
        4
      </div>
      <div class="col col-8 first">
        8
      </div>
    </div>
    <pre class="code skip">&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"row"</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-4"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-8 first"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
  </div>
  <h5>Last</h5>
  <div class="example demo-grid">
    <div class="row">
      <div class="col col-3 last">
        3
      </div>
      <div class="col col-9">
        9
      </div>
    </div>
    <pre class="code skip">&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"row"</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-3 last"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-9"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
  </div>
  <h3 class="section-head" id="h-around"><a href="#h-around">Around</a></h3>
  <p>When you want your columns to be surrounded by an even margins on both sides, just use <code>.around</code> class.</p>
  <div class="example demo-grid">
    <div class="row around">
      <div class="col col-3">
        3
      </div>
      <div class="col col-3">
        3
      </div>
      <div class="col col-3">
        3
      </div>
    </div>
    <pre class="code skip">&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"row around"</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-3"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-3"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-3"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
  </div>
  <h3 class="section-head" id="h-between"><a href="#h-between">Between</a></h3>
  <p>Sometime you just need space between columns, and not around them. Well, there's a class for that: <code>.between</code></p>
  <div class="example demo-grid">
    <div class="row between">
      <div class="col col-3">
        3
      </div>
      <div class="col col-3">
        3
      </div>
      <div class="col col-3">
        3
      </div>
    </div>
    <pre class="code skip">&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"row between"</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-3"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-3"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-3"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
  </div>
  <h3 class="section-head" id="h-alignment"><a href="#h-alignment">Alignment</a></h3>
  <h5>Align Center</h5>
  <div class="example demo-grid">
    <div class="row align-center">
      <div class="col col-6">
        6
      </div>
    </div>
    <pre class="code skip">&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"row align-center"</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-6"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
  </div>
  <h5>Align Right</h5>
  <div class="example demo-grid">
    <div class="row align-right">
      <div class="col col-3">
        3
      </div>
      <div class="col col-3">
        3
      </div>
    </div>
    <pre class="code skip">&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"row align-right"</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-3"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-3"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
  </div>
  <h5>Align Middle</h5>
  <div class="example demo-grid">
    <div class="row align-middle" style="height: 104px; border: 1px solid #eee;">
      <div class="col col-4" style="height: 36px;">
        4
      </div>
      <div class="col col-4" style="height: 56px;">
        4
      </div>
      <div class="col col-4" style="height: 36px;">
        4
      </div>
    </div>
    <pre class="code skip">&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"row align-middle"</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-4"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-4"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-4"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
  </div>
  <h3 class="section-head" id="h-push"><a href="#h-push">Push</a></h3>
  <h5>Push Center</h5>
  <div class="example demo-grid">
    <div class="row">
      <div class="col col-4 push-center">
        .push-center
      </div>
    </div>
    <pre class="code skip">&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"row"</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-4 push-center"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
  </div>
  <h5>Push Right</h5>
  <div class="example demo-grid">
    <div class="row">
      <div class="col col-6">
        ...
      </div>
      <div class="col col-4 push-right">
        .push-right
      </div>
    </div>
    <pre class="code skip">&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"row"</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-6"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-4 push-right"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
  </div>
  <h5>Push Middle</h5>
  <div class="example demo-grid">
    <div class="row" style="height: 104px;">
      <div class="col col-8 push-middle">
        .push-middle
      </div>
    </div>
    <pre class="code skip">&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"row"</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-8 push-middle"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
  </div>
  <h5>Push Mixed</h5>
  <div class="example demo-grid">
    <div class="row" style="height: 104px;">
      <div class="col col-8 push-middle push-center">
        .push-middle.push-center
      </div>
    </div>
    <pre class="code skip">&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"row"</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-8 push-middle push-center"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
  </div>
  <h5>Push Bottom</h5>
  <div class="example demo-grid">
    <div class="row" style="height: 104px;">
      <div class="col col-8 push-bottom">
        .push-bottom
      </div>
    </div>
    <pre class="code skip">&lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"row"</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-8 push-bottom"</span>&gt;...&lt;/<span class="hljs-keyword">div</span>&gt;
&lt;/<span class="hljs-keyword">div</span>&gt;
</pre>
  </div>
  <h3 class="section-head" id="h-fixed-column"><a href="#h-fixed-column">Fixed Width Column</a></h3>
  <p>Here's one more bit of magic. You can set a fixed width column and still have a responsive and flexible layout right beside this fixed width column.</p>
  <div class="example demo-grid">
    <div id="demo-container">
      <div id="demo-sidebar">
        sidebar
      </div>
      <div id="demo-content">
        content
      </div>
    </div>
    <pre class="code skip"><span class="hljs-comment">&lt;!-- scss --&gt;</span>
#container {
    @include grid-row;
}
#sidebar {
    @include flex-item-width(300px);
}
#content {
    @include flex-item-auto;
}

<span class="hljs-comment">&lt;!-- html --&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"container"</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"sidebar"</span>&gt;</span>Sidebar<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"content"</span>&gt;</span>Content<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>

</pre>
  </div>
  <h3 class="section-head" id="h-media-grid"><a href="#h-media-grid">Media Grid</a></h3>
  <p>Media grid is a great example of Kube's flexibility and versatility. Whenever you have some media to display, you can choose to use Media grid. This type of grid is ideal for featuring photos; it adds some visual music and slight randomness to your layout.</p>
  <div class="example demo-grid">
    <div id="demo-media-grid">
      <div>
        1
      </div>
      <div>
        2
      </div>
      <div>
        3
      </div>
      <div>
        4
      </div>
      <div>
        5
      </div>
      <div>
        6
      </div>
      <div>
        7
      </div>
      <div>
        8
      </div>
      <div>
        9
      </div>
      <div>
        10
      </div>
    </div>
    <pre class="code skip"><span class="hljs-comment">&lt;!-- scss --&gt;</span>
#media-grid {

    @include grid-media-columns(2);

    &amp; &gt; div {
        margin-bottom: 20px;
        height: 80px;
    }
    &amp; &gt; div:nth-child(2n) {
      height: 200px;
    }
    &amp; &gt; div:nth-child(5n) {
      height: 120px;
    }
}

<span class="hljs-comment">&lt;!-- html --&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"media-grid"</span>&gt;</span>
  <span class="hljs-tag">&lt;<span class="hljs-name">div</span>&gt;</span>1<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
  <span class="hljs-tag">&lt;<span class="hljs-name">div</span>&gt;</span>2<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
  <span class="hljs-tag">&lt;<span class="hljs-name">div</span>&gt;</span>3<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
  <span class="hljs-tag">&lt;<span class="hljs-name">div</span>&gt;</span>4<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
  <span class="hljs-tag">&lt;<span class="hljs-name">div</span>&gt;</span>5<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
  <span class="hljs-tag">&lt;<span class="hljs-name">div</span>&gt;</span>6<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
  <span class="hljs-tag">&lt;<span class="hljs-name">div</span>&gt;</span>7<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
  <span class="hljs-tag">&lt;<span class="hljs-name">div</span>&gt;</span>8<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
  <span class="hljs-tag">&lt;<span class="hljs-name">div</span>&gt;</span>9<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
  <span class="hljs-tag">&lt;<span class="hljs-name">div</span>&gt;</span>10<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
</pre>
  </div>
</div>
