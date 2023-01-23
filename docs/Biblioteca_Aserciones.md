# Biblioteca aserciones

## Criterios

- Mejores prácticas para Go (Bibliotecas referenciadas como paquetes disponibles por la página oficial de [Go](https://pkg.go.dev/)).
- Biblioteca que permita testear los errores.
- Freshness.
- Valoración en [Go Report Card](https://goreportcard.com/).

## Opciones

Tras hacer una búsqueda de las diferentes opciones de bibliotecas de aserciones para Go, nos quedamos con estas 5 opciones:

- [Biblioteca estandar de testing para Go](https://pkg.go.dev/testing). La primera opción que tenemos es realizar las aserciones de la manera estandar que Go propone. De esta manera estaríamos cumpliendo sin duda las mejores prácticas y nos asegurariamos la frescura de la biblioteca al ser la propia de Go. El problema que se encuentra es que no dispone de funciones de asercion propias. En este caso habría que hacer las aserciones manualmente y lanzar un error en el caso de que la aserción no fuera la esperada. En mi opinión es una herramienta muy válida, pero considero que fusionar el testing nativo de Go con una biblioteca de aserciones externa que nos facilite y nos de la posibilidad de hacer varios tipos de aserciones, aparte de gestionar los errores cuando un test falla, es la mejor opción para este proyecto.

- [BE](https://github.com/carlmjohnson/be): Herramienta con [su propia página](https://pkg.go.dev/github.com/carlmjohnson/be) en la lista de paquetes disponibles para Go. Tiene una puntuación de [A+](https://goreportcard.com/report/github.com/carlmjohnson/be) en Go Report Card. Debido a ser un paquete muy simplificado considero que no tiene las funciones suficientes para poder testear correctamente los errores que se puedan producir. En cuanto a frescura, al ser una biblioteca creada en 2022, no se pueden sacar conclusiones del ritmo de las actualizaciones, siendo la última hace 4 meses. Pero al ser una biblioteca recientemente nueva, se considera apta en el criterio de frescura.

- [Testify](https://github.com/stretchr/testify): Esta herramienta también tiene [su propia página](https://pkg.go.dev/github.com/stretchr/testify) en la lista de paquetes disponibles, cosa obvia ya que es la biblioteca de aserciones más utilizada en el testing con Go debido a su amplia cantidad de funciones. Tiene una puntuación de [A+](https://goreportcard.com/report/github.com/stretchr/testify) en Go Report Card, aunque no obtiene la máxima puntuación en cuanto a complejidad, pero su puntuación es un 89/100 en ese apartado, lo que sigue siendo válido. En este caso, su amplio abanico de funciones nos aporta las herramientas necesarias para poder testear los errores correctamente. Es la herramienta que más frescura aporta, siendo su última actualización hace 2 semanas y se esta actualizando continuamente.

- [Assert](https://github.com/go-playground/assert): Paquete básico de aserciones que tambien está [referenciado por Go](https://pkg.go.dev/gopkg.in/go-playground/assert.v1) y tiene una puntuación de [A+](https://goreportcard.com/report/github.com/go-playground/assert) en Go Report Card. El hecho de ser un paquete básico tiene el problema de no tener un amplio abanico de funciones y esto repercute en la limitación de poder testear los errores. En cuanto a frescura no me termina de convencer, ya que la última actividad en su GitHub fue hace 4 meses y lo actualizan cada bastante tiempo, aunque entiendo que esto se puede deber a que al ser un paquete tan simplificado, necesita menos actualizaciones.

- [Test helper de GopherCloud](https://github.com/gophercloud/gophercloud/tree/master/testhelper): Es un paquete que contiene funciones que ayudan a la hora de escribir test. Como el resto de opciones, esta [referenciado por Go](https://pkg.go.dev/github.com/gophercloud/gophercloud/testhelper). El repositorio completo de github tiene una puntuación de [A+](https://goreportcard.com/report/github.com/gophercloud/gophercloud) en Go Report Card. Se puede comprobar que los paquetes se siguen actualizando en el repositorio, pero justamente este paquete lleva +1 año sin actualizarse. Las funciones de asercion que incluye las considero escasas pero suficientes, ya que incluyen algunas destinadas al testeo de errores.


## Elección

Debido a que es la opción más completa y tiene una buena cantidad de funciones enfocada al testeo de errores, se optará por usar Testify como bibioteca de aserciones. Se puede acceder a la guía de instalación [en el siguiente enlace](https://pkg.go.dev/github.com/stretchr/testify#section-readme)