import os
from base64 import urlsafe_b64encode, urlsafe_b64decode

from flask import Flask, request, jsonify
app = Flask(__name__)

from learn import AI


def to_base64(family):
    return urlsafe_b64encode(family.encode('utf-8')).decode('utf-8')

@app.route('/classify', methods=['POST'])
def classify():
    payload = request.get_json()
    if 'family' not in payload:
        return jsonify({'success': False, 'message': 'must provide family'})
    data_folder = '.'
    if 'data_folder' in payload:
        data_folder = payload['data_folder']
    fname = os.path.join(data_folder, to_base64(payload['family']) + ".de0gee.ai")
    ai = AI()
    try:
        ai.load(fname)
    except FileNotFoundError:
        return jsonify({"success": False, "message": "could not find '{p}'".format(p=fname)})
    classified = ai.classify()
    return jsonify({"success": True, "message": "data classified",'data':classified})

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
    fname = os.path.join(data_folder, to_base64(payload['family']) + ".de0gee.ai")

    ai = AI()
    try:
        ai.load_data(payload['csv_file'])
    except FileNotFoundError:
        return jsonify({"success": False, "message": "could not find '{p[csv_file]}s'".format(p=payload)})
    ai.learn()
    ai.save(fname)
    return jsonify({"success": True, "message": "saved as '{p}'".format(p=fname)})

