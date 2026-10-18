# Upstream source and licenses

This repository is an independent integration maintained under Xauryan's personal account. It is not an official Nezha or Certimate distribution.

| Component | Source location | Baseline | License |
|---|---|---|---|
| Nezha Dashboard | Root Go tree | v2.3.18 | Apache-2.0, root `LICENSE` |
| Nezha admin frontend | `web/admin` | v2.3.8 | Apache-2.0, `web/admin/LICENSE` |
| Certimate | `third_party/certimate` | v0.4.35 | MIT, `third_party/certimate/LICENSE` |
| Certimate SDK dependency snapshot | `third_party/certimate/pkg/sdk3rd-forked` | Locked submodule commit | Preserve each contained SDK's license |
| 彩虹 DNS manager | `reference/dnsmgr` | Locked main commit | MIT, `reference/dnsmgr/LICENSE` |

Exact repositories and commit IDs are recorded in `upstream.lock.json`. Baseline file hashes in `docs/upstream/files.json` distinguish retained source from declared integration patches (CRLF is normalized to LF for cross-platform checkouts). Original Nezha README/workflow files are retained under `docs/upstream/nezha`; upstream workflows are not active workflows of this project. Certimate's UI source and all provider source remain in the tree for the full migration, but only `web/admin` is built as the management app.

Public monitoring themes retain the original configured versions. Their release archives are locked by URL and SHA-256. Their sources and licenses remain with the authors identified in `service/singleton/frontend-templates.yaml`; downloaded build assets are not committed as original work.

New integration code is licensed under Apache-2.0 unless it resides in a retained MIT component, where that component's MIT license applies. Existing copyright notices and per-file licensing remain in effect.
