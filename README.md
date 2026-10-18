# Nezha Unified

以 Nezha Dashboard 为主体，在一个 Go 进程、一个管理前端和一套哪吒登录下整合 DNS、DDNS、凭据与 Certimate 证书业务。

**开发中：目前是首批可运行的集成实现，尚未达到全量方案的交付标准。** 自动签发、工作流设计器与调度、集中 DNS 提供商管理、统一 DDNS 和 Agent 证书部署仍待接入，不能把保留上游源码视为完成这些能力。

## 当前可用

- 保留 Nezha v2.3.18 后端、v2.3.8 管理前端源码及原有导航；统一构建管理前端，固定并校验原公开主题。
- 同进程初始化 Certimate v0.4.35 的 PocketBase 内部数据库和迁移，不启动独立端口、管理页面或第二套账号。
- 原生凭据录入和元数据列表；配置使用 AES-256-GCM 加密，密文绑定资源 ID 和版本。
- 原生证书导入：校验证书与私钥匹配，从实际 X.509 SAN 提取名称，私钥加密存储；独立授权下载公开证书链。
- 哪吒账号下的精确资源授权、授权后的计数和分页、撤权、事务内审计；新模块暂不接受 PAT，避免旧 scope 意外扩权。
- DNS 名称授权基础库：IDNA、标签边界、根/后代开关、独立通配权限、账户/Zone/记录身份以及改名前后校验。

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
make race        # 集成模块与 API 控制器 race 检查
make test-web    # 原管理前端单测与 lint
make build      # 原生管理前端 + 校验过的公开主题 + 单个 Dashboard
python3 script/smoke-unified.py  # 临时库上的真实 HTTP 登录及资源流程
```

GitHub Actions 只使用公开仓库的标准 `ubuntu-24.04` GitHub 托管 runner，不使用 self-hosted/ARC 或 larger runner。初始 CI 不上传 artifact，不启用 Actions cache。

## 数据与边界

哪吒数据仍在 `-db` 指定的位置；同级 `certimate/` 目录存放证书模块的数据库、文件及 `unified.key`。密钥文件权限为 `0600`，目录为 `0700`。备份必须同时保留哪吒数据、整个 `certimate/` 目录和版本清单；当前仅支持停机一致备份，在线协调备份/恢复尚待实现。已有证书库缺少原密钥时启动会拒绝自动生成替代密钥。

本版本仅支持新建集成库。将现有独立 Certimate 的 `pb_data` 直接复制进来不属于已验证的迁移路径，不能保证原凭据的加密迁移、授权映射或工作流执行。恢复、轮换、调度和完整提供商适配将在对应验收通过后开放。

[API 文档](docs/integration/api.md) · [来源和许可](THIRD_PARTY_NOTICES.md) · [固定版本](upstream.lock.json)
