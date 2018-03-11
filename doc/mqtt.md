# MQTT

## Introduction

MQTT makes it very easy to recieve push notifications of new data generated from FIND. The MQTT protocl is very useful for integrating for embedded electronics or home automation software like [OpenHAB](https://www.openhab.org/) or [HomeAssistant](https://home-assistant.io/).

## Getting started 

FIND has first-class support for MQTT. To get started with MQTT you first need to generate a password associated with your family. You can get your password, `XX` by make a request to the server.

```
$ curl https://cloud.internalpositioning.com/api/v1/mqtt/FAMILY
{"message":"Added 'FAMILY' for mqtt. Your passphrase is 'XX'","success":true}
```

## Subscribe to messages

If don't already, install the `mosquitto` client or similar MQTT broker client to read message,

```
$ sudo apt-get install mosquitto-clients
```

Then you can subscribe to the latest data updates using the password "`XX`" provided by the server.

```
$ mosquitto_sub -h cloud.internalpositioning.com -p 1883 \
    -u FAMILY -P XX -t 'FAMILY/location/#'
```

The above will automatically listen to every device. To listen to a particular device you can change "`#`" to the device you want to listen to.

```
$ mosquitto_sub -h cloud.internalpositioning.com -p 1883 \
    -u FAMILY -P XX -t 'FAMILY/location/DEVICE'
```

The corresponding response when new sensor data is to show the raw sensor data ("`sensors`") and the corresponding guesses ("`guesses`") where the first guess is always the best guess.

```
{
    "sensors": {
        "t": 1439596533831,
        "f": "testdb",
        "d": "zack",
        "l": "",
        "s": {
            "wifi": {
                "00:1a:1e:46:cd:10": -82,
                "00:1a:1e:46:cd:11": -84,
                "00:23:69:d4:47:9f": -75,
                "20:aa:4b:b8:31:c8": -82,
                "2c:b0:5d:36:e3:b8": -81,
            }
        }
    },
    "guesses": [{
        "location": "zakhome floor 2 office",
        "probability": 0.98
    }, {
        "location": "zakhome floor 1 kitchen",
        "probability": 0.01
    }, {
        "location": "zakhome floor 2 bedroom",
        "probability": 0
    }]
}
```



