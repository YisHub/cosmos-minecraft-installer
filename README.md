# Cosmos 26.3 — instalador para Windows

Instala en el **launcher oficial de Minecraft** un perfil listo para jugar en el server:

- Minecraft **26.3** con **Fabric Loader 0.20.1**
- Fabric API, Sodium, Iris Shaders, Mod Menu
- Cinematic Interactions, AFK Cinematics
- [Cosmos Tabs](https://github.com/YisHub/cosmos-tabs): ventanas de cofres movibles, pestañas
  entre cofres pegados (con páginas si son muchos), dos ventanas a la vez y cortar/pegar
  (Ctrl+clic, Ctrl+X, Ctrl+V)
- Animaciones: Punchy (primera persona), Not Enough Animations y **Fresh Animations** (con EMF + ETF), ya activado
- Shaders **Complementary Reimagined** activados
- El servidor agregado a la lista de multijugador

## Instalar

1. Ve a **[Releases](../../releases/latest)** y descarga `Instalador-Cosmos-26.3.exe`.
2. Cierra el launcher de Minecraft y ejecuta el `.exe`.
   Si aparece "Windows protegió tu PC": **Más información → Ejecutar de todas formas**.
3. Abre el launcher, elige el perfil **Cosmos 26.3** y pulsa **Jugar**.

Requisitos: haber abierto el launcher oficial al menos una vez y tener una cuenta de Minecraft Java.
Para entrar al servidor hace falta [Tailscale](https://tailscale.com/download) y estar en la red de Jesus.

## El servidor

- Corre en **Crafty** con Fabric 26.3 y estos mods: Fabric API, Cinematic Interactions y
  Cosmos Tabs.
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
- Los mods **no están en este repositorio**: el instalador los descarga de Modrinth y CurseForge
  y comprueba el hash SHA-1 de cada archivo.
- Volver a ejecutarlo actualiza el pack: baja lo que cambió y borra versiones viejas que haya puesto él.
  Los mods que agregues a mano no se tocan.

## Desinstalar

Borra la carpeta `%APPDATA%\.minecraft\cosmos-26.3`, la carpeta
`%APPDATA%\.minecraft\versions\fabric-loader-0.20.1-26.3` y el perfil "Cosmos 26.3" desde el launcher.

## Idioma

Sodium, Punchy, Not Enough Animations y AFK Cinematics solo vienen en inglés. Jesus armó un
paquete de recursos, **Cosmos Español**, que los traduce (524 textos, para todas las variantes
de español del juego); por ahora está solo en su instalación, no en este instalador.
Los ajustes del shader Complementary quedan en inglés: su licencia no permite modificarlo.

## Actualizar el pack (para Jesus)

1. Edita `manifest.json` (URL y `sha1` de cada archivo).
2. `git tag v1.x.x && git push --tags` — GitHub Actions compila el `.exe` y crea la release.

Para actualizar Cosmos Tabs: primero se publica su versión nueva (tag en
[cosmos-tabs](https://github.com/YisHub/cosmos-tabs)), después se cambian URL y `sha1` aquí.
