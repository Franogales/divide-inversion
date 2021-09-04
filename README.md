# divide-inversion

## Descripcion General

Este es un programa que me pidieron para una entrevista

El paquete first contiene solamente un algorithmo para encontrar la combinacion posible de asignacion de 3 tipos de inversiones

El paquete second solo contiene una restAPI al agoritmo.

## Paquete second
El paquete second como comenté anteriormente contiene una api para la asignacion de inversiones.
La posible ruta para ejecutar el programa es un POST Request  en formato json como el siguente

"curl --location --request POST '52.27.42.57:8080/credit-asigment' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Investment": 700
}"

Donde "Investment" es la cantidad a invertir

El paquete second esta hecho con gin por usar fasthttp que es mas rapido que el paquete net/http que viene por defecto en golang. Ademas, ya cuenta con el midleware para recuperarse de los panic, tiene buen sistema de rutas y tiene buenos logs.

## Cómo usarlo
Para poder ejecutar el paquete second debes de hacer lo siguiente. 
- 1.- entrar dentro de second/cmd/ y ejecutar el comando "./server" y se iniciara el servicio
