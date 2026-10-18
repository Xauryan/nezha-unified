# Nezha Unified implementation rules

- Preserve the complete scope in `docs/integration/plan.zh-CN.md`; keep `docs/integration/status.md` accurate. Source retention, build success, contract tests and real provider acceptance are distinct.
- Keep Nezha's original Go package/module paths for upstream maintenance. Only import Certimate through `github.com/certimate-go/certimate/pkg/embedded`.
- One management app (`web/admin`), one Nezha session, one Dashboard process. Do not expose PocketBase routes or bootstrap a separate certificate administrator.
- Authorization applies before count/pagination and to detail/download/events/side effects. Check restricted tokens even for administrators. Never enable executors with only UI-level authorization.
- Keep credential and key material out of metadata, logs, source, fixtures and command output. Use only synthetic secrets in tests. Preserve local `outputs/` and runtime data.
- Use `make` so Go 1.26.6 is selected. For integration changes run relevant tests, then `make check test`; run `make race` for state/lifecycle/auth changes and frontend tests/build for UI changes.
- All active CI jobs must use standard GitHub-hosted `ubuntu-24.04`. No self-hosted/ARC/larger runners or production-provider credentials.
- Update fixed versions, baseline manifests and allowed patch paths together. Do not fetch an independently released upstream admin frontend over the integrated app.
