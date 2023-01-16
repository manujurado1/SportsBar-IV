# Biblioteca aserciones

## Criterios

- Mejores prácticas para Go (Bibliotecas referenciadas como paquetes disponibles por la página oficial de [Go](https://pkg.go.dev/)).
- Biblioteca que permita testear los errores.
- Freshness.
- Valoración en [Go Report Card](https://goreportcard.com/).

## Opciones

### Biblioteca estándar de testing para Go

- [Aquí](https://pkg.go.dev/testing) podemos ver la documentación acerca de la forma de testear nativa de Go, la cuál no incluye funciones de aserción y se deben comparar los resultados manualmente.
- Al hacer las aserciones mediante comparaciones, si permite el testeo de errores.
- Al ser el paquete de testing oficial del lenguaje, cumple el requisito de freshness.
- De la misma manera, al ser un paquete propio de Go, no tiene reporte en Go Report Card.

### BE

- [Aquí](https://pkg.go.dev/github.com/carlmjohnson/be) se puede ver la documentación sobre BE en la página de Go.
- Al ser un paquete muy minimalista, no es una biblioteca apta para el testeo de errores.
- Cumple el criterio de frescura ya que es una biblioteca nueva, estable y en continua actualización.
- Tiene una puntuación de [A+](https://goreportcard.com/report/github.com/carlmjohnson/be) en Go Report Card.

### Testify

- [Aquí](https://pkg.go.dev/github.com/stretchr/testify) se puede ver la documentación sobre Testify en la página de Go.
- Debido a su alto abanico de funciones tiene una gran parte de estas dedicada al testeo de errores.
- Es una herramienta activa y en continua actualización.
- Tiene una puntuación de [A+](https://goreportcard.com/report/github.com/stretchr/testify) en Go Report Card.

### Assert

- [Aquí](https://pkg.go.dev/gopkg.in/go-playground/assert.v1) se puede ver la documentación sobre Assert en la página de Go.
- El hecho de ser un paquete básico tiene el problema de no tener un amplio abanico de funciones y esto repercute en la limitación de poder testear los errores.
- Es una herramienta activa y estable.
- Tiene una puntuación de [A+](https://goreportcard.com/report/github.com/go-playground/assert) en Go Report Card.

### Gomega

- [Aquí](https://pkg.go.dev/github.com/onsi/gomega) se puede ver la documentación sobre Gomega en la página de Go. Aquí podemos ver que este framework está orientado a BDD, y no a TDD como el resto de opciones que se han contemplado.
- Al tener otro enfoque de testeo, y en este caso se busca testear el comportamiento de un usuario promedio y no entrar en implementación, no está enfocado para testear errores.
- Es una biblioteca actualizada y activa.
- Tiene una puntuación de [A+](https://goreportcard.com/report/github.com/onsi/gomega) en Go Report Card.

## Elección

Debido a que es la opción más completa, tiene una buena cantidad de funciones enfocada al testeo de errores y está enfocada en TDD, se optará por usar Testify como bibioteca de aserciones. Se puede acceder a la guía de instalación [en el siguiente enlace](https://pkg.go.dev/github.com/stretchr/testify#section-readme).
