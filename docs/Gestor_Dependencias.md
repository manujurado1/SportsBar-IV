# Gestor Dependencias

## Criterios

- Mejores prácticas del lenguaje (GO)
- Freshness
- Comunidad activa

## Posibles opciones y elección

En el caso de los proyectos en GO, los diferentes criterios para elegir un gestor de dependencias llevan siempre a la misma opción: [Go Modules](https://go.dev/blog/using-go-modules), el propio gestor de dependencias de Go, lanzado en 2019 y que solucionaba la problemática que tenían los proyectos en Go con la gestión de dependencias.

El único gestor de dependencias que podría intentar ser una opción frente a Go Modules es [Dep](https://github.com/golang/dep), gestor de dependencias que se usaba antes de la existencia de Go Modules, pero que dejaron de mantener cuando este salió a la luz y actualmente está deprecado.

Por lo tanto, en este proyecto se ha elegido usar Go Modules como gestor de dependencias.