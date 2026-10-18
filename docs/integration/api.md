# Native integration API

Session paths begin `/api/v1/unified`. They inherit Nezha session authentication and CSRF protection. Cookie-based writes require `X-CSRF-Token` matching `nz-csrf`; the normal frontend API helper supplies it. Legacy Nezha `nzp_` PATs never imply unified-module permissions. The separate `/api/v1/unified-api` surface accepts only the new scoped `Authorization: Bearer nzu_<public-id>.<secret>` credential and applies the same resource handlers without cookie/JWT/PAT fallback. Responses are never cacheable.

JSON responses use the existing `{ "success": true, "data": ... }` envelope. Resource handler errors use HTTP 400/403/404/500/503 and the upstream error envelope `{ "error": ... }` (`success` is omitted when false), without raw SQL, provider secrets or private keys. Authentication failure retains Nezha's HTTP 200 / missing-or-false `success` / `ApiErrorUnauthorized` contract. JSON bodies are limited to 256 KiB.

| Method/path | Access | Behavior |
|---|---|---|
| GET `/status` | Authenticated principal | Storage readiness and runtime capability flags |
| GET `/credentials` | Resource `credential:read` | Metadata only; never config/envelope |
| POST `/credentials` | Administrator | `{name, provider, config}`; encrypt config and write audit atomically |
| GET `/certificates` | Resource `certificate:read` | Metadata and actual SAN names; no private key |
| POST `/certificates/import` | Administrator | `{certificate, private_key}` PEM pair; validate match and SAN; encrypt private key |
| GET `/certificates/:id/pem` | Resource `certificate:download` | Public PEM chain; does not export private key |
| GET `/workflows` | Resource `workflow:read` | Metadata only; does not return workflow graphs/configuration |
| GET `/grants` | Session | Own grants; administrators see all |
| POST `/grants` | Administrator | `{user_id, kind, resource_id, action}`; requires existing user/resource |
| DELETE `/grants/:id` | Administrator | Revoke and audit atomically |

These metadata lists accept `page` (1–100000, default 1) and `per_page` (1–100, default 20). `data` is `{items,total,page,per_page}`; `total` and pagination use only authorized resource IDs. Resource IDs retain their original type: Certimate string IDs and Nezha numeric IDs are not interchangeable. No wildcard resource scopes are accepted. Metadata read, credential use, export, rotation, public-chain download, private-key download and workflow execution are independent actions.

The `config` request field carries provider-specific JSON; credential creation does not contact a provider. Passing syntax validation does not prove valid credentials or provider connectivity. Certificate imports derive names from parsed SANs and reject CA-only/no-SAN certificates and mismatched keys; administrator imports are not an ACME issuance or a trust/chain validation claim.

## Native Certimate surface

These paths are relative to `/api/v1/unified/certimate` (or the scoped-token equivalent). The adapter does not mount PocketBase routes or issue a PocketBase login session.

| Method/path | Behavior |
|---|---|
| GET `/resources/:collection` | Authorized list: `access`, `certificate`, `workflow`, `workflow_run`, `workflow_output`, `workflow_logs`, `settings` |
| GET `/resources/:collection/:id` | Same resource policy as list; runs/logs/outputs resolve their workflow authority |
| POST `/resources/:collection` | Create through application services, ownership intents and encrypted persistence |
| PATCH `/resources/:collection/:id` | Application-service update; scheduling/settings side effects follow the service |
| DELETE `/resources/:collection/:id` | Authorized removal; active-run deletion is rejected |
| POST `/batch` | Up to 200 `{collection,method,id?,body?}` operations; per-item authorization/commits, returns `[{status,body}]` |
| POST `/credentials/:id/duplicate` | `{name}`; server-side duplication returning metadata only |
| POST `/credentials/:id/export` | Explicit export permission/audit; returns `{config}`, never a metadata field |
| GET `/statistics` | Counts over visible certificates/workflows |
| GET `/workflows/stats` | Authorized dispatcher pending/processing state |
| POST `/workflows/:id/runs` | `{trigger:"manual"}`; current authority and workflow policy; returns `{runId}` |
| POST `/workflows/:id/runs/:run/cancel` | Verify workflow/run association, persist cancellation and stop cancellable work |
| POST `/certificates/:id/download` | Independent private-key permission; `{fileFormat:"PEM"\|"PFX"\|"JKS",...}`; returns base64 `{zipBytes}` |
| POST `/certificates/:id/revoke` | Independent revoke permission and applicable CA service action |
| POST `/notifications/test` | `{accessId,provider}`; current credential use permission and provider call |
| GET `/system/crons` | Restricted scheduler diagnostics |
| GET `/system/logs` | Restricted sanitized logs; `page` and `perPage` |

Native lists use `page` (1–100000), `perPage` (1–200), and validated `filter`, `sort`, `fields`, `expand`. Their envelope is `{page,perPage,totalItems,totalPages,items}`. Authorization and activation filtering occur before count/pagination/expansion. Credential metadata never includes config; graph/settings responses redact sensitive values and preserve them server-side during an authorized partial edit.

