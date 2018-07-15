# API

## Introduction

Use the following API calls to make your own front-end, or get specific data that you want from the FIND3 data.

In all of the following examples **FAMILY** refers to your specific family and **DEVICE** refers to a device. All of the endpoints are relative to the main server.

At this stage the front-end is very minimal. The only front-end available right now is to show the location of a single device in realtime. Just browse to `https://cloud.internalpositioning.com/view/location/FAMILY/DEVICE`. If you want some sort of information from FIND, this API is the best place to get it.


## General

> ### Ping server  {#ping}
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

&nbsp;


> ### Current time {#time}
> 
> This is useful for seeing if the server is up.
> 
> **Request**
```
GET /now
```
> 
> **Response**
> 
> This route simply returns the current UTC epoch time in milliseconds.
>
```
1522116478604
```
>


&nbsp;


> ### MQTT setup {#mqtt}
> 
> This is the command to setup MQTT on FIND3 for your family. For more information see [the MQTT document](/doc/mqtt.md)
> 
> **Request**
```
GET /api/v1/mqtt/FAMILY
```
> 
> **Response**
> 
> This route returns the passphrase (`XXX`) that you can use to bind to MQTT.
>
```
Added 'FAMILY' for mqtt. Your passphrase is 'XXX'
```
>



&nbsp;


> ### Dump database {#dump-database}
> 
> This will return the SQL data for the database which can be used to backup the current state of the entire database.
> 
> **Request**
```
GET /api/v1/database/FAMILY
```
> 
> **Response**
> 
> If successful it returns the SQL for the current database. It will return an error message if unsuccesful.
>
```
BEGIN TRANSACTION;
CREATE TABLE devices (id TEXT PRIMARY KEY, name TEXT);
...
```
>

&nbsp;



> ### Delete all data  {#delete}
> 
> **Request**
```
DELETE /api/v1/database/FAMILY
```
> 
> The FAMILY is the name of your family used for your recordings. Making this request will delete all your data, and it is not recoverable.
> 
> **Response**
> 
```
{
    "message": "deleted FAMILY",
    "success": true
}
```
>

&nbsp;



> ### Delete location  {#delete-location}
> 
> **Request**
```
DELETE /api/v1/location/FAMILY/LOCATION
```
> 
> The FAMILY is the name of your family used for your recordings. Making this request will delete all your data learned for LOCATION, and it is not recoverable.
> 
> **Response**
> 
```
{
    "message": "deleted location 'LOCATION' for FAMILY",
    "success": true
}
```
>


## General scanning

