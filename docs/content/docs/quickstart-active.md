+++
title = "Quickstart for active scanning"
description = "Guide to quickstart setup of the FIND3 services"
weight = 10
draft = false
toc = true
bref = "You can position yourself in about 10 minutes using this guide"
+++

<h2>Introduction</h2>

In this tutorial you will learn how to do internal positioning with either an Android phone or a laptop/device. In *FIND3* there are two types of scanning: *active* and *passive*. In *active* scanning, the device that is being tracked is actively scanning all nearby APs and Bluetooth discoverables for their distance (via the RSSI measurement). In *passive* scanning, their are several devices that are scanning for all nearby APs and Bluetooth, and those signals are inverted so that it appears that the measurement comes from the scanned device.

For this tutorial I will only cover the *passive* scanning. For *active* scanning see the [Quickstart for passive scanning](/docs/quickstart-passive).

<h3 class="marker" id="download">Choose your scanner</h3>


The first thing to do is to download FIND3. The easisest way to do this is to use Docker. Do not use **apt-get** to install Docker, just use

```bash
$ curl -sSL https://get.docker.com | sh
```

This command will work (and has been tested) on Raspberry Pis. If you are not on a Raspberry Pi, then you can just pull the latest image using:

```bash
$ docker pull schollz/find3
```

However, if you are using a Raspberry Pi, you'll need to build the **armf** version yourself. Then you should get the latest *Dockerfile*:

```bash
$ wget https://raw.githubusercontent.com/schollz/find3/master/Dockerfile
$ docker build -t schollz/find3 .
```

That's it! Now FIND3 should be installed and read to go. To start it, make a directory to store the data, say **/home/$USER/FIND_DATA** and then start the Docker process in the background.

```bash
$ docker run -p 11883:1883 -p 8003:8003 \
	-v /home/$USER/FIND_DATA:/data \
	--name find3server -d -t schollz/find3
```

Now the server will be running on port **8003** and have an MQTT instance running on port **11883**.


<h3 class="marker"  id="testit">Run the test suite</h3>

To test that things are working you can submit some test data to the server. Download a test script which will make requests to the server:

```bash
$ wget https://raw.githubusercontent.com/schollz/find3/master/server/main/testing/learn.sh
$ chmod +x learn.sh
$ ./learn.sh
```

You have just submitted about 300 fingerprints for three different locations for the family **testdb** for the device **zack**. 

This test data had **location** associated with it, so you can use it for learning. To do the learning just do 

```bash
$ http GET localhost:8003/api/v1/calibrate/testdb
```

Now you should be able to see your location data. You can get the data from the command line doing:

```
$ http GET localhost:8003/api/v1/location/testdb/zack
```

You can also see the data, in realtime, by going to **localhost:8003/view/location/testdb/zack**. If you run the test suite again you should see the values change (albeit very quickly).