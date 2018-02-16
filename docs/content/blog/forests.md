+++
title = "Random Forests now available"
description = "A third machine learning algorithm is now available for classifying locations"
tags = [
    "programming",
    "python",
    "golang"
]
date = "2016-12-03 17:56:08"
thumbnail = "https://www.internalpositioning.com/random_forest.png"
keywordlist = "random forest, machine learning, programming, open-source, app, home automation, home assistant, openhab, particle, esp8266, internal nagivation, indoor positioning"
+++


The Random Forest algorithm is much described <a href="http://www.kdnuggets.com/2016/12/random-forests-python.html">elsewhere</a>, but in short it is a very good choice for prediction problems that involves ensemble learning (aggregating a combination of several models to solve a prediction problem). It is very easy to setup and use, and can be great for classification.


I was inspired by <a href="https://kootenpv.github.io/2016-09-19-predict-where-you-are-indoors">Pascal van Kooten's whereami package</a> (which was actually inspired by FIND!), to implement Random Forests as one of the machine learning algorithms available.

Random forests have been implemented as a separate Python TCP server that uses `sklearn` routines to generate learning models and classify the fingerprints. So far it works very well and I'm impressed with the results!
