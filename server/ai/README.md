# de0gee-ai

[![travis](https://travis-ci.org/de0gee/de0gee-ai.svg?branch=master)](https://travis-ci.org/de0gee/de0gee-ai) 
[![coverage](https://img.shields.io/badge/coverage-94%25-brightgreen.svg)](https://gocover.io/github.com/de0gee/de0gee-ai)


Server for computing machine learning on the de0gee data

## Ideas

- Ask for the # of rooms. Use KNN to determine the best hyperplane seperation for this # and then try to use that to classify a RF.

## testing

```
cd testing
http localhost:5000/classify < testdb_single_rec.json
http --json POST localhost:5000/learn family=testdb csv_file=../testing/testdb.csv
```