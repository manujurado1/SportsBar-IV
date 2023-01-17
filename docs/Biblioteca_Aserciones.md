# Biblioteca aserciones

## Criterios

- Mejores prácticas para Go (Bibliotecas referenciadas por la página oficial de Go).
- Biblioteca amplia con múltiples opciones.
- Biblioteca que aporte legibilidad a los tests.

## Opciones

Tras hacer una búsqueda de las opciones que cumplan los criterios se plantean las siguientes opciones, las cuales todas estan referenciadas por Go.

- [BE](https://github.com/carlmjohnson/be?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego): Es un paquete minimalista de aserciones. Este paquete tiene a su favor que es muy simple y fácil de usar, lo que tambien juega en su contra ya que al ser tan minimalista el abanico de opciones que ofrece es escaso. Se echa en falta funciones que permitan trabajar con errores.

- [Testify](https://github.com/stretchr/testify): Es el paquete más completo de los 3 ya que contiene paquetes que permiten realizar aserciones de una forma sencilla y con una variedad de funciones mucho más amplia que la competencia. Además, incluye paquetes para mockear información y crear suite de tests.

- [Assert](https://github.com/go-playground/assert?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego): Está inspirada en testify pero con funciones reducidas. Al igual que la primera opción, no encontramos funciones para trabajar con los errores.


## Elección

Debido a que es la opción más completa y la que más documentación tiene al respecto, se optará por usar Testify como bibioteca de aserciones. Se puede acceder a la guía de instalación [en el siguiente enlace](https://pkg.go.dev/github.com/stretchr/testify#section-readme)