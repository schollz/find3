import operator
import hashlib
import sys
import random

import randomcolor
import numpy
from matplotlib import pyplot
from scipy.stats import gaussian_kde
from sklearn.decomposition import PCA
from sklearn.preprocessing import StandardScaler


def getcolor(s):
    random.seed(int(hashlib.sha256(s.encode('utf-8')).hexdigest(), 16) % 10**8)
    return randomcolor.RandomColor().generate()[0]

locationSensors = {}
with open('data2.csv', 'r') as f:
    for line in f:
        line = line.strip()
        if len(line) == 0:
            continue
        location, sensors = line.split(',', 1)
        if location not in locationSensors:
            locationSensors[location] = {}
        sensorData = {}
        for _, s in enumerate(sensors[1:-1].split(",")):
            sensorID, value = s.split(":")
            sensorID = sensorID[2:-2]
            value = int(value)
            if sensorID not in locationSensors[location]:
                locationSensors[location][sensorID] = []
            locationSensors[location][sensorID].append(value)
            locationSensors[location]['mm'] = [-1,-2]
            locationSensors[location]['mn'] = [-2,-3]
            locationSensors[location]['am'] = [-50,-50,-52]


# find largest variance
sensorIndex = []
locationIndex = []
for location in locationSensors:
    locationIndex.append(location)
    for sensorID in locationSensors[location]:
        if sensorID not in sensorIndex:
            sensorIndex.append(sensorID)
num_locations = len(locationIndex)
num_sensors = len(sensorIndex)
print(num_locations,num_sensors)
X = numpy.zeros([len(sensorIndex),len(locationSensors)])

for i,location in enumerate(locationIndex):
    for j,sensorID in enumerate(sensorIndex):
        if sensorID not in locationSensors[location]:
            continue
        X[j,i] = numpy.median((locationSensors[location][sensorID]))


varianceOfSensorID = {}
for i,row in enumerate(X):
    data = []
    for v in row:
        if v == 0:
            continue
        data.append(v)
    varianceOfSensorID[sensorIndex[i]] = numpy.var(data)

print(varianceOfSensorID)

# collect sensor ids that are most meaningful
sensorIDs = []
for i, data in enumerate(
        sorted(varianceOfSensorID.items(), key=operator.itemgetter(1),reverse=True)):
    if data[1] == 0:
        continue
    sensorIDs.append(data[0])
    if len(sensorIDs) == 10:
        break


bins = numpy.linspace(-100, 0, 100)
for location in locationSensors:
    pyplot.figure(figsize=(10,4))

    for sensorID in sensorIDs:
        if sensorID not in locationSensors[location]:
            continue
        print(location,sensorID,locationSensors[location][sensorID])
        try:
            density = gaussian_kde(locationSensors[location][sensorID])
        except:
            continue
        density.covariance_factor = lambda : .5
        density._compute_covariance()
        pyplot.fill(bins,density(bins),alpha=0.2,label=sensorID,facecolor=getcolor(sensorID))
        # pyplot.hist(
        #     locationSensors[location][sensorID],
        #     bins,
        #     alpha=0.5,
        #     label=sensorID)
        if i == 10:
            break
    pyplot.title(location)
    pyplot.legend(loc='upper right')
    pyplot.savefig(location + ".png")
    pyplot.close()
