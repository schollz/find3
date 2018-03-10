# Passive tracking 

## Introduction 

It is possible to also use FIND3 to setup a system that can do *passive scanning*. In *passive scanning*, you setup multiple computers which capture packets from phones and use those to classify their location. In this mode the packets used for classification are only broadcast packets, that is packets that originate from a device that are being transmitted to all devices on the network. See [Time resolution](#time-resolution) for information on typical frequencies for these types of broadcasts.

A typical passive scanning system uses a network of Raspberry Pis which sniff the WiFi broadcast requests from WiFi-enabled devices and sends these parcels to a central server which compiles and forwards the fingerprint to the FIND server which then uses machine learning to classify the location based on the unique WiFi fingerprints.

This system does not require being logged into a particular WiFi - it will track any phone/device with WiFi enabled! (Caveat: for iOS devices it will only track if Wi-Fi is associated with a network - any network, though - because of MAC spoofing it uses for security). This system also does not require installing any apps on a phone.

*Note: It may be illegal to monitor networks for MAC addresses, especially on networks that you do not own. Please check your country's laws (for [US Section 18 U.S. Code § 2511](https://www.law.cornell.edu/uscode/text/18/2511)) - [discussion](https://github.com/schollz/howmanypeoplearearound/issues/4).*

## Time resolution {#time-resolution}

The time resolution of passive tracking revolves around the frequency that the device makes WiFi broadcasts. To get an idea of the frequency of these broadcasts, and thus the minimum time separation you will get with passive tracking, here are some measurements I made.

### Affiliated devices

These experiments were done with devices that were affiliated with a WiFi network.

- When left unattended my laptop running Ubuntu 17 emits a WiFi broadcast every 2.1 ± 0.4 minutes. A desktop computer running Ubuntu 17 is similar, sending out a broadcast every 2.0 ± 0.4 minutes. 
- A Google Home emits a WiFi broadcast every 6.5 ± 3.5 minutes. 
- A Pixel2 running the [`find3-android-scanner`](https://play.google.com/store/apps/details?id=com.internalpositioning.find3.find3app) in the background will be detected every 2.1 ± 0.8 minutes.
- A Samsung phone without any active scanning seems to run only when its screen is unlocked (i.e phone is used). When the screen is on, and it is being used, it scans every 2.0 ± 0.1 minutes.

### Unaffiliated devices

TBD.

## Prerequisites

You will need 1+ scanner computers. Raspberry Pis with built-in Wifi work best:

- [Raspberry Pi Zero W](https://www.amazon.com/gp/product/B071L2ZQZX/ref=as_li_tl?ie=UTF8&tag=scholl-20&camp=1789&creative=9325&linkCode=as2&creativeASIN=B071L2ZQZX&linkId=ab2f9d564a4f517c5b004a760d0d6e29)
- [Raspberry Pi 3](https://www.amazon.com/gp/product/B01C6EQNNK/ref=as_li_tl?ie=UTF8&tag=scholl-20&camp=1789&creative=9325&linkCode=as2&creativeASIN=B01C6EQNNK&linkId=805012388be781415a6be827b50c76ac)

You will need a monitor-mode enabled wifi USB adapter. There are a number of possible USB WiFi adapters that support monitor mode. Here's a list that are popular:

- [USB Rt3070 $14](https://www.amazon.com/gp/product/B00NAXX40C/ref=as_li_tl?ie=UTF8&tag=scholl-20&camp=1789&creative=9325&linkCode=as2&creativeASIN=B00NAXX40C&linkId=b72d3a481799c15e483ea93c551742f4)
- [Panda PAU5 $14](https://www.amazon.com/gp/product/B00EQT0YK2/ref=as_li_tl?ie=UTF8&tag=scholl-20&camp=1789&creative=9325&linkCode=as2&creativeASIN=B00EQT0YK2&linkId=e5b954672d93f1e9ce9c9981331515c4)
- [Panda PAU6 $15](https://www.amazon.com/gp/product/B00JDVRCI0/ref=as_li_tl?ie=UTF8&tag=scholl-20&camp=1789&creative=9325&linkCode=as2&creativeASIN=B00JDVRCI0&linkId=e73e93e020941cada0e64b92186a2546)
- [Panda PAU9 $36](https://www.amazon.com/gp/product/B01LY35HGO/ref=as_li_tl?ie=UTF8&tag=scholl-20&camp=1789&creative=9325&linkCode=as2&creativeASIN=B01LY35HGO&linkId=e63f3beda9855abd59009d6173234918)
- [Alfa AWUSO36NH $33](https://www.amazon.com/gp/product/B0035APGP6/ref=as_li_tl?ie=UTF8&tag=scholl-20&camp=1789&creative=9325&linkCode=as2&creativeASIN=B0035APGP6&linkId=b4e25ba82357ca6f1a33cb23941befb3)
- [Alfa AWUS036NHA $40](https://www.amazon.com/gp/product/B004Y6MIXS/ref=as_li_tl?ie=UTF8&tag=scholl-20&camp=1789&creative=9325&linkCode=as2&creativeASIN=B004Y6MIXS&linkId=0277ca161967134a7f75dd7b3443bded)
- [Alfa AWUS036NEH $40](https://www.amazon.com/gp/product/B0035OCVO6/ref=as_li_tl?ie=UTF8&tag=scholl-20&camp=1789&creative=9325&linkCode=as2&creativeASIN=B0035OCVO6&linkId=bd45697540120291a2f6e169dcf81b96)
- [Sabrent NT-WGHU $15 (b/g) only](https://www.amazon.com/gp/product/B003EVO9U4/ref=as_li_tl?ie=UTF8&tag=scholl-20&camp=1789&creative=9325&linkCode=as2&creativeASIN=B003EVO9U4&linkId=06d4784d38b6bcef5957f3f6e74af8c8)

Namely you want to find a USB adapter with one of the following chipsets: Atheros AR9271, Ralink RT3070, Ralink RT3572, or Ralink RT5572.

The commands I'll describe use `httpie` which can be installed with Python.

```
$ sudo python3 -m pip install httpie
```

## Setup a scanner computer

For each scanner computer you will need to use the scanning software. Follow the instructions in the [Command-line scanner](/doc/cli-scanner.md) document to install the FIND3 command-line scanner.

As before, determine your **family name** (here `FAMILY`), your **device name** (here `DEVICE`) and your WiFi interface (here `wlan0`). Make sure that the WiFi interface that you specify supports promiscuous mode.


### Start scanning passively

Choose a **device name**, like the name of your computer. We will use `DEVICE` for the rest of this document. The device name should be unique to each scanning computer.

Choose a **family name** which is a unique namespace that you can use to store data for all your devices. Each scanning computer should have the *same* family name. We will use `FAMILY` for the rest of this document.

You need to run the scanner commands using `sudo` to have privileges to modify the WiFi card. However, if you are using Docker you don't need the `sudo` command.

```
$ sudo ./find3-cli-scanner -i wlan0 -device DEVICE -family FAMILY \
    -server https://cloud.internalpositioning.com \
    -scantime 40 -forever -passive
```

This command-line flag `-passive` tells the scanner to capture the packets with `tshark`. This command will start a scanner that submits to the main server (https://cloud.internalpositioning.com). If you set the `-forever` flag it will also continue running forever.

In this command  the WiFi chip set/unset the promiscuous mode after every scan so that it can connect to the internet to upload the packets. This process takes about 10 seconds, so it is useful to set it permanently if you don't need to connect to the internet with the scanning interface (i.e. you have two WiFi interfaces).

If you have two WiFi interfaces, you can set one to be promiscuous permanently.

```
$ sudo ./find3-cli-scanner -i wlan0 -monitor-mode
```

Then, add the `-no-modify` flag to tell the tool not to alter the promiscuousness of the interface

```
$ sudo ./find3-cli-scanner -i wlan0 -device DEVICE -family FAMILY \
    -server https://cloud.internalpositioning.com \
    -scantime 40 -forever -passive -no-modify
```


## Learning

Unlike the active scanning, to do learning on the passive scanning mode you must tell the server which device to learn on. *You should not stop the scanning tool that is running on the scanning computers*. 

Choose the name of the family you are using, here we will use "`FAMILY`".

Then choose the device you would like to learn on. You can use multiple devices, computers or smartphones. You will need to get the MAC address of the device you are using. In smartphones this is under settings, and in computers it is located in the `ifconfig`. The device name in passive mode is also prefixed by `wifi-`, so if your MAC address is `60:57:18:3d:b8:14` your device name should be `wifi-60:57:18:3d:b8:14`. *Note:* If you are doing learning on your phone you can [download the `find3-android-scanner`](https://play.google.com/store/apps/details?id=com.internalpositioning.find3.find3app) to speed up the number of broadcasts the phone makes.

Once you have the device name you can tell the server where that device is, say the "living room". 

```
$ http POST https://cloud.internalpositioning.com/api/v1/settings/passive \
     family=FAMILY device=wifi-60:57:18:3d:b8:14 location="living room"
```

The family (`FAMILY`) and device (for example, `wifi-60:57:18:3d:b8:14`) are specified.

Leave the device in the location for about 30 minutes to collect a good amount of fingerprints. This will allow you to collect about 15 pieces of sensor data per location.

It is *very important* to stop learning before you move the device away from that location. To stop learning, use the same command as above but without the `location` parameter:

```
$ http POST https://cloud.internalpositioning.com/api/v1/settings/passive \
   family=FAMILY device=wifi-60:57:18:3d:b8:14
```

## Tracking

When scanner computers are running the `find3-cli-scanner` tool, then all devices are always being tracked. The tracking information has no value until you are finished learning. Once you have finished learning, then you can gather information about the devices using data gathering specified in the [API](/doc/api.md#tracking) document.

## Optional customization

### Custom scan times

Each scanning computer submits the data point for one device. The server synchronizes all the scanning computers by waiting a specified amount of time (the time window) for collecting the data point for each device. This time window is 90 seconds by default which is enough time to guarantee that the server will hear from every scanning computer (that have a scan time of 40 seconds). You can change these parameters.

To change the scantime on a scanning computer just use the flag `-scantime`. 

To change the window to `X` seconds, use:

```
$ http POST https://cloud.internalpositioning.com/api/v1/settings/passive \
    family=FAMILY window:=X
```

Make sure to change the time window on the server for collecting data to at least twice the scan time of a scanning computer. For instance, if you set the scantime on the scanning computers to 10 seconds, you should change the server window to about 25 seconds.

### Minimum passive scan points

After the window of data collection, all fingerprints with at least 1 data point will be processed. To change this you can set `miniumum_passive`. For instance, if you have three scanning computers, and you want to ensure that all passive scanning data processed has data from at least 2 scanning computers you should do

```
$ http POST https://cloud.internalpositioning.com/api/v1/settings/passive \
    family=FAMILY minimum_passive:=2
```

## Issues?

If you have issues, please file one on Github at https://github.com/schollz/find3/issues.

## Source

If you are interested, the app is completely open-source and available at  https://github.com/schollz/find3.