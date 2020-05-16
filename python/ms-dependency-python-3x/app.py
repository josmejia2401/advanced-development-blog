#!/usr/bin/python3
from flask import Flask, request
from flask_injector import FlaskInjector
from injector import inject, singleton, Binder

from api import api
from provider.injectx import InjectProvider, DependencyProvider

app = Flask(__name__)


def configure(binder: Binder) -> Binder:
    binder.bind(interface=DependencyProvider, to=DependencyProvider, scope=singleton)
    binder.bind(interface=InjectProvider, to=InjectProvider, scope=request)

@app.route('/ms-inject/v1/get', methods=["GET"])
@inject
def get(data_provider: InjectProvider):
    return api.get(data_provider)

@app.route('/ms-inject/v1/get/dependency', methods=["GET"])
@inject
def get_dependency(data_provider: InjectProvider):
    return api.get_dependency(data_provider)

@app.after_request
def after_request(response):
    response.headers.add('Access-Control-Allow-Origin', '*')
    response.headers.add('Access-Control-Allow-Headers', '*')
    response.headers.add('Access-Control-Allow-Methods', 'GET,OPTIONS')
    return response


if __name__ == '__main__':
    FlaskInjector(app=app, modules=[configure])
    app.run(port=9010, host="0.0.0.0", debug=True)
