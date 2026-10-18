# Implementation and acceptance status

Status is recorded by working behavior, not by the presence of a source directory. This initial development slice is **not a release or production acceptance**. All items in the full plan remain in scope.

| Area | Initial implementation | Remaining acceptance |
|---|---|---|
| Upstreams | Dashboard, admin, Certimate, SDK snapshot and DNS reference are pinned and retained | Full page/API/node/provider/field inventories and upgrade automation |
| Host runtime | One Dashboard owns Nezha and embedded Certimate databases; shutdown closes both module resources | Executor/scheduler lifecycle and bounded queues |
| Identity | Nezha JWT/session and CSRF; no PocketBase superuser bootstrap or router | New scoped PAT/service identities and resource constraints |
| Resources | Metadata projection, exact per-user resource grants, filtered pagination, immediate revoke | Delegated workflows, credential use scopes and live event filtering |
| Credentials | Create and list; authenticated encryption, resource-bound ciphertext, atomic local audit | Provider-specific forms/validation, resolver, version rotation, disable, plaintext export policy, migration |
| Certificates | Import matching X.509/key pair; derive actual SANs; encrypted key; separately authorized public-chain download | ACME issuance/renewal, format conversion, revoke/delete, delegated SAN validation at every side-effect boundary |
| Certimate workflows | Full upstream source retained, metadata list connected | Full designer/pages/repositories, settings, all nodes/providers, application-service side effects and cancellation |
| DNS | Tested name/record authorization primitives | Cloudflare, DNSPod and EdgeOne adapters, real Zone identity, CRUD/batch/import/export, audit and resource UI |
| DDNS | Original Nezha implementation retained | Credential references, single writer, IP freshness, scoped device tokens and unified execution |
| Agent deployment | Original Nezha RPC retained | General AgentRPC, capability acceptance, durable leases/tasks, supersession, unknown-result reconciliation, rollback |
| Frontend | Native credentials/certificates/permissions routes and shared session/theme/i18n | Complete Certimate UI and DNS/targets/aggregated runs; broader browser regression |
| Backup/recovery | Recovery key is generated once and required with an existing DB; new-store reopen tested | Coordinated online backup, cross-store intents/outbox and real restored-state reconciliation |
| CI | Standard GitHub-hosted Ubuntu runners; pinned actions and toolchains; tests/build/HTTP smoke | Agent compatibility stress on fixed agent baseline, provider contracts and authorized live acceptance |

## Next implementation sequence

1. Extend the `access` envelope adapter into a single credential resolver with explicit purpose, resource and version checks; define operation-intent activation and recovery before cross-store writes.
2. Add authoritative DNS account/Zone models and adapters, persist DNS subtree grants and enforce them on provider-fetched original records; integrate existing DDNS through the same entry point.
3. Refactor Certimate singleton/request-hook dependencies into the public embedding boundary. Put current-authority checks at manual, scheduled, retry and recovery execution, then enable the scheduler and migrate the full designer/API surface.
4. Extract AgentRPC with independent purpose cancellation, add durable versioned deployments and reconcile unknown remote outcomes before retries.
5. Implement coordinated backup/restore and finish the full acceptance matrix against fixed upstream inventories.

## Important architectural choices

- `third_party/certimate/pkg/embedded` is the only cross-module import boundary. The initial adapter uses PocketBase core directly; it deliberately does not initialize Certimate's old singleton or scheduler before those paths honor host authorization.
- Nezha GORM owns grants and grant audit. Certimate owns access/certificate records and their audit. Credential/certificate create and audit are one PocketBase transaction; grant/revoke and audit are one GORM transaction. No GORM writes target PocketBase tables.
- Original secret fields stay empty for new encrypted resources. The future workflow resolver must decrypt at the execution boundary; pointing unmodified Certimate executors at these rows would be incorrect.
- New-module PAT requests fail closed, including admin PATs. Original Nezha PAT behavior remains unchanged. Resource-scoped PAT support is a separate explicit feature.
- Metadata read, public certificate download, secret use, secret export and execution are distinct actions. Only the first two are exposed for delegation in this slice.
- Known toolchain incompatibility: Go 1.27 JSON v2 causes recursion in PocketBase v0.39.10 `Collection.UnmarshalJSON`. Builds/tests are pinned to Go 1.26.6; changing the pin requires an actual migration/reopen test.

## Verification evidence

`make test`, `make race`, `make test-web`, `make build` and `script/smoke-unified.py` are reproducible local/CI entry points. Check the GitHub Actions run for a given commit for its actual outcome. No cloud DNS, CA, real host, production database or external deployment credentials are required by these checks, and no live provider acceptance is claimed.
