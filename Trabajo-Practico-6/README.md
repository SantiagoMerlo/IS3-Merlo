# Trabajo Practico n6 Construccion de imagenes Docker
## 1 - Conceptos de Docker
### Descripcion de Instrucciones:
* `FROM`: Nos permite especificar la imagen por la cual nuestro docker va a tener de base. Normalmente esta conformado por la siguiente sintaxis: `<imagen+:<tag+`
* `RUN`: El comando RUN se ejecuta cuando se está construyendo una imagen personalizada para realizar una acción, creando una capa nueva. Este comando tiene el siguiente formato `RUN comando` (preferentemente para ejecutar comandos de shell en /bin/sh), otra forma de utilizar este comando es `RUN ['COMANDOA', 'COMANDOB']` Con esta forma no se invoca a la Shell del sistema por defecto, sino que se permite ejecutar programas que la imagen contenga. Example: `RUN [“-c”, “set -o pipefail && wget –o https://misitio/”]`
* `COPY`: Permite copiar archivos y directorios y pegarlo dentro del contenedor
* `ADD`: Es lo mismo que copy
* `EXPOSE`: Establece en el puerto el cual se va a estar exponiendo el docker.
* `CMD`: Este comando se encarga de pasar valores por defecto a un contenedor. Entre estos valores se pueden pasar ejecutables. Este comando tiene tres posibles formas de pasar los parámetros: `CMD ['ejectuable', 'parametro2']`
* `ENTRYPOINT`: Este comando se ejecuta cuando se quiere ejecutar un ejecutable en el contenedor en su arranque.
## 2 - Generar imagen de docker
* Compilacion del archivo:
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-6/images/1.png)
* Compilacion del archivo dockerizado y como es de esperar, java no anda
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-6/images/2.png)
## 3 - Java
* Por cuestion de comodidad para el trabajo practico numero 7, se creo un archivo realizando esta actividad. Link: https://github.com/SantiagoMerlo/spring-boot-is3
## 4 - Python Fask
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-6/images/3.png)
### Que ocurrio en nuestro Dockerfile?
`FROM python:3.6.3` # Se trajo la imagen de la version de python 3.6.3

`ENV BIND_PORT 5000` # se establecio la varible de entorno BIND_PORT y el de REDIS
`ENV REDIS_HOST localhost`
`ENV REDIS_PORT 6379`

`COPY ./requirements.txt /requirements.txt` # se copian las librerias que se utilizaran en el docker

`RUN pip install -r /requirements.txt` # se instalan las librerias en el docker

`COPY ./app.py /app.py` # se copia el archivo app en el docker

`EXPOSE $BIND_PORT` # se utiliza la variable BIND_PORT como puerto que se expone para el docker

`CMD [ "python", "/app.py" ]` # se corre el siguiente comando dentro del contenedor python /app.py

### Que ocurrio en nuestro dockercompose?
`version: '3.6'` # version del docker

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`services:` # servicios que utilizaremos en nuestro docker

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`app:` # nombre del servicio

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`build:` # construir el docker

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`context: ./` # donde se encuentra el dockerfile

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`depends_on:` # de que archivo depende (primero se ejecuta el redis y luego este, en caso de que no funcione redis este no se levanta)

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`- redis:`

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`environment:` # declaraciones de las variables de entorno

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`- REDIS_HOST=redis`

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`ports:`

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`- "5000:5000"`

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`redis:`

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`image: redis:3.2-alpine` # imagen para la construccion del docker

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`volumes:`

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`- redis_data:/data`

`volumes:` # especificacion del volumenes

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`redis_data:`

## 5 - Imagen para aplicacion web en NodeJs
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-6/images/4.png)
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-6/images/5.png)
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-6/images/6.png)
