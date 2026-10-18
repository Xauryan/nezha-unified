此目录下存放对依赖包的分叉。通常因为它们位于私有仓库，对 CI/CD 不便；或存在依赖冲突问题。

目录结构与原依赖包的目录结构保持一致，并在模块名结尾追加 `@<version>`。

---

This directory stores forks of third-party dependencies.​ These are typically required when packages reside in private repositories (which complicate CI/CD pipelines), or exhibit dependency conflicts.

Maintains a 1:1 mapping to the original dependency's directory tree, tagging the path with `@<version>`.
