+++
title = "Typography"
weight = 25
draft = false
description = "Always precise spacing and perfect font size"
bref = "Typography is perhaps one of the most important and most visible things on a web page. Even slightest imperfection can ruin otherwise perfect website. With Kube, you will have perfect typography with ideal spacing, font sizes and proportions, regardless of the exact style or font you choose for your site."
toc = true
+++

<h3 class="section-head" id="h-get-started"><a href="#h-get-started">Get Started</a></h3>
<p>Typography of Kube based on the 4px horizontal grid, it means that for the headers, paragraphs, quotes and any other texts chosen are a combination of line-height, that fit into the 4px grid and help set the horizontal rhythm of design by default.</p>
<figure>
  <img alt="4px horizontal rhythm" height="180" src="/img/kube/typography/01.png" width="676">
  <figcaption>
    4px grid for typography - the magic of the horizontal rhythm
  </figcaption>
</figure>
<p>With this feature, you can quickly and easily make a solid and harmonious-looking website and UI design. You do not need to do complicated calculations to find the size and proportions of the text baseline, no need to look for the magic formula to build a horizontal rhythm. Magic is already in Kube.</p>
<p>You can use all the default settings of typography and it will always look balanced. But also you can easily change any font sizes, just try to set the value of line-height fold for 4px and your texts will still look great.</p>
<p>Additional balance and harmony, creates a classic typography scale, used in Kube.</p>
<figure>
  <img alt="the typography scale" height="140" src="/img/kube/typography/02.png" width="520">
  <figcaption>
    The typography scale helps to build a balance between the size of headings and text elements
  </figcaption>
</figure>
<h3 class="section-head" id="h-headings"><a href="#h-headings">Headings</a></h3>
<p>Use h1-h6 tags or <code>.h1-.h6</code> classes to define headers and <code>class="title"</code> for the title, which is suitable for the most important inscriptions, for example, in the hero or covers.</p>
<div class="example">
  <h1 class="title">Title</h1>
  <h1>Heading 1</h1>
  <h2>Heading 2</h2>
  <h3>Heading 3</h3>
  <h4>Heading 4</h4>
  <h5>Heading 5</h5>
  <h6>Heading 6</h6>
  <pre class="code"><span class="hljs-tag">&lt;<span class="hljs-name">h1</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"title"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">h1</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">h1</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">h1</span>&gt;</span>
...
<span class="hljs-tag">&lt;<span class="hljs-name">h6</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">h6</span>&gt;</span></pre>
</div>
<h3 class="section-head" id="h-paragraphs"><a href="#h-paragraphs">Paragraphs</a></h3>
<p>The base pragraph has an ideal ratio of font size and baseline. This text is easy to read in most cases.</p>
<div class="example">
  <p class="section-item-desc">16px/24px</p>
  <p>By the same illusion which lifts the horizon of the sea to the level of the spectator on a hillside, the sable cloud beneath was dished out, and the car seemed to float in the middle of an immense dark sphere, whose upper half was strewn with silver.</p>
  <pre class="code"><span class="hljs-tag">&lt;<span class="hljs-name">p</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">p</span>&gt;</span></pre>
</div>
<p>In special cases, you can use modifiers which increase or decrease the text size. It is useful for building a variety of websites and UI, when you need to make a lead text or signatures with a small font size .</p>
<div class="example">
  <p class="section-item-desc">20px/32px</p>
  <p class="large">By the same illusion which lifts the horizon of the sea to the level of the spectator on a hillside, the sable cloud beneath was dished out, and the car seemed to float in the middle of an immense dark sphere, whose upper half was strewn with silver.</p>
  <pre class="code"><span class="hljs-tag">&lt;<span class="hljs-name">p</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"large"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">p</span>&gt;</span></pre>
</div>
<div class="example">
  <p class="section-item-desc">18px/28px</p>
  <p class="big">By the same illusion which lifts the horizon of the sea to the level of the spectator on a hillside, the sable cloud beneath was dished out, and the car seemed to float in the middle of an immense dark sphere, whose upper half was strewn with silver.</p>
  <pre class="code"><span class="hljs-tag">&lt;<span class="hljs-name">p</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"big"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">p</span>&gt;</span></pre>
</div>
<div class="example">
  <p class="section-item-desc">14px/20px</p>
  <p class="small">By the same illusion which lifts the horizon of the sea to the level of the spectator on a hillside, the sable cloud beneath was dished out, and the car seemed to float in the middle of an immense dark sphere, whose upper half was strewn with silver.</p>
  <pre class="code"><span class="hljs-tag">&lt;<span class="hljs-name">p</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"small"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">p</span>&gt;</span></pre>
