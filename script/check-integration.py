#!/usr/bin/env python3
"""Structural gates; these do not claim functional upstream parity."""
import hashlib
import json
from pathlib import Path
import re

root = Path(__file__).resolve().parents[1]
baseline = json.loads((root / "docs/upstream/files.json").read_text())
patches = set(json.loads((root / "docs/upstream/patches.json").read_text()))
errors = []
for item in baseline:
    p = root / item["path"]
    if not p.is_file():
        errors.append(f"Missing upstream source: {item['path']}")
    elif item["path"] not in patches and hashlib.sha256(p.read_bytes().replace(b"\r\n", b"\n")).hexdigest() != item["sha256"]:
        errors.append(f"Undeclared upstream change: {item['path']}")
for p in (root / ".github/workflows").glob("*.yml"):
    text = p.read_text()
    runners = re.findall(r"^\s*runs-on:\s*(.+)$", text, re.M)
    if not runners or any(r != "ubuntu-24.04" for r in runners):
        errors.append(f"Non-standard/unsupported runner: {p.name}")
    if "pull_request_target:" in text:
        errors.append(f"Privileged PR trigger: {p.name}")
for directory in ["cmd", "model", "pkg", "service"]:
    for p in (root / directory).rglob("*.go"):
        if '"github.com/certimate-go/certimate/internal/' in p.read_text(encoding="utf-8-sig"):
            errors.append(f"Private Certimate import: {p.relative_to(root)}")
for path in ["LICENSE", "web/admin/LICENSE", "third_party/certimate/LICENSE", "reference/dnsmgr/LICENSE"]:
    if not (root / path).is_file():
        errors.append(f"Missing license: {path}")
if errors:
    raise SystemExit("\n".join(errors))
print(f"Checked {len(baseline)} retained upstream files, {len(patches)} declared patch paths, licenses, import boundary and GitHub runners")
