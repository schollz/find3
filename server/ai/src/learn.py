#!/usr/bin/python3

import json
import csv
from random import shuffle
import warnings
import pickle
import gzip
import operator
import time
import logging
import math

# create logger with 'spam_application'
logger = logging.getLogger('learn')
logger.setLevel(logging.DEBUG)
fh = logging.FileHandler('learn.log')
fh.setLevel(logging.DEBUG)
ch = logging.StreamHandler()
ch.setLevel(logging.DEBUG)
formatter = logging.Formatter(
    '%(asctime)s - [%(name)s/%(funcName)s] - %(levelname)s - %(message)s')
fh.setFormatter(formatter)
ch.setFormatter(formatter)
logger.addHandler(fh)
logger.addHandler(ch)

import numpy
from sklearn.ensemble import RandomForestClassifier
from sklearn.feature_extraction import DictVectorizer
from sklearn.pipeline import make_pipeline
from sklearn.neural_network import MLPClassifier
from sklearn.neighbors import KNeighborsClassifier
from sklearn.svm import SVC
from sklearn.gaussian_process import GaussianProcessClassifier
from sklearn.gaussian_process.kernels import RBF
from sklearn.tree import DecisionTreeClassifier
from sklearn.ensemble import RandomForestClassifier, AdaBoostClassifier
from sklearn.naive_bayes import GaussianNB
from sklearn.discriminant_analysis import QuadraticDiscriminantAnalysis
from sklearn import cluster, mixture
from sklearn.neighbors import kneighbors_graph
from naive_bayes import ExtendedNaiveBayes
from naive_bayes2 import ExtendedNaiveBayes2


