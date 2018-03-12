# Setting up the server

## Introduction

This document explains how to setup the `FIND3` server on your own setup.

## Pre-requisites

Make sure you have about 1 GB of RAM and a Linux computer (unless you are using Docker).

## Install and run 

There are two ways to run - the easy way, using Docker, or the hard way, installing the source on the computer.

### The easy way

The easiest way to make your own FIND3 instance is to use Docker. Do not use `apt-get` to install Docker, just use

```bash
$ curl -sSL https://get.docker.com | sh
```

This command will work (and has been tested) on Raspberry Pis. If you are not on a Raspberry Pi, then you can just pull the latest image using:

```bash
$ docker pull schollz/find3
```

However, if you are using a Raspberry Pi, you'll need to build the `armf` version yourself. Then you should get the latest *Dockerfile*:

```bash
$ wget https://raw.githubusercontent.com/schollz/find3/master/Dockerfile
$ docker build -t schollz/find3 .
```

That's it! Now FIND3 should be installed and read to go. To start it, make a directory to store the data, say `/home/$USER/FIND_DATA` and then start the Docker process in the background.

```bash
$ docker run -p 11883:1883 -p 8005:8003 \
	-v /home/$USER/FIND_DATA:/data \
    -e MQTT_ADMIN=ADMIN \
    -e MQTT_PASS=PASSWORD \
    -e MQTT_SERVER='localhost:1883' \
	--name find3server -d -t schollz/find3
```

Now the server will be running on port `8005` and have an MQTT instance running on port `11883`. Make sure to change `ADMIN` and `PASSWORD` to a admin user name and password. Do not change `MQTT_SERVER`, as it runs on the Docker image.

### The hard way

The hard way is to run FIND3 from the source. 

There are a couple of prerequisites before installing from source. First install Python 3.5+ and Go 1.6+. Then install a C compiler for SQLite.

```
$ sudo apt-get install g++
```

You'll also need `mosquitto` if using `MQTT`.

```
$ sudo apt-get install mosquitto-clients mosquitto
```

Then get the latest source and Go dependencies.

```
$ go get -u -v github.com/schollz/find3/...
```

Then install the Python dependencies.

```
$ cd $GOPATH/src/github.com/schollz/find3/server/ai
$ sudo python3 -m pip install -r requirements.txt
```

Now there are two pieces of the server to start. In one terminal you can run the AI server.

```
$ cd $GOPATH/src/github.com/schollz/find3/server/ai
$ make
```

In the other terminal you can run the main data storage server.

```
$ cd $GOPATH/src/github.com/schollz/find3/server/main
$ go build -v
$ ./main -port 8005 
```

## Run the test suite

To test that things are working you can submit some test data to the server. Download a test script which will make requests to the server:

```bash
$ wget https://raw.githubusercontent.com/schollz/find3/master/server/main/testing/learn.sh
$ chmod +x learn.sh
$ ./learn.sh
```

You have just submitted about 300 fingerprints for three different locations for the family `testdb` for the device `zack`.

This test data had `location` associated with it, so you can use it for learning. To do the learning just do 

```bash
$ http GET localhost:8005/api/v1/calibrate/testdb
```

Now you should be able to see your location data. You can get the data from the command line doing:

```
$ http GET localhost:8005/api/v1/location/testdb/zack
```

You can also see the data, in realtime, by going to `localhost:8005/view/location/testdb/zack`.If you run the test suite again you should see the values change (albeit very quickly).