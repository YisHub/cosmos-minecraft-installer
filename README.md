# Cosmos 26.3 — instalador para Windows

Instala en el **launcher oficial de Minecraft** un perfil listo para jugar en el server:

- Minecraft **26.3** con **Fabric Loader 0.20.1**
- Fabric API, Sodium, Iris Shaders
- Cinematic Interactions, AFK Cinematics
- Animaciones: Not Enough Animations y **Fresh Animations** (con EMF + ETF), ya activado
- Shaders **Complementary Reimagined** activados
- El servidor agregado a la lista de multijugador

## Instalar

1. Ve a **[Releases](../../releases/latest)** y descarga `Instalador-Cosmos-26.3.exe`.
2. Cierra el launcher de Minecraft y ejecuta el `.exe`.
   Si aparece "Windows protegió tu PC": **Más información → Ejecutar de todas formas**.
3. Abre el launcher, elige el perfil **Cosmos 26.3** y pulsa **Jugar**.

Requisitos: haber abierto el launcher oficial al menos una vez y tener una cuenta de Minecraft Java.
Para entrar al servidor hace falta [Tailscale](https://tailscale.com/download) y estar en la red de Jesus.

## Qué hace

- Todo se instala en `%APPDATA%\.minecraft\cosmos-26.3`, separado de tu Minecraft normal.
- Los mods **no están en este repositorio**: el instalador los descarga de Modrinth y CurseForge
  y comprueba el hash SHA-1 de cada archivo.
- Volver a ejecutarlo actualiza el pack: baja lo que cambió y borra versiones viejas que haya puesto él.
  Los mods que agregues a mano no se tocan.

## Desinstalar

Borra la carpeta `%APPDATA%\.minecraft\cosmos-26.3`, la carpeta
`%APPDATA%\.minecraft\versions\fabric-loader-0.20.1-26.3` y el perfil "Cosmos 26.3" desde el launcher.

## Actualizar el pack (para Jesus)

1. Edita `manifest.json` (URL y `sha1` de cada archivo).
2. `git tag v1.x.x && git push --tags` — GitHub Actions compila el `.exe` y crea la release.
