# Biblioteca aserciones

## Criterios

- Mejores prácticas para Go (Bibliotecas referenciadas como paquetes disponibles por la página oficial de [Go](https://pkg.go.dev/)).
- Biblioteca que permita testear los errores.
- Freshness.
- Valoración en [Go Report Card](https://goreportcard.com/).

## Opciones

Tras hacer una búsqueda de las diferentes opciones de bibliotecas de aserciones para Go, nos quedamos con estas 3 opciones:

- [BE](https://github.com/carlmjohnson/be): Herramienta con [su propia página](https://pkg.go.dev/github.com/carlmjohnson/be) en la lista de paquetes disponibles para Go. Tiene una puntuación de [A+](https://goreportcard.com/report/github.com/carlmjohnson/be) en Go Report Card. Debido a ser un paquete muy simplificado considero que no tiene las funciones suficientes para poder testear correctamente los errores que se puedan producir.

- [Testify](https://github.com/stretchr/testify): Esta herramienta también tiene [su propia página](https://pkg.go.dev/github.com/stretchr/testify) en la lista de paquetes disponibles, cosa obvia ya que es la biblioteca de aserciones más utilizada en el testing con Go debido a su amplia cantidad de funciones. Tiene una puntuación de [A+](https://goreportcard.com/report/github.com/stretchr/testify) en Go Report Card, aunque no obtiene la máxima puntuación en cuanto a complejidad, pero su puntuación es un 89/100 en ese apartado, lo que sigue siendo válido. En este caso, su amplio abanico de funciones nos aporta las herramientas necesarias para poder testear los errores correctamente.

- [Assert](https://github.com/go-playground/assert): Paquete básico de aserciones que tambien está [referenciado por Go](https://pkg.go.dev/gopkg.in/go-playground/assert.v1) y tiene una puntuación de [A+](https://goreportcard.com/report/github.com/go-playground/assert) en Go Report Card. El hecho de ser un paquete básico tiene el problema de no tener un amplio abanico de funciones y esto repercute en la limitación de poder testear los errores.


## Elección

Debido a que es la opción más completa y tiene una buena cantidad de funciones enfocada al testeo de errores, se optará por usar Testify como bibioteca de aserciones. Se puede acceder a la guía de instalación [en el siguiente enlace](https://pkg.go.dev/github.com/stretchr/testify#section-readme)