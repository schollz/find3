# Frequently Asked Questions

### Introduction

This is a compilation of the most frequently asked questions. If you do not find your question here, please feel free to [ask in the Slack channel](https://join.slack.com/t/find3/shared_invite/enQtMzI0MjkwMjc3MDYzLWJiZWEzZjU5NTljM2JlYmE1MDY0NThiYmY2NDYwNGYxNTNmNTJjZjFmNjMwNGMwY2UyNzczNzZhZTIxZWY3ODQ) or [send me a message](https://www.internalpositioning.com/#cta-5).

## The project

### What is FIND3? {#what}

FIND3 simplifies internal positioning. 

The name, "FIND" stands for *the Framework for Internal Navigation and Discovery*. The number 3 specifically indicates the 3rd and latest version.

Using FIND, and only your smartphone or laptop, you will be able to pinpoint your position in your home or office by evaluating surrounding radio waves (Bluetooth or WiFi).  Unlike GPS on your smartphone, this system has a geo-location precision of below 10 sq ft. Also unlike GPS, the app utilizes pre-existing operations so it has no significant cost to your battery. 

This system is open-source and fully configurable but also easy to use and simple to set up. It is suited to small-scale home applications as well as large-scale business and school applications.


### What are the goals of the project? {#goals}

FIND started out as a way to replace motion sensors.

The point is to eventually incorporate FIND into home automation and lifestyle tracking. FIND can replace motion sensors to provide positional and user-specific information. Many things that you would do with a motion sensor you can do with FIND. Also, many things that you can do with GPS information you can do with FIND information. Except here you get internal positioning so you could tell apart one table from another in a cafeteria, or one bookshelf from another in a library.

### Why create something new rather than work on an existing system? {#why-new}

I started working on FIND in 2009 when there were very few, if any solutions for internal positioning. Currently, there are more (even the latest Android P is offering support using RTT now). Unfortunately, still, most solutions are not open-source, or they require external hardware (beacons, etc.), or they are expensive, or they just don’t work very well. 


### Can I use FIND3 on my iPhone? {#mobile}

FIND does not support iPhones. 

Unfortunately, the information about the WiFi scanning has to come from the use of the [Apple80211 library](https://stackoverflow.com/questions/9684341/iphone-get-a-list-of-all-ssids-without-private-library/9684945#9684945). This is private library which means that a user would have to [jail break](https://stackoverflow.com/questions/6341547/ios-can-i-manually-associate-wifi-network-with-geographic-location/6341893#6341893) their device in order to use it. We do not want to distribute an app that would require users to jailbreak their phones, so we will have to avoid developing for iOS until Apple removes this restriction. Sorry!

### Does it use a WiFi location database?

There is no dependency on external resources like [WiFi location databases](https://en.wikipedia.org/wiki/Wi-Fi_positioning_system#Public_Wi-Fi_location_databases). However, these type of databases can add additional information that might be worthwhile to explore to also integrate into FIND.

### What is the minimum distance that can be resolved?

It depends. This system harnesses the available WiFi routers and Bluetooth devices. If you have very few WiFi routers in the vicinity (i.e. <3 in 50 meters) then your resolution will suffer. Otherwise, you can typically get less than 10 square feet in location resolution.

### Can it pick up locations between floors?

Yes, because floors tend to attenuate the signal, so there is a noticeable difference when you are in the same position, but on different floors. 

### How long does it take to learn a location? {#training-time}

At a minimum you you should do learning in each location for about 5 minutes. After that, you can go to the dashboard to see the results of your training. To see the dashboard, goto [cloud.internalpositioning.com](https://cloud.internalpositioning.com) and sign in with your family name. At the dashboard you will the results from the last calibration, which may look like:

<center>
<img src="/images/accuracy2.png">
</center>

The **Overall** gives the average accuracy, while the accuracy for other locations are also shown. These accuracies are determined by cross-validation, so they their representation of reality will correlate with the amount of data available.

If you see that one of the locations has low accuracy, then you should do more learning. To do more learning, simply take your device and send sensor data with the given location. Then, when you are done you can re-calibrate by hitting the calibrate button.

<center>
<img src="/images/calibrate.png">
</center>

As an example, when starting learning I noticed I had low accuracy for the "Bathroom" location. I opened the dashboard to see that the accuracy was 50% (not very good).

<center>
<img src="/images/accuracy50_ss.png" width="70%">
</center>

To amend this I did more learning at that location and then I reloaded the browser with the dashboard again to see the results. After inserting about 70 more data, I had increased the accuracy to 87%.

<center>
<img src="/images/accuracy87.png" width="70%">
</center>

At this point, the accuracy had improved enough for me to move on to learn other locations. *Note:* as you learn new locations, they might end up being too similar to previous locations which could decrease the accuracy of previously learned locations. This is dependent on the number of available sensor points in the vicinity.