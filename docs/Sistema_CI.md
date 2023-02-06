# Elección del Sistema de CI

## Criterios

- Sistema que tenga la posibilidad de usar GitHub Checks API (Si para esto se necesita instalar una aplicación en GitHub, buscar opciones cuyas aplicaciones tengan un número de descargas aceptable para aumentar la fiabilidad)
- Que permita usar el sistema de manera gratuita para proyectos públicos
- Sistemas que permitan trabajar con Docker

## Opciones

Se ha descartado Jenkins ya que aunque es una de las opciones más populares, no tiene una versión gratuita.

### GitHub Actions

- Al ser propio de GitHub usa Checks API
- Es gratuito para proyectos públicos
- Permite trabajar con Docker

### CircleCi

- Usa Checks API
- Es gratuito
- Permite trabajar con Docker

### SemaphoreCi

- Usa checks API 
- Tiene una versión gratuita pero no se puede instalar a GitHub como aplicación gratuitamente
- Permite trabajar con Docker

### CirrusCi

- Usa checks API (Instalando CirrusCi como aplicación en Github gratuitamente)
- Es gratis para repositorios públicos
- Permite trabajar con Docker

### TravisCI
- Usa checks API (Instalando Travis como aplicación en Github)
- Se puede obtener una prueba gratis, pero es necesario introducir la tarjeta de crédito
- Permite trabajar con Docker

## Elección

Entre las opciones plasmadas, se ha descartado CircleCi para evitar la opción más básica, TravisCi por el hecho de tener que introducir la tarjeta de crédito y SemaphoreCi debido a que no he llegado a saber a ciencia cierta si usa GitHub Checks API y si se puede usar de forma gratuita.

Debido a esto, se usarán las propias GitHub Actions y CirrusCi como sistemas de CI.
