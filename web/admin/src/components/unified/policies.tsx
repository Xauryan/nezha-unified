import { FetcherMethod as M, fetcher, swrFetcher } from "@/api/api"
import { Button } from "@/components/ui/button"
import { useState } from "react"
import useSWR from "swr"

import {
    Check,
    ErrorNotice,
    Field,
    JSONField,
    Panel,
    parseArray,
    parseObject,
    useAction,
    useText,
} from "./common"

type Binding = {
    kind: string
    resource_id: string
    owner_id: number
    active: boolean
    version: number
}
type Policy = {
    workflow_id: string
    version: number
    run_as_user_id: number
    service_id: number
    identity_version: number
    enabled: boolean
    allow_inline: boolean
    allow_import: boolean
    certificate_read: boolean
    certificate_download: boolean
    names: unknown[]
    credentials: Record<string, number>
    targets: Record<string, number>
}
const initial = {
    workflow_id: "",
    expected_version: "0",
    run_as_user_id: "",
    service_id: "",
    identity_version: "0",
    enabled: false,
    allow_inline: false,
    allow_import: false,
    certificate_read: false,
    certificate_download: false,
    names: "[]",
    credentials: "{}",
    targets: "{}",
}
export default function ExecutionPolicies() {
    const t = useText(),
        action = useAction(),
        [selected, setSelected] = useState<Binding>(),
        [newOwner, setNewOwner] = useState(""),
        [form, setForm] = useState(initial)
    const {
        data: bindings,
        mutate: refreshBindings,
        error,
    } = useSWR<Binding[]>("/api/v1/unified/ownership", swrFetcher)
    const {
        data: policies,
        mutate: refreshPolicies,
        error: policyError,
    } = useSWR<Policy[]>("/api/v1/unified/workflow-policies", swrFetcher)
    const field = (key: keyof typeof initial, value: string | boolean) =>
        setForm((p) => ({ ...p, [key]: value }))
    return (
        <>
            <ErrorNotice error={action.error || error || policyError} />
            <Panel title={t("资源归属", "Resource ownership")}>
                <div className="overflow-x-auto">
                    <table className="w-full text-left text-sm">
                        <thead>
                            <tr>
                                <th className="p-2">{t("资源", "Resource")}</th>
                                <th className="p-2">{t("所属用户", "Owner")}</th>
                                <th className="p-2">{t("状态 / 版本", "Status / version")}</th>
                                <th className="p-2">{t("操作", "Actions")}</th>
                            </tr>
                        </thead>
                        <tbody>
                            {bindings?.map((b) => (
                                <tr className="border-t" key={b.kind + ":" + b.resource_id}>
                                    <td className="p-2">
                                        {b.kind}
                                        <p className="font-mono text-xs">{b.resource_id}</p>
                                    </td>
                                    <td className="p-2">#{b.owner_id}</td>
                                    <td className="p-2">
                                        {b.active ? t("启用", "Active") : t("停用", "Inactive")} · v
                                        {b.version}
                                    </td>
                                    <td className="p-2">
                                        <Button
                                            variant="outline"
                                            onClick={() => {
                                                setSelected(b)
                                                setNewOwner("")
                                            }}
                                        >
                                            {t("转移归属", "Transfer ownership")}
                                        </Button>
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </div>
                {selected && (
                    <form
                        className="space-y-3 rounded-md border p-4"
                        onSubmit={(e) => {
                            e.preventDefault()
                            void action.run(async () => {
                                await fetcher(M.POST, "/api/v1/unified/ownership/transfer", {
                                    kind: selected.kind,
                                    resource_id: selected.resource_id,
                                    owner_id: Number(newOwner),
                                    expected_version: selected.version,
                                })
                                setSelected(undefined)
                                await refreshBindings()
                                await refreshPolicies()
                            })
                        }}
                    >
                        <p>
                            {selected.kind} / {selected.resource_id}
                        </p>
                        <Field
                            id="ownership-user"
                            label={t("新的所属用户 ID", "New owner user ID")}
                            type="number"
                            min={1}
                            value={newOwner}
                            onChange={setNewOwner}
                            required
                        />
                        <p className="text-sm text-muted-foreground">
                            {t(
                                "转移后旧工作流执行授权需要重新核验。",
                                "Previous workflow execution authority must be revalidated after transfer.",
                            )}
                        </p>
                        <div className="flex gap-2">
                            <Button disabled={action.busy}>{t("确认转移", "Transfer")}</Button>
                            <Button
                                type="button"
                                variant="outline"
                                onClick={() => setSelected(undefined)}
                            >
                                {t("取消", "Cancel")}
                            </Button>
                        </div>
                    </form>
                )}
            </Panel>
            <Panel title={t("工作流执行策略", "Workflow execution policies")}>
                <p className="text-sm text-muted-foreground">
                    {t(
                        "将已保存的工作流图绑定到执行身份、域名、凭据版本与目标版本。策略变更采用版本检查。",
                        "Bind the saved workflow graph to an identity, names, credential versions and target versions. Policy updates are version checked.",
                    )}
                </p>
                <div className="space-y-2">
                    {policies?.map((p) => (
                        <div
                            key={p.workflow_id}
                            className="flex flex-wrap items-center justify-between gap-2 rounded-md border p-3"
                        >
                            <span className="font-mono text-xs">
                                {p.workflow_id} · v{p.version} ·{" "}
                                {p.enabled ? t("启用", "Enabled") : t("停用", "Disabled")}
                            </span>
                            <Button
                                variant="outline"
                                onClick={() =>
                                    setForm({
                                        workflow_id: p.workflow_id,
                                        expected_version: String(p.version),
                                        run_as_user_id: String(p.run_as_user_id),
                                        service_id: String(p.service_id || ""),
                                        identity_version: String(p.identity_version || 0),
                                        enabled: p.enabled,
                                        allow_inline: p.allow_inline,
                                allow_import: p.allow_import,
                                        certificate_read: p.certificate_read,
                                        certificate_download: p.certificate_download,
                                        names: JSON.stringify(p.names || [], null, 2),
                                        credentials: JSON.stringify(p.credentials || {}, null, 2),
                                        targets: JSON.stringify(p.targets || {}, null, 2),
                                    })
                                }
                            >
                                {t("编辑", "Edit")}
                            </Button>
                        </div>
                    ))}
                </div>
                <form
                    className="space-y-4"
                    onSubmit={(e) => {
                        e.preventDefault()
                        void action.run(async () => {
                            await fetcher(
                                M.PUT,
                                `/api/v1/unified/workflow-policies/${encodeURIComponent(form.workflow_id)}`,
                                {
                                    expected_version: Number(form.expected_version),
                                    run_as_user_id: Number(form.run_as_user_id),
                                    service_id: Number(form.service_id),
                                    identity_version: Number(form.identity_version),
                                    enabled: form.enabled,
                                    allow_inline: form.allow_inline,
                                allow_import: form.allow_import,
                                    certificate_read: form.certificate_read,
                                    certificate_download: form.certificate_download,
                                    names: parseArray(form.names),
                                    credentials: parseObject(form.credentials),
                                    targets: parseObject(form.targets),
                                },
                            )
                            await refreshPolicies()
                            setForm(initial)
                        })
                    }}
                >
                    <div className="grid gap-4 sm:grid-cols-2">
                        <Field
                            id="policy-workflow"
                            label={t("工作流 ID", "Workflow ID")}
                            value={form.workflow_id}
                            onChange={(v) => field("workflow_id", v)}
                            required
                        />
                        <Field
                            id="policy-version"
                            label={t(
                                "当前策略版本（新建填 0）",
                                "Current policy version (0 for new)",
                            )}
                            type="number"
                            min={0}
                            value={form.expected_version}
                            onChange={(v) => field("expected_version", v)}
                            required
                        />
                        <Field
                            id="policy-user"
                            label={t("执行用户 ID", "Run as user ID")}
                            type="number"
                            min={1}
                            value={form.run_as_user_id}
                            onChange={(v) => field("run_as_user_id", v)}
                            required
                        />
                        <Field
                            id="policy-identity"
                            label={t("执行身份 ID（可选）", "Execution identity ID (optional)")}
                            value={form.service_id}
                            onChange={(v) => field("service_id", v)}
                        />
                        <Field
                            id="policy-identity-version"
                            label={t("身份版本", "Identity version")}
                            type="number"
                            min={0}
                            value={form.identity_version}
                            onChange={(v) => field("identity_version", v)}
                        />
                    </div>
                    <JSONField
                        id="policy-names"
                        label={t("允许的证书名称范围", "Allowed certificate name scopes")}
                        value={form.names}
                        onChange={(v) => field("names", v)}
                        help={
                            '[{"root":"example.com","include_root":true,"descendants":false,"allow_wildcard":false}]'
                        }
                    />
                    <div className="grid gap-4 sm:grid-cols-2">
                        <JSONField
                            id="policy-credentials"
                            label={t("凭据 ID 与固定版本", "Credential IDs and pinned versions")}
                            value={form.credentials}
                            onChange={(v) => field("credentials", v)}
                            help={'{"credential-id":1}'}
                        />
                        <JSONField
                            id="policy-targets"
                            label={t("部署目标 ID 与固定版本", "Target IDs and pinned versions")}
                            value={form.targets}
                            onChange={(v) => field("targets", v)}
                            help={'{"target-id":1}'}
                        />
                    </div>
                    <div className="grid gap-2 sm:grid-cols-2">
                        {(
                            [
                                "enabled",
                                "allow_inline",
                                "allow_import",
                                "certificate_read",
                                "certificate_download",
                            ] as const
                        ).map((key, i) => (
                            <Check
                                key={key}
                                label={
                                    [
                                        t("启用执行", "Enable execution"),
                                        t("允许内联材料", "Allow inline material"),
                                    t("允许导入证书", "Allow certificate import"),
                                        t("允许证书读取", "Allow certificate read"),
                                        t("允许证书下载", "Allow certificate download"),
                                    ][i]
                                }
                                checked={form[key]}
                                onChange={(v) => field(key, v)}
                            />
                        ))}
                    </div>
                    <Button disabled={action.busy}>{t("保存策略", "Save policy")}</Button>
                </form>
            </Panel>
        </>
    )
}