Resource writes and duplication accept `Idempotency-Key`, at most 64 ASCII letters, digits, underscores or hyphens. Cross-store creation and cloning bind that key to actor/resource/body through operation intents and reject a different request using the same key. Existing-resource updates use their service's version/authorization checks. An absent key is generated server-side. Batch items are independent commits, not a cross-store atomic transaction.

PEM/PFX/JKS downloads may include private keys; `/certificates/:id/pem` continues to return only a public chain. Format options are `pfxPassword`, `pfxEncoder`, `jksAlias`, `jksKeypass`, `jksStorepass`. Downloads must not be logged or cached. Notification tests and revocation are real external actions and are excluded from synthetic smoke tests.

## Identity, ownership and workflow policy

| Method/path under `/unified` | Contract |
|---|---|
| GET/POST `/tokens`, DELETE `/tokens/:id` | Session user manages own tokens; creation `{name,expires_at,scopes:[{kind,resource_id,action}]}` returns `{token,secret}` once |
| GET/POST `/identities`, PATCH `/identities/:id` | Admin creates/edits service identities; patch requires current `version`; owner can list own |
| GET `/ownership` | Authoritative resource owner/version/activation |
| POST `/ownership/transfer` | `{kind,resource_id,owner_id,expected_version}`; versioned transfer |
| GET `/workflow-policies` | Visible execution policies |
| PUT `/workflow-policies/:id` | Expected policy version plus names, credentials, targets and run identity |

Tokens require a future RFC3339 expiry at most 366 days away. Exact scopes intersect the current user's role, grants, account token version and resource constraints on each request/side effect, including administrator-owned tokens. Expired/revoked tokens and deleted owners fail closed. Token/service callers cannot manage tokens or identities. Formal service identities are explicitly active and versioned; policy/owner changes invalidate captured run authority.

Manual runs, schedules, retries and recovery use the same current-authority checks. Workflow read does not grant edit, execute, certificate download or deployment. The server binds the workflow version, trigger, owner/service identity, policy version, credential versions and target versions to each run. Dynamic SANs and deployment resources are checked at the actual side-effect boundary.

## DNS, DDNS and targets

DNS routes include `/dns/accounts`, `/dns/accounts/:id/refresh`, `/dns/zones`, `/dns/zones/:id/records`, `/dns/zones/:id/changes`, `/dns/zones/:id/options`, `/dns/zones/:id/export`, `/dns/zones/:id/batch`, `/dns/zones/:id/import`, `/dns/policies`, `/dns/operations` and `/dns/operations/:id/reconcile`. Provider-fetched account/zone/record identity and old/new names are checked server-side. Lists/export filter before returning data; batches authorize each item. Unknown external outcomes require reconciliation before replay.

DDNS routes include `/ddns/bindings`, `/ddns/bindings/:id/run`, `/ddns/hosts/:id/addresses` and `/ddns/devices`. The separate `POST /api/v1/unified/ddns/device/update` accepts only its binding-scoped bearer token and `{ip}`. It cannot select another name, account, record type or host and returns no credential/account metadata.

Target routes include `/targets`, `/targets/:id`, `/targets/:id/capabilities`, `POST /targets/:id/deployments`, `/deployments`, `/deployments/:id`, `POST /deployments/:id/cancel`, `POST /deployments/:id/reconcile` and `POST /deployments/:id/retire`. Requests fix `operation_id`, `certificate_id`, `certificate_version` and `target_version`. Target configuration owns paths, commands and health checks; delegated callers cannot override them. Current authorization and Agent capability checks precede calls. A signed certificate and a remotely served certificate remain separate states.

Targets support `kind=nezha` and `kind=certimate`. Agent targets use the versioned host protocol. Certimate targets bind a registered deployment provider, an exact credential version and encrypted provider configuration. Blank configuration on a metadata edit preserves the saved configuration under the new target version; metadata responses never contain it. Execution checks the current target, certificate, credential and initiating identity again. An accepted cloud API operation records `applied`; only an Agent service fingerprint check can record `verified`. An interrupted or uncertain cloud mutation requires explicit reconciliation and is never automatically replayed.

Retirement requires `acknowledge_unknown=true` and the exact `expected_fingerprint`. It records an operator decision to stop tracking/replaying an uncertain operation; it does not assert remote deployment success. Agent reconciliation uses a validated operation/target/version/sequence receipt, preserving newer deployment observations and treating missing first-install recovery material as uncertain.

Legacy `/api/v1/ddns` entry points retain their original profile IDs and server bindings. Startup migration moves provider/webhook material into versioned encrypted credentials and clears the old columns. Profiles with missing owners or material remain fenced. Editing requires the current `config_version`; blank secret fields preserve existing material only for the same provider. Each worker write resolves the exact saved profile and current user authority again, and refuses a record already owned by a unified DDNS binding.

Backups and restore are local operator commands in the same executable (`backup.md`); raw archives and recovery keys are not exposed through resource APIs.
