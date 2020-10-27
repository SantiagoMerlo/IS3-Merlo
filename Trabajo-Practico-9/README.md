# Trabajo Practico 9
## Test Unitarios con Javascrip
Para este trabajo Utilizamos un [fork](https://github.com/SantiagoMerlo/js-unit-testing-examples) de este [repositorio](https://github.com/MarcL/js-unit-testing-examples)

## Jenkinks
### Antes
* [tutorial](https://medium.com/@gustavo.guss/jenkins-starting-with-pipeline-doing-a-node-js-test-72c6057b67d4)
* Necesitamos instalar el plugin de nodejs
* Instalarlo globalmente con el nombre de `node`
```
pipeline {

    agent any

    tools {nodejs "node"}

    stages {
        stage('Cloning Git') {
            steps {
                git 'https://github.com/SantiagoMerlo/js-unit-testing-examples'
            }
        }

        stage('Install dependencies') {
            steps {
                sh 'npm install'
            }
        }

        stage('Test') {
            steps {
                sh 'npm test'
            }
        }
  }
}
```
### Salida:
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-9/images/1.png)
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-9/images/2.png)


