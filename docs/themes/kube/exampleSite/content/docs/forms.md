+++
date = "2017-04-10T16:40:50+01:00"
title = "Forms"
draft = false
weight = 60
description = "Rows and columns for all sorts of tables"
bref = 'Forms come in all forms and shapes in Kube, and you can do all sorts of things with them, especially when combining with <a href="/docs/custom-plugins/">custom plugins</a> for extra interactivity. These forms are ideal building material for your awesome projects!'
toc = true
+++

<h3 class="section-head" id="h-base"><a href="#h-base">Base</a></h3>
<p>This is the most basic form with all the basic inputs.</p>
<div class="example">
  <form action="" autocomplete="off" class="form" method="post">
    <input name="authorize-token" type="hidden" value="c6851d805bd96d2c91d52574c65d3ae26ce4b6bca43f560518a0bea61335f9d52acc6055807bf8616ff3d13d3882d5aa2f4b15a046562c1e6a9d932e369e9fad"> <input id="fake_user_name" name="fake_user[name]" style="position:absolute; top:-10000px;" type="text" value="Safari Autofill Me">
    <div class="form-item">
      <label>Email</label> <input class="w50" name="email" type="email">
    </div>
    <div class="form-item">
      <label>Country</label> <select>
        <option value="">
          ---
        </option>
      </select>
    </div>
    <div class="form-item">
      <label class="checkbox"><input type="checkbox"> Check me</label> <label class="checkbox"><input type="radio"> Radio me</label>
    </div>
    <div class="form-item">
      <textarea rows="6"></textarea>
    </div>
    <div class="form-item form-buttons">
      <button>Log in</button> <button class="button secondary outline">Cancel</button>
    </div>
  </form>

<pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">form</span> <span class="hljs-attr">method</span>=<span class="hljs-string">"post"</span> <span class="hljs-attr">action</span>=<span class="hljs-string">""</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form"</span>&gt;</span>

    <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">label</span>&gt;</span>Email<span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">input</span> <span class="hljs-attr">type</span>=<span class="hljs-string">"email"</span> <span class="hljs-attr">name</span>=<span class="hljs-string">"email"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"w50"</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>

    <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">label</span>&gt;</span>Country<span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">select</span>&gt;</span>
            <span class="hljs-tag">&lt;<span class="hljs-name">option</span> <span class="hljs-attr">value</span>=<span class="hljs-string">""</span>&gt;</span>---<span class="hljs-tag">&lt;/<span class="hljs-name">option</span>&gt;</span>
        <span class="hljs-tag">&lt;/<span class="hljs-name">select</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>

    <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">label</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"checkbox"</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">input</span> <span class="hljs-attr">type</span>=<span class="hljs-string">"checkbox"</span>&gt;</span> Check me<span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">label</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"checkbox"</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">input</span> <span class="hljs-attr">type</span>=<span class="hljs-string">"radio"</span>&gt;</span> Radio me<span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>

    <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">textarea</span> <span class="hljs-attr">rows</span>=<span class="hljs-string">"6"</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">textarea</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>

    <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">button</span>&gt;</span>Log in<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button secondary outline"</span>&gt;</span>Cancel<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>

<span class="hljs-tag">&lt;/<span class="hljs-name">form</span>&gt;</span>
</pre>

</div>

<h3 class="section-head" id="h-inputs"><a href="#h-inputs">Inputs</a></h3>
<p>Here's a standard input field with type set as text. Label serves as input's label, and the following div with class desc serves as an optional description.</p>
<div class="example">
  <form class="form">
    <div class="form-item">
      <label>City</label> <input type="text">
      <div class="desc">
        This information helps us deliver orders on time.
      </div>
    </div>
  </form>
  <pre class="code skip">&lt;form <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"form"</span>&gt;
    <span class="xml"><span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">label</span>&gt;</span>City<span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">input</span> <span class="hljs-attr">type</span>=<span class="hljs-string">"text"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"desc"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">form</span>&gt;</span></span>
</pre>
</div>
<p>Here are two text input fields, one with a placeholder and another one disabled, both set to be 6 grid columns wide.</p>
<div class="example">
  <form class="form">
    <div class="row gutters">
      <div class="col col-6">
        <div class="form-item">
          <input placeholder="Email" type="text">
        </div>
      </div>
      <div class="col col-6">
        <div class="form-item">
          <input disabled="true" type="text" value="Disabled">
        </div>
      </div>
    </div>
  </form>
  <pre class="code skip">&lt;form <span class="hljs-built_in">class</span>=<span class="hljs-string">"form"</span>&gt;
    &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"row gutters"</span>&gt;
        &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-6"</span>&gt;
            &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"form-item"</span>&gt;
                &lt;input type=<span class="hljs-string">"text"</span> placeholder=<span class="hljs-string">"Email"</span>&gt;
            &lt;/<span class="hljs-keyword">div</span>&gt;
        &lt;/<span class="hljs-keyword">div</span>&gt;
        &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"col col-6"</span>&gt;
            &lt;<span class="hljs-keyword">div</span> <span class="hljs-built_in">class</span>=<span class="hljs-string">"form-item"</span>&gt;
                &lt;input type=<span class="hljs-string">"text"</span> disabled=<span class="hljs-string">"true"</span> value=<span class="hljs-string">"Disabled"</span>&gt;
            &lt;/<span class="hljs-keyword">div</span>&gt;
        &lt;/<span class="hljs-keyword">div</span>&gt;
    &lt;/<span class="hljs-keyword">div</span>&gt;
