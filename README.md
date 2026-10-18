# Nezha Unified

以 Nezha Dashboard 为主体，在一个 Go 进程、一个管理前端和一套哪吒登录下整合 DNS、DDNS、凭据与 Certimate 证书业务。

原生集成代码已接入完整工作流引擎、统一 DNS/DDNS、共享部署目标与离线加密备份恢复。源码清单、编译结果、合成契约测试和真实提供商验收分别记录在 [验收台账](docs/integration/acceptance.md)；外部 DNS、CA 和目标主机的实际效果需要独立验收。

## 当前实现

- 保留 Nezha v2.3.18 后端、v2.3.8 管理前端及 Agent v2.3.0 基线和原有功能；一个管理应用、一个 Dashboard、一个 Nezha 会话。
- Certimate v0.4.35 完整引擎、全部保留的提供商、13 类工作流节点、设计器、预设、调度、日志、运行状态和证书业务通过公开 `pkg/embedded` 入口接入。PocketBase 仅用于内部存储。
- 统一凭据中心支持加密存储、版本、启停、轮换与单独授权的导出；证书、ACME 密钥、工作流图、运行快照和部署目标配置加密保存。
- 用户授权、精确资源 Token 和版本化服务身份在请求与后台副作用前重新校验；列表先授权再计数/分页，元数据、凭据使用、私钥导出和部署执行分别授权。
- Cloudflare、DNSPod 和 EdgeOne NS 集中 DNS：Zone 发现、记录操作、批量和导入导出、厂商字段、名称子树委托、DDNS 单写者和未知结果核验。
- 主机 IPv4/IPv6 与网卡来源、设备 Token、旧 DDNS ID/绑定迁移、加密凭据引用、配置版本与执行时授权。
- 兼容 Agent 的能力协商、证书暂存/检查/切换/重载/健康检查/指纹核验，以及使用原 Certimate 提供商的共享云目标；持久任务支持防旧版本覆盖、取消和结果未知时冻结。
- 单个可执行程序提供停机一致加密备份、独立恢复密钥、双库/资产/配置/日志清单校验、离线演练、恢复隔离和逐操作核验。

完整范围和剩余工作见 [实施状态](docs/integration/status.md)；输入方案见 [全量方案](docs/integration/plan.zh-CN.md)（公开文档中的业务域名已改为 example.com）。

## 本地构建

需要 Go 工具链下载能力、Node.js 26.10.0、npm、Python 3、make 和 C 编译器（GORM SQLite 使用 CGO）。

```sh
git clone https://github.com/Xauryan/nezha-unified.git
cd nezha-unified
make build
mkdir -p data
printf 'listen_host: 127.0.0.1\nlisten_port: 8008\n' > data/config.yaml
./bin/nezha-unified -c data/config.yaml -db data/sqlite.db
```

打开 `http://127.0.0.1:8008/dashboard/`。首次创建空库沿用哪吒的 `admin` / `admin` 初始化行为，完成登录后在个人设置修改密码，再配置对外访问。

`make` 会明确选择 **Go 1.26.6**。不要使用系统 Go 1.27 直接构建此基线：其默认 JSON v2 行为与 PocketBase v0.39.10 的 collection 反序列化不兼容，迁移会栈溢出。前端使用已提交的 npm lockfile。

## 验证

```sh
make check       # 上游来源、基线文件、集成边界及 runner 约束
make test        # 哪吒原测试与集成存储/授权测试
make race        # 集成、生命周期、RPC 与 API 控制器 race 检查
make vet         # 完整引擎静态检查
make test-agent build-agent # 固定 Agent 回归与构建
make test-web    # 原管理前端单测与 lint
make build       # 原生管理前端 + 校验过的公开主题 + 单个 Dashboard
python3 script/smoke-unified.py  # 临时库上的 HTTP、工作流和离线恢复
make test-e2e    # 临时 Dashboard + Agent 上的完整浏览器用例
```

GitHub Actions 只使用公开仓库的标准 `ubuntu-24.04` GitHub 托管 runner，不使用 self-hosted/ARC 或 larger runner。初始 CI 不上传 artifact，不启用 Actions cache。

## 数据与边界

哪吒数据仍在 `-db` 指定的位置；同级 `certimate/` 目录存放证书模块的数据库、文件及 `unified.key`。密钥文件权限为 `0600`，目录为 `0700`。使用 [备份恢复命令](docs/integration/backup.md) 在停机并获得独占数据锁后创建一致归档，覆盖两套数据库、资产、配置和版本清单；恢复时先隔离自动执行，再逐操作提供核验记录。已有证书库缺少原密钥时启动会拒绝自动生成替代密钥。

旧 Nezha DDNS 配置提供保留原 ID/绑定的本地凭据迁移，缺失所有者或必要材料时禁止执行。现有独立 Certimate 的 `pb_data` 不能直接复制替代集成库：存量凭据转换、所有权映射和完整数据迁移仍需使用获得授权的真实备份单独演练。任何本地或 CI 通过都不表示生产迁移、CA 签发或远程部署已完成。

[API 文档](docs/integration/api.md) · [来源和许可](THIRD_PARTY_NOTICES.md) · [固定版本](upstream.lock.json)
