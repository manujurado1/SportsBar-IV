# Elección imagen base

## Criterios

Para la elección de la imagen base se seguirán las mejores prácticas, lo que consiste en:

- Usar una versión estable lo más actualizada posible.
- Basada en Linux.
- Priorizar la imagen de menor peso posible que satisfaga nuestras necesidades.

## Opciones 

Tras ver las imagenes oficiales que nos proporciona DockerHub con la versión de Go 1.20 y tras descartar ciertas como WindowsServer y NanoServer debido a que están basadas en windows tenemos las siguientes opciones:

- Alpine : Tamaño comprimido = 100 MB , Tamaño de la imagen una vez pulleada = 253 MB.
- Buster : Tamaño comprimido = 282 MB , Tamaño de la imagen una vez pulleada = 720 MB.
- Bullseye : Tamaño comprimido = 300 MB , Tamaño de la imagen una vez pulleada = 777 MB.

Pasando a las no oficiales, encontramos las siguientes opciones (descartando las que apenas tienen descargas o llevan mucho tiempo sin ser actualizadas):

- CircleCi : Su tamaño comprimido es 1GB, por lo que ni se acerca a las imágenes oficiales. Su última versión es para Go 1.17, lo que lo hace tambien estar peor posicionado que las oficiales elegidas.
- Okteto : Su tamaño comprimido es de 400MB, que aunque mejora al de CircleCi, está aún por encima de las imágenes oficiales. Su última versión es para Go 1.18, lo que lo hace tambien estar peor posicionado que las oficiales elegidas.
- Bitnami (by VMware): Tamaño comprimido de 289MB y última versión 1.19
- Antrea: También pertenece a VMware y presenta los mismos problemas que Bitnami encuanto a versión y peso.
- CorpusOps: Encontramos una versión con 1.19 de Alpine la cual mejora a algunas de las imágenes oficiales en cuanto a peso (171MB), pero es más pesada que la imagen de alpine oficial, la cuál llega a la version 1.20 de Go.
- ClearLinux: Tamaño comprimido de +1GB.
- AlpineLinux: Opción interesante debido a su poco peso, aunque más pesado que la imagen de Alpine oficial.

## Elección

Debido al tamaño y al que las mejores prácticas sugieren usar esta imagen base se ha elegido la versión oficial de Golang sobre Alpine.