#!/usr/bin/python3
from injector import inject

class DependencyProvider(object):

    def __init__(self):
        pass

    def get(self) -> object:
        try:
            return {"message" : "DependencyProvider"}, 200
        except Exception as e:
            print(e)
            return { "error" : "Internal error" }, 500

class InjectProvider(object):

    @inject
    def __init__(self, dependencyProvider: DependencyProvider):
        self.dependencyProvider = dependencyProvider 

    def get_dependency(self) -> object:
        try:
            return self.dependencyProvider.get()
        except Exception as e:
            print(e)
            return { "error" : "Internal error" }, 500

    def get(self) -> object:
        try:
            return {"message" : "InjectProvider"}, 200
        except Exception as e:
            print(e)
            return { "error" : "Internal error" }, 500

