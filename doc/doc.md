# FIND Documentation 

<img src="/images/find_logo.png" width="180px" alt="Home"/>

Have you ever wanted to...

- do home automation without motion sensors?
- track your phone or laptop *indoors* without GPS?
- monitor the whereabouts and trajectories of cell phones?

FIND stands for the **Framework for Internal Navigation and Discovery**, which is an attempt to implement a solution to these questions. This version, FIND3, is v3.X of FIND (the latest version).

To get started, read the documentation and take a look at the repo. The easiest way to start is to read [Tracking your phone](/doc/tracking_your_phone.md).

## Introduction

- To start understanding FIND and get basic question answered, see the [FAQ](/doc/faq.md).

## User guide

- The [Tracking your phone](/doc/tracking_your_phone.md) document describes how to get started tracking your phone (most people should start here).

- The [Tracking your computer](/doc/tracking_your_computer.md) explains how to install and run the software for using your computer for internal positioning.

- The [Passive tracking](/doc/passive_tracking.md) document instructs how to get started to track any phones in a vicinity.

- The [API](/doc/api.md) document has a detailed guide on how to access the FIND3 data after tracking/learning.

- The [MQTT](/doc/mqtt.md) document will explain how to subscribe to FIND using MQTT.


## System

- The [`find3-cli-scanner`](/doc/cli-scanner.md) is a CLI tool for scanning fingerprints from your computer.

- The [`find3-server`](/doc/server_setup.md) document explains how
  to set up your own FIND3 installation on a Linux server.
  
- The [`esp-client`](https://github.com/DatanoiseTV/esp-find3-client) is a repo that can be used to setup a scanner using ESP8266/ESP32.

## Applications

- The [Home Automation](/doc/automation.md) document describes the basics for getting started with FIND and openHAB or Home Assistant.


## Community

FIND3 is an open source project with a growing community
of users and contributors.
These resources support the open source project and
point to things such as the GitHub repository,
mailing lists, user forums, and so on.

- [The FIND3 project on GitHub](https://github.com/schollz/find3)
- [Issue tracker](https://github.com/schollz/find3/issues)
- Official discussion forums / mailing lists:
  - [Slack channel](https://join.slack.com/t/find3/shared_invite/enQtMzU4MjY0NjE1NjU0LWRkY2JhNWFkM2U3Y2JhY2RlZTQ5ZTdmZTQ2M2UzMjI2MGVmMjZlOWQyZmU3MzM5YzIzOTM0YmYzYmQ3NTQzNjQ),
    for discussion FIND3 users and developers.
  - [FIND mailing list](http://eepurl.com/bhfFI1),
    a low-traffic list for important announcements about the project;
    all FIND3 users should subscribe.
- [Contribution guidelines](https://github.com/schollz/find3/blob/master/CONTRIBUTING.md)