class AI(object):

    def __init__(self, family, path_to_data):
        self.logger = logging.getLogger('learn.AI')
        self.naming = {'from': {}, 'to': {}}
        self.family = family
        self.path_to_data = path_to_data

    def classify(self, sensor_data):
        header = self.header[1:]
        csv_data = numpy.zeros(len(header))
        for sensorType in sensor_data['s']:
            for sensor in sensor_data['s'][sensorType]:
                sensorName = sensorType + "-" + sensor
                if sensorName in header:
                    csv_data[header.index(sensorName)] = sensor_data[
                        's'][sensorType][sensor]
        return self.do_classification(header, csv_data)

    def do_classification(self, header, csv_data):
        """
        header = ['wifi-a', 'wifi-b']
        csv_data = [-67 0]
        """

        t = time.time()
        payload = {'location_names': self.naming['to'], 'predictions': []}
        for name in self.algorithms:
            try:
                prediction = self.algorithms[
                    name].predict_proba(csv_data.reshape(1, -1))
            except:
                continue
            predict = {}
            for i, pred in enumerate(prediction[0]):
                predict[i] = pred
            predict_payload = {'name': name,
                               'locations': [], 'probabilities': []}
            badValue = False
            for tup in sorted(predict.items(), key=operator.itemgetter(1), reverse=True):
                predict_payload['locations'].append(str(tup[0]))
                predict_payload['probabilities'].append(
                    round(float(tup[1]), 2))
                if math.isnan(tup[1]):
                    badValue = True
                    break
            if badValue:
                continue
            payload['predictions'].append(predict_payload)

        # try:
        #     t2 = time.time()
        #     name = "Extended Naive Bayes"
        #     clf = ExtendedNaiveBayes(self.family,path_to_data=self.path_to_data)
        #     predictions = clf.predict_proba(header,csv_data)
        #     predict_payload = {'name': name,'locations': [], 'probabilities': []}
        #     for tup in predictions:
        #         predict_payload['locations'].append(str(self.naming['from'][tup[0]]))
        #         predict_payload['probabilities'].append(round(tup[1],2))
        #     payload['predictions'].append(predict_payload)
        #     self.logger.debug("{} {:d} ms".format(name,int(1000 * (t2 - time.time()))))
        # except Exception as e:
        #     self.logger.error(str(e))

        # try:
        #     t2 = time.time()
        #     name = "Extended Naive Bayes2"
        #     clf = ExtendedNaiveBayes2(self.family, path_to_data=self.path_to_data)
        #     predictions = clf.predict_proba(header, csv_data)
        #     predict_payload = {'name': name, 'locations': [], 'probabilities': []}
        #     for tup in predictions:
        #         predict_payload['locations'].append(
        #             str(self.naming['from'][tup[0]]))
        #         predict_payload['probabilities'].append(round(tup[1], 2))
        #     payload['predictions'].append(predict_payload)
        #     self.logger.debug("{} {:d} ms".format(
        #         name, int(1000 * (t2 - time.time()))))
        # except Exception as e:
        #     self.logger.error(str(e))

        self.logger.debug("{:d} ms".format(int(1000 * (t - time.time()))))
        return payload

    def learn(self, fname):
        t = time.time()
        # load CSV file
        self.header = []
        rows = []
        naming_num = 0
        with open(fname, 'r') as csvfile:
            reader = csv.reader(csvfile, delimiter=',')
            for i, row in enumerate(reader):
                if i == 0:
                    self.header = row
                else:
                    for j, val in enumerate(row):
                        if val == '':
                            row[j] = 0
                            continue
                        try:
                            row[j] = float(val)
                        except:
                            if val not in self.naming['from']:
                                self.naming['from'][val] = naming_num
                                self.naming['to'][naming_num] = val
                                naming_num += 1
                            row[j] = self.naming['from'][val]
                    rows.append(row)

        # first column in row is the classification, Y
        y = numpy.zeros(len(rows))
        x = numpy.zeros((len(rows), len(rows[0]) - 1))

        # shuffle it up for training
        record_range = list(range(len(rows)))
        shuffle(record_range)
        for i in record_range:
            y[i] = rows[i][0]
            x[i, :] = numpy.array(rows[i][1:])

        names = [
            "Nearest Neighbors",
            "Linear SVM",
            "RBF SVM",
            # "Gaussian Process",
            "Decision Tree",
            "Random Forest",
            "Neural Net",
            "AdaBoost",
            "Naive Bayes",
            "QDA"]
        classifiers = [
            KNeighborsClassifier(3),
            SVC(kernel="linear", C=0.025, probability=True),
            SVC(gamma=2, C=1, probability=True),
            # GaussianProcessClassifier(1.0 * RBF(1.0), warm_start=True),
            # GaussianProcess takes too long!
            DecisionTreeClassifier(max_depth=5),
            RandomForestClassifier(
                max_depth=5, n_estimators=10, max_features=1),
            MLPClassifier(alpha=1),
            AdaBoostClassifier(),
            GaussianNB(),
            QuadraticDiscriminantAnalysis()]
        self.algorithms = {}
        # split_for_learning = int(0.70 * len(y))
        for name, clf in zip(names, classifiers):
            t2 = time.time()
            self.algorithms[name] = clf
            try:
                self.algorithms[name].fit(x, y)
                # self.algorithms[name].fit(x[:split_for_learning],
                #                           y[:split_for_learning])
                # score = self.algorithms[name].score(x[split_for_learning:], y[
                #     split_for_learning:])
                # print(name, score)
            except:
                pass
            self.logger.debug("learned {}, {:d} ms".format(
                name, int(1000 * (t2 - time.time()))))

        # t2 = time.time()
        # name = "Extended Naive Bayes"
        # clf = ExtendedNaiveBayes(self.family, path_to_data=self.path_to_data)
        # try:
        #     clf.fit(fname)
        #     self.logger.debug("learned {}, {:d} ms".format(
        #         name, int(1000 * (t2 - time.time()))))
        # except Exception as e:
        #     self.logger.error(str(e))

        # t2 = time.time()
        # name = "Extended Naive Bayes2"
        # clf = ExtendedNaiveBayes2(self.family, path_to_data=self.path_to_data)
        # try:
        #     clf.fit(fname)
        #     self.logger.debug("learned {}, {:d} ms".format(
        #         name, int(1000 * (t2 - time.time()))))
        # except Exception as e:
        #     self.logger.error(str(e))

        self.logger.debug("{:d} ms".format(int(1000 * (t - time.time()))))

    def save(self, save_file):
        t = time.time()
        f = gzip.open(save_file, 'wb')
        pickle.dump(self.header, f)
        pickle.dump(self.naming, f)
        pickle.dump(self.algorithms, f)
        pickle.dump(self.family, f)
        f.close()
        self.logger.debug("{:d} ms".format(int(1000 * (t - time.time()))))

    def load(self, save_file):
        t = time.time()
        f = gzip.open(save_file, 'rb')
        self.header = pickle.load(f)
        self.naming = pickle.load(f)
        self.algorithms = pickle.load(f)
        self.family = pickle.load(f)
        f.close()
        self.logger.debug("{:d} ms".format(int(1000 * (t - time.time()))))


