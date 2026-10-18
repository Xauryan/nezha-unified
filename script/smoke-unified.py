#!/usr/bin/env python3
"""Real HTTP smoke test against a disposable, loopback-only Dashboard.

No production config, cloud accounts, ACME order or Agent is used. Cookies and
secrets remain in memory. Set --keep to keep a local test server for browser QA.
"""
import argparse
import http.cookiejar
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
import urllib.request

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

    def request(self, path, data=None, method=None, expected=200, csrf=True, success=True):
        headers = {"Content-Type": "application/json"}
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
            for path in ["/dashboard/credentials", "/dashboard/certificates", "/dashboard/certificates/workflows", "/dashboard/permissions", "/"]:
                assert b"<html" in anonymous.request(path).lower(), f"No built frontend for {path}"
            # Nezha's shared authentication middleware uses a 200/false error
            # envelope. Preserve that contract; resource handlers use HTTP 4xx.
            anonymous.request("/api/v1/unified/credentials", success=False)
            anonymous.request("/api/collections/access/records", expected=404)
            admin = Client(base)
            admin.request("/api/v1/login", {"username": "admin", "password": "admin"})
            admin.request("/api/v1/profile")
            status = admin.request("/api/v1/unified/status")
            assert status["storage"] == "ready" and status["workflow_execution"] is False
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
            with sqlite3.connect(cwd / "certimate/data.db") as db:
                assert db.execute("SELECT COUNT(*) FROM _superusers").fetchone()[0] == 0
                raw = db.execute("SELECT config, unified_envelope FROM access WHERE id = ?", (resource["id"],)).fetchone()
                assert secret not in " ".join(raw) and raw[1].startswith("v1:")
            print("PASS: built frontends, single login, embedded migrations, no PocketBase API, CSRF, encrypted credential, filtered pagination, grant/revoke, certificate import and independent public-chain download permission")
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
