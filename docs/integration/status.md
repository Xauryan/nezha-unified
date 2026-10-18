# Implementation and acceptance status

The complete scope in `plan.zh-CN.md` remains authoritative. The implementation now includes the integration surfaces below; final exact-commit build, full regression, browser, CI and real-provider acceptance are separate evidence. No release, production deployment or live cloud/CA/host acceptance is claimed by source retention or individual test results.

| Area | Implemented integration | Acceptance evidence still required |
|---|---|---|
| Upstreams | Dashboard/admin/Certimate/SDK/DNS reference and Agent pinned; 5059 retained files; page/API/node/provider/configuration/migration inventory; declared patches; shared Dashboard/Agent protocol equality gate | Final reviewed diff and exact-commit structural check |
| Host runtime | One Dashboard owns both stores, certificate runtime, DDNS and Agent integration; exclusive data lock; cancellation/drain/shutdown and restore fence | Full runtime startup/shutdown/race gates; Agent compatibility stress |
| Identity | One session; legacy PATs rejected for new powers; exact-resource scoped tokens and versioned service identities; current authority at requests/side effects | Final token/session/scoped HTTP and revoke-during-execution regressions |
| Resource authority | Nezha owns bindings/grants/versions; cross-store intents and activation; stable typed IDs; filtered child runs/logs/outputs; explicit transfer | Interruption/reopen/repair tests and existing-installation migration rehearsal |
| Credentials | Authenticated encryption, version/purpose/resource-bound resolver, enable/disable/rotation, explicit export, clone, shared provider forms | Actual credentials against authorized providers |
| Certificate assets | Matching X.509/key import/SAN derivation, encrypted keys/ACME accounts, independent metadata/public/private export, format conversion/revoke APIs | Complete runtime format/revoke tests and real issuance/renewal/revocation |
| Certimate workflows | Full source/providers/nodes/forms; native settings/statistics/run APIs; designer/repositories; scheduling/cache services; versioned run authorization; encrypted state and sanitized logs | Full-provider compilation, local runtime tests, browser interactions and authorized provider operations |
| DNS | Cloudflare, DNSPod legacy/cloud and EdgeOne NS adapters; Zone discovery; CRUD/batch/import/export; provider fields/options; subtree policy and audit/reconciliation | Final fixture tests plus real account/Zone/record/propagation acceptance |
| DDNS | Credential refs, record single writer, current grants, host/interface reports, freshness/IP selection, pause-on-manual-change, device tokens, legacy migration adapter | Fixed Agent report compatibility, real reconnect/propagation and authorized old-data migration |
| Agent deployment | Purpose-scoped Agent call layer; compatible Agent source/capability protocol; durable versioned targets/tasks; staged bundle/check/reload/health/fingerprint; supersession/reconciliation | Exact host+Agent build/tests and live served-fingerprint/rollback acceptance |
| Frontend | One admin build/root/session/theme; adapted Certimate designer/pages/forms; DNS/DDNS/targets/permissions/ownership screens | Final tests/lint/build and browser scenarios |
| Backup/recovery | Encrypted archive with DBs/assets/config/journals; independent wrapped key; hash/schema/reference/version validation; CLI/rehearsal; quarantine and exact operation evidence | Final CLI rehearsal on disposable DBs, then authorized existing-installation recovery |
| CI/release | Standard GitHub-hosted `ubuntu-24.04`, pinned actions/toolchains, unified admin/release build and synthetic checks | Actual complete final-commit CI, release, deployment and customer acceptance |

## Evidence recorded during implementation

- `python3 script/check-integration.py` passed after retaining the exact Agent v2.3.0 snapshot and reconciling intended patch paths: 5059 files, 50 declared patches, licenses, module boundary, capability inventory and shared protocol checks. Later edits require another run.
- `service/unifiedbackup` passed Go 1.26.6 `-race` tests covering archive confidentiality/authentication/path safety, references, independent keys, offline SQLite integrity/ownership validation, exclusive lock, configuration/TSDB path hardening, restore quarantine, exact operation evidence and chained backup/recovery-history preservation. Later code changes require another run before the final report.
- Other module owners run focused synthetic checks during development. Final project-level command/CI results must be recorded against the completed commit; unrun or failing checks are not passes.

The full matrix/provider counts are in `acceptance.md`; backup commands/no-replay recovery are in `backup.md`; resource/token/execution contracts are in `api.md`.

## Architectural boundaries

- `github.com/certimate-go/certimate/pkg/embedded` is the only host import boundary. No Certimate internal imports, PocketBase login/admin/public routes or second management application are exposed.
- Full releases include the `certimate_full` engine and all upstream providers. Explicit storage-only development mode reports unavailable execution and is not the release acceptance target.
- GORM owns users, grants, bindings, policies and intents. Certimate owns records, graph/run state, keys and audit. Local transactions and inactive-until-bound resources prevent partially linked resources from executing.
- Credential config, certificate/ACME keys, graphs, run snapshots, node configuration/output and sensitive settings are encrypted at rest. Provider logs/errors are sanitized before persistence. Metadata access never implies secret export or execution.
- Restored stores are fenced until operation-specific reconciliation completes. Workflows, DDNS and targets stay disabled until explicitly reviewed/re-enabled. Old backups never justify blind remote-side-effect replay.
- Go 1.26.6 preserves PocketBase v0.39.10 JSON behavior. Toolchain changes require actual migration/reopen and complete workflow tests.

## Remaining release gates

Run `make check test`, `make race`, `make test-web`, `make build`, the disposable native HTTP smoke, fixed Agent compatibility and browser scenarios for the final head. Confirm complete remote CI for that same commit. Track authorized real DNS/CA/notification/host acceptance separately; no production credentials belong in tests or CI.
