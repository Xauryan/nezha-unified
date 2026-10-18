#!/usr/bin/env python3
"""Reproducible retained-source inventory, not a functional/provider acceptance.

The immutable provenance/hash authority remains files.json + upstream.lock.json.
This index locates pages, API declarations, nodes, provider IDs, config fields and
migrations for review when upgrading that baseline. --check never rewrites it.
"""
import argparse
from collections import Counter, defaultdict
import json
from pathlib import Path
import re


ROOT = Path(__file__).resolve().parents[1]
OUTPUT = ROOT / "docs/upstream/capability-inventory.json"


def build_inventory():
    baseline = json.loads((ROOT / "docs/upstream/files.json").read_text())
    lock = json.loads((ROOT / "upstream.lock.json").read_text())
    files = sorted(item["path"] for item in baseline)
    cert = "third_party/certimate/"

    def select(prefix, suffix=None):
        return [name for name in files if name.startswith(prefix)
                and (suffix is None or name.endswith(suffix)) and not name.endswith("_test.go")]

    def content(name):
        return (ROOT / name).read_text(encoding="utf-8-sig")

    providers = defaultdict(list)
    constants = {}
    for constant, kind, expression in re.findall(r'^\s*(\w+)\s*=\s*(\w*ProviderType)\(([^)\n]*)\)',
                                                content(cert + "internal/domain/provider.go"), re.M):
        pieces = []
        for term in expression.split("+"):
            term = term.strip()
            if term.startswith('"'):
                pieces.append(json.loads(term))
            elif term in constants:
                pieces.append(constants[term])
            else:
                raise RuntimeError(f"Provider identifier expression changed at {constant}; review the inventory extractor")
        identifier = "".join(pieces)
        constants[constant] = identifier
        providers[kind].append(identifier)

    provider_paths = {}
    config_fields = {}
    for category, prefix in {
        "deployer": cert + "pkg/core/deployer/providers/",
        "notifier": cert + "pkg/core/notifier/providers/",
        "certmgr": cert + "pkg/core/certmgr/providers/",
        "dns01": cert + "pkg/core/certifier/challengers/dns01/",
        "http01": cert + "pkg/core/certifier/challengers/http01/",
    }.items():
        grouped = defaultdict(list)
        for name in select(prefix, ".go"):
            grouped[name.removeprefix(prefix).split("/")[0]].append(name)
            for struct_name, body in re.findall(r"type\s+(\w*Config\w*)\s+struct\s*\{(.*?)^\}", content(name), re.S | re.M):
                fields = sorted(set(re.findall(r'`json:"([^",]+)', body)) - {"-"})
                if fields:
                    config_fields[f"{name}#{struct_name}"] = fields
        provider_paths[category] = dict(sorted(grouped.items()))

    workflow_source = cert + "ui/src/domain/workflow.ts"
    workflow = content(workflow_source)
    node_block = re.search(r"WORKFLOW_NODE_TYPES = Object.freeze\(\{(.*?)\}\s*as const", workflow, re.S)
    if not node_block:
        raise RuntimeError("workflow node declaration shape changed; review the inventory extractor")
    nodes = dict(re.findall(r'(\w+):\s*"([^"]+)"', node_block[1]))
    node_fields = {}
    for name, body in re.findall(r"export type (WorkflowNodeConfigFor\w+) = \{(.*?)^\};", workflow, re.M | re.S):
        node_fields[name] = sorted(set(re.findall(r"^\s*(\w+)\??:", body, re.M)))

    api_sources = select("cmd/dashboard/controller/", ".go") + select(cert + "internal/rest/handlers/", ".go")
    api_declarations = {}
    route_pattern = re.compile(r'\.(GET|POST|PUT|PATCH|DELETE|HEAD|OPTIONS|Get|Post|Put|Patch|Delete|Head|Options)\(\s*"([^"]+)"')
    for name in api_sources:
        declarations = [{"method": method.upper(), "path_fragment": path} for method, path in route_pattern.findall(content(name))]
        if declarations:
            api_declarations[name] = declarations

    return {
        "schema": 1,
        "evidence_level": "retained-source declarations only; compilation, contracts and live provider acceptance are independent",
        "sources": [{"name": source["name"], "revision": source["revision"], "commit": source["commit"], "path": source["path"]} for source in lock["sources"]],
        "retained_file_counts": dict(sorted(Counter(item["source"] for item in baseline).items())),
        "nezha_admin_pages": select("web/admin/src/routes/", ".tsx"),
        "nezha_api_sources": select("cmd/dashboard/controller/", ".go"),
        "certimate_pages": select(cert + "ui/src/pages/", ".tsx"),
        "certimate_repositories": select(cert + "ui/src/repository/", ".ts"),
        "certimate_api_sources": select(cert + "internal/rest/handlers/", ".go"),
        "api_declaration_note": "Literal source fragments; parent groups and middleware determine effective routes. Patched files may include integrated routes. HTTP acceptance is checked separately.",
        "api_declarations": api_declarations,
        "workflow_nodes": {"source": workflow_source, "types": nodes, "configuration_fields": node_fields,
                           "executors": select(cert + "internal/workflow/engine/executor", ".go")},
        "provider_identifiers": {kind: sorted(set(ids)) for kind, ids in sorted(providers.items())},
        "provider_sources": provider_paths,
        "provider_configuration_fields": dict(sorted(config_fields.items())),
        "configuration_field_note": "Named Go *Config structs' JSON fields and workflow node TypeScript fields; source paths retain comments, nested types and validation rules. Dynamic provider maps must be reviewed in source.",
        "credential_forms": select(cert + "ui/src/components/access/forms/", ".tsx"),
        "workflow_forms": select(cert + "ui/src/components/workflow/", ".tsx"),
        "settings_sources": select(cert + "internal/settings/", ".go") + select(cert + "ui/src/pages/settings/", ".tsx"),
        "certimate_migrations": select(cert + "migrations/", ".go"),
        "nezha_migration_sources": [name for name in select("service/singleton/", ".go") if "config" in name],
        "dns_reference_sources": select("reference/dnsmgr/app/lib/dns/", ".php"),
        "dns_reference_ui": select("reference/dnsmgr/app/view/domain/"),
        "dns_reference_services": select("reference/dnsmgr/app/service/", ".php"),
        "agent_baseline": lock["agent"],
    }


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--check", action="store_true")
    args = parser.parse_args()
    inventory = build_inventory()
    serialized = json.dumps(inventory, indent=2, ensure_ascii=False) + "\n"
    if args.check:
        if not OUTPUT.is_file() or OUTPUT.read_text() != serialized:
            raise SystemExit("Capability inventory is stale; review source changes, run python3 script/inventory-upstream.py, and commit the reviewed result")
        print("Retained page/API/node/provider/configuration/migration inventory matches source declarations")
    else:
        OUTPUT.write_text(serialized)
        print(f"Wrote {OUTPUT.relative_to(ROOT)}")


if __name__ == "__main__":
    main()
