# Bot comunicación

## Criterios

Para la elección del bot que comunique en telegram los cambios efectuados en este repositorio tendremos en cuenta los siguientes criterios:
- Debe ser un bot gratuito.
- Posibilidad de notificar tanto de push como de creación de pull requests e issues.
- La puesta a punto y configuración sea la menor posible.
- Se pueda contestar a los issues directamente desde Telegram.

## Opciones

Las 2 opciones que se han contemplado han sido las siguientes:

### Crear un bot de telegram y hacer que este notifique los cambios usando GitHub Actions

- Es gratuito
- Puede notificar de todas las acciones
- Se necesita configurar desde 0
- No se pueden contestar a los issues desde Telegram

### GitHubBot

- Es gratuito
- Puede notificar de todas las acciones indicadas
- Ya viene configurado y su puesta a punto es sencilla.
- Se pueden contestar a los issues desde Telegram

## Elección

Debido a la facilidad de tenerlo configurado para todos los tipos de eventos y junto a la posibilidad de responder directamente desde Telegram usaré GitHubBot como bot que notifique los cambios del repositorio.