&lt;/form&gt;
</pre>
</div>
<p>Here's how you denote required fields with req class, and add descriptions to labels using span with class desc.</p>
<div class="example">
  <form class="form">
    <div class="row gutters">
      <div class="col col-6">
        <div class="form-item">
          <label>City <span class="req">*</span></label> <input type="text">
        </div>
      </div>
      <div class="col col-6">
        <div class="form-item">
          <label>City <span class="desc">Just curious.</span></label> <input type="text">
          <div class="desc">
            This information helps us deliver orders on time.
          </div>
        </div>
      </div>
    </div>
  </form>
  <pre class="code skip">&lt;form <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"form"</span>&gt;
    <span class="xml"><span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"row gutters"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"col col-6"</span>&gt;</span>
            <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
                <span class="hljs-tag">&lt;<span class="hljs-name">label</span>&gt;</span>City <span class="hljs-tag">&lt;<span class="hljs-name">span</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"req"</span>&gt;</span>*<span class="hljs-tag">&lt;/<span class="hljs-name">span</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
                <span class="hljs-tag">&lt;<span class="hljs-name">input</span> <span class="hljs-attr">type</span>=<span class="hljs-string">"text"</span>&gt;</span>
            <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
        <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"col col-6"</span>&gt;</span>
            <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
                <span class="hljs-tag">&lt;<span class="hljs-name">label</span>&gt;</span>City <span class="hljs-tag">&lt;<span class="hljs-name">span</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"desc"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">span</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
                <span class="hljs-tag">&lt;<span class="hljs-name">input</span> <span class="hljs-attr">type</span>=<span class="hljs-string">"text"</span>&gt;</span>
                <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"desc"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
            <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
        <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">form</span>&gt;</span>
</span></pre>
</div>
<h3 class="section-head" id="h-search"><a href="#h-search">Search</a></h3>
<div class="example">
  <form class="form">
    <div class="form-item">
      <input class="search" type="text">
    </div>
  </form>
  <pre class="code skip">&lt;form <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"form"</span>&gt;
    &lt;div <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"form-item"</span>&gt;
        &lt;input <span class="hljs-class"><span class="hljs-keyword">type</span></span>=<span class="hljs-string">"text"</span> <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"search"</span>&gt;
    &lt;/div&gt;
