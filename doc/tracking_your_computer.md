# Tracking your computer 


## Introduction

In this tutorial you will learn how to do internal positioning on a computer using FIND3.


## Install the scanning tool

First you will need the FIND3 command-line scanner. You can install this the easy way, using Docker, or the hard way by compiling from source.

### Easy way


Install Docker:

```
$ curl -sSL https://get.docker.com | sh
```

If *not* using a Raspberry Pi, fetch the latest image.

```
$ docker pull schollz/find3-cli-scanner
```

If you are using a Raspberry Pi (`armf` arch), you need to build the image yourself.

```
$ wget https://raw.githubusercontent.com/schollz/find3/master/scanner/Dockerfile
$ docker build -t schollz/find3-cli-scanner .
```

Now you can start the scanning image in the background.

```
$ docker run --net="host" --privileged --name scanner -d -i -t schollz/find3-cli-scanner
```

To use the scanner, your syntax will be

```
$ docker exec scanner sh -c "X"
```

where "`X`" is the command for the command-line tool, as specified below.

You can start/stop the image using

```
$ docker start scanning
$ docker stop scanning
```

> Note, you can jump inside the image and play if you are curious of trying new things.
```
$ docker run --net="host" --privileged --name scanning -i -t scanner /bin/bash
```
> 


### Hard way

I don't recommed this because I can't gaurantee that all the processes that the scanner calls will work in every OS. I can tell you that these instructions will work on Ubuntu16/Raspbian though.

Install the dependencies.

```
$ sudo apt-get install wireless-tools iw net-tools
```

(Optional) If you want to do Bluetooth scanning too, then also:

```
$ sudo apt-get install bluetooth
```

(Optional) If you want to do Passive scanning, then do:

```
$ sudo apt-get install tshark
```

Now, you can  [download the command-line scanner](https://github.com/schollz/find3-cli-scanner/releases/latest).

Or you can build from source. First [install Go](https://golang.org/dl/) and pull the latest:

```
$ go get -u -v github.com/schollz/find3-cli-scanner
```

Then you can install it using

```
$ go install github.com/schollz/find3-cli-scanner
```



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

If you have issues, please file one on Github at https://github.com/schollz/find3-android-scanner/issues.

## Source

If you are interested, the app is completely open-source and available at  https://github.com/schollz/find3-android-scanner.