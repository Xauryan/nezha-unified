#!/usr/bin/env python3
"""Real HTTP smoke test against a disposable, loopback-only Dashboard.

No production config, cloud accounts, ACME order or Agent is used. Cookies and
secrets remain in memory. Set --keep to keep a local test server for browser QA.
"""
import argparse
import base64
import hashlib
import http.cookiejar
import io
import json
from pathlib import Path
import secrets
import signal
import socket
import sqlite3
import subprocess
import tempfile
import time
import urllib.error
import urllib.parse
import urllib.request
import zipfile

root = Path(__file__).resolve().parents[1]
parser = argparse.ArgumentParser()
parser.add_argument("--keep", action="store_true")
parser.add_argument("--port", type=int, default=0)
args = parser.parse_args()


class Client:
    def __init__(self, base):
        self.base = base
        self.jar = http.cookiejar.CookieJar()
        self.opener = urllib.request.build_opener(urllib.request.HTTPCookieProcessor(self.jar))

    def request(self, path, data=None, method=None, expected=200, csrf=True, success=True, extra_headers=None):
        headers = {"Content-Type": "application/json"}
        headers.update(extra_headers or {})
        if csrf:
            for cookie in self.jar:
                if cookie.name == "nz-csrf":
                    headers["X-CSRF-Token"] = cookie.value
        request = urllib.request.Request(self.base + path, data=json.dumps(data).encode() if data is not None else None, headers=headers, method=method)
        try:
            response = self.opener.open(request, timeout=10)
        except urllib.error.HTTPError as error:
            response = error
        body = response.read()
        assert response.code == expected, f"{path}: expected {expected}, got {response.code}"
        if "application/json" in response.headers.get("Content-Type", ""):
            result = json.loads(body)
            if expected == 200:
                assert bool(result.get("success")) is success, f"{path}: unexpected authentication/result status"
            return result.get("data")
        return body