&lt;/form&gt;
</pre>
</div>
<h3 class="section-head" id="h-checkboxes"><a href="#h-checkboxes">Checkboxes &amp; Radio</a></h3>
<p>Kube features full variety of stylish checkboxes and radio buttons. You can feature them on your pages by defining input type as checkbox or radio. There's also a neat option to place checkboxes inline by adding <code>form-checkboxes</code> class to the container (works for both checkboxes and radio buttons)</p>
<div class="example">
  <form action="" class="form" method="post">
    <input name="authorize-token" type="hidden" value="c6851d805bd96d2c91d52574c65d3ae26ce4b6bca43f560518a0bea61335f9d52acc6055807bf8616ff3d13d3882d5aa2f4b15a046562c1e6a9d932e369e9fad">
    <div class="form-item form-checkboxes">
      <label class="checkbox"><input type="checkbox"> Check 1</label> <label class="checkbox"><input type="checkbox"> Check 2</label> <label class="checkbox"><input type="checkbox"> Check 3</label> <label class="checkbox"><input type="checkbox"> Check 4</label>
    </div>
  </form>
  <pre class="code skip">&lt;form method=<span class="hljs-string">"post"</span> action=<span class="hljs-string">""</span> <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"form"</span>&gt;
    <span class="xml"><span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item form-checkboxes"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">label</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"checkbox"</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">input</span> <span class="hljs-attr">type</span>=<span class="hljs-string">"checkbox"</span>&gt;</span> Check 1<span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
        ...
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">form</span>&gt;</span></span>
</pre>
</div>
<div class="example">
  <form action="" class="form" method="post">
    <input name="authorize-token" type="hidden" value="c6851d805bd96d2c91d52574c65d3ae26ce4b6bca43f560518a0bea61335f9d52acc6055807bf8616ff3d13d3882d5aa2f4b15a046562c1e6a9d932e369e9fad">
    <div class="form-item">
      <label class="checkbox"><input type="checkbox"> Check 1</label> <label class="checkbox"><input type="checkbox"> Check 2</label> <label class="checkbox"><input type="checkbox"> Check 3</label> <label class="checkbox"><input type="checkbox"> Check 4</label>
    </div>
  </form>
  <pre class="code skip">&lt;form method=<span class="hljs-string">"post"</span> action=<span class="hljs-string">""</span> <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"form"</span>&gt;
    <span class="xml"><span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">label</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"checkbox"</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">input</span> <span class="hljs-attr">type</span>=<span class="hljs-string">"checkbox"</span>&gt;</span> Check 1<span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
        ...
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">form</span>&gt;</span></span>
</pre>
</div>
<div class="example">
  <form action="" class="form" method="post">
    <input name="authorize-token" type="hidden" value="c6851d805bd96d2c91d52574c65d3ae26ce4b6bca43f560518a0bea61335f9d52acc6055807bf8616ff3d13d3882d5aa2f4b15a046562c1e6a9d932e369e9fad">
    <div class="form-item form-checkboxes">
      <label class="checkbox"><input name="check" type="radio"> Check 1</label> <label class="checkbox"><input name="check" type="radio"> Check 2</label> <label class="checkbox"><input name="check" type="radio"> Check 3</label> <label class="checkbox"><input name="check" type="radio"> Check 4</label>
    </div>
  </form>
  <pre class="code skip">&lt;form method=<span class="hljs-string">"post"</span> action=<span class="hljs-string">""</span> <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"form"</span>&gt;
    <span class="xml"><span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item form-checkboxes"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">label</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"checkbox"</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">input</span> <span class="hljs-attr">type</span>=<span class="hljs-string">"radio"</span>&gt;</span> Check 1<span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
        ...
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">form</span>&gt;</span></span>
</pre>
</div>
<div class="example">
  <form action="" class="form" method="post">
    <input name="authorize-token" type="hidden" value="c6851d805bd96d2c91d52574c65d3ae26ce4b6bca43f560518a0bea61335f9d52acc6055807bf8616ff3d13d3882d5aa2f4b15a046562c1e6a9d932e369e9fad">
    <div class="form-item">
      <label class="checkbox"><input name="check" type="radio"> Check 1</label> <label class="checkbox"><input name="check" type="radio"> Check 2</label> <label class="checkbox"><input name="check" type="radio"> Check 3</label> <label class="checkbox"><input name="check" type="radio"> Check 4</label>
    </div>
  </form>
  <pre class="code skip">&lt;form method=<span class="hljs-string">"post"</span> action=<span class="hljs-string">""</span> <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"form"</span>&gt;
    <span class="xml"><span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">label</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"checkbox"</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">input</span> <span class="hljs-attr">type</span>=<span class="hljs-string">"radio"</span>&gt;</span> Check 1<span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
        ...
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">form</span>&gt;</span></span>
</pre>
</div>
<h3 class="section-head" id="h-fieldset"><a href="#h-fieldset">Fieldset</a></h3>
<p>Fieldsets in Kube are nicely pre-formatted and ready to go. Just wrap your form in a fieldset tag and you're good to go.</p>
<form action="" class="form" method="post">
  <input name="authorize-token" type="hidden" value="c6851d805bd96d2c91d52574c65d3ae26ce4b6bca43f560518a0bea61335f9d52acc6055807bf8616ff3d13d3882d5aa2f4b15a046562c1e6a9d932e369e9fad">
  <fieldset>
    <legend>Login data</legend>
    <div class="form-item">
      <label>Email</label> <input class="w50" name="user-email" type="email">
    </div>
    <div class="form-item">
      <label>Password</label> <input class="w50" name="user-password" type="password">
    </div>
  </fieldset>
  <pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">fieldset</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">legend</span>&gt;</span>Login data<span class="hljs-tag">&lt;/<span class="hljs-name">legend</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">label</span>&gt;</span>Email<span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">input</span> <span class="hljs-attr">type</span>=<span class="hljs-string">"email"</span> <span class="hljs-attr">name</span>=<span class="hljs-string">"user-email"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"w50"</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">label</span>&gt;</span>Password<span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">input</span> <span class="hljs-attr">type</span>=<span class="hljs-string">"password"</span> <span class="hljs-attr">name</span>=<span class="hljs-string">"user-password"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"w50"</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">fieldset</span>&gt;</span>
</pre>
  <fieldset>
    <legend>About</legend>
    <div class="form-item">
      <textarea name="user-about" rows="5"></textarea>
    </div>
    <div class="form-item">
      <button>Submit</button>
    </div>
  </fieldset>
</form>
<pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">fieldset</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">legend</span>&gt;</span>About<span class="hljs-tag">&lt;/<span class="hljs-name">legend</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">textarea</span> <span class="hljs-attr">name</span>=<span class="hljs-string">"user-about"</span> <span class="hljs-attr">rows</span>=<span class="hljs-string">"5"</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">textarea</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">button</span>&gt;</span>Submit<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">fieldset</span>&gt;</span>
</pre>
<h3 class="section-head" id="h-small"><a href="#h-small">Small</a></h3>
<p>Simple class <code>.small</code> makes your selects and fields, well, smaller.</p>
<div class="example">
  <form class="form">
    <div class="form-item">
      <input class="small" placeholder="Title" type="text">
    </div>
    <div class="form-item">
      <select class="small">
        <option>
          Choose me
        </option>
      </select>
    </div>
  </form>
  <pre class="code skip">&lt;form <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"form"</span>&gt;
    <span class="xml"><span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">input</span> <span class="hljs-attr">type</span>=<span class="hljs-string">"text"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"small"</span> <span class="hljs-attr">placeholder</span>=<span class="hljs-string">"Title"</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">select</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"small"</span>&gt;</span>
            <span class="hljs-tag">&lt;<span class="hljs-name">option</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">option</span>&gt;</span>
        <span class="hljs-tag">&lt;/<span class="hljs-name">select</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">form</span>&gt;</span></span>
