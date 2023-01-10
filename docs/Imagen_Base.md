# Elección imagen base

## Criterios

Para la elección de la imagen base se seguirán las mejores prácticas, lo que consiste en:

- Usar una versión estable lo más actualizada posible.
- Basada en Linux.
- Priorizar la imagen de menor peso posible que satisfaga nuestras necesidades.
- En caso de que sea posible intentar simplificar una imagen oficial para eliminar herramientas que nos sean inútiles y así poder disminuir el peso de la imagen.

## Opciones 

Tras ver las imagenes que nos proporciona DockerHub con la versión de Go 1.20 y tras descartar ciertas como WindowsServer y NanoServer debido a que están basadas en windows tenemos las siguientes opciones:

- Alpine : Tamaño comprimido = 100 MB , Tamaño de la imagen una vez pulleada = 253 MB.
- Buster : Tamaño comprimido = 282 MB , Tamaño de la imagen una vez pulleada = 720 MB.
- Bullseye : Tamaño comprimido = 300 MB , Tamaño de la imagen una vez pulleada = 777 MB.

## Elección

Debido al tamaño y al que las mejores prácticas sugieren usar esta imagen base se ha elegido Alpine. Una vez tomada esta decisión, se plantea la duda de que versión usar, teniendo como opciones las versiones 3.16 y 3.17, ambas con el mismo peso y consideradas estables.

Siguiendo las mejores prácticas, usaremos la [versión de alpine más actual](https://github.com/docker-library/golang/blob/af7579626a74bc783a4f511a4951955390ef8c95/1.20-rc/alpine3.17/Dockerfile), para así asegurarnos tener una versión que tendrá soporte el mayor tiempo posible. En este caso, hasta 22/11/2024, lo que són más de 5 meses de diferencia si nos decantáramos por alpine3.16 (23/05/2024)