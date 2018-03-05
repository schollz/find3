+++
title = "Quickstart for active scanning on a phone"
description = "Guide to quickstart setup of the FIND3 services"
weight = 10
draft = false
toc = true
bref = "You can position yourself in about 10 minutes using this guide."
+++

<h2>Introduction</h2>

In this tutorial you will learn how to do internal positioning an Android phone. Here you will track your Android device using a simple app that scans all nearby APs and Bluetooth discoverables for their distance (via the RSSI measurement). 


<h3 class="marker" id="download">Download the app</h3>

First download the latest version of the Android app from Google Play. [Click here to download](https://play.google.com/store/apps/details?id=com.internalpositioning.find3.find3app) the latest FIND3 app.

<center><img src="/img/snap3.PNG"></center>

<h3 class="marker" id="start">Start the app</h3>

Find the app on your phone and start it up. When you start you will encounter a prompt about accessing the device's location. Though FIND3 does *not* use GPS, Android devices require location permissions in order to access WiFi and Bluetooth settings. **Press "ALLOW" to continue.**

<center><img src="/img/snap1.PNG"></center>

<h3 class="marker" id="learn">Enter information</h3>

When you open the app for the first time you will encounter empty textfields that require data.

<center><img src="/img/snap2.PNG"></center>

To get started, enter in a **family** name. The **family** is used to distinguish your group of devices. It can be anything you want, but remember it because you will need it to see your results.

Then enter in a **device** name. The **device** name is used to distinguish this particular device. This can also be the name of the person carrying the device, if that helps you when you see the results.

The **server** is already specified as **find3.internalpositioning.com**, the public server. If you are hosting your own server, you can change this to the address of your self-hosted server. [See here if you want to setup your own server](/docs/server_setup/).

<h3 class="marker" id="learn">Learn a location</h3>

The first thing you need to do after entering data is to **learn the locations for tracking.** This requires walking to each room and doing a scan for about 5 minutes. 

Go to a location, like your *kitchen*, *bathroom* or *living room* and enter the name of the location where it says **location (optional)**. Then hit **START SCAN** and wait about 5 minutes. Then press **STOP SCAN** and repeat this process in each room, for every room you want to learn.

<h3 class="marker" id="track">Track yourself</h3>

Once you are done learning, simply delete the location