</pre>
</div>
<h3 class="section-head" id="h-big"><a href="#h-big">Big</a></h3>
<p>Simple class <code>.big</code> makes your selects and fields, well, bigger.</p>
<div class="example">
  <form class="form">
    <div class="form-item">
      <input class="big" placeholder="Title" type="text">
    </div>
    <div class="form-item">
      <select class="big">
        <option>
          Choose me
        </option>
      </select>
    </div>
  </form>
  <pre class="code skip">&lt;form <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"form"</span>&gt;
    <span class="xml"><span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">input</span> <span class="hljs-attr">type</span>=<span class="hljs-string">"text"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"big"</span> <span class="hljs-attr">placeholder</span>=<span class="hljs-string">"Title"</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">select</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"big"</span>&gt;</span>
            <span class="hljs-tag">&lt;<span class="hljs-name">option</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">option</span>&gt;</span>
        <span class="hljs-tag">&lt;/<span class="hljs-name">select</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">form</span>&gt;</span></span>
</pre>
</div>
<h3 class="section-head" id="h-width"><a href="#h-width">Width</a></h3>
<p>Following framework-wide rules, it is very easy to manipulate form fields sizes. Just use <code>w50</code> class to make a field 50% wide or <code>w25</code> to a 25% wide.</p>
<div class="example">
  <form action="" class="form" method="post">
    <input name="authorize-token" type="hidden" value="c6851d805bd96d2c91d52574c65d3ae26ce4b6bca43f560518a0bea61335f9d52acc6055807bf8616ff3d13d3882d5aa2f4b15a046562c1e6a9d932e369e9fad">
    <div class="form-item">
      <input class="w25" type="text" value=".w25">
    </div>
    <div class="form-item">
      <input class="w50" type="text" value=".w50">
    </div>
    <div class="form-item">
      <input class="w75" type="text" value=".w75">
    </div>
    <div class="form-item">
      <input type="text" value="100% by default">
    </div>
  </form>
  <pre class="code skip">&lt;form method=<span class="hljs-string">"post"</span> action=<span class="hljs-string">""</span> <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"form"</span>&gt;
    &lt;div <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"form-item"</span>&gt;
        &lt;input <span class="hljs-class"><span class="hljs-keyword">type</span></span>=<span class="hljs-string">"text"</span> <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"w25"</span>&gt;
    &lt;/div&gt;
    &lt;div <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"form-item"</span>&gt;
        &lt;input <span class="hljs-class"><span class="hljs-keyword">type</span></span>=<span class="hljs-string">"text"</span> <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"w50"</span>&gt;
    &lt;/div&gt;
    &lt;div <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"form-item"</span>&gt;
        &lt;input <span class="hljs-class"><span class="hljs-keyword">type</span></span>=<span class="hljs-string">"text"</span> <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"w75"</span>&gt;
    &lt;/div&gt;
    &lt;div <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"form-item"</span>&gt;
        &lt;input <span class="hljs-class"><span class="hljs-keyword">type</span></span>=<span class="hljs-string">"text"</span> value=<span class="hljs-string">"100% by default"</span>&gt;
    &lt;/div&gt;
&lt;/form&gt;
</pre>
</div>
<h3 class="section-head" id="h-states"><a href="#h-states">States</a></h3>
<p>By default, Kube features two different states: success and failure.</p>
<div class="example">
  <form action="" class="form" method="post">
    <input name="authorize-token" type="hidden" value="c6851d805bd96d2c91d52574c65d3ae26ce4b6bca43f560518a0bea61335f9d52acc6055807bf8616ff3d13d3882d5aa2f4b15a046562c1e6a9d932e369e9fad">
    <div class="form-item">
      <label>Your height <span class="success">Looks like a valid value</span></label> <input class="success" type="text">
    </div>
  </form>
  <pre class="code skip">&lt;form method=<span class="hljs-string">"post"</span> action=<span class="hljs-string">""</span> <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"form"</span>&gt;
    <span class="xml"><span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">label</span>&gt;</span>Your height <span class="hljs-tag">&lt;<span class="hljs-name">span</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"success"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">span</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">input</span> <span class="hljs-attr">type</span>=<span class="hljs-string">"text"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"success"</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">form</span>&gt;</span></span>
</pre>
</div>
<div class="example">
  <form action="" class="form" method="post">
    <input name="authorize-token" type="hidden" value="c6851d805bd96d2c91d52574c65d3ae26ce4b6bca43f560518a0bea61335f9d52acc6055807bf8616ff3d13d3882d5aa2f4b15a046562c1e6a9d932e369e9fad">
    <div class="form-item">
      <label>Enter your weight <span class="error">No negative numbers, please</span></label> <input class="error" type="text">
    </div>
  </form>
  <pre class="code skip">&lt;form method=<span class="hljs-string">"post"</span> action=<span class="hljs-string">""</span> <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"form"</span>&gt;
    <span class="xml"><span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">label</span>&gt;</span>Enter your weight <span class="hljs-tag">&lt;<span class="hljs-name">span</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"error"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">span</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">input</span> <span class="hljs-attr">type</span>=<span class="hljs-string">"text"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"error"</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">form</span>&gt;</span></span>
