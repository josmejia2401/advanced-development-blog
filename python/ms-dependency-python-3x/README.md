# MS-DEPENDENCY
Ejemplo del uso de dependencias en Python 3.x
# Nuevas caracteristicas
  - Ninguna
# Construcción
El servicio fue construido en Python 3.7.
# Dependencias
Para instalar las dependencias:
```
pip3 install -r requirements.txt 
```
# Ejecución
Para ejecutar al API sincrona:
```
python3 app.py
```
Para ejecutar al API asincrona:
```
nohup python3 -u ./app.py > ./output.log &
```
# Métodos permitidos
- GET
- OPTIONS
# RECURSOS
| ID | DESCRIPCIÓN | RECURSO | MÉTODO |
| ------ | ------ | ------ |------ |
|1.|Obtener información del empleado|[/ms-inject/v1/get][PlDb]|GET|
|1.|Obtener información del empleado|[/ms-inject/v1/get/dependency][PlDb]|GET|
# DETALLE TECNICO

Tecnologías usadas:

* [Python3] - Python3
* [Flask] - Flask
* [FlaskInjector] - FlaskInjector
* [Inject] - Inject