</div>
<div class="example">
  <p class="section-item-desc">12px/20px</p>
  <p class="smaller">By the same illusion which lifts the horizon of the sea to the level of the spectator on a hillside, the sable cloud beneath was dished out, and the car seemed to float in the middle of an immense dark sphere, whose upper half was strewn with silver.</p>
  <pre class="code"><span class="hljs-tag">&lt;<span class="hljs-name">p</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"smaller"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">p</span>&gt;</span></pre>
</div>
<h3 class="section-head" id="h-columns"><a href="#h-columns">Columns</a></h3>
<p>Kube has three classes of <code>.columns-2</code> through <code>.columns-4</code> to create a multi-column layout. All multi-column layouts will be in a single column on mobile.</p>
<div class="example">
  <p class="section-item-desc">2 columns</p>
  <div class="columns-2">
    <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
  </div>
  <pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"columns-2"</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">p</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">p</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
</pre>
</div>
<div class="example">
  <p class="section-item-desc">3 columns</p>
  <div class="columns-3">
    <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
  </div>
  <pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"columns-3"</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">p</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">p</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
</pre>
</div>
<div class="example">
  <p class="section-item-desc">4 columns</p>
  <div class="columns-4">
    <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
  </div>
  <pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"columns-4"</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">p</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">p</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
</pre>
</div>
<h3 class="section-head" id="h-quotes"><a href="#h-quotes">Quotes</a></h3>
<p>Quotes in the text stand out for contrast and added variety to the text. In the quotes you can mark text as a paragraph tag and without it.</p>
<div class="example">
  <blockquote>
    <p>No, she'll probably make me do it. Goodbye, friends. I never thought I'd die like this. But I always really hoped. I saw you with those two "ladies of the evening" at Elzars. Explain that. I never loved you.</p>
  </blockquote>
  <pre class="code"><span class="hljs-tag">&lt;<span class="hljs-name">blockquote</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">blockquote</span>&gt;</span></pre>
</div>
<p>Use <code>small</code> tag for attribution. It's a prefer way for semantic code.</p>
<div class="example">
  <blockquote>
    <p>Who's brave enough to fly into something we all keep calling a death sphere? Yes. You gave me a dollar and some candy. I just want to talk. It has nothing to do with mating. Fry, that doesn't make sense.</p><small>— Bender</small>
  </blockquote>
  <pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">blockquote</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">p</span>&gt;</span>Quotation content<span class="hljs-tag">&lt;/<span class="hljs-name">p</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">small</span>&gt;</span>Author attribution<span class="hljs-tag">&lt;/<span class="hljs-name">small</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">blockquote</span>&gt;</span>
</pre>
</div>
<h3 class="section-head" id="h-pre-formatted"><a href="#h-pre-formatted">Pre-formatted</a></h3>
<p>Pre-formatted text most often used to display code snippets or other text elements, for which you wish to preserve the exact formatting they've got.</p>
<div class="example">
  <pre>Function.prototype.inherits = function(parent)
{
    for (var key in parent.prototype)
    {
        this.prototype[key] = parent.prototype[key];
    }
};
</pre>
  <pre class="code"><span class="hljs-tag">&lt;<span class="hljs-name">pre</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">pre</span>&gt;</span></pre>
</div>
<p>You feel free to use <code>code</code> tag inside <code>pre</code>.</p>
<div class="example">
  <pre class="demo-pre"><code>Function.prototype.inherits = function(parent)
{
    for (var key in parent.prototype)
    {
        this.prototype[key] = parent.prototype[key];
    }
};
</code></pre>
  <pre class="code">&lt;pre&gt;<span class="xml"><span class="hljs-tag">&lt;<span class="hljs-name">code</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">code</span>&gt;</span></span><span class="xml"><span class="hljs-tag">&lt;/<span class="hljs-name">pre</span>&gt;</span></span></pre>