</pre>
</div>
<h3 class="section-head" id="h-required"><a href="#h-required">Required</a></h3>
<p>Along with making a field actually required, you can add a visual clue for the users using a span with req class.</p>
<div class="example">
  <form action="" class="form" method="post">
    <input name="authorize-token" type="hidden" value="c6851d805bd96d2c91d52574c65d3ae26ce4b6bca43f560518a0bea61335f9d52acc6055807bf8616ff3d13d3882d5aa2f4b15a046562c1e6a9d932e369e9fad">
    <div class="form-item">
      <label>Email <span class="req">*</span></label> <input name="user-email" type="email">
    </div>
  </form>
  <pre class="code skip">&lt;form method=<span class="hljs-string">"post"</span> action=<span class="hljs-string">""</span> <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"form"</span>&gt;
    <span class="xml"><span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">label</span>&gt;</span>Email <span class="hljs-tag">&lt;<span class="hljs-name">span</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"req"</span>&gt;</span>*<span class="hljs-tag">&lt;/<span class="hljs-name">span</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">input</span> <span class="hljs-attr">type</span>=<span class="hljs-string">"email"</span> <span class="hljs-attr">name</span>=<span class="hljs-string">"user-email"</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">form</span>&gt;</span></span>
</pre>
</div>
<h3 class="section-head" id="h-descriptions"><a href="#h-descriptions">Descriptions</a></h3>
<p>Descriptions are simple: as long as form's element has desc class, Kube will treat it as a description.</p>
<div class="example">
  <form action="" class="form" method="post">
    <input name="authorize-token" type="hidden" value="c6851d805bd96d2c91d52574c65d3ae26ce4b6bca43f560518a0bea61335f9d52acc6055807bf8616ff3d13d3882d5aa2f4b15a046562c1e6a9d932e369e9fad">
    <div class="form-item">
      <label>Name <span class="desc">Always nice to feel important.</span></label> <input name="user-name" type="text">
    </div>
    <div class="form-item">
      <label>Email</label> <input name="user-email" type="email">
      <div class="desc">
        Please enter your email.
      </div>
    </div>
  </form>
  <pre class="code skip">&lt;form method=<span class="hljs-string">"post"</span> action=<span class="hljs-string">""</span> <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"form"</span>&gt;
    <span class="xml"><span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">label</span>&gt;</span>Name <span class="hljs-tag">&lt;<span class="hljs-name">span</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"desc"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">span</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">input</span> <span class="hljs-attr">type</span>=<span class="hljs-string">"text"</span> <span class="hljs-attr">name</span>=<span class="hljs-string">"user-name"</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">label</span>&gt;</span>Email<span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">input</span> <span class="hljs-attr">type</span>=<span class="hljs-string">"email"</span> <span class="hljs-attr">name</span>=<span class="hljs-string">"user-email"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"desc"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">form</span>&gt;</span>
</span></pre>
</div>
<h3 class="section-head" id="h-disabled"><a href="#h-disabled">Disabled</a></h3>
<p>Not only text inputs can be disabled, but checkboxes and radio buttons as well. Just add <code>disabled</code> attribute or <code>.disabled</code> class to the input.</p>
<div class="example">
  <p><input disabled type="text" value="Text"></p>
  <p><input class="disabled" type="checkbox"></p>
  <p><input disabled type="radio"></p>
  <p>
  <textarea disabled>Text</textarea></p>
  <p><select disabled>
    <option>
      Item
    </option>
  </select></p>
  <pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">input</span> <span class="hljs-attr">type</span>=<span class="hljs-string">"text"</span> <span class="hljs-attr">value</span>=<span class="hljs-string">""</span> <span class="hljs-attr">disabled</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">input</span> <span class="hljs-attr">type</span>=<span class="hljs-string">"checkbox"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"disabled"</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">input</span> <span class="hljs-attr">type</span>=<span class="hljs-string">"radio"</span> <span class="hljs-attr">disabled</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">textarea</span> <span class="hljs-attr">disabled</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">textarea</span>&gt;</span>
<span class="hljs-tag">&lt;<span class="hljs-name">select</span> <span class="hljs-attr">disabled</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">option</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">option</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">select</span>&gt;</span>
</pre>
</div>
<h3 class="section-head" id="h-append"><a href="#h-append">Append &amp; Prepend</a></h3>
<p>You can prepend or append certain elements within your input fields, such as currency characters.</p>
<div class="example">
  <form action="" class="form" method="post">
    <input name="authorize-token" type="hidden" value="c6851d805bd96d2c91d52574c65d3ae26ce4b6bca43f560518a0bea61335f9d52acc6055807bf8616ff3d13d3882d5aa2f4b15a046562c1e6a9d932e369e9fad">
    <div class="form-item">
      <label>How much is it?</label>
      <div class="prepend w50">
        <span>$</span> <input type="text">
      </div>
    </div>
  </form>
  <pre class="code skip">&lt;form method=<span class="hljs-string">"post"</span> action=<span class="hljs-string">""</span> <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"form"</span>&gt;
    <span class="xml"><span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">label</span>&gt;</span>How much is it?<span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"prepend w50"</span>&gt;</span>
            <span class="hljs-tag">&lt;<span class="hljs-name">span</span>&gt;</span>$<span class="hljs-tag">&lt;/<span class="hljs-name">span</span>&gt;</span>
            <span class="hljs-tag">&lt;<span class="hljs-name">input</span> <span class="hljs-attr">type</span>=<span class="hljs-string">"text"</span>&gt;</span>
        <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">form</span>&gt;</span></span>
