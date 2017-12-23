# datastore

[![travis](https://travis-ci.org/de0gee/datastore.svg?branch=master)](https://travis-ci.org/de0gee/datastore) 
[![go report card](https://goreportcard.com/badge/github.com/de0gee/datastore)](https://goreportcard.com/report/github.com/de0gee/datastore) 
[![coverage](https://img.shields.io/badge/coverage-94%25-brightgreen.svg)](https://gocover.io/github.com/de0gee/datastore)
[![godocs](https://godoc.org/github.com/de0gee/datastore?status.svg)](https://godoc.org/github.com/de0gee/datastore) 

Datastorage for sensor data from de0gee apps.

## Posting data

Sensor data should be sent in the format:

```json
{
   "t":1514034330040,
   "g":"groupname",
   "u":"username",
   "a":"asdlkjf.alsdkfj.aiwejciwe234",
   "s":{
      "location":{
        "living room":1
      },
      "wifi":{
         "aa:bb:cc:dd:ee":-20,
         "ff:gg:hh:ii:jj":-80
      },
      "bluetooth":{
         "aa:00:cc:11:ee":-42,
         "ff:22:hh:33:jj":-50        
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

The initial keys are shorthand to reduce the size of the payload. The `t` is the timestamp - the current epoch time in milliseconds - which is used as a primary key, since no piece of data can be sent simultaneously. `g` is the groupname, `u` is the username, and `a` is the encrypted username and groupname. You can include just `a` and allow the server to authenticate you with the encrypted group+username (encrypted using [schollz/encryption](https://github.com/schollz/encryption)).

`s` is for the sensors. This is a map of maps for all kinds of sensor data. Here the keys are the names of the *type of sensor data*. The types of sensor data can be anything you want - they will dynamically alter the SQLITE table. Each type of sensor (e.g. "wifi", "bluetooth", etc.) refers to a map of its sensor data. This map contains the name of the sensor and its data. For example, "wifi" refers to two access points, "aa:bb:cc:dd:ee" and "ff:gg:hh:ii:jj" which have strengths -20 and -80 respectively. You can also specify information for learning, like the current location. The values should always be floats.