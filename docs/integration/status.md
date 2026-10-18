# Implementation and acceptance status

The complete scope in `plan.zh-CN.md` remains authoritative. The native integration is implemented in this repository. Source retention, executable integration, local synthetic validation, exact-commit CI, release/deployment and real provider acceptance are separate evidence. No production deployment or live cloud/CA/host acceptance is claimed by the results below.

| Area | Implemented integration | Validation boundary |
|---|---|---|
| Upstreams | Dashboard/admin/Certimate/SDK/DNS reference and Agent pinned; 5059 retained files, capability inventory and 65 declared patch paths; Dashboard/Agent protocol equality gate | Structural gate checks actual retained hashes, module boundary, inventories and runner labels |
| Host runtime | One Dashboard owns both stores, certificate runtime, DDNS and Agent integration; exclusive data lock; cancellation/drain/shutdown and restore fence; streaming HTTP remains usable without a global request deadline | Synthetic startup/shutdown, HTTP/2 streaming, lifecycle and race tests |
| Identity | One Nezha session; legacy PATs rejected for new powers; exact-resource scoped tokens and versioned service identities; current authority at requests and side effects | Scoped HTTP, revoke-during-execution, ownership/transfer and service-identity regressions |
| Resource authority | Nezha owns bindings/grants/versions; cross-store intents and activation; stable typed IDs; filtered child runs/logs/outputs; explicit transfer | Interruption/reopen/repair and restored-reference tests; existing-installation migration remains a separate rehearsal |
| Credentials | Authenticated encryption, version/purpose/resource-bound resolver, enable/disable/rotation, explicit export, clone and shared provider forms | Synthetic confidentiality and authorization tests; actual credentials still require authorized provider validation |
| Certificate assets | Matching X.509/key import/SAN derivation, encrypted keys/ACME accounts, separate metadata/public/private export, format conversion and revoke APIs | Synthetic crypto/storage/download tests; real issuance/renewal/revocation remains external acceptance |
| Certimate workflows | Full engine/providers/nodes/forms; native settings/statistics/run APIs; designer/repositories; scheduling/cache services; versioned authorization; encrypted state and sanitized logs | Full-provider build, engine/runtime/race tests, native HTTP and browser workflow checks; live challenges/deployment/notifications remain separate |
| DNS | Cloudflare, DNSPod legacy/cloud and EdgeOne NS adapters; Zone discovery; CRUD/batch/import/export; provider fields/options; subtree policy, audit and reconciliation | Loopback provider contracts and authorization tests; real accounts/zones/propagation remain separate |
| DDNS | Credential refs, record single writer, current grants, host/interface reports, freshness/IP selection, pause-on-manual-change, device tokens; ID-preserving legacy credential migration and current host/domain reauthorization before legacy effects | Synthetic DDNS and fixed Agent report tests; live reconnect/propagation and authorized old-data migration remain separate |
| Agent deployment | Compatible retained Agent; durable versioned targets/tasks; staged bundle/check/reload/health/fingerprint; supersession, leases, cancellation and reconciliation | Local executor/protocol/Agent build tests and browser Agent interactions; actual served fingerprint/rollback needs an authorized target |
| Shared cloud deployment | Encrypted target configuration; version-bound credentials and identity; current authorization at execution; API acceptance recorded separately from observed fingerprint; unresolved operations freeze the target until reconciled or explicitly retired | Synthetic native-provider and durable-state/race tests; real remote upload/result remains separate |
| Frontend | One admin build/root/session/theme; complete adapted designer/pages/forms; DNS/DDNS/targets/permissions/ownership screens; all emitted assets embedded | 51 unit tests, lint/type/build and all 13 browser cases, including session preservation, mobile/deep links, CSRF and actual loopback Agent terminal/file interactions and designer duplicate/drag/delete/save/reload/history navigation |
| Backup/recovery | Encrypted DB/assets/config/journal archive with independent wrapped key; hash/schema/reference/version validation; offline CLI/rehearsal; quarantine and exact operation evidence; retained host schema used to disable legacy DDNS | Crypto/path/race tests and disposable CLI create/rehearse/restore/reconcile; existing-installation and external-outcome reconciliation remain separate |
| CI/release | Standard GitHub-hosted `ubuntu-24.04` only; pinned toolchains/actions; unified admin/release build and synthetic checks | Exact-revision CI is reported by PR #1; tags, artifacts, production deployment and customer acceptance are independent gates |

## Local validation record

The completed source tree was validated on 2026-10-18 with the pinned Go 1.26.6 Make toolchain. The reproducible gates are:

```sh
make check test
make race
make vet
make test-agent build-agent
make test-web
make test-e2e
python3 script/smoke-unified.py
```

`make test-e2e` depends on the full `make build` and `make build-agent` targets. It launches an isolated loopback Dashboard and the retained compatibility Agent, runs every browser test, and cleans up its disposable processes/data. The HTTP smoke also creates, rehearses, restores and reconciles an encrypted backup without provider traffic. A failed or interrupted command is not evidence for its later stages; the final PR/CI report records the actual outcomes for the pushed revision.

The designer deletion regression was first reproduced against the prior build: its delayed hover callback accessed a removed node. The current designer cancels node hover timers on disposal/unmount and ignores disposed-node mouse events. The browser case retains duplicate parameters, drags the copy, deletes the original, saves/reloads the graph and navigates backward/forward without browser errors.

Frontend lint succeeds with inherited hook-dependency warnings in the retained designer components. The build also reports large-bundle warnings; neither warning class is represented as an error-free lint/build log.

The full inventory/provider counts are in `acceptance.md`; backup commands and no-replay recovery are in `backup.md`; resource/token/execution contracts are in `api.md`.

## Architectural boundaries

- `github.com/certimate-go/certimate/pkg/embedded` is the only host import boundary. No Certimate internal imports, PocketBase login/admin/public routes or second management application are exposed.
- Full releases include `certimate_full` and all retained upstream providers. Explicit storage-only development mode reports unavailable execution and is not the release acceptance target.
- GORM owns users, grants, bindings, policies and intents. Certimate owns records, graph/run state, keys and audit. Local transactions and inactive-until-bound resources prevent partially linked resources from executing.
- Credentials, certificate/ACME keys, graphs, run snapshots, node configuration/output and sensitive settings are encrypted at rest. Provider logs/errors are sanitized before persistence. Metadata access never implies secret export or execution.
- Restored stores are fenced until operation-specific reconciliation completes. Workflows, DDNS and targets stay disabled until explicitly reviewed/re-enabled. Old backups never justify blind remote-side-effect replay.
- Go 1.26.6 preserves the retained PocketBase v0.39.10 behavior. Toolchain changes require migration/reopen and complete workflow tests.

## Remaining release and external acceptance gates

Review the final diff and exact-revision remote CI before merging or publishing a release. The implementation does not authorize production rollout. Separately record authorized real DNS/CA/notification/cloud-target/host acceptance and an existing-installation migration/restore rehearsal. Production credentials do not belong in tests, fixtures, CI or logs.