</pre>
</div>
<div class="example">
  <form action="" class="form" method="post">
    <input name="authorize-token" type="hidden" value="c6851d805bd96d2c91d52574c65d3ae26ce4b6bca43f560518a0bea61335f9d52acc6055807bf8616ff3d13d3882d5aa2f4b15a046562c1e6a9d932e369e9fad">
    <div class="form-item">
      <label>It is much how?</label>
      <div class="append">
        <input type="text"><span>$</span>
      </div>
    </div>
  </form>
  <pre class="code skip">&lt;form method=<span class="hljs-string">"post"</span> action=<span class="hljs-string">""</span> <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"form"</span>&gt;
    <span class="xml"><span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">label</span>&gt;</span>It is much how?<span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"append"</span>&gt;</span>
            <span class="hljs-tag">&lt;<span class="hljs-name">input</span> <span class="hljs-attr">type</span>=<span class="hljs-string">"text"</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">span</span>&gt;</span>$<span class="hljs-tag">&lt;/<span class="hljs-name">span</span>&gt;</span>
        <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">form</span>&gt;</span></span>
</pre>
</div>
<h3 class="section-head" id="h-append-button"><a href="#h-append-button">Button Append</a></h3>
<p>Appending works for buttons as well. Here you can see a button that has been appended and included withing a text input field.</p>
<div class="example">
  <form action="" class="form" method="post">
    <input name="authorize-token" type="hidden" value="c6851d805bd96d2c91d52574c65d3ae26ce4b6bca43f560518a0bea61335f9d52acc6055807bf8616ff3d13d3882d5aa2f4b15a046562c1e6a9d932e369e9fad">
    <div class="form-item">
      <label>Label</label>
      <div class="append w50">
        <input name="search" placeholder="Search" type="text"> <button class="button outline">Go</button>
      </div>
    </div>
  </form>
  <pre class="code skip">&lt;form method=<span class="hljs-string">"post"</span> action=<span class="hljs-string">""</span> <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"form"</span>&gt;
    <span class="xml"><span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">label</span>&gt;</span>Label<span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"append w50"</span>&gt;</span>
            <span class="hljs-tag">&lt;<span class="hljs-name">input</span> <span class="hljs-attr">type</span>=<span class="hljs-string">"text"</span> <span class="hljs-attr">name</span>=<span class="hljs-string">"search"</span> <span class="hljs-attr">placeholder</span>=<span class="hljs-string">"Search"</span>&gt;</span>
            <span class="hljs-tag">&lt;<span class="hljs-name">button</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"button outline"</span>&gt;</span>Go<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span>
        <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">form</span>&gt;</span></span>
</pre>
</div>
<h3 class="section-head" id="h-more-examples"><a href="#h-more-examples">More Examples</a></h3>
<p>The examples below are quite self explanatory and they cover multiple selections, selecting dates, filling in phone numbers, a ready feedback form template and more. Feel free to use any of these examples as templates for your sites.</p>
<div class="example">
  <form action="" class="form" method="post">
    <input name="authorize-token" type="hidden" value="c6851d805bd96d2c91d52574c65d3ae26ce4b6bca43f560518a0bea61335f9d52acc6055807bf8616ff3d13d3882d5aa2f4b15a046562c1e6a9d932e369e9fad">
    <div class="form-item">
      <label>Select multiple</label> <select class="w50" multiple="multiple" size="10">
        <optgroup label="Group 1">
          <option value="1">
            Some text goes here
          </option>
          <option value="2">
            Another choice could be here
          </option>
          <option value="3">
            Yet another item to be chosen
          </option>
        </optgroup>
        <optgroup label="Group 2">
          <option value="4">
            Some text goes here
          </option>
          <option value="5">
            Another choice could be here
          </option>
          <option value="6">
            Yet another item to be chosen
          </option>
        </optgroup>
        <optgroup label="Group 3">
          <option value="7">
            Some text goes here
          </option>
          <option value="8">
            Another choice could be here
          </option>
          <option value="9">
            Yet another item to be chosen
          </option>
        </optgroup>
      </select>
    </div>
  </form>
  <pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">form</span> <span class="hljs-attr">method</span>=<span class="hljs-string">"post"</span> <span class="hljs-attr">action</span>=<span class="hljs-string">""</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form"</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">label</span>&gt;</span>Select multiple<span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">select</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"w50"</span> <span class="hljs-attr">multiple</span>=<span class="hljs-string">"multiple"</span> <span class="hljs-attr">size</span>=<span class="hljs-string">"10"</span>&gt;</span>
            <span class="hljs-tag">&lt;<span class="hljs-name">optgroup</span> <span class="hljs-attr">label</span>=<span class="hljs-string">"Group 1"</span>&gt;</span>
                <span class="hljs-tag">&lt;<span class="hljs-name">option</span> <span class="hljs-attr">value</span>=<span class="hljs-string">""</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">option</span>&gt;</span>
            <span class="hljs-tag">&lt;/<span class="hljs-name">optgroup</span>&gt;</span>
        <span class="hljs-tag">&lt;/<span class="hljs-name">select</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">form</span>&gt;</span>
