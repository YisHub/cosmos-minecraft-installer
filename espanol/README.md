# Cosmos Español

Paquete de recursos que traduce al español los mods del pack que solo vienen en inglés. Va
dentro del instalador: no se publica aparte.

- `traducciones/<mod>.json`: los textos en español de cada mod (la clave es el nombre interno
  que usa el mod; solo hay que cambiar el texto de la derecha).
- `pack.mcmeta`: nombre y descripción del paquete (formatos 84 en adelante, Minecraft 26.x).
- `build.py`: comprueba que todo sea JSON válido sin textos vacíos y arma
  `dist/Cosmos-Espanol.zip`, copiando cada archivo a todas las variantes de español del juego
  (es_ar, es_cl, es_ec, es_es, es_mx, es_uy, es_ve). Con el mismo contenido da siempre el mismo zip.

## Corregir o agregar textos

1. Edita el archivo del mod en `traducciones/`.
2. Respeta los huecos `%s`, `%d`, `%1$s`… tal cual: ahí el juego pone números o nombres.
3. `python3 espanol/build.py` para comprobarlo, y publica una versión nueva del instalador.

Si un mod se actualiza y agrega textos, esos se verán en inglés hasta agregarlos aquí. Para
traducir un mod nuevo, crea `traducciones/<id del mod>.json` (el id es la carpeta de
`assets/` dentro del `.jar` del mod).

Se traducen solo los textos; no incluye código ni otros recursos de los mods.
