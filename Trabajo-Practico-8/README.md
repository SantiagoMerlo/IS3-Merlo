# Trabajo Practico 8 - Herramientas de construccion de software en la nube
## Pros y Contras de este tipo de herramientas
### Pros
* La mayoria de las tecnologias cuentan con un periodo de prueba.
* Permite un despliege rapido cuando no se poseen muchos recursos.
* Escalabilidad.
* La mayoria de las herramientas cuentan con soportes para muchas tecnologias.
* Despreocupacion de algunas actividades relacionadas a la infraestructura.
* Baja latencia.
* Todo lo anterior permite la concentrancio de esfuerzos unicamente en el Negocio
### Contras
* La informacion, a veces informacion sensible, esta fuera de la organizacion.
* La disponibilidad no depende de la organizacion.
* Falta de control de como se estan administrando los recursos.
* El uso de estas tecnologias provoca una dependencia de como funcione la organizacion.
### Conclusiones
Personalmente por lo que estuve investigando es una muy buena alternativa cuando no se cuenta con los recursos suficientes para poder desplegar y organizar tus tecnologias. El hecho de que la mayoria cuenten con servicio gratuito hasta 5 usuarios o cierta cantidad de tiempo de compilacion o un tiempo free , etc. Permite que se pueda probarlas o empezar a trabajar en el negocio y una vez que escale bastante comenzar a pagar por el uso.
Sin embargo como super junior en el mundo de la informatica lo veo como una desventaja para organizaciones mucho mas grande o que ya revalsaron la primera etapa de generar ingresos, ya que ocurre un poco de los problemas de informatica de antes, la dependencia hacia otra organizacion. Cualquier tipo de movimiento (Cobrar mas caro los servicios, eliminar funcionalidad, desarticular herramientas) genera incertidumbre sobre el futuro de la misma organizacion. Ademas el hecho de contar con tus propios recursos y profesionales realicionados en esa area no solo provoca que funcione correctamente, sino que da la posibilidad de desarrollar nuevas tecnologias optimizadas a las necesidades de la organizacion.
## Configurando AppVeyor
**Utilizando el siguiente repositorio** https://github.com/SantiagoMerlo/spring-boot-is3
### Creamos nuevo proyecto
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-8/imagenes/1.png)
### Configuramos el Build
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-8/imagenes/2.png)
### Configuramos el entorno
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-8/imagenes/3.png)
### COnfiguramos los artefactos
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-8/imagenes/4.png)
### COnfiguramos los test
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-8/imagenes/5.png)
### Probamos que todo funcione correctamente
![alt text](https://github.com/SantiagoMerlo/IS3-Merlo/blob/master/Trabajo-Practico-8/imagenes/6.png)
### YAML
./spring-boot/appveyor.yml
```
version: 1.0.{build}
build_script:
- cmd: >-
    java -version

    docker info

    dir

    mvn clean package
test_script:
- sh: >-
    $url = "https://ci.appveyor.com/api/testresults/junit/$($env:APPVEYOR_JOB_ID)"
        $wc = New-Object 'System.Net.WebClient'
        $dirs = Get-ChildItem -Filter surefire-reports -Recurse
        ForEach ($dir in $dirs)
        {
          $files = Get-ChildItem -Path $dir.FullName -Filter TEST-*.xml
          ForEach ($file in $files)
          {
            $wc.UploadFile($url, (Resolve-Path $file.FullName))
          }
        }
artifacts:
- path: .\target\*.jar
```
### TravisCI
