#!/usr/bin/python3
def get_dependency(data_provider) -> object:
    try:
        return data_provider.get_dependency()
    except Exception as e:
        return {"error": str(e)}, 400

def get(data_provider) -> object:
    try:
        return data_provider.get()
    except Exception as e:
        return {"error": str(e)}, 400