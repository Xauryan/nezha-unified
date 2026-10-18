import { FetcherMethod as M, fetcher, swrFetcher } from "@/api/api"
import { Button } from "@/components/ui/button"
import {
    Check,
    ErrorNotice,
    Field,
    JSONField,
    Panel,
    UnifiedNav,
    parseArray,
    useAction,
    useText,
} from "@/components/unified/common"
import { useMainStore } from "@/hooks/useMainStore"
import { useState } from "react"
import useSWR from "swr"

type Target = {
    id: number
    name: string
    kind: string
    server_id: number
    allowed_names: string[]
    bundle_path: string
    file_mode: number
    owner_uid?: number
    owner_gid?: number
    check_command: string[]
    reload_command: string[]
    health_command: string[]
    tls_address: string
    tls_server_name: string
    version: number
    enabled: boolean
    desired_version: number
    applied_version: number
}
type Deployment = {
    id: number
    operation_id: string
    target_id: number
    certificate_id: string
    state: string
    code?: string
    created_at: string
    fingerprint?: string
    receipt?: {
        fingerprint?: string
        verified?: boolean
        rolled_back?: boolean
        phase?: string
        certificate_observed?: boolean
    }
}
const initial = {
    name: "",
    server_id: "",
    allowed_names: "",
    bundle_path: "",
    file_mode: "384",
    owner_uid: "",
    owner_gid: "",
    check_command: "[]",
    reload_command: "[]",
    health_command: "[]",
    tls_address: "",
    tls_server_name: "",
    enabled: true,
}
export default function TargetsPage() {
    const t = useText(),
        admin = useMainStore((s) => s.profile?.role === 0),
        action = useAction()
    const { data, mutate, error } = useSWR<Target[]>("/api/v1/unified/targets", swrFetcher)
    const {
        data: deployments,
        mutate: refreshDeployments,
        error: deploymentError,
    } = useSWR<Deployment[]>("/api/v1/unified/deployments", swrFetcher, { refreshInterval: 5000 })
    const [form, setForm] = useState(initial),
        [edit, setEdit] = useState<Target>(),
        [selected, setSelected] = useState<Target>(),
        [certificate, setCertificate] = useState("")
    const { data: capabilities, error: capabilityError } = useSWR<{
        online: boolean
        protocol: number
        certificate_deploy: boolean
    }>(selected ? `/api/v1/unified/targets/${selected.id}/capabilities` : null, swrFetcher)
    const field = (name: keyof typeof initial, value: string | boolean) =>
        setForm((p) => ({ ...p, [name]: value }))
    function startEdit(target: Target) {
        setEdit(target)
        setForm({
            name: target.name,
            server_id: String(target.server_id),
            allowed_names: target.allowed_names.join("\n"),
            bundle_path: target.bundle_path,
            file_mode: String(target.file_mode || 384),
            owner_uid: target.owner_uid == null ? "" : String(target.owner_uid),
            owner_gid: target.owner_gid == null ? "" : String(target.owner_gid),
            check_command: JSON.stringify(target.check_command || []),
            reload_command: JSON.stringify(target.reload_command || []),
            health_command: JSON.stringify(target.health_command || []),
            tls_address: target.tls_address || "",
            tls_server_name: target.tls_server_name || "",
            enabled: target.enabled,
        })
    }
    return (
        <main className="space-y-6 px-4 py-8">
            <UnifiedNav />
            <h1 className="text-2xl font-semibold">{t("部署目标", "Deployment targets")}</h1>
            <p className="text-muted-foreground">
                {t(
                    "关联哪吒主机、证书名称与部署目录，查看 Agent 能力和部署结果。",
                    "Associate Nezha hosts, certificate names and deployment directories; inspect Agent capabilities and deployment results.",
                )}
            </p>
            <ErrorNotice error={action.error || error || deploymentError} />
            {admin && (
                <Panel
                    title={
                        edit
                            ? t("编辑部署目标", "Edit deployment target")
                            : t("新增部署目标", "Create deployment target")
                    }
                    open={!!edit}
                >
                    <form
                        className="grid gap-4 sm:grid-cols-2"
                        onSubmit={(e) => {
                            e.preventDefault()
                            void action.run(async () => {
                                await fetcher(
                                    edit ? M.PATCH : M.POST,
                                    `/api/v1/unified/targets${edit ? "/" + edit.id : ""}`,
                                    {
                                        ...form,
                                        kind: "nezha",
                                        server_id: Number(form.server_id),
                                        file_mode: Number(form.file_mode),
                                        owner_uid:
                                            form.owner_uid === ""
                                                ? undefined
                                                : Number(form.owner_uid),
                                        owner_gid:
                                            form.owner_gid === ""
                                                ? undefined
                                                : Number(form.owner_gid),
                                        allowed_names: form.allowed_names
                                            .split(/[\n,]+/)
                                            .map((s) => s.trim())
                                            .filter(Boolean),
                                        check_command: parseArray<string>(form.check_command),
                                        reload_command: parseArray<string>(form.reload_command),
                                        health_command: parseArray<string>(form.health_command),
                                        version: edit?.version,
                                    },
                                )
                                setForm(initial)
                                setEdit(undefined)
                                await mutate()
                            })
                        }}
                    >
                        <Field
                            id="target-name"
                            label={t("名称", "Name")}
                            value={form.name}
                            onChange={(v) => field("name", v)}
                            required
                        />
                        <Field
                            id="target-server"
                            label={t("主机 ID", "Host ID")}
                            type="number"
                            min={1}
                            value={form.server_id}
                            onChange={(v) => field("server_id", v)}
                            required
                        />
                        <Field
                            id="target-names"
                            label={t("允许的证书名称", "Allowed certificate names")}
                            value={form.allowed_names}
                            onChange={(v) => field("allowed_names", v)}
                            required
                            help={t(
                                "用逗号分隔；通配符名称必须明确列出。",
                                "Comma-separated; wildcard names must be explicitly listed.",
                            )}
                        />
                        <Field
                            id="target-bundle"
                            label={t("部署目录", "Deployment directory")}
                            value={form.bundle_path}
                            onChange={(v) => field("bundle_path", v)}
                            required
                            help={t(
                                "绝对路径；Agent 原子切换 current 目录中的证书和私钥。",
                                "Absolute path; Agent atomically switches certificate and key under current.",
                            )}
                        />

                        <label className="space-y-2">
                            <span className="block">{t("文件权限", "File permissions")}</span>
                            <select
                                className="h-11 w-full rounded-md border bg-background px-3"
                                value={form.file_mode}
                                onChange={(e) => field("file_mode", e.target.value)}
                            >
                                <option value="384">0600</option>
                                <option value="416">0640</option>
                            </select>
                        </label>
                        <Field
                            id="target-uid"
                            label={t("所属用户 UID（可选）", "Owner UID (optional)")}
                            type="number"
                            min={0}
                            value={form.owner_uid}
                            onChange={(v) => field("owner_uid", v)}
                        />
                        <Field
                            id="target-gid"
                            label={t("所属组 GID（可选）", "Owner GID (optional)")}
                            type="number"
                            min={0}
                            value={form.owner_gid}
                            onChange={(v) => field("owner_gid", v)}
                        />
                        <Check
                            label={t("启用", "Enabled")}
                            checked={form.enabled}
                            onChange={(v) => field("enabled", v)}
                        />
                        {(["check_command", "reload_command", "health_command"] as const).map(
                            (name, index) => (
                                <JSONField
                                    key={name}
                                    id={name}
                                    label={
                                        [
                                            t("配置检查命令参数", "Config check argv"),
                                            t("重载命令参数", "Reload argv"),
                                            t("健康检查命令参数", "Health check argv"),
                                        ][index]
                                    }
                                    value={form[name]}
                                    onChange={(v) => field(name, v)}
                                    help={t(
                                        "必填 JSON 参数数组，首项为可执行文件绝对路径；支持 {directory}、{certificate}、{key}",
                                        "Required JSON argv; first item is an absolute executable. Supports {directory}, {certificate}, {key}.",
                                    )}
                                />
                            ),
                        )}
                        <Field
                            id="target-tls"
                            label={t("TLS 验证地址", "TLS verification address")}
                            value={form.tls_address}
                            onChange={(v) => field("tls_address", v)}
                            help="example.com:443"
                        />
                        <Field
                            id="target-sni"
                            label="TLS SNI"
                            value={form.tls_server_name}
                            onChange={(v) => field("tls_server_name", v)}
                        />
                        <div className="flex gap-2 sm:col-span-2">
                            <Button disabled={action.busy}>{t("保存", "Save")}</Button>
                            {edit && (
                                <Button
                                    type="button"
                                    variant="outline"
                                    onClick={() => {
                                        setEdit(undefined)
                                        setForm(initial)
                                    }}
                                >
                                    {t("取消编辑", "Cancel edit")}
                                </Button>
                            )}
                        </div>
                    </form>
                </Panel>
            )}
            {!data && !error && <p role="status">{t("正在加载…", "Loading…")}</p>}
            {data?.length === 0 && (
                <p className="rounded-lg border border-dashed p-8 text-center">
                    {t("还没有可见的部署目标。", "No deployment targets are available.")}
                </p>
            )}
            <div className="space-y-3">
                {data?.map((target) => (
                    <article className="rounded-lg border p-4" key={target.id}>
                        <div className="flex flex-wrap items-center justify-between gap-3">
                            <div>
                                <h2 className="font-medium">
                                    {target.name}{" "}
                                    <span className="text-muted-foreground">#{target.id}</span>
                                </h2>
                                <p className="mt-1 text-sm text-muted-foreground">
                                    {t("主机", "Host")} #{target.server_id} · {t("版本", "Version")}{" "}
                                    {target.version} ·{" "}
                                    {target.enabled ? t("启用", "Enabled") : t("停用", "Disabled")}
                                </p>
                                <p className="mt-1 break-all font-mono text-xs">
                                    {target.allowed_names.join(", ")} · {target.bundle_path}
                                </p>
                                <p className="mt-1 text-xs">
                                    {t("期望 / 已应用", "Desired / applied")}:{" "}
                                    {target.desired_version || 0} / {target.applied_version || 0}
                                </p>
                            </div>
                            <div className="flex flex-wrap gap-2">
                                <Button variant="outline" onClick={() => setSelected(target)}>
                                    {t("检查与部署", "Inspect & deploy")}
                                </Button>
                                {admin && (
                                    <>
                                        <Button variant="outline" onClick={() => startEdit(target)}>
                                            {t("编辑", "Edit")}
                                        </Button>
                                        <Button
                                            variant="outline"
                                            disabled={action.busy}
                                            onClick={() =>
                                                void action.run(async () => {
                                                    await fetcher(
                                                        M.DELETE,
                                                        `/api/v1/unified/targets/${target.id}`,
                                                    )
                                                    if (selected?.id === target.id)
                                                        setSelected(undefined)
                                                    await mutate()
                                                })
                                            }
                                        >
                                            {t("删除", "Delete")}
                                        </Button>
                                    </>
                                )}
                            </div>
                        </div>
                    </article>
                ))}
            </div>
            {selected && (
                <section className="space-y-4 rounded-lg border p-4">
                    <h2 className="font-semibold">
                        {selected.name} · {t("部署证书", "Deploy certificate")}
                    </h2>
                    <ErrorNotice error={capabilityError} />
                    <p role="status">
                        {!capabilities
                            ? t("正在检查 Agent…", "Checking Agent…")
                            : !capabilities.online
                              ? t("Agent 离线", "Agent offline")
                              : capabilities.certificate_deploy
                                ? t("Agent 支持证书部署", "Agent supports certificate deployment")
                                : t(
                                      "Agent 不支持证书部署，请升级 Agent。",
                                      "Agent does not support certificate deployment; upgrade the Agent.",
                                  )}
                    </p>
                    <form
                        className="grid gap-3 sm:grid-cols-2"
                        onSubmit={(e) => {
                            e.preventDefault()
                            void action.run(async () => {
                                await fetcher(
                                    M.POST,
                                    `/api/v1/unified/targets/${selected.id}/deployments`,
                                    {
                                        operation_id: crypto.randomUUID(),
                                        certificate_id: certificate,
                                        target_version: selected.version,
                                    },
                                )
                                await refreshDeployments()
                            })
                        }}
                    >
                        <Field
                            id="deploy-cert"
                            label={t("证书 ID", "Certificate ID")}
                            value={certificate}
                            onChange={setCertificate}
                            required
                        />
                        <Button
                            disabled={
                                action.busy ||
                                !capabilities?.online ||
                                !capabilities.certificate_deploy ||
                                !selected.enabled
                            }
                        >
                            {t("开始部署", "Start deployment")}
                        </Button>
                    </form>
                </section>
            )}
            <section className="space-y-3">
                <h2 className="text-lg font-semibold">{t("部署记录", "Deployment history")}</h2>
                {deployments?.map((d) => (
                    <article key={d.id} className="space-y-2 rounded-lg border p-4">
                        <div className="flex flex-wrap justify-between gap-3">
                            <div>
                                <p>
                                    #{d.id} · {t("目标", "Target")} #{d.target_id} · {d.state}
                                </p>
                                <p className="font-mono text-xs text-muted-foreground">
                                    {d.certificate_id} · {d.created_at}
                                </p>
                            </div>
                            <div className="flex gap-2">
                                {["cancel", "reconcile"].map((op) => (
                                    <Button
                                        key={op}
                                        variant="outline"
                                        disabled={action.busy}
                                        onClick={() =>
                                            void action.run(async () => {
                                                await fetcher(
                                                    M.POST,
                                                    `/api/v1/unified/deployments/${d.id}/${op}`,
                                                )
                                                await refreshDeployments()
                                            })
                                        }
                                    >
                                        {op === "cancel"
                                            ? t("取消", "Cancel")
                                            : t("核验结果", "Reconcile")}
                                    </Button>
                                ))}
                            </div>
                        </div>
                        {d.code && <p className="text-destructive">{d.code}</p>}
                        {d.fingerprint && (
                            <p className="break-all text-xs">
                                {t("期望指纹", "Expected fingerprint")}: {d.fingerprint}
                            </p>
                        )}
                        {d.receipt?.certificate_observed && (
                            <p>
                                {t(
                                    "已通过 TLS 验证所部署的证书",
                                    "Deployed certificate observed over TLS",
                                )}
                            </p>
                        )}
                        {d.receipt?.verified && !d.receipt.certificate_observed && (
                            <p>{t("配置的健康检查通过", "Configured health check passed")}</p>
                        )}
                        {d.receipt?.fingerprint && (
                            <p className="break-all text-xs">
                                {t("回执指纹", "Receipt fingerprint")}: {d.receipt?.fingerprint}
                            </p>
                        )}
                    </article>
                ))}
            </section>
        </main>
    )
}
