#!/usr/bin/env python3
"""Fetch immutable public themes; this script cannot replace the admin app."""
import hashlib
import io
import json
from pathlib import Path, PurePosixPath
import shutil
import tempfile
import urllib.request
import zipfile

root = Path(__file__).resolve().parents[1]
lock = json.loads((root / "upstream.lock.json").read_text())
for theme in lock["public_themes"]:
    path = theme["path"]
    if path == "admin-dist" or not path.endswith("-dist") or "/" in path or ".." in path:
        raise SystemExit("Invalid public theme path")
    with urllib.request.urlopen(theme["url"], timeout=60) as response:
        blob = response.read(32 * 1024 * 1024 + 1)
    if hashlib.sha256(blob).hexdigest() != theme["sha256"]:
        raise SystemExit(f"Checksum mismatch for {path}")
    with zipfile.ZipFile(io.BytesIO(blob)) as archive, tempfile.TemporaryDirectory() as tmp:
        if sum(entry.file_size for entry in archive.infolist()) > 128 * 1024 * 1024:
            raise SystemExit("Theme expands beyond allowed size")
        for entry in archive.infolist():
            name = PurePosixPath(entry.filename)
            if name.is_absolute() or ".." in name.parts or not name.parts or name.parts[0] != "dist":
                raise SystemExit("Theme contains an unexpected path")
            if (entry.external_attr >> 16) & 0o170000 == 0o120000:
                raise SystemExit("Theme contains a symbolic link")
        archive.extractall(tmp)
        source = Path(tmp) / "dist"
        if not (source / "index.html").is_file():
            raise SystemExit("Theme has no index.html")
        target = root / "cmd/dashboard" / path
        if target.exists():
            shutil.rmtree(target)
        shutil.copytree(source, target)
        if path == "user-dist":
            (target / ".gitkeep").touch()
    print(f"Verified public theme: {path}")
