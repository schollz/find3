import os
import time
from base64 import urlsafe_b64encode, urlsafe_b64decode
import logging


# create logger with 'spam_application'
logger = logging.getLogger('server')
logger.setLevel(logging.DEBUG)
fh = logging.FileHandler('server.log')
fh.setLevel(logging.DEBUG)
ch = logging.StreamHandler()
ch.setLevel(logging.DEBUG)
formatter = logging.Formatter(
    '%(asctime)s - [%(name)s/%(funcName)s] - %(levelname)s - %(message)s')
fh.setFormatter(formatter)
ch.setFormatter(formatter)
logger.addHandler(fh)
logger.addHandler(ch)


from flask import Flask, request, jsonify
app = Flask(__name__)


from learn import AI

cache = {}


def to_base64(family):
    return urlsafe_b64encode(family.encode('utf-8')).decode('utf-8')


@app.route('/classify', methods=['POST'])
def classify():
    t = time.time()

    payload = request.get_json()
    if 'sensor_data' not in payload:
        return jsonify({'success': False, 'message': 'must provide sensor data'})
    data_folder = '.'
    if 'data_folder' in payload:
        data_folder = payload['data_folder']
    fname = os.path.join(data_folder, to_base64(
        payload['sensor_data']['f']) + ".de0gee.ai")
    ai = AI()
    logger.debug("loading {}".format(fname))
    try:
        ai.load(fname)
    except FileNotFoundError:
        return jsonify({"success": False, "message": "could not find '{p}'".format(p=fname)})
    classified = ai.classify(payload['sensor_data'])
    logger.debug(classified)
    logger.debug("{:d} ms".format(int(1000 * (t - time.time()))))
    return jsonify({"success": True, "message": "data analyzed", 'analysis': classified})


@app.route('/learn', methods=['POST'])
def learn():
    payload = request.get_json()
    if 'family' not in payload:
        return jsonify({'success': False, 'message': 'must provide family'})
    if 'csv_file' not in payload:
        return jsonify({'success': False, 'message': 'must provide CSV file'})
    data_folder = '.'
    if 'data_folder' in payload:
        data_folder = payload['data_folder']

    ai = AI()
    try:
        ai.learn(os.path.join(data_folder, payload['csv_file']))
    except FileNotFoundError:
        return jsonify({"success": False, "message": "could not find '{p[csv_file]}'".format(p=payload)})

    ai.save(os.path.join(data_folder, to_base64(
        payload['family']) + ".de0gee.ai"))
    return jsonify({"success": True, "message": "calibrated data"})

if __name__ == "__main__":
    app.run(host='0.0.0.0')
