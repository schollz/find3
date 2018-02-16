[![](https://raw.githubusercontent.com/schollz/find/master/static/splash.gif)](https://www.internalpositioning.com/)

[![Version 3.0](https://img.shields.io/badge/version-3.0.0-brightgreen.svg)](https://www.internalpositioning.com/guide/development/) 
[![Donate](https://img.shields.io/badge/donate-$1-brown.svg)](https://www.paypal.me/ZackScholl/1.00)
 
**The Framework for Internal Navigation and Discovery** (_FIND_) is like GPS, but for your every room in your house/business, with using only a simple smartphone or laptop.

> This version, 3.X, is a complete re-write of the [previous versions 2.x](https://github.com/schollz/find).

There are two modes of localization that you can implement with FIND:

- **Active scanning**: Your device (laptop or a smartphone) tracks itself by actively scanning for nearby Bluetooth/WiFi devices and records their signal strengths and classifying them based on known signal strengths for a given location. *The software for localization must run on the device*.

- **Passive scanning**: Your device (laptop or smartphone) is tracked by 2+ computers that are sniffing for Bluetooth/WiFi broadcasts from your device and using those to classify its location. *The software for localization must run on multiple other computers*.

Documentation: [www.internalpositioning.com](https://www.internalpositioning.com)

## About the project

This repository is a complete re-write of the previous version of FIND ([github.com/schollz/find](https://github.com/schollz/find)). The API for sending fingerprints (`/track` and `/learn`) is backward compatible. There are several notable improvements on the previous version:

- Support for any data source, Bluetooth / WiFi / magnetic fields / etc. (previously just WiFi)
- Passive scanning built-in (previously required a [separate server](https://github.com/schollz/find-lf))
- Support for Bluetooth scanning in scanning utility (previously just WiFi)
- Meta-learning with 10 different machine learning classifiers (previously just three)
- Client uses Websockets+React which reduces bandwidth (and coding complexity)

# Quickstart

```
$ docker build -t find3 .
$ mkdir /tmp/find3
$ docker run -p 11883:1883 -p 8003:8003 -v /tmp/find3:/data --name find3server -d -t find3
```

Then to start/stop

```
$ docker start/stop find3server
```
