# API

## Introduction

Use the following API calls to make your own front-end, or get specific data that you want from the FIND3 data.

In all of the following examples **FAMILY** refers to your specific family and **DEVICE** refers to a device. All of the endpoints are relative to the main server.

At this stage the front-end is very minimal. The only front-end available right now is to show the location of a single device in realtime. Just browse to `https://cloud.internalpositioning.com/view/location/FAMILY/DEVICE`. If you want some sort of information from FIND, this API is the best place to get it.


## General

> ### Ping server
> 
> This is useful for seeing if the server is up.
> 
> **Request**
```
GET /ping
```
> 
> **Response**
> 
```
pong
```
>


## Posting and learning

> ### Post sensor data
> 
> **Request**
```
POST /data
```
```
{  
   "d":"DEVICE",
   "f":"FAMILY",
   "t":1520424248897,
   "l":"LOCATION",
   "s":{  
      "bluetooth":{  
         "20:25:64:b7:91:42":-72,
         "20:25:64:b8:06:38":-81,    
      },
      "wifi":{  
         "20:25:64:b7:91:40":-73,
         "70:4d:7b:11:3a:c8":-81,
         "88:d7:f6:a7:2a:4c":-39,
         "8c:0f:6f:e7:2b:78":-42,
         "8c:0f:6f:e7:2b:80":-43,
         "92:0f:6f:e7:2b:80":-43,
         "96:0f:6f:e7:2b:78":-39,
         "9e:0f:6f:e7:2b:80":-43,
         "ac:9e:17:7f:38:a4":-55,
         "dc:fe:07:79:aa:c0":-90,
         "dc:fe:07:79:aa:c3":-89
      }
   }
}
```

> 
> When posting you must include a JSON body that specifies the family name ("`f`"),  the device name ("`d`"), and the current timestamp specified as the Epoch time in milliseconds at UTC ("`t`").
> 
> The sensor data ("`s`") is a map where the keys are the type of the data. You can insert *any* type of data, but `wifi` and `bluetooth` are most common. These types of data are keys to a map of all the devices and their signals associated with that signal type.
>
> **Important:** The location("`l`") is optional. If it is specified it designates that sensor data to be used for learning. If it is not specified it designates that the sensor data will be used for only tracking. 
> 
> **Response**
> 
```
{
    "message": "posted data [need GPS]",
    "success": true
}
```
> After posting you'll recieve "`success`" boolean. If false it relies the error in the message. There is a special message for a successful request - it may include "`[need GPS]`" which is specified when the posted MAC addresses have no GPS coordinates associated with them (this is used for war-driving in the Android app but may be otherwise ignored).
>

&nbsp; 

> ### Post passive sensor data
> 
> This endpoint is used for passive scanning. It will alert the server to effectively holdover these data for a certain amount of time (default 25 seconds) and then reverse the sensor data and put into the database.
> 
> **Request**
```
POST /passive
```
> Requires same JSON body as `POST /data`.
> 
> **Response**
> 
```
{
    "message": "posted data",
    "success": true
}
```
>

&nbsp; 

> ### Toggle learning for passive scanning
> 
> This endpoint is used for passive scanning. It will tell the server to filter out specified mac addresses for learning specified locations. 
> 
> Yes, it is reusing the same endpoint for posting data, but it is set using a special `"t":1` flag to indicate toggling of learning.
>   
> **Request**
```
POST /passive
```
```
{
    "t":1,
    "f":"FAMILY",
    "d":"DEVICE",
    "l":"LOCATION"
}
```
> 
> **Important:** Learning for **DEVICE** is turned **on** if the location ("`l`") is specified. Learning for **DEVICE** is turned **off** is the location ("`l`") is empty.
>
> You can learn on multiple devices in multiple locations simultaneously. *Always make sure to turn off learning before moving a device!*
>
> **Response**
> 
> When turning on:
```
{
    "message": "set location to 'LOCATION' for FAMILY for learning with device 'DEVICE', now learning on 1 devices: map[DEVICE:LOCATION]",
    "success": true
}
```
> When turning off:
```
{
    "message": "switched to tracking for FAMILY, now learning on 0 devices: map[]",
    "success": true
}
```
> 

&nbsp; 

