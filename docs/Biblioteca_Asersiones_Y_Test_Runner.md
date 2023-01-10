# Biblioteca asersiones

## Criterios

- Mejores prácticas para GO
- Biblioteca amplia con múltiples opciones
- Aporten legibilidad a los tests

## Opciones

Tras hacer una búsqueda de las opciones que cumplan los criterios se plantean las siguientes opciones, las cuales todas estan referencias por Go.

- [BE](https://github.com/carlmjohnson/be?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego): Es una paquete minimalista de assersiones. A su favor es que es muy simple y fácil de usar, lo que tambien juega en su contra ya que al ser tan minimalista el abanico de opciones que ofrece es escaso.

- [Testify](https://github.com/stretchr/testify): Es el paquete más completo de las 3 ya que contiene paquetes que permiten realizar asersiones de una forma sencilla y con una variedad de funciones mucho más amplia que la competencia. Además incluye paquetes para mockear información y crear suite de tests.

- [Assert](https://github.com/go-playground/assert?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego): Está inspirada en testify pero con funciones reducidas.


## Elección

Debido a que es la opción más completa y la que más documentación tiene al respecto, se optará por usar Testify como bibioteca de asersiones. Se puede acceder a la guía de instalación [en el siguiente enlace](https://pkg.go.dev/github.com/stretchr/testify#section-readme)

# Test runner

## Criterios

- Mejores prácticas para GO
- Facilidad para lanzar y gestionar los tests

## Elección

La elección del test runner en este caso es sencilla ya que el propio Go tiene un test runner propio y no se ha encontrado ninguna alternativa real que ofrezca algo por lo que usarla en vez del test runner propio.