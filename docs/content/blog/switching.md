+++
title = "Python -> Go"
description = "Porting FIND to Golang for future development"
tags = [
    "programming",
    "python",
    "golang"
]
date = "2016-04-16 10:08:00"
thumbnail = "https://www.internalpositioning.com/upright.svg"
keywordlist = "python, golang, programming, open-source, app, home automation, home assistant, openhab, particle, esp8266, internal nagivation, indoor positioning, positioning"
+++


I've rewritten FIND in Go because the server is a bit faster (see below), but the real reason is that it saves me $$ because I can run it on the cheapest Digital Ocean (DO) droplet. I run the FIND server on a droplet with along with a half dozen other services.


I only have about 20% of a 500MB of memory to use on my DO machine. I've never been able to write a Python server than can use less then 100MB, so by writing it in Golang it can run pretty easily on cheapest DO!


# Abandoned, but not forgotten

My contributions to this project were stalled out around December, 2015. I was far from done working on FIND though. Instead I began teaching myself Golang in order to rewrite FIND to be faster, and also have a smaller memory footprint. However, if you looked at my Github commits you would have seen quite an empty graph.

<center>
<img class="pure-img" src="/contributions1.PNG"></img>
</center>

# Better, stronger, faster.

The git commit history actually does not show previous versions of FIND before I knew about source control. The first iteration was written in PHP! Then it was written in [Python using Flask, and Tornado](/post/the-find-stack/). However, each time it was rewritten it got a little bit better.

This time took several months, but I finally got it working well in Golang. The speedup results are very promising:
<center>
<table class="pure-table">
<thead>
<tr class="header">
<th>Version</th>
<th>Fingerprints sent to /learn</th>
<th>Optimizing priors through /calculate</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td><a href="https://github.com/schollz/find/tree/python3">Python</a></td>
<td>15 fingerprints/sec</td>
<td>3 calculations/min</td>
</tr>
<tr class="even">
<td>Go</td>
<td>76 fingerprints/sec</td>
<td>619 calculations/min</td>
</tr>
</tbody>
</table>
</center>

There is still a lot of features I'd like to implement, but I will start working more closely with version control so my work is more transparent.