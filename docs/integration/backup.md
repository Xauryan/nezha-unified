# Coordinated backup and restore

`service/unifiedbackup` owns the archive and restore safety boundary. Its Dashboard adapter validates the two databases, stable references and operation journals. `cmd/dashboard/unified_backup.go` provides the single executable's offline commands. An archive is complete only when both layers succeed. All examples and tests use synthetic data; no provider traffic is required.

## Dashboard commands

Stop the Dashboard before creating a backup. Both Dashboard startup and backup creation take the same exclusive data-directory lock; a backup request cannot silently copy a running instance. The configuration file must reside inside the database's parent directory so that the complete configuration is included. The key file must contain exactly 32 bytes and have no group/other permissions. Keep a protected copy separate from the archive.

```sh
nezha-unified --backup-create /backups/unified.nzb -db /var/lib/nezha/sqlite.db -c /var/lib/nezha/config.yaml --backup-key-file /secure/nezha.backup-key
nezha-unified --backup-rehearse /backups/unified.nzb --backup-key-file /secure/nezha.backup-key
nezha-unified --backup-restore /backups/unified.nzb --backup-destination /var/lib/nezha-restored --backup-key-file /secure/nezha.backup-key
nezha-unified --backup-inspect-recovery /var/lib/nezha-restored
nezha-unified --backup-reconcile /var/lib/nezha-restored --backup-evidence /secure/reconciliation.json
```

Restore requires the same Dashboard/module versions recorded by the archive. Its offline checks include SQLite integrity/foreign keys, schema fingerprints, stable resource references and opening the Certimate store with the recovered key. Restore then disables workflow policies, workflows, DDNS bindings, deployment targets and legacy server DDNS execution. It revokes restored run bindings and quarantines ambiguous DNS/deployment operations. No external provider is contacted by these steps.

The recovery inspection command prints a public report containing archive/reference IDs and operations requiring a decision. The reconciliation evidence file must match the archive ID and reference digest and list each reported operation exactly once:

```json
{
  "archive_id": "ID from the recovery report",
  "reference_digest": "digest from the recovery report",
  "boundary_evidence": "Operator's evidence for external changes since the backup boundary",
  "operations": [
    {
      "kind": "kind from the recovery report",
      "id": "ID from the recovery report",
      "outcome": "retired_without_replay",
      "evidence": "Specific observation or decision for this operation"
    }
  ]
}
```

Allowed outcomes are `confirmed_applied`, `confirmed_not_applied`, and `retired_without_replay`. Evidence must not contain tokens, credentials, certificate private keys or raw provider responses. The boundary evidence also covers operations that happened after the backup and therefore cannot appear in the restored database. The application validates completeness and identity of the evidence; the operator remains responsible for the external observations.

Reconciliation preserves an audit file and clears the startup fence. Existing operations are never blindly replayed, and automation remains disabled after the fence is removed. Review each workflow, DDNS binding and deployment target before explicitly enabling it. An operation retired without replay is not reported as an externally successful deployment.

## Consistency boundary

`Create(ctx, archivePath, recoveryKey, Hooks)` requires a `Quiesce` callback. Before returning a `Snapshot`, the host must stop new API writes and task enqueue, stop schedulers, drain in-flight mutation calls, then checkpoint and close **both** Nezha and Certimate database handles. A process-level lock must also exclude another Dashboard or backup writer. Taking two independent live file copies is not a supported snapshot.

`Snapshot.Root` contains the Nezha database, Certimate database, file assets, stable references and cross-store operation journals. Every regular file below that root is inventoried and copied. Its `MasterKey` is read separately and never added as an ordinary archive file. The output archive must be outside the snapshot root.

The host supplies:

| Snapshot field | Meaning |
|---|---|
| `NezhaDB`, `CertimateDB`, `MasterKey` | Distinct, safe relative paths below one snapshot root |
| `ConfigFile` | Dashboard configuration path, verified against the inventory; optional only for generic non-Dashboard callers |
| `Versions` | Fixed application/module versions |
| `Schemas` | Database migration/schema identifiers |
| `ReferenceDigest` | SHA-256 of canonical public owner, grant, resource, version and operation-state references |
| `Boundary` | Public identifier for this coordinated, stopped-writer boundary |

