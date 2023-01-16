# Principios a seguir para la creación de tests

Para la creación de los test se han seguido los principios F.I.R.S.T.

- Fast (Rápido): Los test se ejecutan en milisegundos, lo que hace que lanzarlos no sea una pérdida de tiempo.

- Independent (Independiente): Ninguno de los test tienen dependencias y pueden ser lanzados en cualquier orden sin que esto repercuta a los resultados.

- Repeteable (Repetible): Los test se pueden lanzar múltiples veces en múltiples entornos y el resultado seguirá siendo el mismo, ya que no dependen de la configuración del usuario.

- Self-validating (Auto evaluable): Los test muestran claramente si el test ha sido pasado correctamente o muestra un error. En este caso, muestra una comparación del resultado esperado con el obtenido.

- Timely (A tiempo): Los test se han escrito justo a la par que la funciónes que testean. De esta manera, ninguna función ha sido mergeada sin haber pasado los respectivos tests.