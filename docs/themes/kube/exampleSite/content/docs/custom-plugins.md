+++
title = "Custom Plugins"
description = "Flexible and extensible way to add custom features"
weight = 20
draft = false
toc = true
bref = "Kube has a lot to offer in terms of extensibility and flexibility, and plugins are the way of doing incredible things without bloating the core. With plugins, you can extend existing features, make them more interactive, or you can crate completely new functionality."
+++

<h3 class="section-head" id="h-plugin"><a href="#h-plugin">Plugin Template</a></h3>
<p>Here's what a generic plugin looks like. This template gives an overall idea of what you can do with plugins in Kube. Feel free to use this one as a boilerplate for your custom plugins.</p>
<pre class="code skip">(<span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params">Kube</span>)
</span>{
    Kube.Myplugin = <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params">element, options</span>)
    </span>{
        <span class="hljs-keyword">this</span>.namespace = <span class="hljs-string">'myplugin'</span>;

        <span class="hljs-comment">// default settings</span>
        <span class="hljs-keyword">this</span>.defaults = {
            mysetting: <span class="hljs-literal">true</span>
        };

        <span class="hljs-comment">// Parent Constructor</span>
        Kube.apply(<span class="hljs-keyword">this</span>, <span class="hljs-built_in">arguments</span>);

        <span class="hljs-comment">// Initialization</span>
        <span class="hljs-keyword">this</span>.start();
    };

    <span class="hljs-comment">// Functionality</span>
    Kube.Myplugin.prototype = {
        start: <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
        </span>{
            <span class="hljs-comment">// plugin element</span>
            <span class="hljs-built_in">console</span>.log(<span class="hljs-keyword">this</span>.$element);

            <span class="hljs-comment">// call options</span>
            <span class="hljs-built_in">console</span>.log(<span class="hljs-keyword">this</span>.opts.mysetting);

            <span class="hljs-comment">// call methods</span>
            <span class="hljs-keyword">this</span>.method();
        },
        method: <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
        </span>{
            <span class="hljs-comment">// do something...</span>

            <span class="hljs-comment">// callback</span>
            <span class="hljs-keyword">this</span>.callback(<span class="hljs-string">'show'</span>);

            <span class="hljs-comment">// callback with arguments</span>
            <span class="hljs-keyword">this</span>.callback(<span class="hljs-string">'show'</span>, value1, value2);
        }
    };

    <span class="hljs-comment">// Inheritance</span>
    Kube.Myplugin.inherits(Kube);

    <span class="hljs-comment">// Plugin</span>
    Kube.Plugin.create(<span class="hljs-string">'Myplugin'</span>);
    Kube.Plugin.autoload(<span class="hljs-string">'Myplugin'</span>);

}(Kube));
</pre>
<h3 class="section-head" id="h-call"><a href="#h-call">Call</a></h3>
<p>Calling a plugin is very easy, just add the <code>data-component</code> with the name of your plugin and it will start automatic.</p>
<pre class="code">&lt;<span class="hljs-keyword">div</span> data-component=<span class="hljs-string">"myplugin"</span>&gt;&lt;/<span class="hljs-keyword">div</span>&gt;</pre>
<p>Or call manually</p>
<pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"my-element"</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>

<span class="hljs-tag">&lt;<span class="hljs-name">script</span>&gt;</span><span class="javascript">
    $(<span class="hljs-string">'#my-element'</span>).myplugin();
</span><span class="hljs-tag">&lt;/<span class="hljs-name">script</span>&gt;</span>
</pre>
<h3 class="section-head" id="h-callbacks"><a href="#h-callbacks">Callbacks</a></h3>
<p>Kube plugins can react on events with callbacks. Whenever you need your plugin to do something in response to an action or an event, just use a callback.</p>
<pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">div</span> <span class="hljs-attr">id</span>=<span class="hljs-string">"myplugin"</span> <span class="hljs-attr">data-component</span>=<span class="hljs-string">"myplugin"</span>&gt;</span><span class="hljs-tag">&lt;/<span class="hljs-name">div</span>&gt;</span>

<span class="hljs-tag">&lt;<span class="hljs-name">script</span>&gt;</span><span class="javascript">
$(<span class="hljs-string">'#myplugin'</span>).on(<span class="hljs-string">'show.myplugin'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// do something...</span>
});
</span><span class="hljs-tag">&lt;/<span class="hljs-name">script</span>&gt;</span>
</pre>
<p>All plugin methods and variables are available within the plugin via <var>this</var>:</p>
<pre class="code skip"><span class="hljs-tag">&lt;<span class="hljs-name">script</span>&gt;</span><span class="javascript">
$(<span class="hljs-string">'#myplugin'</span>).on(<span class="hljs-string">'show.myplugin'</span>, <span class="hljs-function"><span class="hljs-keyword">function</span>(<span class="hljs-params"></span>)
</span>{
    <span class="hljs-comment">// plugin element</span>
    <span class="hljs-built_in">console</span>.log(<span class="hljs-keyword">this</span>.$element);

    <span class="hljs-comment">// call plugin method</span>
    <span class="hljs-keyword">this</span>.method();
});
</span><span class="hljs-tag">&lt;/<span class="hljs-name">script</span>&gt;</span>
</pre>