# FIND3 Overview

## Introduction

The Framework for Internal Navigation and Discovery (FIND) allows you to use your (Android) smartphone or WiFi-enabled computer (laptop or Raspberry Pi or etc.) to determine your position within your home or office. You can easily use this system in place of motion sensors as its resolution will allow your phone to distinguish whether you are in the living room, the kitchen or the bedroom, etc. The position information can then be used in a variety of ways including home automation, way-finding, or tracking!

Simply put, FIND will allow you to replace tons of motion sensors with a single smartphone!

## How does it work?

Each time a Bluetooth/WiFi-enabled device conducts a scan of nearby devices, it will recieve a unique identifier and a signal strength that correlates with the distance to that device. A compilation of these different signals can be compiled into a fingerprint which can be used to uniquely classify the current location of that device.

The access points can be anything - routers, Rokus, Raspberry Pis. They also can be anywhere - since they only need to be seen and not connected to, it will successfully use routers that are in a different building.

The basis of this system is to catalog all the fingerprints about the Wifi routers in the area (MAC addresses and signal values) and then classify them according to their location. This is done using a Android App, or computer program, that collects the fingerprints, and then sends them on to the FIND server which can compute the location. 

Locations are determined on the FIND server using classification. Currently the server supports several different machine learning algorithms. Positioning by classification is accomplished by first learning the distributions of WiFi signals for a given location and then classifying it during tracking. Learning only takes ~10 minutes and will last almost indefinitely. The WiFi fingerprints are also the same across all devices so that learning using one device is guaranteed to work across all devices.

## Framework 

FIND3 is an [open-source project](https://github.com/schollz/find3) that comprises several main components:


1. Data storage server [(github.com/schollz/find3/server/main)](https://github.com/schollz/find3/tree/master/server/main)
2. Machine learning server [(github.com/schollz/find3/server/ai)](https://github.com/schollz/find3/tree/master/server/ai)
3. Command-line tool for gathering fingerprints [(schollz/find3-cli-scanner)](https://github.com/schollz/find3-cli-scanner)
4. Android app for gathering fingerprints [(schollz/find3-android-scanner)](https://github.com/schollz/find3-android-scanner)

Using these elements as building blocks, FIND3 provides the ability to track devices indoors, without motion sensors.


## Active and Passive scanning 

There are two modes of localization that you can implement with FIND:

- **Active scanning**: In *active scanning* the scanner will report the classified location of the device that is doing the scanning. *Requires installing software on the device being tracked*.

- **Passive scanning**: In *passive scanning* the scanner will report the classified location of the devices that it scans. This mode requires having a WiFi card that supports monitor mode on the scanning device. No software is needed on the device that is being tracked.

You can use *active scanning* to track your own phone or computer, as outlined in the [Tracking your phone](/doc/tracking_your_phone.md) and [Tracking your computer](/doc/tracking_your_computer.md) documents. The *passive sacnning* can be used to track others that are nearby, as outline in the [Passive tracking](/doc/passive_tracking.md).

## FIND3 vs FIND

FIND3 is a complete re-write of the previous version of FIND ([github.com/schollz/find](https://github.com/schollz/find)). The API for sending fingerprints (`/track` and `/learn`) is backward compatible. There are several notable improvements on the previous version:

- Support for any data source - Bluetooth / WiFi / magnetic fields / etc. (previously just WiFi).
- Passive scanning built-in (previously required a [separate server](https://github.com/schollz/find-lf)).
- Support for Bluetooth scanning in scanning utility (previously just WiFi).
- Meta-learning with 10 different machine learning classifiers (previously just three).
- Client uses Websockets+React which reduces bandwidth (and coding complexity).
- Rolling compression of MAC addresses for much smaller on-disk databases (see [stringsizer](https://github.com/schollz/stringsizer))
- Data storage in SQLITE-database (previously it was BoltDB).
- Support for WiFi wardriving.


