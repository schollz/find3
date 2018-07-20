# FIND3 Overview

## Introduction

The Framework for Internal Navigation and Discovery (FIND) allows you to use your (Android) smartphone or WiFi-enabled computer (laptop or Raspberry Pi or etc.) to determine your position within your home or office. You can easily use this system in place of motion sensors as its resolution will allow your phone to distinguish whether you are in the living room, the kitchen or the bedroom, etc. The position information can then be used in a variety of ways including home automation, way-finding, or tracking!

Simply put, FIND will allow you to replace tons of motion sensors with a single smartphone!

## How does it work?  {#how-does-it-work}


Each time a Bluetooth/WiFi-enabled device conducts a scan of nearby devices, it will receive a unique identifier and a signal strength that correlates with the distance to that device. A compilation of these different signals can be compiled into a fingerprint which can be used to uniquely classify the current location of that device.

The access points can be anything - routers, Rokus, Raspberry Pis. They also can be anywhere - since they only need to be seen and not connected to, it will successfully use routers that are in a different building.

The basis of this system is to catalog all the fingerprints about the Wifi routers in the area (MAC addresses and signal values) and then classify them according to their location. (More information about classification is in the [Machine learning](#machine-learning) document) This is done using a Android App, or computer program, that collects the fingerprints, and then sends them on to the FIND server which can compute the location. 

Locations are determined on the FIND server using classification. Currently the server supports several different machine learning algorithms. Positioning by classification is accomplished by first learning the distributions of WiFi signals for a given location and then classifying it during tracking. Learning only takes ~10 minutes and will last almost indefinitely. The WiFi fingerprints are also the same across all devices so that learning using one device is guaranteed to work across all devices.


## Framework 

FIND3 is an [open-source project](https://github.com/schollz/find3) that comprises several main components:


1. Data storage server [(github.com/schollz/find3/server/main)](https://github.com/schollz/find3/tree/master/server/main)
2. Machine learning server [(github.com/schollz/find3/server/ai)](https://github.com/schollz/find3/tree/master/server/ai)
3. Command-line tool for gathering fingerprints [(schollz/find3-cli-scanner)](https://github.com/schollz/find3-cli-scanner)
4. Android app for gathering fingerprints [(schollz/find3-android-scanner)](https://github.com/schollz/find3-android-scanner)

Using these elements as building blocks, FIND3 provides the ability to track devices indoors, without motion sensors.


## Active and Passive scanning {#active-passive}

There are two modes of localization that you can implement with FIND:

- **Active scanning**: In *active scanning* the scanner will report the classified location of the device that is doing the scanning. *Requires installing software on the device being tracked*.

- **Passive scanning**: In *passive scanning* the scanner will report the classified location of the devices that it scans. This mode requires having a WiFi card that supports monitor mode on the scanning device. No software is needed on the device that is being tracked.

You can use *active scanning* to track your own phone or computer, as outlined in the [Tracking your phone](/doc/tracking_your_phone.md) and [Tracking your computer](/doc/tracking_your_computer.md) documents. The *passive scanning* can be used to track others that are nearby, as outline in the [Passive tracking](/doc/passive_tracking.md).

## FIND3 vs FIND   {#new-version}


FIND3 is a complete re-write of the previous version of FIND ([github.com/schollz/find](https://github.com/schollz/find)). The API for sending fingerprints (`/track` and `/learn`) is backward compatible. There are several notable improvements on the previous version:

- Support for any data source - Bluetooth / WiFi / magnetic fields / etc. (previously just WiFi).
- Passive scanning built-in (previously required a [separate server](https://github.com/schollz/find-lf)).
- Support for Bluetooth scanning in scanning utility (previously just WiFi).
- Meta-learning with 10 different machine learning classifiers (previously just three).
- Client uses Websockets+React which reduces bandwidth (and coding complexity).
- Rolling compression of MAC addresses for much smaller on-disk databases (see [stringsizer](https://github.com/schollz/stringsizer))
- Data storage in SQLITE-database (previously it was BoltDB).
- Support for WiFi wardriving.



## Machine learning {#machine-learning}

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
10. [Extended naive bayes](#extended-naive-bayes2)


### Extended Naive Bayes{#extended-naive-bayes2}

*In progress*

The basic question we want to answer is what is $P(location_y)$ for each of the $N$ possible locations that have been learned? In each location there are $M$ sensor data that is specified by $mac_x$. Assuming each device is independent, the probability of the location can be given by the product.

<div>
$$P(location_y) = \Pi_{i=1}^{M} P(location_y | mac_x)$$
</div>

In pratice though, its easier to take the log and compute the sum of logs. We just need to determine $P(location_y | mac_x)$ for each $y$ location and $mac_x$.


<div>
$$ P(location_y | mac_x)  = \frac{P(mac_x | location_y) P(location_y)}{P(mac_x | location_y) P(location_y) + P(mac_x | \neg location_y) P(\neg location_y)}$$
</div>


## Meta-learning {#meta-learning}

Random forests is arguably the best algorithm for doing this type of classification as it has the highest specificity and sensitivity of any other algorithms listed above. On one of my given datasets it scores an average of 75% accuracy across all locations (most locations have > 90% accuracy and some have less). Its fine to just use Random forests then to do the classification.

However, using 10 different machine learning algorithms will eek out a little more accuracy - in my particular example above I was able to get 87% average accuracy (all locations have >90% except one) using all 10 machine learning algorithms. How?

### Calculations

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











<script type="text/x-mathjax-config">
MathJax.Hub.Config({
  tex2jax: {inlineMath: [['$','$'], ['\\(','\\)']]}
});
</script>
<script src='https://cdnjs.cloudflare.com/ajax/libs/mathjax/2.7.2/MathJax.js?config=TeX-MML-AM_CHTML'></script>