> ### Calibrate machine learning algorithms
> 
> This endpoint is used for calibrating and will cause the server to update all the machine learning algorithms with the latest learning data. Normally this endpoint will automatically run after aquiring ~20 fingerprints, but you can manually run it to make sure you get the most up-to-date calibration.
> 
> **Request**
```
GET /api/v1/calibrate/FAMILY
```
>
> **Response**
> 
```
{
    "message": "calibrated data",
    "success": true
}
```
>

## Tracking and getting information

The following API calls are useful for getting information after the server has been taught about locations.

> ### Get a list of all devices
> **Request**
```
GET /api/v1/devices/FAMILY
```
>
> **Response**
> 
```
{
    "devices": [
        "device1",
        "device2"
    ],
    "message": "got devices",
    "success": true
}
```
>

&nbsp; 

> ### Get the last known location for a device
> **Request** 
> 
```
GET /api/v1/location/FAMILY/DEVICE
```
>
> **Response**
> 
> JSON with several components. The `analysis` the probability of each guess and the location, along with a breakdown of the probabilities associated with each machine learning algorithm (note most algorithms omitted for brevity).
> 
> The `sensors` is the original sensor data sent to the server.
```
{
    "analysis": {
        "guesses": [
            {
                "location": "living room",
                "probability": 0.7555629615587942
            },
            {
                "location": "kitchen",
                "probability": 0.23040164675357372
            },
            {
                "location": "bathroom",
                "probability": 0.014035391687632025
            }
        ],
        "location_names": {
            "0": "guest room",
            "1": "kitchen",
            "2": "living room",
            "3": "bathroom",
            "4": "bedroom"
        },
        "predictions": [
            {
                "locations": [
                    "1",
                    "2",
                    "0",
                    "3",
                    "4"
                ],
                "name": "Nearest Neighbors",
                "probabilities": [
                    0.67,
                    0.33,
                    0,
                    0,
                    0
                ]
            },
            {
                "locations": [
                    "2",
                    "1",
                    "3",
                    "4",
                    "0"
                ],
                "name": "Extended Naive Bayes2",
                "probabilities": [
                    1,
                    0,
                    0,
                    0,
                    0
                ]
            }
        ]
    },
    "message": "got location",
    "sensors": {
        "d": "android",
        "f": "schollz",
        "l": "",
        "s": {
            "bluetooth": {},
            "wifi": {
                "20:25:64:b7:91:40": -73,
                "20:25:64:b7:91:42": -72,
                "20:25:64:b8:06:38": -81,
                "70:4d:7b:11:3a:c8": -81,
                "88:d7:f6:a7:2a:4c": -39,
                "8c:0f:6f:e7:2b:78": -42,
                "8c:0f:6f:e7:2b:80": -43,
                "92:0f:6f:e7:2b:80": -43,
                "96:0f:6f:e7:2b:78": -39,
                "9e:0f:6f:e7:2b:80": -43,
                "ac:9e:17:7f:38:a4": -55,
                "dc:fe:07:79:aa:c0": -90,
                "dc:fe:07:79:aa:c3": -89
            }
        },
        "t": 1520424248897
    },
    "success": true
}
```
>

&nbsp; 

> ### Get a list of all location data for all devices
> **Request**
```
GET /api/v1/locations/FAMILY
```
>
> **Response**
> 
> Same as previous, except it is an array of the latest location for each device in the family.
>


&nbsp; 

> ### Get simple list of devices grouped by location
> **Request**
```
GET /api/v1/by_location/FAMILY
```
>
> **Response**
> 
> Returns a map of `locations` where each key is a `location` that contains an array of devices. The devices are specified by `device` and the probability that they are in that location (`probability`). The `random_mac` specifies, if it is a mac address, whether or not it is randomized (i.e. is the 2nd LSB not set). The `timestamp` specifies the last time they were found in that location.
> 
```
{
    "locations": {
        "floor 1 kitchen": [
            {
                "device": "device2",
                "probability": 0.5,
                "random_mac": false,
                "timestamp": "2015-08-15T00:04:25.993Z"
            }
        ],
        "floor 2 office": [
            {
                "device": "device1",
                "probability": 0.75,
                "random_mac": false,
                "timestamp": "2015-08-14T23:55:33.831Z"
            },
            {
                "device": "device3",
                "probability": 0.9,
                "random_mac": false,
                "timestamp": "2015-08-14T23:55:43.831Z"
            },
        ]
    },
    "message": "got locations",
    "success": true
}
```
>>

## API requests?

If you have API requests, please file an idea on Github at https://github.com/schollz/find3/issues.