The returned resume callback always runs, including when quiescence partially fails, archive generation fails, or the context is canceled. The package requires this callback; the host decides whether a deliberately offline CLI resumes nothing or reopens a previously running runtime. A resume failure is returned even if archive publication succeeded.

## Confidentiality and integrity

The backup recovery key is 32 random bytes; the package rejects reusing the stored Certimate master key as the recovery key. Supply it from a protected file or secret-management input outside the source data tree. It is not a password; no password derivation or command-line secret value is supported.

Each backup gets a random 256-bit data key and a random backup identity. The recovery key wraps the data key with AES-256-GCM. The data key encrypts the complete tar stream, including the manifest and database bytes, in 64 KiB authenticated chunks. Chunk sequence numbers and the complete header digest are authenticated. An authenticated terminal record detects truncation at a chunk boundary; extra trailing data is rejected. The master key is additionally wrapped with an authenticated `backup ID:master-key` context inside the encrypted manifest.

The manifest records each file's size and SHA-256. Files are hashed before inclusion and while copied; a changed snapshot is rejected. Backup publication uses an atomic hard link and refuses an existing output path. Archive files and restored files have mode `0600`; newly created directories have mode `0700`.

The package rejects absolute paths, traversal, backslashes, duplicate entries, symlinks, special files, an unwrapped master-key entry and undeclared archive entries. Root-relative file operations prevent paths from escaping the selected tree. Missing master keys, either missing database, missing version/schema/reference metadata and pending recovery markers fail closed. File count, individual file size, total size and manifest size are bounded.

These properties protect backup contents; they do not replace application-level encryption of credentials, keys, workflow inputs and provider logs in the live stores.

## Offline restore and rehearsal

`Restore(ctx, archivePath, destination, recoveryKey, RestoreHooks)` requires a new destination. It never overwrites a runtime directory. Authentication, file inventory and per-file hashes must pass before application validation. A failed restore removes only its newly reserved destination.

The callbacks run with no scheduler, queue, network side effects or HTTP listener:

1. `Validate` opens the restored stores offline, checks SQLite integrity, schema/version compatibility, matching key decryption and the stable-reference digest, then closes its handles.
2. `PrepareRecovery` repairs pending cross-store intents and changes unfinished external operations to states requiring reconciliation. It must not retry DNS changes, ACME issuance, reload commands or remote deployments.

Restore leaves `.unified-restore-pending.json` in the destination. The host **must call `CheckReady` before any normal initialization that can mutate storage or enable execution**. The marker blocks startup until `CompleteRecovery` succeeds. Its reconciliation callback must confirm current external outcomes, settle local operation journals and verify the resulting references before clearing the fence. Returning success solely because the databases reopen is insufficient.

`Rehearse` runs the identical restore path in a private temporary directory and removes the rehearsal afterward. Rehearsal callbacks must remain offline. A successful rehearsal proves archive readability and the application's supplied checks; it does not prove that external DNS, CA or target-host state matches an old backup.

## Contract evidence

`service/unifiedbackup/backup_test.go` covers encrypted round-trip of both stores/assets/journals, preserved reference identity, protected file modes, recovery fencing, missing key/database rejection, quiesce/resume failures, cancellation, existing-data preservation, source and destination symlinks, wrong keys, header/chunk/footer tampering, truncation, trailing data, authenticated traversal/duplicate/symlink/raw-key members, file-hash mismatch, mandatory restore hooks and rehearsal cleanup.

Run the package through the pinned Make toolchain, for example:

```sh
make --eval 'test-backup: ; go test -mod=readonly -p 2 -race ./service/unifiedbackup' test-backup
```

`dashboard_test.go` adds disposable SQLite integrity/reference/schema checks, exclusive-lock behavior, restore quarantine, exact recovery evidence and chained backup/recovery-history preservation. The restored configuration disables its original TSDB path so that starting the restored instance cannot silently write the old runtime's time-series directory. Operators must explicitly rehome and enable that path after review. Configured external TSDB files outside the selected data root are rejected as an incomplete backup.

The CLI additionally opens the restored Certimate store offline to validate encrypted synthetic credentials/keys against the restored master key. Production restore requires an operator-selected backup, matching recovery key and operation-specific external reconciliation evidence.
