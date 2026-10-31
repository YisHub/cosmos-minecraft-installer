#!/usr/bin/env python3
"""Arma dist/Cosmos-Espanol.zip a partir de traducciones/<mod>.json (va dentro del instalador).

Cada archivo se copia a assets/<mod>/lang/ para todas las variantes de español del juego, así
que se traduce una sola vez. Antes comprueba que todo sea JSON válido con textos no vacíos.
"""
import json
import sys
import zipfile
from pathlib import Path

ROOT = Path(__file__).resolve().parent
LOCALES = ["es_ar", "es_cl", "es_ec", "es_es", "es_mx", "es_uy", "es_ve"]
# Fecha fija: el mismo contenido da siempre el mismo zip (y el mismo sha1).
DATE = (2026, 1, 1, 0, 0, 0)


def entry(zf: zipfile.ZipFile, name: str, data: bytes) -> None:
    info = zipfile.ZipInfo(name, DATE)
    info.compress_type = zipfile.ZIP_DEFLATED
    info.external_attr = 0o644 << 16
    zf.writestr(info, data)


def main() -> int:
    errors = []
    mods = {}
    for path in sorted((ROOT / "traducciones").glob("*.json")):
        try:
            texts = json.loads(path.read_text(encoding="utf-8"))
        except json.JSONDecodeError as e:
            errors.append(f"{path.name}: JSON inválido ({e})")
            continue
        empty = [k for k, v in texts.items() if not isinstance(v, str) or not v.strip()]
        if empty:
            errors.append(f"{path.name}: textos vacíos: {', '.join(empty)}")
        mods[path.stem] = path.read_bytes()
    json.loads((ROOT / "pack.mcmeta").read_text(encoding="utf-8"))
    if errors:
        print("\n".join(errors), file=sys.stderr)
        return 1

    out = ROOT / "dist" / "Cosmos-Espanol.zip"
    out.parent.mkdir(exist_ok=True)
    with zipfile.ZipFile(out, "w") as zf:
        entry(zf, "pack.mcmeta", (ROOT / "pack.mcmeta").read_bytes())
        for mod, data in mods.items():
            for locale in LOCALES:
                entry(zf, f"assets/{mod}/lang/{locale}.json", data)
    total = sum(len(json.loads(d)) for d in mods.values())
    print(f"{out.relative_to(ROOT)}: {len(mods)} mods, {total} textos")
    return 0


if __name__ == "__main__":
    sys.exit(main())
