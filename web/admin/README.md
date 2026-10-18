# Nezha Admin Frontend

## Bootstrap

```bash
npx swagger-typescript-api -p http://localhost:8008/swagger/doc.json -o ./src/types -n api.ts --no-client --union-enums
```

## End-to-end tests

The complete Playwright suite lives in `tests/e2e/` and runs in the application
job of `.github/workflows/ci.yml` on standard `ubuntu-24.04` runners.

Run it from the repository root:

```bash
make test-e2e
```

This builds the current integrated application and local Agent, installs
Chromium, and starts `script/e2e-unified.py`. The runner creates a private
temporary directory, SQLite database and configuration, starts Dashboard on a
random loopback port, waits for the Agent to register and report host state,
then runs every test against the built management app. Both processes and the
temporary directory are cleaned up on success or failure. Agent automatic
updates, remote queries, NAT and external IP discovery are disabled. No
production configuration or CA, DNS or cloud provider account is used.

The suite contains real local session/password, scoped API token/MCP, CSRF,
Agent-reporting/file-manager and guest-visibility checks. The certificate UI
workflow, permission and mobile-layout tests use synthetic API responses; these
checks establish browser behavior separately from real provider acceptance.

To rerun after both binaries have already been built:

```bash
python3 script/e2e-unified.py
```

To use an existing compatible local test stack, start Vite separately and run:

```bash
cd web/admin
E2E_SKIP_WEBSERVER=1 E2E_BASE_URL=http://localhost:5173 \
  E2E_AGENT_UUID=<connected-agent-uuid> E2E_AGENT_SERVER_ID=<connected-server-id> \
  npm run e2e
```

`E2E_ADMIN_USER` and `E2E_ADMIN_PASS` select the disposable test account when
reusing a local stack; the defaults are `admin` and `admin`. These tests change
and restore that account's password and create/revoke scoped API tokens.
`E2E_CHROMIUM_EXECUTABLE` can select an already installed Chromium executable.
The isolated certificate UI suite can run without Dashboard or Agent using
`npm run e2e -- tests/e2e/unified-ui.spec.ts`; Playwright starts Vite for it.

When `CI` is set, Playwright writes an HTML report and retains traces, screenshots
and videos for failed tests. The temporary Dashboard and Agent logs are private
fixture data and are not published as CI artifacts.
