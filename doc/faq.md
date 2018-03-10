# Frequently Asked Questions

### Introduction

In command-line examples on this page, commands to be typed to the shell begin
with a dollar sign "`$`".
Lines that do not begin with "`$`" show command output.

The examples are Unix-oriented but it should be easy to adapt them to a Windows environment.

## The project

### What is FIND3? {#what}

FIND3 attempts to simplify internal positioning.

Internal positioning, simplified
Using FIND, and only your smartphone or laptop, you will be able to pinpoint your position in your home or office. 

Unlike GPS on your smartphone, this system has a geo-location precision of below 10 sq ft. Also unlike GPS, the app utilizes pre-existing operations so it has no significant cost to your battery. 

This system is open-source and fully configurable but also easy to use and simple to set up. It is suited to small-scale home applications as well as large-scale business and school applications.



### What are the goals of the project? {#goals}

The point is to eventually incorporate FIND into home automation and lifestyle tracking. FIND can replace motion sensors to provide positional and user-specific information. Anything that you would do with a motion sensor you can do with FIND. Anything you can do with GPS information you can do with FIND information. Except here you get internal positioning so you could tell apart one table from another in a cafeteria, or one bookshelf from another in a library.

### Why create something new rather than work on an existing system? {#why-new}

Most solutions are not open-source, or they require external hardware (beacons, etc.), or they are expensive, or they just don’t work very well. But don’t take my word for it, try it yourself. Here are some of the programs I found that are similar:

If you are looking for a more commercial, large-scale deployable application, look at these up-and-coming solutions:

- MazeMap Indoor Navigation - a Norway-based and Cisco-partnered enterprise that takes your CAD floor plans and generates a nice user-interface with similar indoor-positioning capabilities.
- Meridian Kits - a SF and Portland based company (part of Aruba Networks) that offers specialized App SDK environments for building internal positioning systems into workplaces, businesses and hospitals
- MPact Platform - Motorola is working on a internal positioning system that takes advantage of BlueTooth beacons and Wi-Fi for internal positioning for large applications

If you are looking for a free single-user, single-home application, perhaps you can find solutions with these apps:

- Google Maps Floor Plan Maker - not sure how it works (and have not tested) but claims to be able to navigate within small businesses. Reviewed okay.
- WiFi Indoor Localization - single-floor grid-based learning system that uses Wi-Fi to train on the names of mac addresses. In my hands it did not work well below 20ft resolution. Reviewed okay.
- Indoor Positioning - Selective learning, not tested by me, but also grid-based. Not reviewed.
- BuildNGO - Indoor Navi - Offers Android app that requires online service for uploading floor plans to their server and uses learning based on Max signal, may require Bluetooth as well.
- Wifarer - Uses Beacons and WiFi for Indoor positioning, but trainable and limited to select museums. Reviewed well, but no training available.
- Indoor GPS - Perfunctory application that trains on a route, instead of a location and offers SDK but still lots of work to be done. Reviewed okay.


### Can I use FIND3 on my iPhone? {#mobile}

We currently do not support iPhone. 

Unfortunately, the information about the WiFi scanning has to come from the use of the [Apple80211 library](https://stackoverflow.com/questions/9684341/iphone-get-a-list-of-all-ssids-without-private-library/9684945#9684945). This is private library which means that a user would have to [jail break](https://stackoverflow.com/questions/6341547/ios-can-i-manually-associate-wifi-network-with-geographic-location/6341893#6341893) their device in order to use it. We do not want to distribute an app that would require users to jailbreak their phones, so we will have to avoid developing for iOS until Apple removes this restriction. Sorry!

### Does it use a WiFi location database?

There is no dependency on external resources like [WiFi location databases](https://en.wikipedia.org/wiki/Wi-Fi_positioning_system#Public_Wi-Fi_location_databases). However, these type of databases can add additional information that might be worthwhile to explore to also integrate into FIND.

### What is the minimum distance that can be resolved?

It depends. This system harnesses the available WiFi routers and Bluetooth devices. If you have very few WiFi routers in the vicinity (i.e. <3 in 50 meters) then your resolution will suffer. Otherwise, you can typically get less than 10 square feet in location resolution.

### Can it pick up locations between floors?

Yes, because floors tend to attenuate the signal, so there is a noticeable difference when you are in the same position, but on different floors. 

### What is a good amount of time to train a location?

Optimally you should do about 10 minutes per location.