</pre>
</div>
<div class="example">
  <form action="" class="form" method="post">
    <input name="authorize-token" type="hidden" value="c6851d805bd96d2c91d52574c65d3ae26ce4b6bca43f560518a0bea61335f9d52acc6055807bf8616ff3d13d3882d5aa2f4b15a046562c1e6a9d932e369e9fad"> <label>Choose date</label>
    <div class="row gutters">
      <div class="col col-3">
        <div class="form-item">
          <select>
            <option>
              ---
            </option>
          </select>
          <div class="desc">
            Month
          </div>
        </div>
      </div>
      <div class="col col-3">
        <div class="form-item">
          <select>
            <option>
              ---
            </option>
          </select>
          <div class="desc">
            Day
          </div>
        </div>
      </div>
      <div class="col col-6">
        <div class="form-item">
          <select>
            <option>
              ---
            </option>
          </select>
          <div class="desc">
            Year
          </div>
        </div>
      </div>
    </div>
  </form>
  <pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">form</span> <span class="hljs-attr">method</span>=<span class="hljs-string">"post"</span> <span class="hljs-attr">action</span>=<span class="hljs-string">""</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form"</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">label</span>&gt;</span>Choose date<span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"row gutters"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"col col-3"</span>&gt;</span>
            <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
                <span class="hljs-tag">&lt;<span class="hljs-name">select</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">option</span>&gt;</span>---<span class="hljs-tag">&lt;/<span class="hljs-name">option</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">select</span>&gt;</span>
                <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"desc"</span>&gt;</span>Month<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
            <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
        <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"col col-3"</span>&gt;</span>
            <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
                <span class="hljs-tag">&lt;<span class="hljs-name">select</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">option</span>&gt;</span>---<span class="hljs-tag">&lt;/<span class="hljs-name">option</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">select</span>&gt;</span>
                <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"desc"</span>&gt;</span>Day<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
            <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
        <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"col col-6"</span>&gt;</span>
            <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
                <span class="hljs-tag">&lt;<span class="hljs-name">select</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">option</span>&gt;</span>---<span class="hljs-tag">&lt;/<span class="hljs-name">option</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">select</span>&gt;</span>
                <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"desc"</span>&gt;</span>Year<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
            <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
        <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">form</span>&gt;</span>
</pre>
</div>
<div class="example">
  <form action="" class="form form-inline" method="post">
    <input name="authorize-token" type="hidden" value="c6851d805bd96d2c91d52574c65d3ae26ce4b6bca43f560518a0bea61335f9d52acc6055807bf8616ff3d13d3882d5aa2f4b15a046562c1e6a9d932e369e9fad">
    <div class="form-item">
      <label>Phone number</label> ( <input class="small" name="phone-prefix" size="3" type="text"> ) <input class="small w50" name="phone-number" type="text"> &nbsp;ext: <input class="small" name="phone-ext" size="3" type="text">
      <div class="desc">
        Needed if there are questions about your order
      </div>
    </div>
  </form>
  <pre class="code skip">&lt;form method=<span class="hljs-string">"post"</span> action=<span class="hljs-string">""</span> <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"form form-inline"</span>&gt;
    &lt;div <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"form-item"</span>&gt;
        &lt;label&gt;<span class="hljs-type">Phone</span> number&lt;/label&gt;
        ( &lt;input <span class="hljs-class"><span class="hljs-keyword">type</span></span>=<span class="hljs-string">"text"</span> name=<span class="hljs-string">"phone-prefix"</span> size=<span class="hljs-string">"3"</span> <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"small"</span>&gt; )
        &lt;input <span class="hljs-class"><span class="hljs-keyword">type</span></span>=<span class="hljs-string">"text"</span> name=<span class="hljs-string">"phone-number"</span> <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"small w50"</span>&gt;
        &amp;nbsp;ext: &lt;input <span class="hljs-class"><span class="hljs-keyword">type</span></span>=<span class="hljs-string">"text"</span> name=<span class="hljs-string">"phone-ext"</span> size=<span class="hljs-string">"3"</span> <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"small"</span>&gt;
        &lt;div <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"desc"</span>&gt;...&lt;/div&gt;
    &lt;/div&gt;
