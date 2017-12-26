# datastore

[![travis](https://travis-ci.org/de0gee/datastore.svg?branch=master)](https://travis-ci.org/de0gee/datastore) 
[![go report card](https://goreportcard.com/badge/github.com/de0gee/datastore)](https://goreportcard.com/report/github.com/de0gee/datastore) 
[![coverage](https://img.shields.io/badge/coverage-94%25-brightgreen.svg)](https://gocover.io/github.com/de0gee/datastore)
[![godocs](https://godoc.org/github.com/de0gee/datastore?status.svg)](https://godoc.org/github.com/de0gee/datastore) 

<h3 class="section-head" id="intro"><a href="#intro">Introduction</a></h3>

The datastore server is a Go server that allows you to store sensor data. In production this server sits behind an authentication server that first authenticates users before allowing access. You can also run the server publicly, without authentication.

<h3 class="section-head" id="sensor"><a href="#sensor">Sensor data</a></h3>

The main element of the datastore server is the **SensorData**. **SensorData** is sent to the datastore server via JSON. The most basic JSON for **SensorData** is the following:

```json
{
    "t":1514034330040,
    "f":"fido and friends",
    "d":"fido's phone",
    "s":{     
    }
 }
```

The keys in this **SensorData** are shorthand (single characters) to cut down on bandwidth for sending/receiving JSON. They characters are "t" for "timestamp", "f" for "family", "d" for "device" and "s" is for the "sensor readings." 

A timestamp ("t") uniquely identifies a piece of **SensorData**, as it is the UNIX epoch time *in milliseconds*, making it highly unlikely for clashes to exist. 

The family ("f") is the group in which the device belongs. A family can have many devices associated with it (your phone, your computer, your dog's collar), but each device can only be associated with one family.

The device ("d") uniquely identifies the current device in that particular family.

The sensor readings ("s") here is empty. The sensor readings in this most basic JSON are blank, as they are optional, although they are the most important part of the **SensorData**. Sensor readings are added to the JSON as maps of sensor data. For example, if you are taking WiFi data from access points, you would format your sensor readings as:

``` 
"wifi": {
  "aa:bb:cc:dd:ee":-20,
  "ff:gg:hh:ii:jj":-80
}
```

The first key explains the sensor type ("wifi") and the key and values inside the map explain the sensor name (a MAC address) and the value (the signal dBm). The same format is followed for *any kind of sensor*. For example, here is the sensor readings formatted for accelerometer data:

```
"accelerometer": {
  "x":-1.11,
  "y":2.111,
  "z":1.23   
}
```

In this case the sensor name is the coordinate and the value is the acceleration in that direction. The only special sensor reading is the location, which you also input as sensor data (since it is sensed by you):

```
"location": {
    "living room":1
}
```

The "location" is useful for machine learning and classification further down the road.

So in the end, if your phone collects a lot of data you will end up sending the following **SensorData** JSON to the datastore server:

```json
{
    "t":1514034330040,
    "f":"fido and friends",
    "d":"fido's phone",
    "s":{
         "location":{
             "living room":1
         },
         "wifi":{
                "aa:bb:cc:dd:ee":-20,
                "ff:gg:hh:ii:jj":-80
         },
         "bluetooth":{
                "kk:ll:mm:nn:oo":-42,
                "pp:qq:rr:ss:tt":-50        
         },
         "temperature":{
                "sensor1":12,
                "sensor2":20       
         },
         "accelerometer":{
                "x":-1.11,
                "y":2.111,
                "z":1.23   
         }      
    }
 }
```

<h3 class="section-head" id="find"><a href="#find">FIND compatibility</a></h3>

The datastore server does have [FIND compatibility](https://www.internalpositioning.com/api/#post-learn). That is, the POST /track and POST /learn routes are still available, where the payload that is sent is the same as in FIND:

```json
{
   "group":"some group",
   "username":"some user",
   "location":"some place",
   "timestamp":12309123,
   "wifi-fingerprint":[
      {
         "mac":"AA:AA:AA:AA:AA:AA",
         "rssi":-45
      },
      {
         "mac":"BB:BB:BB:BB:BB:BB",
         "rssi":-55
      }
   ]
}
```

However, FIND only supports WiFi, so you cannot use these routes for sending other kinds of data.

<h3 class="section-head" id="quick-start"><a href="#quick-start">Quick start</a></h3>

First make sure that [you have installed Go](/docs/quick-start/#install-go).