</div>
<h3 class="section-head" id="h-inline-elements"><a href="#h-inline-elements">Inline Elements</a></h3>
<p>Kube features various inline formatting elements. They all have their own semantic meaning, but you're free to use them just for their style.</p>
<table>
  <thead>
    <tr>
      <th>Example</th>
      <th>Tag</th>
      <th>Example</th>
      <th>Tag</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><abbr title="Cascading Style Sheets">CSS</abbr></td>
      <td class="col-6"><code>&lt;abbr&gt;</code></td>
      <td><cite>Cite</cite></td>
      <td class="col-6"><code>&lt;cite&gt;</code></td>
    </tr>
    <tr>
      <td><code>Code</code></td>
      <td><code>&lt;code&gt;</code></td>
      <td><samp>Sample</samp></td>
      <td><code>&lt;samp&gt;</code></td>
    </tr>
    <tr>
      <td><var>Variable</var></td>
      <td><code>&lt;var&gt;</code></td>
      <td><mark>Mark</mark></td>
      <td><code>&lt;mark&gt;</code></td>
    </tr>
    <tr>
      <td><kbd>Shortcut</kbd></td>
      <td><code>&lt;kbd&gt;</code></td>
      <td>
        <del>Deleted</del>
      </td>
      <td><code>&lt;del&gt;</code></td>
    </tr>
    <tr>
      <td><i>Italic</i></td>
      <td><code>&lt;i&gt;</code></td>
      <td><em>Emphasis</em></td>
      <td><code>&lt;em&gt;</code></td>
    </tr>
    <tr>
      <td><strong>Highlighted</strong></td>
      <td><code>&lt;strong&gt;</code></td>
      <td><b>Bold</b></td>
      <td><code>&lt;b&gt;</code></td>
    </tr>
    <tr>
      <td>x<sup>superscript</sup></td>
      <td><code>&lt;sup&gt;</code></td>
      <td>x<sub>subscript</sub></td>
      <td><code>&lt;sub&gt;</code></td>
    </tr>
    <tr>
      <td><small>Small</small></td>
      <td><code>&lt;small&gt;</code></td>
      <td>
        <ins>Inserted</ins>
      </td>
      <td><code>&lt;ins&gt;</code></td>
    </tr>
  </tbody>
</table>
<h3 class="section-head" id="h-modifiers"><a href="#h-modifiers">Modifiers</a></h3>
<p>Modifiers - a set of helper classes for creating accents in the text and for the solution of useful tasks without writing CSS code, for example, the alignment of the text in the center.</p>
<p>Modifiers can be applied to any tags as inline elements, links or block tags.</p>
<table>
  <thead>
    <tr>
      <th>Example</th>
      <th>Modifier</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><span class="muted">Muted</span></td>
      <td class="col-12"><code>.muted</code></td>
    </tr>
    <tr>
      <td>
        <a class="muted" href="#">Muted link</a>
      </td>
      <td><code>.muted</code></td>
    </tr>
    <tr>
      <td>
        <a class="black" href="#">Black link</a>
      </td>
      <td><code>.black</code></td>
    </tr>
    <tr>
      <td><span class="highlight">Highlight</span></td>
      <td><code>.highlight</code></td>
    </tr>
    <tr>
      <td><span class="upper">upper</span></td>
      <td><code>.upper</code></td>
    </tr>
    <tr>
      <td><span class="lower">LOWER</span></td>
      <td><code>.lower</code></td>
    </tr>
    <tr>
      <td><span class="italic">Italic</span></td>
      <td><code>.italic</code></td>
    </tr>
    <tr>
      <td><span class="strong">Strong</span></td>
      <td><code>.strong</code></td>
    </tr>
    <tr>
      <td><strong class="normal">Not strong</strong></td>
      <td><code>.normal</code></td>
    </tr>
    <tr>
      <td><span class="monospace">Monospace</span></td>
      <td><code>.monospace</code></td>
    </tr>
    <tr>
      <td><span class="nowrap">Nowrap</span></td>
      <td><code>.nowrap</code></td>
    </tr>
    <tr>
      <td class="nowrap">Remove margin bottom</td>
      <td><code>.end</code></td>
    </tr>
    <tr>
      <td><span class="highlight">Highlight</span></td>
      <td><code>.highlight</code></td>
    </tr>
    <tr>
      <td><span class="small">Small</span></td>
      <td><code>.small</code></td>
    </tr>
    <tr>
      <td><span class="smaller">Smaller</span></td>
      <td><code>.smaller</code></td>
    </tr>
    <tr>
      <td><span class="large">Large</span></td>
      <td><code>.large</code></td>
    </tr>
    <tr>
      <td><span class="big">Big</span></td>
      <td><code>.big</code></td>
    </tr>
    <tr>
      <td class="text-left">Text left</td>
      <td><code>.text-left</code></td>
    </tr>
    <tr>
      <td class="text-center">Text center</td>
      <td><code>.text-center</code></td>
    </tr>
    <tr>
      <td class="text-right">Text right</td>
      <td><code>.text-right</code></td>
    </tr>
  </tbody>