&lt;/form&gt;
</pre>
</div>
<div class="example">
  <form class="form" method="post">
    <input name="authorize-token" type="hidden" value="c6851d805bd96d2c91d52574c65d3ae26ce4b6bca43f560518a0bea61335f9d52acc6055807bf8616ff3d13d3882d5aa2f4b15a046562c1e6a9d932e369e9fad">
    <div class="row gutters">
      <div class="col col-6">
        <div class="form-item">
          <label>Email</label> <input type="email">
        </div>
      </div>
      <div class="col col-6">
        <div class="form-item">
          <label>Topic</label> <select>
            <option value="">
              Questions
            </option>
          </select>
        </div>
      </div>
    </div>
    <div class="form-item">
      <label>Message</label> 
      <textarea rows="6"></textarea>
    </div>
  </form>
  <pre class="code skip">&lt;form method=<span class="hljs-string">"post"</span> <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"form"</span>&gt;
    <span class="xml"><span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"row gutters"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"col col-6"</span>&gt;</span>
            <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
                <span class="hljs-tag">&lt;<span class="hljs-name">label</span>&gt;</span>Email<span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
                <span class="hljs-tag">&lt;<span class="hljs-name">input</span> <span class="hljs-attr">type</span>=<span class="hljs-string">"email"</span>&gt;</span>
            <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
        <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"col col-6"</span>&gt;</span>
            <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
                <span class="hljs-tag">&lt;<span class="hljs-name">label</span>&gt;</span>Topic<span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
                <span class="hljs-tag">&lt;<span class="hljs-name">select</span>&gt;</span>
                    <span class="hljs-tag">&lt;<span class="hljs-name">option</span> <span class="hljs-attr">value</span>=<span class="hljs-string">""</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">option</span>&gt;</span>
                <span class="hljs-tag">&lt;/<span class="hljs-name">select</span>&gt;</span>
            <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
        <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">label</span>&gt;</span>Message<span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">textarea</span> <span class="hljs-attr">rows</span>=<span class="hljs-string">"6"</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">textarea</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">form</span>&gt;</span></span>
</pre>
</div>
<div class="example">
  <form action="" class="form" method="post">
    <input name="authorize-token" type="hidden" value="c6851d805bd96d2c91d52574c65d3ae26ce4b6bca43f560518a0bea61335f9d52acc6055807bf8616ff3d13d3882d5aa2f4b15a046562c1e6a9d932e369e9fad">
    <div class="row gutters">
      <div class="col col-6">
        <div class="form-item">
          <label>Country</label> <select>
            <option>
              ---
            </option>
          </select>
          <div class="desc">
            Where are you from?
          </div>
        </div>
      </div>
      <div class="col col-6">
        <div class="form-item">
          <label><br></label> <button>Submit</button>
        </div>
      </div>
    </div>
  </form>
  <pre class="code skip">&lt;form method=<span class="hljs-string">"post"</span> action=<span class="hljs-string">""</span> <span class="hljs-class"><span class="hljs-keyword">class</span></span>=<span class="hljs-string">"form"</span>&gt;
    <span class="xml"><span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"row gutters"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"col col-6"</span>&gt;</span>
            <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
                <span class="hljs-tag">&lt;<span class="hljs-name">label</span>&gt;</span>Country<span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
                <span class="hljs-tag">&lt;<span class="hljs-name">select</span>&gt;</span>
                    <span class="hljs-tag">&lt;<span class="hljs-name">option</span>&gt;</span>---<span class="hljs-tag">&lt;/<span class="hljs-name">option</span>&gt;</span>
                <span class="hljs-tag">&lt;/<span class="hljs-name">select</span>&gt;</span>
                <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"desc"</span>&gt;</span>...<span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
            <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
        <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"col col-6"</span>&gt;</span>
            <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
                <span class="hljs-tag">&lt;<span class="hljs-name">label</span>&gt;</span><span class="hljs-tag">&lt;<span class="hljs-name">br</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
                <span class="hljs-tag">&lt;<span class="hljs-name">button</span>&gt;</span>Submit<span class="hljs-tag">&lt;/<span class="hljs-name">button</span>&gt;</span>
            <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
        <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">form</span>&gt;</span></span>
</pre>
</div>
<div class="example">
  <form action="" class="form" method="post">
    <input name="authorize-token" type="hidden" value="c6851d805bd96d2c91d52574c65d3ae26ce4b6bca43f560518a0bea61335f9d52acc6055807bf8616ff3d13d3882d5aa2f4b15a046562c1e6a9d932e369e9fad">
    <div class="form-item">
      <label>Text</label> 
      <textarea rows="4"></textarea>
    </div>
  </form>
  <pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">form</span> <span class="hljs-attr">method</span>=<span class="hljs-string">"post"</span> <span class="hljs-attr">action</span>=<span class="hljs-string">""</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form"</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">label</span>&gt;</span>Text<span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">textarea</span> <span class="hljs-attr">rows</span>=<span class="hljs-string">"4"</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">textarea</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">form</span>&gt;</span>
</pre>
</div>
<div class="example">
  <form action="" class="form" method="post">
    <input name="authorize-token" type="hidden" value="c6851d805bd96d2c91d52574c65d3ae26ce4b6bca43f560518a0bea61335f9d52acc6055807bf8616ff3d13d3882d5aa2f4b15a046562c1e6a9d932e369e9fad">
    <div class="form-item">
      <label>Text</label> 
      <textarea class="w50" rows="4"></textarea>
    </div>
  </form>
  <pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">form</span> <span class="hljs-attr">method</span>=<span class="hljs-string">"post"</span> <span class="hljs-attr">action</span>=<span class="hljs-string">""</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form"</span>&gt;</span>
    <span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"form-item"</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">label</span>&gt;</span>Text<span class="hljs-tag">&lt;/<span class="hljs-name">label</span>&gt;</span>
        <span class="hljs-tag">&lt;<span class="hljs-name">textarea</span> <span class="hljs-attr">rows</span>=<span class="hljs-string">"4"</span> <span class="hljs-attr">class</span>=<span class="hljs-string">"w50"</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">textarea</span>&gt;</span>
    <span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>
<span class="hljs-tag">&lt;/<span class="hljs-name">form</span>&gt;</span>
</pre>
</div>