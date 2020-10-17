# Trabajo Practico n7
## Integracion Continua
### Poniendo en funcionamiento a Jenkins
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-7/imagenes/1.png)
### Probando Jenkins
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-7/imagenes/2.png)
### Instalando Plugins
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-7/imagenes/3.png)
### Configurando Maven en Jenkins
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-7/imagenes/4.png)
### Realizando Hello-word en Jenkins
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-7/imagenes/5.png)
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-7/imagenes/6.png)
#### Analizando la salida de consola
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-7/imagenes/7.png)

**Jenkins** nos da la posibilidad de configurarlo a travez de codigo para lograr los mejores resultados posibles en base de la automatizacion. La salida que estamos apreciando es las respuestas que nos da de la ejecucion de el codigo que pusimos anteriormente.

En el podemos observar:
* `agent any`: que cualquier puede ejecutar el pipeline y se refiere al ambiente, podems especificar un docker con la version especifica que debe utilizar. (Es una declaracion obligatoria)
* `stages`: Es la separacion logica de los steps. Los mas comunes son utilizar Compilar, Probar e Instalar
* `stage('name_stage')`: Es el nombre que le asignamos a este estado.
* `steps`: son las subdiviciviciones de cada uno de los pasos secuanciales que debe realizar.
* `echo ...`: ejecuta comandos de linux para mostrar esa salida por consola.

### Realizando caso con Maven + Git
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-7/imagenes/8.png)

Realizamos el Hola Mundo de Maven con git, el script para ejecutarlo es el siguiente:
```
pipeline {
    agent any
    tools {
        maven "M3"
    }
    stages {
        stage('Build') {
            steps {
                git 'https://github.com/example/without-test.git'
                sh "mvn -Dmaven.test.failure.ignore=true clean package"
            }
            post {
                success {
                    junit '**/target/surefire-reports/TEST-*.xml'
                    archiveArtifacts 'target/*.jar'
                }
            }
        }
    }
}
```
* `tools`: Instalacion de las herramientas.
* `maven 'M3'`: instalar maven del path M3.
* `git`: traer codigo de un repositorio.
* `sh`: ejecutar comandos de bash
* `post`: define acciones necesaria que deben de ejecutarse si se ha podido finalizar el paso anterior correctamente.

![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-7/imagenes/9.png)
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-7/imagenes/10.png)
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-7/imagenes/11.png)
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-7/imagenes/12.png)
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-7/imagenes/13.png)

### Utilizando el proyecto del TP anterior.
**Github del TP6**(https://github.com/SantiagoMerlo/spring-boot-is3)

* Para este proyecto utilizamos el jenkinsfile anterior con una par de cambios. Las salidas se pueden ver a continuacion. El unico problema que se epresento es que existe un test fallido.

![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-7/imagenes/14.png)
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-7/imagenes/15.png)


### Utilianzaod Docker
#### Publicamos en dockerhub
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-7/imagenes/16.png)
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-7/imagenes/17.png)

#### Construyendo jenkinsfile
* Una vez configuradas las credenciales:
* utilizamos el siguiente codigo:
```
// Docker file ejercicio 7 tp7
pipeline {
    environment {
        registry = "santiagomerlo/test-spring-boot"
        registryCredential = 'dockerhub/sam'
        dockerImage = ''
    }
    agent any
    stages {
        stage('Cloning our Git') {
            steps {
                git 'https://github.com/SantiagoMerlo/spring-boot-is3.git'
            }
        }
        stage('Building our image') {
            steps {
                script {
                    dockerImage = docker.build registry + ":$BUILD_NUMBER"
                }
            }
        }
        stage('Deploy our image') {
            steps {
                script {
                    docker.withRegistry( '', registryCredential ) {
                        dockerImage.push()
                    }
                }
            }
        }
        stage('Cleaning up') {
            steps { 
                sh "docker rmi $registry:$BUILD_NUMBER" 
            }
        } 
    }
}
```
**Probamos que funciones correctamente**
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-7/imagenes/18.png)
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-7/imagenes/19.png)