> ### Post sensor data  {#sensor}
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
   },
   "gps":{
       "lat":12.1,
       "lon":10.1,
       "alt":54
   }
}
```
> 
> When posting you must include a JSON body that specifies the family name ("`f`") and the device name ("`d`").
>
> You can include current timestamp specified as the Epoch time in milliseconds at UTC ("`t`"), but this is optional. If it is not included, the server will assign the current time when it is received.
> 
> The sensor data ("`s`") is a map where the keys are the type of the data. You can insert *any* type of data, but `wifi` and `bluetooth` are most common. These types of data are keys to a map of all the devices and their signals associated with that signal type.
>
> **Important:** The location("`l`") is optional. If it is specified it designates that sensor data to be used for learning. If it is not specified it designates that the sensor data will be used for only tracking. 
>
> The GPS coordinates are optional. If submitted, they will be saved in a database with the location (if provided) and the sensor data. 
> 
> Also, optionally you can add the query parameter `?justsave=1` which will prevent the server from doing any classification on the incoming data.
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


## Passive scanning 

> ### Post passive sensor data  {#post-passive}
> 
> This endpoint is used for passive scanning. It will alert the server to effectively holdover these data for a certain amount of time (default 90 seconds) and then reverse the sensor data and put into the database.
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

> ### Customize passive scanning  {#passive}
> 
> This endpoint is used for customizing passive scanning. It will tell the server to filter out specified mac addresses for learning specified locations or change the window for collecting fingerprints.
> 
>   
> **Request**
```
POST /api/v1/settings/passive
```
>
> There are a few parameters to specify. You must always specify "`family`". When you want to toggle learning on/off you must include "`device`", and if you include "`location`" it will automatically toggle learning. **Important note**: The `device` you include must be prefixed by the sensor type. For example, if you are using the WiFi for a phone with mac address `xx:yy:zz` then you must name the device `wifi-xx:yy:zz`. If you are using the bluetooth of a device, similarly it must be `bluetooth-xx:yy:zz`.
>
> The "`minimum_passive`" will create a threshold that will then only accept fingerprints that are collected with at least that many scanners. So if you have three scanning computers and you want to make sure that any device gets data from all three scanners, you can set it to 3.
>
> The "`window`" specifies the amount of time to wait before merging the collected sensor data from different scanning computers. Make sure to set this to twice the scan time. For instance, if you are having several computers scanning at 40 second intervals (the default), then to make sure that all scanning computers submit their data to the server in a single window you must specify a window of 90 seconds (default).
>
```
{
    "family":"FAMILY",
    "device":"DEVICE",
    "location":"LOCATION"
    "minimum_passive": -1,
    "window": 90,
}
```
> 
> **Important:** Learning for **DEVICE** is turned **on** if the location ("`location`") is specified. Learning for **DEVICE** is turned **off** is the location ("`location`") is empty.
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

## Calibration and analysis

> ### Calibrate machine learning algorithms  {#calibration}
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

&nbsp;

> ### Get analysis of calibration {#analysis}
> 
> This endpoint lists a lot of analysis that can give you an idea of how well the calibration did. It returns the `accuracy_breakdown` which is the location-specific correct guess percentage for the testing training set (a sequested 30% of original data not used for learning). 
> 
> The `confusion_metrics` have a lot of metrics determined from a [Confusion Matrix](https://en.wikipedia.org/wiki/Confusion_matrix) from the test data. It is organized by machine learning algorithm. The one that is of use is the `informedness` which is used to determine the end probability for selecting a location guess.
>
> **Request**
```
GET /api/v1/efficacy/FAMILY
```
>
> **Response**
> 
```
{  
   "efficacy":{  
      "accuracy_breakdown":{  
         "bathroom":0.7,
         "bedroom":0.8717948717948718,
      },
      "confusion_metrics":{  
         "AdaBoost":{  
            "bathroom":{  
               "true_positives":20,
               "false_positives":120,
               "true_negatives":1320,
               "false_negatives":120,
               "sensitivity":0.14285714285714285,
               "specificity":0.9166666666666666,
               "informedness":0.059523809523809534
            },
            "bedroom":{  
               "true_positives":36,
               "false_positives":92,
               "true_negatives":621,
               "false_negatives":45,
               "sensitivity":0.4444444444444444,
               "specificity":0.8709677419354839,
               "informedness":0.3154121863799282
            }
         }
      },
      "last_calibration_time":"2018-03-09T21:13:13.300237656-07:00"
   },
   "message":"got stats",
   "success":true
}
```
>



## Tracking and getting information {#tracking}

The following API calls are useful for getting information after the server has been taught about locations.

> ### Get a list of all devices  {#devices}
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

> ### Get the last known location for a device  {#location}
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

> ### Get a list of all location data for all devices  {#locations}
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

> ### Get simple list of devices grouped by location  {#by_location}
> **Request**
```
GET /api/v1/by_location/FAMILY
```
> This route has the following query parameters:
>
> - `history=X` will return the latest X minutes of historical data (default 120)
> - `randomized=1` will return only non-randomized macs if not 1 (default 1)
> - `active_mins=X` will return only devices that have a total active time greater than `X`  (default 0)
> - `num_scanners=X` will return only devices that have seen at least `X` of the scanners (default 0)
> - `probability=X` will return only devices who have a probability of `X` or greater (default 0.00)
>
> **Response**
> 
> Returns a list of `locations` which is a map containing the name of the location ("`location`"), and the total number of devices seen and a list of devices ("`devices`"). 
>
>Each device in the list has the name ("`device`"), the vendor determined from the Mac address ("`vendor`", if applicable), the timestamp that it was seen ("`timestamp`"), the probability it associates with that location ("`probability`"), whether or not the mac address is randomized ("`randomized`"), the number of devices is saw in its last sensor dump (`"num_scanners"`), the total time that the device as been seen in minutes (`"active_mins"`), and the time that the device was first seen ("`first_seen`").
> 
> Example:
>
```
{
    "locations": [
        {
            "devices": [
                {
                    "device": "wifi-88:d7:f6:a7:2a:48",
                    "vendor": "ASUSTek Computer Inc.",
                    "timestamp": "2018-03-10T11:29:33.063Z",
                    "probability": 0.89,
                    "randomized": false,
                    "num_scanners": 3,
                    "active_mins": 1295,
                    "first_seen": "2018-03-09T06:58:21.327Z"
                },
                {
                    "device": "wifi-40:4e:36:89:63:a5",
                    "vendor": "HTC Corporation",
                    "timestamp": "2018-03-10T11:25:34.469Z",
                    "probability": 0.83,
                    "randomized": false,
                    "num_scanners": 3,
                    "active_mins": 815,
                    "first_seen": "2018-03-09T07:16:49.934Z"
                }
            ],
            "location": "desk",
            "total": 2
        },
        {
            "devices": [
                {
                    "device": "wifi-20:df:b9:49:1c:61",
                    "vendor": "Google, Inc.",
                    "timestamp": "2018-03-10T11:29:33.043Z",
                    "probability": 0.88,
                    "randomized": false,
                    "num_scanners": 3,
                    "active_mins": 1123,
                    "first_seen": "2018-03-09T06:59:34.364Z"
                }
            ],
            "location": "kitchen",
            "total": 1
        }
    ],
    "message": "got locations",
    "success": true
}
```
>>



&nbsp; 

> ### Get simple location of a single device {#location-basic}
> **Request**
```
GET /api/v1/location_basic/FAMILY/DEVICE
```
>
> **Response**
> 
> This is a much simplified response for use with embedded evices. The `data` has the latest location (`loc`) and probability (`p`) for the specified deivce.
>
> Additionaly it specifies how long ago the device was last seen at that location, in seconds (`seen`).
>
> Example:
>
```
{  
   "data":{  
      "loc":"zakhome floor 2 office",
      "prob":0.97,
      "seen":1387
   },
   "message":"ok",
   "success":true
}
```
>>


## GPS 

> ### Post GPS coordinate information  {#post-gps}
> 
> This endpoint is used for specifying the GPS coordinates of learned locations.
> 
> **Request**
```
POST /api/v1/gps
```
```
{  
   "f":"FAMILY",
   "l":"LOCATION",
   "gps":{
       "lat":12.1,
       "lon":10.1,
       "alt":54
   }
}
```
> Requires family and a location.
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


## API requests?

If you have API requests, please [file an idea on Github](https://github.com/schollz/find3/issues/new?title=Feature:%20).

