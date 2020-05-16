# MS-EMPLOYEE
Información de un empleado
# Nuevas caracteristicas
  - Ninguna
# Construcción
El servicio fue construido en Python 3.7.
# Entorno virtual
Este proyecto usa dependencias. En el caso se que se necesite usar env:
```
source ms_employee/bin/activate
```
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
nohup python -u ./app.py > ./output.log &
```
# Métodos permitidos
- GET
- POST
- PUT
- DELETE
- OPTIONS
# RECURSOS
| ID | DESCRIPCIÓN | RECURSO | MÉTODO |
| ------ | ------ | ------ |------ |
|1.|Obtener información del empleado|[/ms-employee/v1/get/<id>][PlDb]|GET|
|2.|Crear un empleado|[/ms-employee/v1/create][PlDb]|POST|
|3.|Actualiza la información de un empleado|[/ms-employee/v1/update/<id>][PlDb]|PUT|
|4.|Elimina un empleado|[/ms-employee/v1/delete/<id>][PlDb]|DELETE|

```sh
1. ENTRADA: {"id" : "number"}
1. SALIDA: {}
```
> Basura
# DETALLE TECNICO

Tecnologías usadas:

* [Python3] - Python3
* [MongoDB] - MongoDB
* [https://dillinger.io/] - https://dillinger.io/