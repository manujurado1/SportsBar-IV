# Elección del Sistema de CI

## Criterios

- Sistema que tenga la posibilidad de usar GitHub Checks API
- Que permita usar el sistema de manera gratuita para proyectos públicos
- Sistemas integrados en la nube
- Sistemas que permitan trabajar con Docker

## Opciones

### GitHub Actions

- Al ser propio de GitHub usa Checks API
- Es gratuito para proyectos públicos
- No es necesario descargar nada
- Permite trabajar con Docker

### Jenkins

- Teóricamente puede usar Checks API, pero por el momento no es muy recomendado
- No se ha encontrado pricing, por lo que se ha supuesto que se puede usar de forma gratuida (No se ha indagado más debido a no cumplir algunos del resto de criterios)
- Hace falta descargar una aplicación
- Permite trabajar con Docker

### CircleCi

- Usa Checks API
- Es gratuito
- No es necesario descargar nada
- Permite trabajar con Docker

### SemaphoreCi
- Usa checks API 
- Tiene una versión gratuita pero no se puede instalar a GitHub como aplicación gratuitamente
- No es necesario descargar nada
- Permite trabajar con Docker

### CirrusCi
- Usa checks API (Instalando CirrusCi como aplicación en Github gratuitamente)
- Es gratis para repositorios públicos
- No es necesario descargar nada
- Permite trabajar con Docker

## Elección

Entre las opciones plasmadas, se ha descartado Jenkins por la necesidad de tener que descargar la aplicación, CircleCi para evitar la opción más básica y SemaphoreCi debido a que no he llegado a saber a ciencia cierta si usa GitHub Checks API y si se puede usar de forma gratuita.

Debido a esto, se usarán las propias GitHub Actions y CirrusCi como sistemas de CI.
