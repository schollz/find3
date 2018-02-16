+++
draft= false
title = "FAQ"
description = "Asked and answered"
+++

## Introduction 

In command-line examples on this page, commands to be typed to the shell begin with a dollar sign “$”. Lines that do not begin with “$” show command output.

The text uses you@example.com as a stand-in for your user name. If you run one of these commands, substitute your own name for that one.

```bash
$ run this
```

The examples are Unix-oriented but it should be easy to adapt them to a Windows environment.


## What is FIND?

The Framework for Internal Navigation and Discovery (FIND) allows you to use your (Android) smartphone or WiFi-enabled computer (laptop or Raspberry Pi or etc.) to determine your position within your home or office. You can easily use this system in place of motion sensors as its resolution will allow your phone to distinguish whether you are in the living room, the kitchen or the bedroom, etc. The position information can then be used in a variety of ways including home automation, way-finding, or tracking!

## What's the point of this?

**The point is to eventually incorporate FIND into home automation.** **FIND** can replace motion sensors to provide positional and user-specific information. Anything that you would do with a motion sensor you can do with **FIND**. Anything you can do with GPS information you can do with **FIND** information. Except here you get internal positioning so you could tell apart one table from another in a cafeteria, or one bookshelf from another in a library.

 As Wi-Fi singleboard computers get smaller and smartphones become more ubiquitous there will be more and more opportunities to harness WiFi signals into something useful for other applications.

##  Does FIND use a WiFi location database?

**No.** There is no dependency on external resources like [WiFi location databases](https://en.wikipedia.org/wiki/Wi-Fi_positioning_system#Public_Wi-Fi_location_databases). However, these type of databases can add additional information that might be worthwhile to explore to also integrate into **FIND**.

## How does this work?

It uses already available WiFi information to classify locations. 

Each time a WiFi-enabled device conducts a scan of nearby access points, it will recieve a unique identifier of the access point and a signal strength that correlates with the distance to the access point. A compilation of these different signals can be compiled into a fingerprint which can be used to uniquely classify the current location of that device.

![Room Schematic](/img/room-schematic.png)

The access points can be anything - routers, Rokus, Raspberry Pis. They also can be anywhere - since they only need to be seen and not connected to, it will successfully use routers that are in a different building.

The basis of this system is to catalog all the fingerprints about the Wifi/Bluetooth devices in the area (MAC addresses and signal values) and then classify them according to their location. This is done using a Android App, or computer program, that collects the fingerprints, and then sends them on to the FIND server which can compute the location.

Locations are determined on the FIND server using classification. Currently the server supports a Naive-Bayes implementation, Random Forests, and Support Vector Machines. Positioning by classification is accomplished by first learning the distributions of WiFi signals for a given location and then classifying it during tracking. Learning only takes ~10 minutes and will last almost indefinitely. The WiFi fingerprints are also the same across all devices so that learning using one device is guaranteed to work across all devices.

## Can I use an iPhone for active tracking?

We currently do not support iPhone for active tracking. Unfortunately, the information about the WiFi scanning has to come from the use of the [`Apple80211` library](https://stackoverflow.com/questions/9684341/iphone-get-a-list-of-all-ssids-without-private-library/9684945#9684945). This is private library which means that [a user would have to jail break their device in order to use it](https://stackoverflow.com/questions/6341547/ios-can-i-manually-associate-wifi-network-with-geographic-location/6341893#6341893). We do not want to distribute an app that would require users to jailbreak their phones, so we will have to avoid developing for iOS until Apple removes this restriction. Sorry!



## Doesn't this already exist?

**Yes - but not satisfyingly.** Most solutions are not open-source, or they require external hardware (beacons, etc.), or they are expensive, or they just don't work very well. But don't take my word for it, try it yourself. Here are some of the programs I found that are similar:

If you are looking for a more **commercial, large-scale deployable application**, look at these up-and-coming solutions:

-   [MazeMap Indoor Navigation] - a Norway-based and Cisco-partnered enterprise that takes your CAD floor plans and generates a nice user-interface with similar indoor-positioning capabilities.
-   [Meridian Kits] - a SF and Portland based company (part of Aruba Networks) that offers specialized App SDK environments for building internal positioning systems into workplaces, businesses and hospitals
-   [MPact Platform] - Motorola is working on a internal positioning system that takes advantage of BlueTooth beacons and Wi-Fi for internal positioning for large applications

If you are looking for a **free single-user, single-home application**, perhaps you can find solutions with these apps:

-   [Google Maps Floor Plan Maker] - not sure how it works (and have not tested) but claims to be able to navigate within small businesses. Reviewed okay.
-   [WiFi Indoor Localization] - single-floor grid-based learning system that uses Wi-Fi to train on the names of mac addresses. In my hands it did not work well below 20ft resolution. Reviewed okay.
-   [Indoor Positioning] - Selective learning, not tested by me, but also grid-based. Not reviewed.
-   [BuildNGO - Indoor Navi] - Offers Android app that requires online service for uploading floor plans to their server and uses learning based on Max signal, may require Bluetooth as well.
-   [Wifarer] - Uses Beacons and WiFi for Indoor positioning, but trainable and limited to select museums. Reviewed well, but no training available.
-   [Indoor GPS] - Perfunctory application that trains on a route, instead of a location and offers SDK but still lots of work to be done. Reviewed okay.

  [MazeMap Indoor Navigation]: http://mazemap.com/what-it-is
  [Meridian Kits]: http://www.meridianapps.com
  [MPact Platform]: http://newsroom.motorolasolutions.com/Press-Releases/Communicate-to-Shoppers-at-the-Right-Time-with-First-of-its-Kind-Location-Based-Platform-from-Motor-49e1.aspx
  [Google Maps Floor Plan Maker]: https://play.google.com/store/apps/details?id=com.google.android.apps.insight.surveyor&hl=en
  [WiFi Indoor Localization]: https://play.google.com/store/apps/details?id=com.hfalan.wifilocalization&hl=en
  [Indoor Positioning]: https://play.google.com/store/apps/details?id=com.bombao.projetwifi&hl=en
  [BuildNGO - Indoor Navi]: https://play.google.com/store/apps/details?id=com.sails.buildngo&hl=en
  [Wifarer]: https://play.google.com/store/apps/details?id=com.wifarer.android&hl=en
  [Indoor GPS]: https://play.google.com/store/apps/details?id=com.ladiesman217.indoorgps&hl=en

## Can you run the server on a Raspberry Pi?

**Yes.** Its been tested and runs great on a Raspberry Pi model B+, and model 3. Simply install Docker and build the image:

```bash
$ curl -sSL https://get.docker.com | sh
$ cd /tmp
$ wget https://raw.githubusercontent.com/schollz/find3/master/Dockerfile
$ docker build -t find3 .
$ docker run -p 11883:1883 -p 8003:8003 \
  -v /tmp/find3:/data --name find3server -d -t find3
```

## What is a good amount of time to train a location?

**2 to 5 minutes**. Optimally you want to send ~100 pieces of information to the server. It transmits about 20 per minute, so you should give it some time to train well.

## Can it pick up locations between floors?

**Yes.** Yes it will pick up floors no problem. Floors tend to attenuate the signal, so there is a noticeable difference when you are in the same position, but on different floors. 

## Does it work with Home Assistant?

**Yes.** See [here](https://community.home-assistant.io/t/anyone-seen-this-find-internal-positioning/772/2?u=schollz) for the discussion on how to use it with [home-assistant.io](https://home-assistant.io/).