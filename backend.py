#!/usr/bin/python3
from flask import Flask, jsonify
app = Flask(__name__)

@app.route('/command')
def hello_world():
    return '{"command":"ls","path":"/"}'


if __name__ == "__main__":
    app.run()
