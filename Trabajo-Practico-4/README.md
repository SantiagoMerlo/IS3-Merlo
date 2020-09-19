# Trabajo Practico n4
## 1- Instanciación del sistema
`Clonar el repositorio https://github.com/microservices-demo/microservices-demo`
`mkdir -p socks-demo`
`cd socks-demo`
`git clone https://github.com/microservices-demo/microservices-demo.git`
* Ejecutar lo siguiente
* `cd microservices-demo`
* `docker-compose -f deploy/docker-compose/docker-compose.yml up -d`
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-4/capturas/1.png)
* Una vez terminado el comando docker-compose acceder a http://localhost
* Generar un usuario
* Realizar búsquedas por tipo de media, color, etc.
* Hacer una compra - poner datos falsos de tarjeta de crédito ;)
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-4/capturas/2.png)
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-4/capturas/3.png)
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-4/capturas/4.png)
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-4/capturas/5.png)
## 2- Investigación de los componentes
* Describa los contenedores creados, indicando cuales son los puntos de ingreso del sistema

Como se puede ver en la siguiente imagen, al sistema podemos acceder a traves del puerto 80 directamente a las vistas (lo visto anteriormente) y en el puerto 8080 tenemos la base de datos es la segunda imagen. Los demas puerto visibles son los clientes que podemos acceder y los demas son APIs internas del docker.
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-4/capturas/6.png)
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-4/capturas/7.png)

- Edge-router: unifica a todas las APIs. 
- Carts, Payment, User, Catalogue, Shipping, Orders: API's del Sistema.
- Rabbiting: Es el middleware que conecta las aplicaciones.
- Carts-db, User-db, Catalogue-db=, Orders-db: Base de datos que consume cada una de las api de vistas.
- Queue-master: está en JAVA, "IPAddress": "172.20.0.5".  "Gateway": "172.20.0.1".
- Front-end: API de Front-end, le pega a las demas apis para generar las vistas.

* Clonar algunos de los repositorios con el código de las aplicaciones
* `cd socks-demo`
* `git clone https://github.com/microservices-demo/front-end.git`
* `git clone https://github.com/microservices-demo/user.git`
* `git clone https://github.com/microservices-demo/edge-router.git`
* **¿Por qué cree usted que se está utilizando repositorios separados para el código y/o la configuración del sistema? Explique puntos a favor y en contra.**

Es una perspectiva planteada a traves de microservicios, por lo tanto:

* Puntos a favor: 

Toda aplicación se compone entonces de 3 artefactos:
    * El código fuente.
    * La configuración.
    * Los scripts de despliegue.

Estos artefactos tienen un frecuencia de cambio distinta: el código fuente suele estar en constante modificación, mientras que los scripts de despliegue suelen ser mucho más estables y en algunos casos no recibir cambios por semanas o meses. La frecuencia de cambio de la configuración de la aplicación se encuentra en un punto intermedio, no está en constante cambio pero tampoco es tan estable como los scripts de despliegue.

Por otro lado, dependiendo del contexto organizacional, pueden definirse distintas políticas de acceso a estos artefactos. En algunas organizaciones es muy común que el equipo de desarrolladores no tenga permisos para acceder a los parámetros de configuración del ambiente productivo. También suele ocurrir que los scripts de despliegue de la aplicación sean administrados por personas ajenas al equipo de desarrolladores (usualmente personas del área de operaciones).
Estas dos cuestiones, frecuencia de cambio y permisos de acceso, son las que nos motivan a separar estos tres artefactos del proyecto en diferentes repositorios.

* Puntos en contra:

Separar los repositorios puede complicar a los desarrolladores y el personal de operaciones descubrir el código de alguna función si no hay pautas y los nuevos repositorios no están vinculados entre sí como submódulos. 

También pueden faltar bibliotecas / directorios genéricos donde los desarrolladores y las operaciones terminan compartiendo bibliotecas que crean nuevas dependencias inexplicables inesperadas. 
Pueden generar barreras entre desarrolladores y operaciones La operación y despliegue del código, los entornos permanentes, la configuración de los servicios, todo debe estar al alcance de los desarrolladores y con expectativas explícitas de desarrollar dicho código en coordinación con el equipo de operaciones.

* 4- **¿Cuál contenedor hace las veces de API Gateway?**
El contenedor del Front end es el API Gateway y el Edge-router es el encargado de manejar las APIS.
* 5- **Cuando ejecuto este comando**:
`curl http://localhost/customers`
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-4/capturas/8.png)
* - **¿Cuál de todos los servicios está procesando la operación?**
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-4/capturas/9.png)
La API de Front-end le pega directamente a los servicios de user.
* 6- **¿Y para los siguientes casos?**
* `curl http://localhost/catalogue`
La API de Front-end le pega directamente a los servicios de catalogue.
* `curl http://localhost/tags`
La API de Front-end le pega directamente a los servicios de tags.
* 7- **¿Como perisisten los datos los servicios?**
La contenedores que enumeramos anteriormente como base de datos se encargan  de la persistencia.
* 8- **¿Cuál es el componente encargado del procesamiento de la cola de mensajes?**
Como nombramos anteriormente el `Queue-master`
* 9- **¿Qué tipo de interfaz utilizan estos microservicios para comunicarse?**
El middleware `rabbitmq:3.6.8`