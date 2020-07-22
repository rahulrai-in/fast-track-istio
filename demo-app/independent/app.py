import base64
import copy
import json
import datetime
import flask
from flask import request, jsonify

app = flask.Flask(__name__)
app.config["DEBUG"] = False
ENCODING = 'utf-8'

with open('books.json') as booksdatafile:
    books = json.load(booksdatafile)

with open('movies.json') as moviesdatafile:
    movies = json.load(moviesdatafile)


@app.route('/', methods=['GET'])
def __home():
    return "Hello from independent service at " + datetime.datetime.now().strftime("%c")


@app.route('/books/', methods=['GET'])
@app.route('/books/<int:id>', methods=['GET'])
def __books(id=None):
    payload = copy.deepcopy(books)
    if id is None:
        for item in payload:
            img_toBase64(item)
            item["summary"] = ""
        return jsonify(payload)
    for book in payload:
        if book['id'] == id:
            img_toBase64(book)
            return jsonify(book)
    return jsonify(None)


@app.route('/movies/', methods=['GET'])
@app.route('/movies/<int:id>', methods=['GET'])
def __movies(id=None):
    payload = copy.deepcopy(movies)
    if id is None:
        for item in payload:
            img_toBase64(item)
            item["summary"] = ""
        return jsonify(payload)
    for movie in payload:
        if movie['id'] == id:
            img_toBase64(movie)
            return jsonify(movie)
    return jsonify(None)


def img_toBase64(item):
    with open(item["thumbnail"], "rb") as image_file:
        item["thumbnail"] = base64.b64encode(
            image_file.read()).decode(ENCODING)


app.run(host="0.0.0.0", port="8080")
