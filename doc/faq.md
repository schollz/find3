# Frequently Asked Questions

## Introduction

This FAQ should give you the best overview of FIND3, its goal and how it works. Also a lot of common questions are here, so you might find your answer here if you have a question!

If you do not find your question here, please feel free to [ask in the Slack channel](https://join.slack.com/t/find3/shared_invite/enQtMzU4MjY0NjE1NjU0LWRkY2JhNWFkM2U3Y2JhY2RlZTQ5ZTdmZTQ2M2UzMjI2MGVmMjZlOWQyZmU3MzM5YzIzOTM0YmYzYmQ3NTQzNjQ) or [send me a message](https://www.internalpositioning.com/#cta-5).

## What is FIND? {#what}

The Framework for Internal Navigation and Discovery (FIND) allows you to use your (Android) smartphone or WiFi-enabled computer (laptop or Raspberry Pi or etc.) to determine your position within your home or office. You can easily use this system in place of motion sensors as its resolution will allow your phone to distinguish whether you are in the living room, the kitchen or the bedroom, etc. The position information can then be used in a variety of ways including home automation, way-finding, or tracking!

The name, "FIND" stands for *the Framework for Internal Navigation and Discovery*. The number 3 specifically indicates the 3rd and latest version.

Using FIND, and only your smartphone or laptop, you will be able to pinpoint your position in your home or office by evaluating surrounding radio waves (Bluetooth or WiFi).  Unlike GPS on your smartphone, this system has a geo-location precision of below 10 sq ft. Also unlike GPS, the app utilizes pre-existing operations so it has no significant cost to your battery. 

This system is open-source and fully configurable but also easy to use and simple to set up. It is suited to small-scale home applications as well as large-scale business and school applications.


## What are the goals of the project? {#goals}

FIND started out as a way to replace motion sensors.

The point is to eventually incorporate FIND into home automation and lifestyle tracking. FIND can replace motion sensors to provide positional and user-specific information. Many things that you would do with a motion sensor you can do with FIND. Also, many things that you can do with GPS information you can do with FIND information. Except here you get internal positioning so you could tell apart one table from another in a cafeteria, or one bookshelf from another in a library.

## Why create something new rather than work on an existing system? {#why-new}

I started working on FIND in 2009 when there were very few, if any solutions for internal positioning. Currently, there are more (even the latest Android P is offering support using RTT now). Unfortunately, still, most solutions are not open-source, or they require external hardware (beacons, etc.), or they are expensive, or they just don’t work very well. 


## How does it work?  {#how-does-it-work}


Each time a Bluetooth/WiFi-enabled device conducts a scan of nearby devices, it will receive a unique identifier and a signal strength that correlates with the distance to that device. A compilation of these different signals can be compiled into a fingerprint which can be used to uniquely classify the current location of that device.

The access points can be anything - routers, Rokus, Raspberry Pis. They also can be anywhere - since they only need to be seen and not connected to, it will successfully use routers that are in a different building.

The basis of this system is to catalog all the fingerprints about the Wifi routers in the area (MAC addresses and signal values) and then classify them according to their location. (More information about classification is in the [Machine learning](#machine-learning) document) This is done using a Android App, or computer program, that collects the fingerprints, and then sends them on to the FIND server which can compute the location. 

Locations are determined on the FIND server using classification. Currently the server supports several different machine learning algorithms. Positioning by classification is accomplished by first learning the distributions of WiFi signals for a given location and then classifying it during tracking. Learning only takes ~10 minutes and will last almost indefinitely. The WiFi fingerprints are also the same across all devices so that learning using one device is guaranteed to work across all devices.

FIND3 is an [open-source project](https://github.com/schollz/find3) that comprises several main components:

1. Data storage server [(github.com/schollz/find3/server/main)](https://github.com/schollz/find3/tree/master/server/main)
2. Machine learning server [(github.com/schollz/find3/server/ai)](https://github.com/schollz/find3/tree/master/server/ai)
3. Command-line tool for gathering fingerprints [(schollz/find3-cli-scanner)](https://github.com/schollz/find3-cli-scanner)
4. Android app for gathering fingerprints [(schollz/find3-android-scanner)](https://github.com/schollz/find3-android-scanner)

Using these elements as building blocks, FIND3 provides the ability to track devices indoors, without motion sensors.

## What is active scanning and what is passive scanning? {#active-passive}

There are two modes of localization that you can implement with FIND:

- **Active scanning**: In *active scanning* the scanner will report the classified location of the device that is doing the scanning. *Requires installing software on the device being tracked*.

- **Passive scanning**: In *passive scanning* the scanner will report the classified location of the devices that it scans. This mode requires having a WiFi card that supports monitor mode on the scanning device. No software is needed on the device that is being tracked.

You can use *active scanning* to track your own phone or computer, as outlined in the [Tracking your phone](/doc/tracking_your_phone.md) and [Tracking your computer](/doc/tracking_your_computer.md) documents. The *passive scanning* can be used to track others that are nearby, as outline in the [Passive tracking](/doc/passive_tracking.md).


## How are locations learned? {#machine-learning}

The principal behind FIND is to collect sensor data and then *classify* that sensor data using a machine learning algorithm. Classification is done by splitting the original data into two datasets - 70% of the original data goes towards learning and 30% goes towards testing. The learning data is composed of **unique identifies** (MAC addresses usually) and **signal values** (Bluetooth, WiFi, or whatever other signals) and a label of the **location** that the signals were evaluated at.

The learning data is fed into a machine learning algorithm that can do classification with probability. There are currently 10 classifiers that are enabled. The #1-9 come `sklearn`, and the #10-11 are ones that I implemented in Python.

1. [Classifier implementing the k-nearest neighbors vote](http://scikit-learn.org/stable/modules/generated/sklearn.neighbors.KNeighborsClassifier.html)
2. [Support Vector classification (linear)](http://scikit-learn.org/stable/modules/generated/sklearn.svm.SVC.html)
3. [Support Vector classification (gamma)](http://scikit-learn.org/stable/modules/generated/sklearn.svm.SVC.html)
4. [Decision tree classifiers](http://scikit-learn.org/stable/modules/generated/sklearn.tree.DecisionTreeClassifier.html)
5. [Random forest classifiers](http://scikit-learn.org/stable/modules/generated/sklearn.ensemble.RandomForestClassifier.html)
6. [Multi-layer perceptron classifier](http://scikit-learn.org/stable/modules/generated/sklearn.neural_network.MLPClassifier.html)
7. [AdaBoost classifier](http://scikit-learn.org/stable/modules/generated/sklearn.ensemble.AdaBoostClassifier.html)
8. [Gaussian naive bayes](http://scikit-learn.org/stable/modules/generated/sklearn.naive_bayes.GaussianNB.html)
9. [Quadratic discriminant analysis](http://scikit-learn.org/stable/modules/lda_qda.html).
10. Extended naive bayes


An example of a learning algorithm is the **Extended Naive Bayes**. The basic question we want to answer is what is $P(location_y)$ for each of the $N$ possible locations that have been learned? In each location there are $M$ sensor data that is specified by $mac_x$. Assuming each device is independent, the probability of the location can be given by the product.

<div>
$$P(location_y) = \Pi_{i=1}^{M} P(location_y | mac_x)$$
</div>

In pratice though, its easier to take the log and compute the sum of logs. We just need to determine $P(location_y | mac_x)$ for each $y$ location and $mac_x$.


<div>
$$ P(location_y | mac_x)  = \frac{P(mac_x | location_y) P(location_y)}{P(mac_x | location_y) P(location_y) + P(mac_x | \neg location_y) P(\neg location_y)}$$
</div>


**Meta learning**

Random forests is arguably the best algorithm for doing this type of classification as it has the highest specificity and sensitivity of any other algorithms listed above. On one of my given datasets it scores an average of 75% accuracy across all locations (most locations have > 90% accuracy and some have less). Its fine to just use Random forests then to do the classification.

However, using 10 different machine learning algorithms will eek out a little more accuracy - in my particular example above I was able to get 87% average accuracy (all locations have >90% except one) using all 10 machine learning algorithms. How?

After the learning I use the test suite to generate another metric - the [**informedness**](https://en.wikipedia.org/wiki/Youden%27s_J_statistic) (also called the Youden's J statistic). This metric combines the true positives (tp), false negatives (fn), true negatives (tn) and false positives (fp) into a single value. This metric can be determined for each machine learning algorithm, $w$, on each location, $y$.

<div>
$$ J_{w,y} = \frac{tp}{tp + fn} + \frac{tn}{tn + fp} - 1$$
</div>

The data to compute $J_{w,y}$ comes from the cross-validation learning, where the test data is used to calculate the confusion matrix of true/false positive/negatives.

<div>
New sensor data enters the server and needs to be classifered to determine the location. The location is determined by each of the $N$ machine learning algorithms which each provide a probability for each location $y$, $P_w(y)$. This probability can then be weighted by the informedness statistic for that particular machine learning algorithm and location to return a weighted probability metric, $Q_{w,y}$ for each algorithm $w$ and location $y$.
</div>

<div>
$$Q_{w,y} = J_y  P_w(location_y)$$
</div>

<div>
After computing this for each algorithm, then a total value can be assigned to each location by summation over each algorithm $w$,
</div>

<div>
$$\sum_{w=1}^{N} Q_{w,y}(location_y) = Q_{y}$$
</div>

<div>
The server then returns an ordered set of normalized values of $Q_{y}$, where the highest value is likely the best answer.
</div>


## What is the difference between FIND and FIND3?

FIND3 is a complete re-write of the previous version of FIND ([github.com/schollz/find](https://github.com/schollz/find)). The API for sending fingerprints (`/track` and `/learn`) is backward compatible. There are several notable improvements on the previous version:

- Support for any data source - Bluetooth / WiFi / magnetic fields / etc. (previously just WiFi).
- Passive scanning built-in (previously required a [separate server](https://github.com/schollz/find-lf)).
- Support for Bluetooth scanning in scanning utility (previously just WiFi).
- Meta-learning with 10 different machine learning classifiers (previously just three).
- Client uses Websockets+React which reduces bandwidth (and coding complexity).
- Rolling compression of MAC addresses for much smaller on-disk databases (see [stringsizer](https://github.com/schollz/stringsizer))
- Data storage in SQLITE-database (previously it was BoltDB).
- Support for WiFi wardriving.

## Can I use FIND3 on my iPhone? {#iphone}

FIND does not support iPhones. 

Unfortunately, the information about the WiFi scanning has to come from the use of the [Apple80211 library](https://stackoverflow.com/questions/9684341/iphone-get-a-list-of-all-ssids-without-private-library/9684945#9684945). This is private library which means that a user would have to [jail break](https://stackoverflow.com/questions/6341547/ios-can-i-manually-associate-wifi-network-with-geographic-location/6341893#6341893) their device in order to use it. We do not want to distribute an app that would require users to jailbreak their phones, so we will have to avoid developing for iOS until Apple removes this restriction. Sorry!

## Does it use a WiFi location database?

There is no dependency on external resources like [WiFi location databases](https://en.wikipedia.org/wiki/Wi-Fi_positioning_system#Public_Wi-Fi_location_databases). However, there is a feature to automatically extract GPS coordinates from a location database. 

## What is the minimum distance that can be resolved? {#minimum}

It depends. This system harnesses the available WiFi routers and Bluetooth devices. If you have very few WiFi routers in the vicinity (i.e. <3 in 50 meters) then your resolution will suffer. Otherwise, you can typically get less than 10 square feet in location resolution.

To see whether your fingerprints are resolving well, you should do learning and then check the "Location analysis" at https://cloud.internalpositioning.com. This location analysis shows charts of access point raw data, and if two charts are very similar looking then they are likely too close together to resolve. For example, here is learning data for my desk and my kitchen - which are only about 10 feet apart and there is no wall between them.. You can see that there is little difference between the graphs:

<center>
<img src="/images/desk.png">
</center>

<center>
<img src="/images/kitchen.png">
</center>

This indicates that these two places are probably too close together. Compare this to another room that is 10 feet away but is separated by a wall - you can see that its analysis looks very different from the other two so it will resolve well.

<center>
<img src="/images/guest room.png">
</center>


## Can it pick up locations between floors?

Yes, because floors tend to attenuate the signal, so there is a noticeable difference when you are in the same position, but on different floors. 

## How long does it take to learn a location? {#training-time}

At a minimum you you should do learning in each location for about 5 minutes. After that, you can go to the dashboard to see the results of your training. To see the dashboard, goto [cloud.internalpositioning.com](https://cloud.internalpositioning.com) and sign in with your family name. At the dashboard you will the results from the last calibration, which may look like:

<center>
<img src="/images/accuracy2.png">
</center>

The **Overall** gives the average accuracy, while the accuracy for other locations are also shown. These accuracies are determined by cross-validation, so they their representation of reality will correlate with the amount of data available.

If you see that one of the locations has low accuracy, then you should do more learning. To do more learning, simply take your device and send sensor data with the given location. Then, when you are done you can re-calibrate by hitting the calibrate button.

<center>
<img src="/images/calibrate.png">
</center>

As an example, when starting learning I noticed I had low accuracy for the "Bathroom" location. I opened the dashboard to see that the accuracy was 50% (not very good).

<center>
<img src="/images/accuracy50_ss.png" width="70%">
</center>

To amend this I did more learning at that location and then I reloaded the browser with the dashboard again to see the results. After inserting about 70 more data, I had increased the accuracy to 87%.

<center>
<img src="/images/accuracy87.png" width="70%">
</center>

At this point, the accuracy had improved enough for me to move on to learn other locations. *Note:* as you learn new locations, they might end up being too similar to previous locations which could decrease the accuracy of previously learned locations. This is dependent on the number of available sensor points in the vicinity.

## Can I run FIND3 on a Raspberry Pi?

Yes. Just make sure to build the server and/or client natively from the source code. If you use Docker, make sure to [install Docker correctly](https://github.com/schollz/find3/issues/1#issuecomment-370205508) and build the Docker images yourself and note that you need at least 1GB of RAM and 1GB of Swap to build the images. 

## Why use SQLite vs BoltDB?

I really wanted to have SQL to query things. I know that [Storm, the BoltDB toolkit](https://github.com/asdine/storm) can do this, but with SQLite I knew I wouldn't have to re-write a lot of things (database dumping/loading) and I knew I wouldn't face any problems in querying.


























<script type="text/x-mathjax-config">
MathJax.Hub.Config({
  tex2jax: {inlineMath: [['$','$'], ['\\(','\\)']]}
});
</script>
<script src='https://cdnjs.cloudflare.com/ajax/libs/mathjax/2.7.2/MathJax.js?config=TeX-MML-AM_CHTML'></script>
