# Cosmos 26.3 — instalador para Windows

Instala en el **launcher oficial de Minecraft** un perfil listo para jugar en el server:

- Minecraft **26.3** con **Fabric Loader 0.20.1**, con 6 GB de RAM
- Fabric API, Sodium, Iris Shaders, Mod Menu
- **Distant Horizons**: terreno lejano (LODs) mucho más allá de la distancia de renderizado
- Rendimiento: Lithium, FerriteCore, ImmediatelyFast, Entity Culling
- Cinematic Interactions, AFK Cinematics
- [Cosmos Tabs](https://github.com/YisHub/cosmos-tabs): ventanas de cofres movibles, pestañas
  entre cofres pegados (con páginas si son muchos), dos ventanas a la vez y cortar/pegar
  (Ctrl+clic, Ctrl+X, Ctrl+V)
- Animaciones: Punchy (primera persona), Not Enough Animations y **Fresh Animations** (con EMF + ETF), ya activado
- Mundo vivo: **Fancy World Animations** (puertas, trampillas, palancas…) e **Interactive Foliage**
  (el pasto se aparta al pasar y se mueve con el viento; necesita Sway API)
- **Scribble**: editor de libros con formato
- Shaders **Complementary Reimagined** activados
- **Cosmos Español**: traduce al español los mods que solo vienen en inglés, ya activado
- El servidor agregado a la lista de multijugador

## Instalar

1. Ve a **[Releases](../../releases/latest)** y descarga `Instalador-Cosmos-26.3.exe`.
2. Cierra el launcher de Minecraft y ejecuta el `.exe`.
   Si aparece "Windows protegió tu PC": **Más información → Ejecutar de todas formas**.
3. Abre el launcher, elige el perfil **Cosmos 26.3** y pulsa **Jugar**.

Requisitos: haber abierto el launcher oficial al menos una vez y tener una cuenta de Minecraft Java.
Para entrar al servidor hace falta [Tailscale](https://tailscale.com/download) y estar en la red de Jesus.

## El servidor

- Corre en **Crafty** con Fabric 26.3, 6 GB de RAM y estos mods: Fabric API, Cinematic
  Interactions, Cosmos Tabs, Distant Horizons, Lithium, FerriteCore, ScalableLux, Krypton y
  Trader Mirage (traducido al español: los comerciantes errantes pueden ser ilusionistas disfrazados).
- **Distant Horizons** tiene los LODs precargados: unos 16 km alrededor del spawn en el mundo
  superior (1024 chunks), 256 chunks en el Nether y 384 en el End. El servidor se los manda a
  quien tenga DH, como máximo 500 KB/s por jugador; quien no lo tenga entra igual. La distancia
  de LODs que se ve la elige cada uno en las opciones de DH, según su PC.
- **Sin lista blanca** (`white-list=false`): entra cualquiera que llegue por Tailscale. Ojo: en
  Minecraft 26.x todo servidor nuevo se crea *con* lista blanca; si alguien ve
  *"You are not white-listed"*, es eso.
- Dificultad **difícil**. Para que alcance con que duerma uno solo:
  `/gamerule players_sleeping_percentage 1`.
- **No hace falta este pack para entrar**: con Minecraft 26.3 + Fabric (o incluso sin mods
  del lado del jugador) se puede jugar; solo no tendrás pestañas, animaciones ni shaders.
- Si cambias los mods del servidor, hay que **reiniciarlo desde Crafty** para que los cargue.

## Qué hace el instalador

- Todo se instala en `%APPDATA%\.minecraft\cosmos-26.3`, separado de tu Minecraft normal.
- Los mods **no están en este repositorio**: el instalador los descarga de Modrinth,
  CurseForge y GitHub (Cosmos Tabs) y comprueba el hash SHA-1 de cada archivo. Lo único que
  viaja dentro del `.exe` es el paquete de traducciones Cosmos Español.
- Activa los shaders y los paquetes de recursos del pack (Fresh Animations, Punchy y Cosmos
  Español) sin desactivar los que ya tengas.
- Agrega el servidor a la lista de multijugador.
- Deja configurados dos mods para que se vean bien con Complementary:
  - **Interactive Foliage** con el renderizador por CPU (con el de GPU, el pasto lejano se ve
    blanco). Solo si no tienes una configuración propia.
  - **Punchy** sin animar los 17 tipos de mapa (con sus animaciones el mapa se ve cortado y
    corrido). Si ya tenías Punchy, solo agrega los mapas a su lista de excluidos.
- Volver a ejecutarlo actualiza el pack: baja lo que cambió y borra versiones viejas que haya puesto él.
  Los mods que agregues a mano no se tocan.

## Problemas comunes

| Mensaje o síntoma | Causa y solución |
|---|---|
| *"You are not white-listed on this server!"* | El servidor tiene la lista blanca activada (viene así por defecto en 26.x). Jesus pone `white-list=false` en Crafty o te agrega con `/whitelist add <nombre>`. |
| *"This server requires Fabric API… cosmostabs"* | Era un error de Cosmos Tabs 1.0.0. Con la 1.0.1 (instalador v1.0.5 en adelante) se entra con o sin el mod. Vuelve a ejecutar el instalador. |
| Las pestañas o las dos ventanas no aparecen | El servidor y tu pack tienen que tener la misma versión de Cosmos Tabs. Vuelve a ejecutar el instalador. |
| Cofres que quedan con la tapa abierta | Error de Cosmos Tabs 1.0.0, arreglado en la 1.0.1. Abre y cierra el cofre una vez. |
| Sodium o Punchy siguen en inglés | Pon el juego en español (*Opciones → Idioma*) y comprueba que **Cosmos Español** esté activo en *Opciones → Paquetes de recursos*. |
| No conecta al servidor | Hace falta Tailscale conectado a la red de Jesus. |
| Matas de pasto blancas a lo lejos | Interactive Foliage con el renderizador por GPU. En su configuración (*Mods → Interactive Foliage*) pon el renderizador en CPU. El instalador lo deja así si no tenías configuración propia. |
| El mapa en la mano se ve cortado, corrido o con un rectángulo blanco y azul | Punchy animando el mapa. Vuelve a ejecutar el instalador (v1.0.7 en adelante): agrega los mapas a la lista de excluidos de Punchy. |
| El horizonte lejano tarda en aparecer al entrar | Distant Horizons está bajando los LODs del servidor (hasta 500 KB/s). Solo pasa la primera vez en cada zona: después quedan guardados en tu PC. |
| El agua se ve más opaca al acercarse a la superficie | Es el reflejo del shader (más reflejo cuanto más de costado miras el agua). Mira más hacia abajo para ver el fondo. |

## Desinstalar

Borra la carpeta `%APPDATA%\.minecraft\cosmos-26.3`, la carpeta
`%APPDATA%\.minecraft\versions\fabric-loader-0.20.1-26.3` y el perfil "Cosmos 26.3" desde el launcher.

## Idioma

Muchos mods del pack solo vienen en inglés, o traen un español incompleto o para una sola
variante. El paquete de recursos **Cosmos Español** los traduce completos (1320 textos):

| Mod | Qué traduce | Textos |
|---|---|---|
| Distant Horizons | Toda la configuración, con descripciones y valores | 489 |
| Punchy | Ajustes, editor y mensajes | 280 |
| Mod Menu | Completo en todas las variantes (solo venía completo en es_ar) | 172 |
| Not Enough Animations | Ajustes completos | 107 |
| Sodium | Opciones de video completas, con descripciones y avisos | 105 |
| Fancy World Animations | Ajustes completos | 55 |
| Scribble | Editor de libros, colores, avisos y ajustes | 52 |
| Entity Culling | Ajustes completos | 28 |
| AFK Cinematics | Ajustes y teclas | 16 |
| Iris, EMF, ETF | Solo lo que les faltaba | 16 |

- Viene **dentro del instalador** y queda activado; se ve con el juego en cualquier variante de
  español (*Opciones → Idioma*). Los que ya tenían el pack lo reciben al volver a ejecutarlo.
- Cinematic Interactions, Interactive Foliage y Cosmos Tabs ya vienen en español. Lithium,
  FerriteCore, ImmediatelyFast y Sway no tienen textos.
- Trader Mirage (en el servidor) no usa archivos de idioma: sus mensajes van escritos en el
  datapack, así que el servidor usa una copia con los textos ya en español.
- Los ajustes del shader Complementary quedan en inglés: su licencia no permite modificarlo.

Las traducciones están en [`espanol/`](espanol/README.md), una por mod.

## Versiones

Ver [CHANGELOG.md](CHANGELOG.md).

## Actualizar el pack (para Jesus)

1. Edita `manifest.json` (URL y `sha1` de cada archivo).
2. `git tag v1.x.x && git push --tags` — GitHub Actions compila el `.exe` y crea la release.

Las traducciones no se descargan: GitHub Actions arma `Cosmos-Espanol.zip` con
`espanol/build.py` y queda dentro del `.exe` (`"embedded": true` en el manifiesto). Para compilar
a mano, corre primero `python3 espanol/build.py`.

Para actualizar Cosmos Tabs: primero se publica su versión nueva (tag en
[cosmos-tabs](https://github.com/YisHub/cosmos-tabs)), después se cambian URL y `sha1` aquí.
