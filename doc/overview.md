# FIND3 Overview

## Introduction


## How does it work?

Each time a Bluetooth/WiFi-enabled device conducts a scan of nearby devices, it will recieve a unique identifier and a signal strength that correlates with the distance to that device. A compilation of these different signals can be compiled into a fingerprint which can be used to uniquely classify the current location of that device.

The access points can be anything - routers, Rokus, Raspberry Pis. They also can be anywhere - since they only need to be seen and not connected to, it will successfully use routers that are in a different building.

The basis of this system is to catalog all the fingerprints about the Wifi routers in the area (MAC addresses and signal values) and then classify them according to their location. This is done using a Android App, or computer program, that collects the fingerprints, and then sends them on to the FIND server which can compute the location. 

Locations are determined on the FIND server using classification. Currently the server supports several different machine learning algorithms. Positioning by classification is accomplished by first learning the distributions of WiFi signals for a given location and then classifying it during tracking. Learning only takes ~10 minutes and will last almost indefinitely. The WiFi fingerprints are also the same across all devices so that learning using one device is guaranteed to work across all devices.

## Active and Passive scanning 

There are two modes of localization that you can implement with FIND:

- **Active scanning**: Your device (laptop or smartphone) actively scans for nearby Bluetooth/WiFi devices. It records their signal strengths, and sends these to the FIND3 server. The FIND3 server compares them to its database of known signal strengths, and calculates your device's most likely location. *Requires running a client on your device*.

- **Passive scanning**: The FIND3 scanner runs on 2+ more nearby computers. Each scanner listens for any Bluetooth/Wifi broadcasts from your device. The scanner measures the signal strength of these broadcasts, and sends them to the FIND3 server. The FIND3 server compares them to its database of known signal strengths, and calculates your device's most likely location. *Requires running the FIND3 scanner on 2+ computers. Does not require a client on the device you are trying to locate*.

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


## What is included


## Signing up


## Comparison with existing systems

## Installing and Contributing

The source for FIND3 is hosted on GitHub
([https://github.com/schollz/find3](https://github.com/schollz/find3)).