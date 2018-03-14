# Tracking your phone 


## Introduction

In this tutorial you will learn how to do internal positioning an Android phone using FIND3. Here you will track your Android device using a simple app that scans all nearby APs and Bluetooth discoverables for their distance (via the RSSI measurement). 

## Download the app

First download the latest version of the Android app from Google Play. [Click here to download](https://play.google.com/store/apps/details?id=com.internalpositioning.find3.find3app) the latest FIND3 app.


## Start the app

Find the app on your phone and start it up. When you start you will encounter a prompt about accessing the device's location. Though FIND3 does *not* use GPS, Android devices require location permissions in order to access WiFi and Bluetooth settings. **Press "ALLOW" to continue.**

<center><img src="/images/snap1.PNG" width="30%" height="30%"></center>

### Enter information

When you open the app for the first time you will encounter empty textfields that require data.

<center><img src="/images/1.PNG" width="30%" height="30%"></center>

To get started, enter in a **family** name. The **family** is used to distinguish your group of devices. It can be anything you want, but remember it because you will need it to see your results.

Then enter in a **device** name. The **device** name is used to distinguish this particular device. This can also be the name of the person carrying the device, if that helps you when you see the results.

The **server** is already specified as https://cloud.internalpositioning.com,
the public server. If you are hosting your own server, you can change this
to the address of your self-hosted server. [See here if you want to setup
your own server](/docs/server_setup/).

## Learn some locations

The first thing you need to do after entering data is to **learn the locations for tracking.** This requires walking to each room and doing a scan for about 5 minutes. 

<center><img src="/images/2.PNG" width="30%" height="30%"></center>

Go to a location, like your *kitchen*, *bathroom* or *living room* and enter the name of the location where it says **location (optional)**. Then hit **TRACKING** so it turns to **LEARNING**.  Then hit **START SCAN** and wait about 5 minutes. Then press **STOP SCAN** and repeat this process in each room, for every room you want to learn.

#### *Read the [FAQ](/doc/faq.md#training-time) for more information about how long to do learning in a location.*

## Track yourself

Once you are done learning, simply hit the **LEARNING** button so it toggles back to **TRACKING** and then do **START SCAN**.

<center><img src="/images/3.PNG" width="30%" height="30%"></center>


Tracking will continue in the background and will show a notification that you can use to toggle the tracking off.

<center><img src="/images/backgroundscanning.png" width=70%></center>

The scans will take place approximately 10-30 seconds apart, forever, until you turn off the app. The battery usage is minimal since it is doing only a short WiFi scan and Bluetooth scan.

## Other options

There is an option that says "**Allow GPS**". If activated, the phone will send GPS coordinates if requested by the server. The server only requests GPS coordinates when it has no GPS coordinates for *any* of the recently added scan points. A given location will only generally ever ask for GPS coordinates *once* so it doesn't use any extra battery. In the future these GPS coordinates will be used to do positioning on a map.

## Get data

Once you have learned several locations and are tracking with the computers, you can get data from FIND3 by consulting the [API](/doc/api.md) document.

## Issues?

If you have issues, please file one on Github at https://github.com/schollz/find3.

## Source

If you are interested, the app is completely open-source and available at  https://github.com/schollz/find3-android-scanner.
