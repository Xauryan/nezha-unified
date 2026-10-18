# Native integration API

All paths begin `/api/v1/unified`. They inherit Nezha session authentication and CSRF protection. Cookie-based writes require `X-CSRF-Token` matching `nz-csrf`; the normal frontend API helper supplies it. PATs are rejected until dedicated resource-scoped claims exist. Responses are never cacheable.

JSON responses use the existing `{ "success": true, "data": ... }` envelope. Resource handler errors use HTTP 400/403/404/500/503 and `{ "success": false, "error": ... }`, without raw SQL, provider secrets or private keys. Authentication failure retains Nezha's HTTP 200 / `success: false` / `ApiErrorUnauthorized` contract. JSON bodies are limited to 256 KiB.

| Method/path | Access | Behavior |
|---|---|---|
| GET `/status` | Session | Storage readiness and explicit disabled capability flags |
| GET `/credentials` | Resource `credential:read` | Metadata only; never config/envelope |
| POST `/credentials` | Administrator | `{name, provider, config}`; encrypt config and write audit atomically |
| GET `/certificates` | Resource `certificate:read` | Metadata and actual SAN names; no private key |
| POST `/certificates/import` | Administrator | `{certificate, private_key}` PEM pair; validate match and SAN; encrypt private key |
| GET `/certificates/:id/pem` | Resource `certificate:download` | Public PEM chain; does not export private key |
| GET `/workflows` | Resource `workflow:read` | Metadata only; does not return workflow graphs/configuration |
| GET `/grants` | Session | Own grants; administrators see all |
| POST `/grants` | Administrator | `{user_id, kind, resource_id, action}`; requires existing user/resource |
| DELETE `/grants/:id` | Administrator | Revoke and audit atomically |

Lists accept `page` (1–100000, default 1) and `per_page` (1–100, default 20). `data` is `{items,total,page,per_page}`; `total` and pagination use only authorized resource IDs. Resource kinds are `certificate`, `credential`, `workflow`; actions are `read`, plus `download` for certificates. Resource IDs are the original typed PocketBase IDs, not Nezha numeric IDs. No wildcard resource grants are accepted.

`config` currently stores provider-specific JSON without contacting a provider. Passing syntax validation does not prove valid credentials or provider connectivity. Certificate imports derive names from parsed SANs and reject CA-only/no-SAN certificates and mismatched keys; administrator imports are not an ACME issuance or a trust/chain validation claim.
