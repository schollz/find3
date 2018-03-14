# Tracking your computer 


## Introduction

In this tutorial you will learn how to do internal positioning on a computer using FIND3.

## Pre-requisites

First you will need the FIND3 command-line scanner. You can install this the easy way, using Docker, or the hard way by compiling from source. Follow the instructions in the [Command-line scanner](/doc/cli-scanner.md) document to install the scanner.

## Learning

To begin using FIND, you will need to learn the surroundings by putting your device in a location and gathering the signals around it.


First determine the name of your WiFi interface.

```
$ iwconfig
```

For the rest of this document we will assume its `wlan0`, a common name of the interface. 

Choose a **device name**, like the name of your computer. We will use `DEVICE` for the rest of this document. 

Choose a **family name** which is a unique namespace that you can use to store data for all your devices. We will use `FAMILY` for the rest of this document.

To do learning with the tool, you can set the learning flag `-location`. Say, for instance you have your computer in the "living room", you can run the following command.

To start scanning use the following command (this is the "`X`" if you are using Docker).

```
$ find3-cli-scanner -i wlan0 -device DEVICE -family FAMILY \
    -server https://cloud.internalpositioning.com \
    -scantime 10 -bluetooth -forever -location "living room"
```

This command will start a scanner that submits to the main server (https://cloud.internalpositioning.com). It uses a scan time of 10 seconds, and it scan bluetooth (`-bluetooth`). If you set the "`-forever`" flag it will also continue running forever.

For your tracking scans to work, you must go to each room and run the learning command for about 10 minutes. 


#### *Read the [FAQ](/doc/faq.md#training-time) for more information about how long to do learning in a location.*


Once you have finished learning each room, do a calibration to update the machine learning algorithms.

```
$ http GET https://cloud.internalpositioning.com/api/v1/calibrate/FAMILY
```


## Start tracking

After learning is accomplished, you can track your device.

```
$ find3-cli-scanner -i wlan0 -device DEVICE -family FAMILY \
    -server https://cloud.internalpositioning.com \
    -scantime 10 -bluetooth -forever
```

The command for tracking is the same as for learning, but without the `-location` flag.

## Get data

Once you have learned several locations and are tracking with the computers, you can get data from FIND3 by consulting the [API](/doc/api.md) document.


## Issues?

If you have issues, please file one on Github at https://github.com/schollz/find3-cli-scanner/issues.

## Source

If you are interested, the app is completely open-source and available at  https://github.com/schollz/find3-cli-scanner.