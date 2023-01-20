# Principios F.I.R.S.T

## Fast

Para cumplir este principio se ha buscado la máxima eficiencia de código posible, tanto a la hora de desarrollo como de testing, para que al lanzar los test se pueda tener el resultado de estos de inmediato.

## Independent

Para cumplir este principio se ha procurado que no haya nada externo al propio test que afecte a él, para que de esta manera ningún factor externo pueda influir en el resultado.

## Repeatable

Este principio se ha cumplido haciendo que los tests fabriquen sus propios datos y todo esté igual tanto antes como después de lanzar el test. De esta manera se pueden lanzar todas las veces que se desee y el resultado será siempre el mismo.

## Self-validating

Mediante el uso de la biblioteca de aserciones se ha comprobado que la salida sea la esperada para todos los casos, de esta manera no se necesita de ningún checkeo humano.

## Thorough

Se ha testeado el 100% del código, planteando todos los casos de uso posibles.