with tempfile.TemporaryDirectory(prefix="nezha-unified-smoke-") as tmp:
    cwd = Path(tmp)
    with socket.socket() as sock:
        sock.bind(("127.0.0.1", args.port))
        port = sock.getsockname()[1]
    config = cwd / "config.yaml"
    config.write_text(f"listen_host: 127.0.0.1\nlisten_port: {port}\nsite_name: Nezha Unified Test\n")
    log_path = cwd / "dashboard.log"
    with log_path.open("w") as log:
        process = subprocess.Popen([str(root / "bin/nezha-unified"), "-c", str(config), "-db", str(cwd / "sqlite.db")], cwd=cwd, stdout=log, stderr=subprocess.STDOUT)
        try:
            base = f"http://127.0.0.1:{port}"
            anonymous = Client(base)
            for attempt in range(120):
                if process.poll() is not None:
                    raise RuntimeError("Dashboard exited before startup; inspect local test log")
                try:
                    anonymous.request("/api/v1/setting")
                    break
                except (urllib.error.URLError, ConnectionError):
                    time.sleep(0.25)
            else:
                raise RuntimeError("Dashboard startup timeout")
            for path in ["/dashboard/credentials", "/dashboard/certificates", "/dashboard/certificates/workflows", "/dashboard/certificates/presets", "/dashboard/dns", "/dashboard/ddns", "/dashboard/targets", "/dashboard/identities", "/dashboard/runs", "/dashboard/permissions", "/"]:
                assert b"<html" in anonymous.request(path).lower(), f"No built frontend for {path}"
            # Vite may create chunks beginning with an underscore; verify the
            # complete built asset tree is included by the binary's go:embed.
            for asset in sorted((root / "web/admin/dist/assets").rglob("*")):
                if asset.is_file():
                    path = "/dashboard/" + asset.relative_to(root / "web/admin/dist").as_posix()
                    observed = anonymous.request(urllib.parse.quote(path, safe="/"))
                    assert hashlib.sha256(observed).digest() == hashlib.sha256(asset.read_bytes()).digest(), f"Built asset mismatch: {path}"
            # Nezha's shared authentication middleware uses a 200/false error
            # envelope. Preserve that contract; resource handlers use HTTP 4xx.
            anonymous.request("/api/v1/unified/credentials", success=False)
            anonymous.request("/api/collections/access/records", expected=404)
            admin = Client(base)
            admin.request("/api/v1/login", {"username": "admin", "password": "admin"})
            admin.request("/api/v1/profile")
            status = admin.request("/api/v1/unified/status")
            assert status["storage"] == "ready" and status["workflow_execution"] is True
            for kind in ["credentials", "certificates", "workflows"]:
                assert admin.request(f"/api/v1/unified/{kind}")["total"] == 0
            secret = secrets.token_urlsafe(32)
            credential = {"name": "Smoke credential", "provider": "cloudflare", "config": {"apiToken": secret}}
            admin.request("/api/v1/unified/credentials", credential, expected=403, csrf=False)
            resource = admin.request("/api/v1/unified/credentials", credential)
            assert "config" not in resource and secret not in json.dumps(resource)
            admin.request("/api/v1/unified/credentials?page=0", expected=400)
            password = secrets.token_urlsafe(24)
            uid = admin.request("/api/v1/user", {"username": "smoke-member", "password": password, "role": 1})
            member = Client(base)
            member.request("/api/v1/login", {"username": "smoke-member", "password": password})
            member.request("/api/v1/profile")
            assert member.request("/api/v1/unified/credentials")["total"] == 0
            member.request("/api/v1/unified/credentials", credential, expected=403)
            grant = admin.request("/api/v1/unified/grants", {"user_id": uid, "kind": "credential", "resource_id": resource["id"], "action": "read"})
            visible = member.request("/api/v1/unified/credentials?per_page=1")
            assert visible["total"] == 1 and visible["items"][0]["id"] == resource["id"]
            admin.request(f"/api/v1/unified/grants/{grant['id']}", method="DELETE")
            assert member.request("/api/v1/unified/credentials")["total"] == 0
            member.request("/api/v1/unified/certificates/missing/pem", expected=403)
            cert_path, key_path = cwd / "synthetic.pem", cwd / "synthetic.key"
            subprocess.run(["openssl", "req", "-x509", "-newkey", "ec", "-pkeyopt", "ec_paramgen_curve:prime256v1", "-nodes", "-keyout", str(key_path), "-out", str(cert_path), "-days", "1", "-subj", "/CN=smoke.example.test", "-addext", "subjectAltName=DNS:smoke.example.test", "-addext", "basicConstraints=critical,CA:FALSE"], check=True, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
            cert = admin.request("/api/v1/unified/certificates/import", {"certificate": cert_path.read_text() + key_path.read_text(), "private_key": key_path.read_text()})
            assert cert["names"] == "smoke.example.test"
            admin.request("/api/v1/unified/grants", {"user_id": uid, "kind": "certificate", "resource_id": cert["id"], "action": "read"})
            assert member.request("/api/v1/unified/certificates")["total"] == 1
            pem_path = f"/api/v1/unified/certificates/{cert['id']}/pem"
            member.request(pem_path, expected=403)
            download_grant = admin.request("/api/v1/unified/grants", {"user_id": uid, "kind": "certificate", "resource_id": cert["id"], "action": "download"})
            public_chain = member.request(pem_path)
            assert b"BEGIN CERTIFICATE" in public_chain and b"PRIVATE KEY" not in public_chain
            admin.request(f"/api/v1/unified/grants/{download_grant['id']}", method="DELETE")
            member.request(pem_path, expected=403)
            native = "/api/v1/unified/certimate"
            credential_page = admin.request(native + "/resources/access?page=1&perPage=1")
            assert credential_page["totalItems"] == 1 and len(credential_page["items"]) == 1
            assert "config" not in credential_page["items"][0] and secret not in json.dumps(credential_page)
            # Export is independently authorized from metadata/public-chain read.
            member.request(native + f"/certificates/{cert['id']}/download", {"fileFormat": "PEM"}, expected=403)
            exported = admin.request(native + f"/certificates/{cert['id']}/download", {"fileFormat": "PEM"})
            with zipfile.ZipFile(io.BytesIO(base64.b64decode(exported["zipBytes"]))) as bundle:
                material = b"\n".join(bundle.read(name) for name in bundle.namelist())
                assert b"BEGIN CERTIFICATE" in material and b"PRIVATE KEY" in material
            for file_format in ("PFX", "JKS"):
                converted = admin.request(native + f"/certificates/{cert['id']}/download", {
                    "fileFormat": file_format, "pfxPassword": "synthetic-export-password",
                    "jksAlias": "synthetic", "jksKeypass": "synthetic-export-password",
                    "jksStorepass": "synthetic-export-password"})
                with zipfile.ZipFile(io.BytesIO(base64.b64decode(converted["zipBytes"]))) as bundle:
                    assert bundle.namelist() and all(bundle.read(name) for name in bundle.namelist())
            # Exercise the real embedded dispatcher with a graph having no
            # network, CA, DNS, host or provider side effects.
            graph = {"nodes": [
                {"id": "synthetic-start", "type": "start", "data": {"name": "synthetic-start"}},
                {"id": "synthetic-delay", "type": "delay", "data": {"name": "synthetic-delay", "config": {"wait": 0}}},
                {"id": "synthetic-end", "type": "end", "data": {"name": "synthetic-end"}},
            ]}
            workflow_body = {"name": "Synthetic HTTP workflow", "trigger": "manual", "enabled": True,
                             "hasContent": True, "graphContent": graph, "graphDraft": graph}
            member.request(native + "/resources/workflow", workflow_body, expected=403)
            workflow = admin.request(native + "/resources/workflow", workflow_body,
                                     extra_headers={"Idempotency-Key": secrets.token_hex(16)})
            workflow_id = workflow["id"]
            workflow_path = native + f"/resources/workflow/{workflow_id}"
            member.request(workflow_path, expected=403)
            assert member.request(native + "/resources/workflow?perPage=1")["totalItems"] == 0
            workflow_grant = admin.request("/api/v1/unified/grants", {
                "user_id": uid, "kind": "workflow", "resource_id": workflow_id, "action": "read"})
            assert member.request(workflow_path)["id"] == workflow_id
            visible_workflows = member.request(native + "/resources/workflow?perPage=1")
            assert visible_workflows["totalItems"] == 1 and len(visible_workflows["items"]) == 1
            member.request(native + f"/workflows/{workflow_id}/runs", {"trigger": "manual"}, expected=403)
            started = admin.request(native + f"/workflows/{workflow_id}/runs", {"trigger": "manual"})
            run_id = started["runId"]
            run_path = native + f"/resources/workflow_run/{run_id}"
            for attempt in range(120):
                run = admin.request(run_path)
                if run["status"] in {"succeeded", "failed", "canceled"}:
                    break
                time.sleep(0.1)
            assert run["status"] == "succeeded", "Synthetic workflow did not succeed"
            assert member.request(run_path)["id"] == run_id
            run_filter = urllib.parse.urlencode({"filter": f"workflowRef='{workflow_id}'", "perPage": 1})
            assert member.request(native + "/resources/workflow_run?" + run_filter)["totalItems"] == 1
            admin.request(f"/api/v1/unified/grants/{workflow_grant['id']}", method="DELETE")
            member.request(workflow_path, expected=403)
            member.request(run_path, expected=403)
            assert member.request(native + "/resources/workflow_run?" + run_filter)["totalItems"] == 0
            with sqlite3.connect(cwd / "certimate/data.db") as db:
                assert db.execute("SELECT COUNT(*) FROM _superusers").fetchone()[0] == 0
                raw = db.execute("SELECT config, unified_envelope FROM access WHERE id = ?", (resource["id"],)).fetchone()
                assert secret not in " ".join(raw) and raw[1].startswith("v1:")
                raw_graph = db.execute("SELECT graphContent FROM workflow WHERE id = ?", (workflow_id,)).fetchone()[0]
                assert "synthetic-start" not in raw_graph, "Workflow graph was stored without encryption"
            legacy_id = admin.request("/api/v1/ddns", {"name": "Synthetic legacy DDNS",
                "provider": "cloudflare", "access_secret": "synthetic-legacy-http-secret",
                "domains": ["legacy.example.test"], "max_retries": 3, "enable_ipv4": True})
            legacy = next(item for item in admin.request("/api/v1/ddns") if item["id"] == legacy_id)
            assert legacy["credential_id"] and "synthetic-legacy-http-secret" not in json.dumps(legacy)
            with sqlite3.connect(cwd / "sqlite.db") as db:
                assert db.execute("SELECT access_secret FROM ddns WHERE id=?", (legacy_id,)).fetchone()[0] == ""
            target = admin.request("/api/v1/unified/targets", {"name": "Synthetic shared cloud target", "kind": "certimate",
                "provider": "cloudflare-pages", "credential_id": resource["id"], "credential_version": resource["version"],
                "allowed_names": ["smoke.example.test"], "enabled": False, "config": {"syntheticSetting": "synthetic-private-target-setting"}})
            assert target["configured"] and "config" not in target and "configuration_envelope" not in target
            member.request(f"/api/v1/unified/targets/{target['id']}", expected=403)
            print("PASS: built frontend routes, single login, no PocketBase API, CSRF, encrypted credential/workflow, scoped pagination/detail/run history, immediate revoke, certificate import and separate public/private downloads, and successful embedded workflow execution without external providers")
            if args.keep:
                print(f"Browser QA server: {base}/dashboard/ (disposable admin/admin)", flush=True)
                while process.poll() is None:
                    time.sleep(1)
        finally:
            process.send_signal(signal.SIGTERM)
            try:
                process.wait(timeout=20)
            except subprocess.TimeoutExpired:
                process.kill()
                process.wait(timeout=5)

    with tempfile.TemporaryDirectory(prefix="nezha-unified-recovery-smoke-") as recovery_tmp:
        recovery = Path(recovery_tmp)
        key = recovery / "independent.key"
        key.write_bytes(secrets.token_bytes(32))
        key.chmod(0o600)
        archive = recovery / "synthetic.nzb"
        restored = recovery / "restored"
        binary = str(root / "bin/nezha-unified")
        def offline(*arguments):
            result = subprocess.run([binary, *map(str, arguments)], cwd=cwd, capture_output=True, timeout=45)
            assert result.returncode == 0, "offline backup/recovery command failed"
            return result.stdout
        offline("--backup-create", archive, "-db", cwd / "sqlite.db", "-c", config, "--backup-key-file", key)
        offline("--backup-rehearse", archive, "--backup-key-file", key)
        offline("--backup-restore", archive, "--backup-destination", restored, "--backup-key-file", key)
        report = json.loads(offline("--backup-inspect-recovery", restored))
        evidence = recovery / "evidence.json"
        evidence.write_text(json.dumps({"archive_id": report["archive_id"], "reference_digest": report["reference_digest"],
            "boundary_evidence": "Disposable synthetic fixture; no external provider operations occurred.",
            "operations": [{"kind": op["kind"], "id": op["id"], "outcome": "retired_without_replay",
                "evidence": "Synthetic operation retired without contacting any provider."} for op in report["operations"]]}))
        offline("--backup-reconcile", restored, "--backup-evidence", evidence)
        with sqlite3.connect(restored / "sqlite.db") as db:
            assert db.execute("SELECT count(*) FROM unified_deployment_targets WHERE enabled=1").fetchone()[0] == 0
            assert db.execute("SELECT count(*) FROM unified_workflow_policies WHERE enabled=1").fetchone()[0] == 0
        print("PASS: offline CLI backup, independent-key rehearsal, restore fence and exact operation reconciliation without provider replay")
