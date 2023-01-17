# Elección imagen base

## Criterios

Para la elección de la imagen base se seguirán las mejores prácticas, lo que consiste en:

- Priorizar la imagen de menor peso posible que satisfaga nuestras necesidades. (Evitar imagenes que incluyan paquetes o funcionalidades que no se vayan a usar para el testing)
- Freshness


## Opciones 

Tras ver las [imagenes oficiales](https://hub.docker.com/_/golang) que nos proporciona DockerHub con la versión de Go 1.20 y tras descartar ciertas como WindowsServer y NanoServer debido a que están basadas en windows tenemos las siguientes opciones, las cuales cumplen el criterio de frescura ya que son imágenes oficiales y siempre están en continua actualización:

- Alpine : Tamaño comprimido = 100 MB , Tamaño de la imagen una vez pulleada = 253 MB.
- Buster : Tamaño comprimido = 282 MB , Tamaño de la imagen una vez pulleada = 720 MB.
- Bullseye : Tamaño comprimido = 300 MB , Tamaño de la imagen una vez pulleada = 777 MB.

Pasando a las no oficiales, encontramos las siguientes opciones (descartando las que apenas tienen descargas o llevan mucho tiempo sin ser actualizadas):

- [CircleCi](https://hub.docker.com/r/circleci/golang) : Su tamaño comprimido es 1GB, por lo que ni se acerca a las imágenes oficiales. Su última versión es para Go 1.17, lo que lo hace tambien estar peor posicionado que las oficiales elegidas ya que no cumple con el criterio de frescura al llevar más de un año sin actualizarse.
- [Okteto](https://hub.docker.com/r/okteto/golang) : Su tamaño comprimido es de 400MB, que aunque mejora al de CircleCi, está aún por encima de las imágenes oficiales. Su última versión es para Go 1.18, lo que lo hace tambien estar peor posicionado que las oficiales elegidas. Su ultima actualización fue hace 15 días, por lo que cumple con el criterio de frescura.
- [Bitnami (by VMware)](https://hub.docker.com/r/bitnami/golang): Tamaño comprimido de 289MB y última versión 1.19.5, la misma que las imágenes oficiales, lo que hace que se posicione muy bien en cuanto a frescura, pero es casi 3 veces más pesada que la mejor opción de las imágenes oficiales.
- [Antrea](https://hub.docker.com/r/antrea/golang): También pertenece a VMware, pero la imagen base es aún más pesada que la opción anterior y aunque a día de hoy se sigue actualizando, su versión actual está por detrás que la de Bitnami y las imágenes oficiales.
- [CorpusOps](https://hub.docker.com/r/corpusops/golang): Encontramos una imagen basada en Alpine la cual mejora a algunas de las imágenes oficiales en cuanto a peso (171MB), pero es más pesada que la imagen de alpine oficial, la cuál, pese a ambas estar en continua actualización a día de hoy, está por delante en cuanto a versión de Go.
- [ClearLinux](https://hub.docker.com/r/clearlinux/golang): Tamaño comprimido de +1GB. Cumple con el criterio de frescura ya que se actualizó hace 1 semana.
- [AlpineLinux](https://hub.docker.com/r/alpinelinux/golang): Opción interesante debido a su poco peso (193 MB),pero es más pesada que la imagen oficial de Go sobre Alpine. Cumple con el criterio de frescura ya que su ultima actualización fue hace 3 días.

## Elección

Debido al tamaño y a que las mejores prácticas sugieren usar esta imagen base se ha elegido la versión oficial de Go sobre Alpine. Se usará el tag "alpine", el cual apunta siempre a la última versión oficial de Go sobre Alpine lanzada.