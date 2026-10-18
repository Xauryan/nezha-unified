#!/usr/bin/env python3
"""Structural gates; these do not claim functional upstream parity."""
import hashlib
import json
from pathlib import Path
import re
import runpy

root = Path(__file__).resolve().parents[1]
baseline = json.loads((root / "docs/upstream/files.json").read_text())
patches = set(json.loads((root / "docs/upstream/patches.json").read_text()))
errors = []
if len(patches) != len(json.loads((root / "docs/upstream/patches.json").read_text())):
    errors.append("Duplicate declared upstream patch paths")
if not patches.issubset({item["path"] for item in baseline}):
    errors.append("Declared patch path is not in the retained upstream baseline")
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
        source = p.read_text(encoding="utf-8-sig")
        imports = re.findall(r'"(github\.com/certimate-go/certimate/[^"\n]+)"', source)
        if any(name != "github.com/certimate-go/certimate/pkg/embedded" for name in imports):
            errors.append(f"Certimate import outside public embedded boundary: {p.relative_to(root)}")
for path in ["LICENSE", "web/admin/LICENSE", "third_party/certimate/LICENSE", "reference/dnsmgr/LICENSE", "third_party/agent/LICENSE"]:
    if not (root / path).is_file():
        errors.append(f"Missing license: {path}")
try:
    expected = runpy.run_path(str(root / "script/inventory-upstream.py"))["build_inventory"]()
    actual = json.loads((root / "docs/upstream/capability-inventory.json").read_text())
    if actual != expected:
        errors.append("Stale capability inventory; run python3 script/inventory-upstream.py after reviewing the upstream changes")
    retained_providers = json.loads((root / "docs/upstream/provider-sources.json").read_text())
    for category in ("deployer", "notifier", "certmgr"):
        if sorted(retained_providers[category]) != sorted(expected["provider_sources"][category]):
            errors.append(f"Provider source inventory mismatch: {category}")
except (OSError, ValueError, KeyError, RuntimeError) as error:
    errors.append(f"Cannot validate capability inventory: {error}")
for shared in ["pkg/agentproto/protocol.go", "pkg/certdeploy/executor.go"]:
    host, agent = root / shared, root / "third_party/agent" / shared
    if not host.is_file() or not agent.is_file():
        errors.append(f"Missing shared Dashboard/Agent protocol source: {shared}")
    elif host.read_text().replace("github.com/nezhahq/nezha/", "github.com/nezhahq/agent/") != agent.read_text():
        errors.append(f"Dashboard/Agent protocol implementation drift: {shared}")
if errors:
    raise SystemExit("\n".join(errors))
print(f"Checked {len(baseline)} retained upstream files, {len(patches)} declared patch paths, licenses, import boundary and GitHub runners")
