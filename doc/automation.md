# Home automation


## Introduction

The following tutorials are not soup to nuts for setting up the respective home automation software. Those tutorials are well established on their respective pages. 

Here I've collected the basics for integrating the FIND MQTT endpoints with two very popular open-source home automation frameworks: [Home Assistant](https://home-assistant.io/), and [openHAB](https://www.openhab.org/).


## Pre-requisites


Before beginning you need to get your MQTT password for your FIND family. To get one just do

```
$ curl https://cloud.internalpositioning.com/api/v1/mqtt/FAMILY
{"message":"Added 'FAMILY' for mqtt. Your passphrase is 'MQTT_PASS'","success":true}
```

The `MQTT_PASS` is what you'll need for all of the configurations. This password is unique to you and will change each time you make that request, so do not forget it. For more information, see the [MQTT](/doc/mqtt.md) document.


## Home Assistant

First [get started with Home Assistant](https://home-assistant.io/docs/installation/).

Add to the `configuration.yaml`:

```
mqtt:
  broker: cloud.internalpositioning.com
  port: 1883
  keepalive: 60
  username: FAMILY
  password: MQTT_PASS
  protocol: 3.1

sensor:
  - platform: mqtt
    state_topic: 'FAMILY/location/USER'
    name: USER_location
    value_template: '{{ value_json.guesses[0].location }}'
```

The `FAMILY` is the family name you use for FIND. The `USER` is a given user you have for FIND. You can setup multiple users on Home Assistant. The `MQTT_PASS` is the password generated from FIND. The broker and port is the default server (`cloud.internalpositioning.com:1883`), but you can change those if you are hosting yourself.

There is some more documentation on the [Home Assistant forms](https://community.home-assistant.io/t/anyone-seen-this-find-internal-positioning/772).

## openHAB

First [get started with openHAB](https://docs.openhab.org/tutorials/beginner/).

You need to get the [MQTT Binding](http://docs.openhab.org/addons/bindings/mqtt1/readme.html) to get data from FIND. Once installed, create a service `services/mqtt.cfg` with the following:

```
broker.url=tcp://cloud.internalpositioning.com:1883
broker.clientId=OpenHAB
broker.user=FAMILY
broker.pwd=MQTT_PASS
```

The `FAMILY` is the family name you use for FIND. The `MQTT_PASS` is the password generated from FIND. The broker and port is the default server (`cloud.internalpositioning.com:1883`), but you can change those if you are hosting yourself.

To track the information you need to add the username of the person into the `location.items` file.

```
String	mqttfind_USER			"USER is @ [%s]"	(All)	{mqtt="<[BROKER:FAMILY/location/USER:state:JSONPATH($.gueses[0].location)]"}
```
The `USER` is a given user you have for FIND. Make sure to have JSON transformations enabled (the addon can be added by putting it in the transformation line in `services/addons.cfg`).

There is some more documentation on the [openHAB forms](https://community.openhab.org/t/find-personalized-indoor-localization/35799).