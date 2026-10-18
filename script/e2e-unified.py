#!/usr/bin/env python3
"""Run the complete browser suite against disposable local Dashboard and Agent.

Only locally built binaries, a temporary SQLite database and loopback listeners
are used. Agent updates, remote queries, NAT and external IP discovery are off.
No production configuration or provider account is loaded.
"""
from contextlib import ExitStack
import http.cookiejar
import json
import os
from pathlib import Path
import secrets
import signal
import socket
import subprocess
import tempfile
import time
import urllib.error
import urllib.request
import uuid

ROOT = Path(__file__).resolve().parents[1]


def stop(process):
    if process is None or process.poll() is not None:
        return
    try:
        os.killpg(process.pid, signal.SIGTERM)
    except ProcessLookupError:
        return
    try:
        process.wait(timeout=10)
    except subprocess.TimeoutExpired:
        try:
            os.killpg(process.pid, signal.SIGKILL)
        except ProcessLookupError:
            pass
        process.wait(timeout=5)


def wait_for(stage, processes, check, timeout=60):
    deadline = time.monotonic() + timeout
    while time.monotonic() < deadline:
        if any(process.poll() is not None for process in processes):
            raise RuntimeError(f"{stage}: a local fixture process exited")
        try:
            result = check()
            if result:
                return result
        except (urllib.error.URLError, ConnectionError, TimeoutError):
            pass
        time.sleep(0.25)
    raise RuntimeError(f"{stage}: local fixture readiness timed out")


def main():
    dashboard_binary = ROOT / "bin/nezha-unified"
    agent_binary = ROOT / "bin/nezha-agent"
    if not dashboard_binary.is_file() or not agent_binary.is_file():
        raise RuntimeError("Build both local binaries with make build build-agent first")
    # NZ_* values configure both programs; inherited developer/production values
    # must never override this disposable fixture's files.
    environment = {key: value for key, value in os.environ.items() if not key.startswith("NZ_")}
    environment["GIN_MODE"] = "release"
    environment["NO_PROXY"] = "127.0.0.1,localhost"
    environment["no_proxy"] = environment["NO_PROXY"]
    processes = []
    try:
        with tempfile.TemporaryDirectory(prefix="nezha-unified-e2e-") as temporary:
            directory = Path(temporary)
            with socket.socket() as listener:
                listener.bind(("127.0.0.1", 0))
                port = listener.getsockname()[1]
            base_url = f"http://127.0.0.1:{port}"
            agent_secret = secrets.token_hex(32)
            agent_uuid = str(uuid.uuid4())
            dashboard_config = directory / "dashboard.yaml"
            dashboard_config.write_text(
                f"listen_host: 127.0.0.1\nlisten_port: {port}\n"
                "site_name: Nezha Unified Browser Test\nlanguage: en_US\n"
                f"agent_secret_key: {json.dumps(agent_secret)}\n"
            )
            dashboard_config.chmod(0o600)
            agent_config = directory / "agent.yaml"
            agent_config.write_text(
                f"server: 127.0.0.1:{port}\nclient_secret: {json.dumps(agent_secret)}\n"
                f"uuid: {json.dumps(agent_uuid)}\n"
                "tls: false\ndisable_auto_update: true\ndisable_force_update: true\n"
                "disable_send_query: true\ndisable_nat: true\nreport_delay: 1\n"
                "skip_connection_count: true\nskip_procs_count: true\n"
                # This local JSON endpoint intentionally returns no IP. Replacing
                # the Agent's default IP endpoints prevents external discovery.
                f"custom_ip_api:\n  - {base_url}/api/v1/setting\n"
            )
            agent_config.chmod(0o600)
            with (directory / "dashboard.log").open("w") as dashboard_log, (directory / "agent.log").open("w") as agent_log, ExitStack() as running:
                dashboard = subprocess.Popen(
                    [str(dashboard_binary), "-c", str(dashboard_config), "-db", str(directory / "dashboard.sqlite")],
                    cwd=directory, env=environment, stdout=dashboard_log, stderr=subprocess.STDOUT,
                    start_new_session=True,
                )
                processes.append(dashboard)
                running.callback(stop, dashboard)
                cookies = http.cookiejar.CookieJar()
                client = urllib.request.build_opener(urllib.request.HTTPCookieProcessor(cookies))

                def request(path, data=None):
                    headers = {"Content-Type": "application/json"}
                    for cookie in cookies:
                        if cookie.name == "nz-csrf":
                            headers["X-CSRF-Token"] = cookie.value
                    response = client.open(urllib.request.Request(
                        base_url + path,
                        data=json.dumps(data).encode() if data is not None else None,
                        headers=headers,
                    ), timeout=3)
                    body = json.load(response)
                    if not body.get("success"):
                        return None
                    return body.get("data")

                wait_for("Dashboard startup", processes, lambda: request("/api/v1/setting"))
                request("/api/v1/login", {"username": "admin", "password": "admin"})
                if not request("/api/v1/profile"):
                    raise RuntimeError("Disposable Dashboard login did not succeed")
                agent = subprocess.Popen(
                    [str(agent_binary), "-c", str(agent_config)],
                    cwd=directory, env=environment, stdout=agent_log, stderr=subprocess.STDOUT,
                    start_new_session=True,
                )
                processes.append(agent)
                running.callback(stop, agent)

                def online_agent():
                    servers = request("/api/v1/server") or []
                    return next((server for server in servers if server.get("uuid") == agent_uuid
                                 and server.get("last_active")
                                 and not server["last_active"].startswith("0001-")
                                 and (server.get("host") or {}).get("platform")), None)

                server = wait_for("Agent registration and host report", processes, online_agent)
                # File-manager and UI tests resolve this exact Agent, rather than
                # assuming database IDs happen to start at one.
                browser_environment = {
                    **environment, "E2E_SKIP_WEBSERVER": "1", "E2E_BASE_URL": base_url,
                    "E2E_ADMIN_USER": "admin", "E2E_ADMIN_PASS": "admin",
                    "E2E_AGENT_UUID": agent_uuid, "E2E_AGENT_SERVER_ID": str(server["id"]),
                }
                print("Disposable loopback Dashboard and Agent are ready; running all browser tests.", flush=True)
                browser = subprocess.Popen(
                    ["npm", "run", "e2e"], cwd=ROOT / "web/admin", env=browser_environment,
                    start_new_session=True,
                )
                processes.append(browser)
                running.callback(stop, browser)
                browser_status = browser.wait()
                if browser_status:
                    dashboard_log.flush()
                    # Show only useful error classifications, never raw logs or
                    # authentication/Agent payloads.
                    for phrase in ("server not found", "permission denied", "blocked by nezha WAF", "invalid jwt", "authentication failed"):
                        count = (directory / "dashboard.log").read_text().lower().count(phrase)
                        if count:
                            print(f"Local Dashboard diagnostics: {phrase}: {count}", flush=True)
                return browser_status
    finally:
        for process in reversed(processes):
            stop(process)


if __name__ == "__main__":
    # CI termination must unwind the same private fixture cleanup as Ctrl-C.
    def interrupt_fixture(_signum, _frame):
        raise KeyboardInterrupt

    signal.signal(signal.SIGTERM, interrupt_fixture)
    try:
        raise SystemExit(main())
    except KeyboardInterrupt:
        raise SystemExit(130) from None
    except (RuntimeError, urllib.error.URLError) as error:
        # Startup logs/configuration stay private to the temporary fixture.
        raise SystemExit(str(error)) from None