def do():
    ai = AI()
    ai.load()
    # ai.learn()
    params = {'quantile': .3,
              'eps': .3,
              'damping': .9,
              'preference': -200,
              'n_neighbors': 10,
              'n_clusters': 3}
    bandwidth = cluster.estimate_bandwidth(ai.x, quantile=params['quantile'])
    connectivity = kneighbors_graph(
        ai.x, n_neighbors=params['n_neighbors'], include_self=False)
    # make connectivity symmetric
    connectivity = 0.5 * (connectivity + connectivity.T)
    ms = cluster.MeanShift(bandwidth=bandwidth, bin_seeding=True)
    two_means = cluster.MiniBatchKMeans(n_clusters=params['n_clusters'])
    ward = cluster.AgglomerativeClustering(
        n_clusters=params['n_clusters'], linkage='ward',
        connectivity=connectivity)
    spectral = cluster.SpectralClustering(
        n_clusters=params['n_clusters'], eigen_solver='arpack',
        affinity="nearest_neighbors")
    dbscan = cluster.DBSCAN(eps=params['eps'])
    affinity_propagation = cluster.AffinityPropagation(
        damping=params['damping'], preference=params['preference'])
    average_linkage = cluster.AgglomerativeClustering(
        linkage="average", affinity="cityblock",
        n_clusters=params['n_clusters'], connectivity=connectivity)
    birch = cluster.Birch(n_clusters=params['n_clusters'])
    gmm = mixture.GaussianMixture(
        n_components=params['n_clusters'], covariance_type='full')
    clustering_algorithms = (
        ('MiniBatchKMeans', two_means),
        ('AffinityPropagation', affinity_propagation),
        ('MeanShift', ms),
        ('SpectralClustering', spectral),
        ('Ward', ward),
        ('AgglomerativeClustering', average_linkage),
        ('DBSCAN', dbscan),
        ('Birch', birch),
        ('GaussianMixture', gmm)
    )

    for name, algorithm in clustering_algorithms:
        with warnings.catch_warnings():
            warnings.filterwarnings(
                "ignore",
                message="the number of connected components of the " +
                "connectivity matrix is [0-9]{1,2}" +
                " > 1. Completing it to avoid stopping the tree early.",
                category=UserWarning)
            warnings.filterwarnings(
                "ignore",
                message="Graph is not fully connected, spectral embedding" +
                " may not work as expected.",
                category=UserWarning)
            try:
                algorithm.fit(ai.x)
            except:
                continue

        if hasattr(algorithm, 'labels_'):
            y_pred = algorithm.labels_.astype(numpy.int)
        else:
            y_pred = algorithm.predict(ai.x)
        if max(y_pred) > 3:
            continue
        known_groups = {}
        for i, group in enumerate(ai.y):
            group = int(group)
            if group not in known_groups:
                known_groups[group] = []
            known_groups[group].append(i)
        guessed_groups = {}
        for i, group in enumerate(y_pred):
            if group not in guessed_groups:
                guessed_groups[group] = []
            guessed_groups[group].append(i)
        for k in known_groups:
            for g in guessed_groups:
                print(
                    k, g, len(set(known_groups[k]).intersection(guessed_groups[g])))


# ai = AI()
# ai.learn("../testing/testdb.csv")
# ai.save("dGVzdGRi.find3.ai")
# ai.load("dGVzdGRi.find3.ai")
# a = json.load(open('../testing/testdb_single_rec.json'))
# classified = ai.classify(a)
# print(json.dumps(classified,indent=2))
