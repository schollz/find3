# FIND3 Overview

## Introduction


## Active and Passive scanning 

There are two modes of localization that you can implement with FIND:

- **Active scanning**: Your device (laptop or smartphone) actively scans for nearby Bluetooth/WiFi devices. It records their signal strengths, and sends these to the FIND3 server. The FIND3 server compares them to its database of known signal strengths, and calculates your device's most likely location. *Requires running a client on your device*.

- **Passive scanning**: The FIND3 scanner runs on 2+ more nearby computers. Each scanner listens for any Bluetooth/Wifi broadcasts from your device. The scanner measures the signal strength of these broadcasts, and sends them to the FIND3 server. The FIND3 server compares them to its database of known signal strenghts, and calculates your device's most likely location. *Requires running the FIND3 scanner on 2+ computers. Does not require a client on the device you are trying to locate*.

## FIND3 vs FIND

FIND3 is a complete re-write of the previous version of FIND ([github.com/schollz/find](https://github.com/schollz/find)). The API for sending fingerprints (`/track` and `/learn`) is backward compatible. There are several notable improvements on the previous version:

- Support for any data source - Bluetooth / WiFi / magnetic fields / etc. (previously just WiFi)
- Passive scanning built-in (previously required a [separate server](https://github.com/schollz/find-lf))
- Support for Bluetooth scanning in scanning utility (previously just WiFi)
- Meta-learning with 10 different machine learning classifiers (previously just three)
- Client uses Websockets+React which reduces bandwidth (and coding complexity)
- Rolling compression of MAC addresses for much smaller on-disk databases (see [stringsizer](https://github.com/schollz/stringsizer))
- Data storage in SQLITE-database (previously it was BoltDB)


## What is included


## Signing up


## Comparison with existing systems

## Installing and Contributing

The source for FIND3 is hosted on GitHub
([https://github.com/schollz/find3](https://github.com/schollz/find3)).