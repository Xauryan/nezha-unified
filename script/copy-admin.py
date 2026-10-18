#!/usr/bin/env python3
"""Use only the in-repository management app for the integrated binary."""
from pathlib import Path
import shutil

root = Path(__file__).resolve().parents[1]
source = root / "web/admin/dist"
target = root / "cmd/dashboard/admin-dist"
if not (source / "index.html").is_file():
    raise SystemExit("Build web/admin before copying assets")
if target.exists():
    shutil.rmtree(target)
shutil.copytree(source, target)
(target / ".gitkeep").touch()
