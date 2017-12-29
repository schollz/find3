# de0gee-data

[![travis](https://travis-ci.org/de0gee/de0gee-data.svg?branch=master)](https://travis-ci.org/de0gee/de0gee-data) 
[![go report card](https://goreportcard.com/badge/github.com/de0gee/de0gee-data)](https://goreportcard.com/report/github.com/de0gee/de0gee-data) 
[![coverage](https://img.shields.io/badge/coverage-94%25-brightgreen.svg)](https://gocover.io/github.com/de0gee/de0gee-data)
[![godocs](https://godoc.org/github.com/de0gee/de0gee-data?status.svg)](https://godoc.org/github.com/de0gee/de0gee-data) 

<h3 class="section-head" id="intro"><a href="#intro">Introduction</a></h3>

The datastore server is a Go server that allows you to store sensor data. In production this server sits behind an authentication server that first authenticates users before allowing access. You can also run the server publicly, without authentication.

<h3 class="section-head" id="sensor"><a href="#sensor">Sensor data</a></h3>

The main element of the datastore server is the **SensorData**. **SensorData** is sent to the datastore server via JSON. The most basic JSON for **SensorData** is the following:

```json
{
    "t":1514034330040,
    "f":"fido and friends",
    "d":"fido's phone",
    "l":"dog house",
    "s":{     
    }
 }
```

The keys in this **SensorData** are shorthand (single characters) to cut down on bandwidth for sending/receiving JSON. They characters are "t" for "timestamp", "f" for "family", "d" for "device", "l" for "location" and "s" is for the "sensor readings." 

A timestamp ("t") uniquely identifies a piece of **SensorData**, as it is the UNIX epoch time *in milliseconds*, making it highly unlikely for clashes to exist. 

The family ("f") is the group in which the device belongs. A family can have many devices associated with it (your phone, your computer, your dog's collar), but each device can only be associated with one family.

The device ("d") uniquely identifies the current device in that particular family.

The location ("l") classifies the location of the current **SensorData**. This is optional, and it is used for classifying the location in preparation for machine learning.

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

In this case the sensor name is the coordinate and the value is the acceleration in that direction. So in the end, if your phone collects a lot of data you will end up sending the following **SensorData** JSON to the datastore server:

```json
{
    "t":1514034330040,
    "f":"fido and friends",
    "d":"fido's phone",
    "l":"dog house",
    "s":{
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

## API 

Starting the datastore server gives you several endpoints to insert, delete, or pull information.

<h3 class="section-head" id="post-slash"><a href="#post-slash"><code>POST / (insert data)</code></a></h3>

**Parameters**:

Requires JSON of the sensor data, e.g. 
```json
{
    "t":1514034330040,
    "f":"fido and friends",
    "d":"fido's phone",
    "l":"dog house",
    "s":{
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

**Response**:

```json
{
    "success": true,
    "message": "inserted sensor data",
}
```

<h3 class="section-head" id="get-learningdata"><a href="#get-learningdata"><code>GET /learningdata</code></a></h3>

This route will take all of the records for a particular family that have a "location" in the **SensorData** and output them to a random file `random_file.mldata` in a format that can be used for doing the machine learning.

**Parameters**:


| Name 	| Location 	| Description  	| Required 	|
|------	|----------	|--------------	|----------	|
| family|query     	|determines family| yes 	|

**Response**:

The response will give a `true` success if it succeeds, and it will also tell the name of the file generated with the machine learning data.

```json
{
    "success": true,
    "message": "random_file.mldata"
}
```

It will then generate a file `random_file.mldata` which should correspond to the following:

```
location,wifi-aa:bb:cc:dd:ee,wifi-ff:gg:hh:ii:jj:kk
dog house,-10,-20
dog house,-9,-18
...
```


### Testing


```
# Start machine learning server
cd $GOPATH/de0gee/de0gee-ai/src
export FLASK_DEBUG=1 && export FLASK_APP=server.py && flask run --debugger --port 8002

# Load machine learning data
cd $GOPATH/de0gee/de0gee-ai/testing
http --json POST localhost:8002/learn family='testdb' csv_file='../testing/testdb.csv'

# Test classification
http localhost:8002/classify < testdb_single_rec.json

# Start datastore server
cd $GOPATH/de0gee/de0gee-data
go build && ./de0gee-data

# Test getting the classification of the latest location
http --json GET localhost:8003/location family=testdb device=zack2@gmail.com
```

Supervisord file:

```
[supervisorctl]

[supervisord]

[program:de0gee-data]
directory=/home/zns/go/src/github.com/de0gee/de0gee-data
command=de0gee-data
stdout_logfile: /home/zns/go/src/github.com/de0gee/de0gee-data.std.log
stdout_logfile: /home/zns/go/src/github.com/de0gee/de0gee-data.err.log

[program:de0gee-ai]
directory=/home/zns/go/src/github.com/de0gee/de0gee-ai/src
environment = 
    FLASK_APP=server.py,
    FLASK_DEBUG=1
command=/usr/local/bin/flask run --debugger --port 8002
stdout_logfile: /home/zns/go/src/github.com/de0gee/de0gee-ai.std.log
stdout_logfile: /home/zns/go/src/github.com/de0gee/de0gee-ai.err.log
```

http://www.steves-internet-guide.com/install-mosquitto-linux/

MOSQUITTO

```
# bootstrap
cd $GOPATH/src/github.com/de0gee/de0gee-data/src/mqtt
go test 
pkill -9 mosquitto
mosquitto -c mosquitto_config/mosquitto.conf -d

# this should allow you to subscribe (change password though)
mosquitto_sub -h localhost -p 1883 -u labs -P de7r3 -t 'labs/#'

# labs should see this
mosquitto_pub -u zack -P 1234 -t 'labs/location' -m 'hello'

# labs should not see this
mosquitto_pub -u zack -P 1234 -t 'someother' -m 'hello'

# if you start labs with -t '#' it should not see anything, but admin can
```

Starting up everything

```
cd $GOPATH/src/github.com/de0gee/de0gee-data/src/mqtt && /usr/sbin/mosquitto -c mosquitto_config/mosquitto.conf -d
cd $GOPATH/src/github.com/de0gee/de0gee-data && ./de0gee-data

# for debugging
cd $GOPATH/src/github.com/de0gee/de0gee-ai/src && export FLASK_APP=server.py && export FLASK_DEBUG=1 && flask run --debugger --port 8002

# for production
cd $GOPATH/src/github.com/de0gee/de0gee-ai/src && gunicorn server:app -b 0.0.0.0:8002 -w 8
```

Ideas: use one-time-pass API keys for accessing pieces of the AI server or the datastore server directly (mainly things like websockets).