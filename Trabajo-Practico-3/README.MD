# Trabajo Practico Numero 3
## Desarrollo
### Parte 1 - Sistema distribuido simple
- `docker network create -d bridge mybridge`: Nos va a servir para crear una red docker con el nombre mybridge con la opcion `-d` o `--driver` que es el driver para manejar la red.
   - `docker run -d --net mybridge --name db redis:alpine`: Como vimos en el trabajo practico anterior, lo que estamos realizando es ejecutar un docker, en este caso es redis el cual es un motor de base de datos en memoria. Se va a conectar con los drivers de la red que anteriormente creamos.
   - `docker run -d --net mybridge -e REDIS_HOST=db -e REDIS_PORT=6379 -p 5000:5000 --name web alexisfr/flask-app:latest`: Levantamos la aplicacion web en el puerto 5000. Al mismo tiempo con el comando `-e` se estan estableciendo las variables de contorno gloables. Estos incluso sobre exponen los definidos en un ENV del docker
   - `docker ps`: a traves de este comando vemos un pantallazo del estado de los dockers
   - `docker network ls`: vemos el listado de redes que tenemos configurado en nuestro docker.
   - `docker network inspect mybridge`: con este comando vemos mucha mas informacion de una red en especifica.
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-3/1.png)
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-3/2.png)
### Parte 2 - Analizando el sistema
* Explicar cómo funciona el sistema.
El sistema es un contador de la cantidad de veces que se ingreso a esa pagina.
* ¿Para qué se sirven y porque están los parámetros `-e` en el segundo Docker run del ejercicio 1? 
   - Retomando lo que se escribio anteriormente es para establecer las variables globales del docker.
* ¿Qué pasa si ejecuta `docker rm -f web` y vuelve a correr `docker run -d --net mybridge -e REDIS_HOST=db -e REDIS_PORT=6379 -p 5000:5000 --name web alexisfr/flask-app:latest`?
   - Como podemos apreciar en la siguiente imagen con este comando eliminamos el contenedor. Luego ejecutando el siguiente comando volvemos a levantar uno nuevo.
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-3/3.png)
* ¿Qué occure en la página web cuando borro el contenedor de Redis con `docker rm -f db`? 
   - Como se puede apreciar en la siguiente imagen, lo que ocurre es una excepcion porque no puede conectarse al servicio ya que este es elimando.
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-3/4.png)
* Y si lo levanto nuevamente con docker run -d --net mybridge --name db redis:alpine ?
   - Se crea una nueva instancia, si vamos al programa vamos a notar que se reinicio la memoria.
* ¿Qué considera usted que haría falta para no perder la cuenta de las visitas?
   - Se podria tener un volumen en la aplicacion para desacoplar la base de datos del contador de la aplicacion, de igual manera se podria hacer al reves.
### Parte 3 - Utilizando docker-compose
`docker-compose up -d`: docker-compose es un orquestador de dockers, a travez de este comando va a analizar lo que este dentro del yml para levantar y configurarlos. El comando `-d` se utiliza para realizar todo en segundo plano.
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-3/5.png)
* ¿Qué hizo Docker Compose por nosotros? Explicar con detalle?
   - A traves del yml nos esta levantando dos servicios (uno con nombre app y otro con db), a cada uno de estos especifica en `image` de donde se busca el docker (Se puede utilizar con un mismo proyecto a traves de `build`). En App configuro que depende de otro servicio, el puerto el cual lo estoy levantando y las variables de entorno. Es realmente muy comodo la configuracion de un .yml.
   - Algo para aclarar, es que el docker-compose que estoy trabajando se modifico del de la actividad, se eliminio la parte de asigancion de volumen por un error que tiraba.
`docker-compose down`: es una forma comoda de detener todos los contenedores.
### Parte 4 - Aumentando la complejidad, análisis de otro sistema distribuido.
En la imagen siguiente se puede apreciar como se levanta el archivo.
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-3/6.png)
#### Como funciona la aplicaciones:
* Aclaraciones:
    * Se puede ver mas informacion en la imagen siguiente imagen
    * `build` este comando nos va a especificar que direccorio es el cual se va a dockerizar.
    * `command` especifica el comando que se tiene que ejecutar una ves el docker haya sido levantado.
    * `depends_on` significa que este contenedor depende de otros, asi que primero levanta los que se especifica y si no ocurre ningun error, este.
    ![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-3/7.png)
* Tenemos 2 networks diferentes:
    * `example-voting-app_back-tier`
    * `example-voting-app_front-tier`
* Tenemos **5** servicios diferentes:
    * `vote`: Escrita en python, visible en el puerto `5000` de nuestra maquina para votar con un front-end. Este esta conectado a las dos redes y finalmente se monta dentro del directorio `app` dentro del contenedor.
    * `result`: Escrita en JS, visible en el pueto `5001` de nuestra maquina para ver el front-end y puerto `5828` donde se encontraria la logica del backend. Esta conectado a las dos redes y finalmente se monta dentro del directorio `app` dentro del contenedor.
    * `worker`: Escrita en java, el directorio que busca para levantar el docker es `./worker`, el `context` es lo mismo que hacer un `build` directamente, con la diferencia de que nos permite pasar mas argumentos como por ejemplo un dockerfile. Este servicio depende de que se inicie primero `redis` y `db`.
    * `redis`: busca la imagen en docker.hub de redis, la descarga y la levanta. Se seta el nombre como solo redis y utiliza la red del back.
    * `db`: Busca la imagen en posgres y lo renombra como db. Se configurar las variables globales del user y la contrasenia. Se especifica donde se carga el directorio que va a contener la informacion y que va a utilizar la red del backend.

### Part 5 - Anlisis detallado
#### Revisando el codigo:
`example-voting-app\vote\app.py`:
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-3/8.png)
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-3/9.png)
* Cuando recibe un get analiza las cookies. Si no tiene, genera una automaticamente.
* Si la peticion es un POST, toma el valor de redis asociacio con el id de seccion 
* `resp` envia las opciones, el hostname y el voto del id de seccion.
`example-voting-app\worker\src\main\java\worker\Worker.java`
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-3/10.png)
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-3/11.png)
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-3/12.png)
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-3/13.png)
* Se conecta con la base de datos
* Registra el voto proveniente en la base de datos
* vendria a ser el backend que se conecta con la accion de emitir un voto
`example-voting-app\result\server.js`

![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-3/14.png)
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-3/15.png)
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-3/16.png)
* Se conecta con postgres
* Trae los datos para mostrar la votacion en tiempo real.
* el response en un objeto con la cantidad de votos de `a` y `b`
* Vendria a ser el backend que se conecta con la accion de ver los votos.