</table>
<h3 class="section-head" id="h-figure"><a href="#h-figure">Figure</a></h3>
<p>A figure tag features an image, a video or a code plus a caption. It is a good accessibility practice. It also helps serve responsive video to various devices when you wrap video into <code>.video-container</code> class.</p>
<h5>Images</h5>
<figure>
  <img alt="Image" height="533" src="/img/favicons/logo-384x384.png" width="800">
  <figcaption>
    Lorem ipsum dolor sit amet, consectetur adipisicing elit...
  </figcaption>
</figure>
<p class="end"><var>HTML</var></p>
<pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">figure</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">img</span> <span class="hljs-attr">src</span>=<span class="hljs-string">"image.jpg"</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">figcaption</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">figcaption</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">figure</span>&gt;</span>
</pre>
<h5>Video</h5>
<figure>
  <div class="video-container">
    <iframe allowfullscreen frameborder="0" height="315" src="https://www.youtube.com/embed/nywsA8wCCfY" width="560"></iframe>
  </div>
  <figcaption>
    Journey Through The Ice | National Geographic
  </figcaption>
</figure>
<p class="end"><var>HTML</var></p>
<pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">figure</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"video-container"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">iframe</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">iframe</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">figcaption</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">figcaption</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">figure</span>&gt;</span>
</pre>
<h5>Code</h5>
<figure>
  <pre>Function.prototype.inherits = function(parent)
{
    for (var key in parent.prototype)
    {
        this.prototype[key] = parent.prototype[key];
    }
};
</pre>
  <figcaption>
    One of the implementations of inheritance
  </figcaption>
</figure>
<p class="end"><var>HTML</var></p>
<pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">figure</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">pre</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">pre</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">figcaption</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">figcaption</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">figure</span>&gt;</span>
</pre>
<h3 class="section-head" id="h-lists"><a href="#h-lists">Lists</a></h3>
<p>Although you most likely will not ever need to do so, you can still nest as many list levels as you like.</p>
<div class="row gutters">
  <div class="col col-6">
    <div class="example">
      <ul>
        <li>list item 1</li>
        <li>list item 2
          <ul>
            <li>list item 2.1</li>
            <li>list item 2.2
              <ul>
                <li>list item 2.2.1</li>
                <li>list item 2.2.2</li>
              </ul>
            </li>
            <li>list item 2.3</li>
            <li>list item 2.4</li>
          </ul>
        </li>
        <li>list item 3</li>
        <li>list item 4</li>
      </ul>
      <pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">ul</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">li</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">ul</span>&gt;</span>
</pre>
    </div>
  </div>
  <div class="col col-6">
    <div class="example">
      <ol>
        <li>list item 1</li>
        <li>list item 2
          <ol>
            <li>list item 2.1</li>
            <li>list item 2.2
              <ol>
                <li>list item 2.2.1</li>
                <li>list item 2.2.2</li>
              </ol>
            </li>
            <li>list item 2.3</li>
            <li>list item 2.4</li>
          </ol>
        </li>
        <li>list item 3</li>
        <li>list item 4</li>
      </ol>
      <pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">ol</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">li</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">ol</span>&gt;</span>
</pre>
    </div>
  </div>
</div>
<h5>Unstyled List</h5>
<p>To remove default list styling, use <code>.unstyled</code>.</p>
<div class="example">
  <ul class="unstyled">
    <li>list item 1</li>
    <li>list item 2
      <ul>
        <li>list item 2.1</li>
        <li>list item 2.2</li>
      </ul>
    </li>
    <li>list item 3</li>
    <li>list item 4</li>
  </ul>
  <pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">ul</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"unstyled"</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">li</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">li</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">ul</span>&gt;</span>
</pre>
</div>
<h5>Definition List</h5>
<p>From time to time, you may need to produce a list of definitions, and there's one nicely built-in into Kube. We love semantic things. And lists. We love lists.</p>
<div class="example">
  <dl>
    <dt>Term 1</dt>
    <dd>Description 1</dd>
    <dt>Term 2</dt>
    <dd>Description 2</dd>
    <dt>Term 3</dt>
    <dd>Description 3</dd>
  </dl>
  <pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">dl</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">dt</span>&gt;</span>Term<span class="hljs-tag">&lt;/<span class="hljs-name">dt</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">dd</span>&gt;</span>Description<span class="hljs-tag">&lt;/<span class="hljs-name">dd</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">dl</span>&gt;</span>
</pre>
</div>