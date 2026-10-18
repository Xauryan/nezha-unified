import { FetcherMethod, fetcher, swrFetcher } from "@/api/api"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Textarea } from "@/components/ui/textarea"
import { Check, ErrorNotice, UnifiedNav } from "@/components/unified/common"
import { useMainStore } from "@/hooks/useMainStore"
import { type FormEvent, useState } from "react"
import { useTranslation } from "react-i18next"
import useSWR from "swr"

type Account = {
    id: number
    name: string
    provider: string
    credential_id: string
    credential_version: number
    enabled: boolean
}
type Zone = { id: number; account_id: number; name: string; mode: string }
type RecordValue = {
    id?: string
    name: string
    type: string
    value: string
    ttl: number
    proxied?: boolean
    priority?: number
    line?: string
    weight?: number
    enabled?: boolean
    comment?: string
    data?: Record<string, unknown>
    tags?: string[]
}
type RecordPage = {
    items: RecordValue[]
    total: number
    capabilities: {
        types: string[]
        fields: string[]
        min_ttl: number
        ns_only: boolean
        auto_ttl: boolean
    }
}
type Operation = {
    id: number
    name: string
    action: string
    state: string
    error?: string
    created_at: string
}
type DNSGrant = {
    id: number
    user_id: number
    service_id?: number
    zone_id: number
    root: string
    include_root: boolean
    descendants: boolean
    allow_wildcard: boolean
    types: string[]
    actions: string[]
}

const initialRecord: RecordValue = { name: "", type: "A", value: "", ttl: 600 }

export default function DNSPage() {
    const { i18n } = useTranslation()
    const zh = i18n.language.startsWith("zh")
    const text = (cn: string, en: string) => (zh ? cn : en)
    const admin = useMainStore((s) => s.profile?.role === 0)
    const { data: accounts, mutate: refreshAccounts } = useSWR<Account[]>(
        "/api/v1/unified/dns/accounts",
        swrFetcher,
    )
    const { data: zones, mutate: refreshZones } = useSWR<Zone[]>(
        "/api/v1/unified/dns/zones",
        swrFetcher,
    )
    const { data: operations, mutate: refreshOperations } = useSWR<Operation[]>(
        "/api/v1/unified/dns/operations",
        swrFetcher,
    )
    const { data: grants, mutate: refreshGrants } = useSWR<DNSGrant[]>(
        "/api/v1/unified/dns/policies",
        swrFetcher,
    )
    const [zone, setZone] = useState("")
    const [page, setPage] = useState(1)
    const [pauseDDNS, setPauseDDNS] = useState(false)
    const { data: options, error: optionsError } = useSWR<{
        capabilities: RecordPage["capabilities"]
        lines: { id: string; name: string }[]
        warnings: string[]
    }>(zone ? `/api/v1/unified/dns/zones/${zone}/options` : null, swrFetcher)
    const {
        data,
        error,
        isLoading,
        mutate: refreshRecords,
    } = useSWR<RecordPage>(
        zone ? `/api/v1/unified/dns/zones/${zone}/records?page=${page}&per_page=50` : null,
        swrFetcher,
    )
    const [busy, setBusy] = useState(false)
    const [message, setMessage] = useState("")
    const [account, setAccount] = useState({
        name: "",
        provider: "cloudflare",
        credential_id: "",
        credential_version: 1,
    })
    const [record, setRecord] = useState<RecordValue>(initialRecord)
    const [advanced, setAdvanced] = useState("{}")
    const [editing, setEditing] = useState(false)
    const [importJSON, setImportJSON] = useState("")
    const [grant, setGrant] = useState({
        user_id: "",
        service_id: "",
        root: "",
        include_root: false,
        descendants: true,
        allow_wildcard: false,
        types: "A,AAAA,TXT,CNAME",
        actions: "read,create,update,delete",
    })

    async function action(work: () => Promise<void>) {
        setBusy(true)
        setMessage("")
        try {
            await work()
        } catch (err) {
            setMessage(err instanceof Error ? err.message : text("操作失败", "Operation failed"))
        } finally {
            setBusy(false)
        }
    }
    async function change(kind: string, value: RecordValue) {
        const result = await fetcher<Operation>(
            FetcherMethod.POST,
            `/api/v1/unified/dns/zones/${zone}/changes`,
            {
                operation_key: crypto.randomUUID(),
                action: kind,
                record: value,
                pause_ddns: pauseDDNS,
            },
        )
        await refreshOperations()
        if (result.state !== "succeeded")
            throw new Error(
                `${text("操作", "Operation")} #${result.id}: ${result.state}. ${text("请先核验提供商实际结果，勿重复提交。", "Reconcile provider state before submitting again.")}`,
            )
        await refreshRecords()
    }
    function beginEdit(value: RecordValue) {
        const { id, name, type, value: content, ttl, ...extras } = value
        setRecord({ id, name, type, value: content, ttl })
        setAdvanced(JSON.stringify(extras, null, 2))
        setEditing(true)
    }
    function saveRecord(event: FormEvent) {
        event.preventDefault()
        void action(async () => {
            const extras = JSON.parse(advanced)
            if (!extras || typeof extras !== "object" || Array.isArray(extras))
                throw new Error(
                    text("高级字段必须是 JSON 对象", "Advanced fields must be a JSON object"),
                )
            const allowed = data?.capabilities.fields || []
            if (Object.keys(extras).some((key) => !allowed.includes(key)))
                throw new Error(
                    text(
                        "包含当前提供商不支持的高级字段",
                        "Advanced fields include unsupported provider capabilities",
                    ),
                )
            await change(record.id ? "update" : "create", { ...extras, ...record })
            setEditing(false)
            setRecord(initialRecord)
            setAdvanced("{}")
        })
    }
    async function exportRecords() {
        const values = await fetcher<{ version: number; zone: string; records: RecordValue[] }>(
            FetcherMethod.GET,
            `/api/v1/unified/dns/zones/${zone}/export`,
        )
        const url = URL.createObjectURL(
            new Blob([JSON.stringify(values, null, 2)], { type: "application/json" }),
        )
        const link = document.createElement("a")
        link.href = url
        link.download = "dns-records.json"
        link.click()
        URL.revokeObjectURL(url)
    }

    return (
        <main className="space-y-6 p-4 py-8">
            <UnifiedNav />
            <ErrorNotice error={optionsError} />
            <div>
                <h1 className="text-2xl font-semibold">
                    {text("集中 DNS 管理", "DNS management")}
                </h1>
                <p className="mt-2 text-muted-foreground">
                    {text(
                        "统一管理 Cloudflare、DNSPod 和 EdgeOne 的权威解析与名称授权。",
                        "Manage authoritative DNS and delegated names across Cloudflare, DNSPod and EdgeOne.",
                    )}
                </p>
            </div>
            {(message || error) && (
                <p
                    role="alert"
                    className="rounded-md border border-destructive p-4 text-destructive"
                >
                    {message || error.message}
                </p>
            )}
            {admin && (
                <details className="rounded-lg border p-4">
                    <summary className="cursor-pointer font-medium">
                        {text("账户与凭据绑定", "Accounts and credentials")}
                    </summary>
                    <form
                        className="mt-4 grid gap-3 sm:grid-cols-2"
                        onSubmit={(e) => {
                            e.preventDefault()
                            void action(async () => {
                                await fetcher(
                                    FetcherMethod.POST,
                                    "/api/v1/unified/dns/accounts",
                                    account,
                                )
                                await refreshAccounts()
                            })
                        }}
                    >
                        <div>
                            <Label htmlFor="dns-account-name">
                                {text("账户名称", "Account name")}
                            </Label>
                            <Input
                                id="dns-account-name"
                                required
                                value={account.name}
                                onChange={(e) => setAccount({ ...account, name: e.target.value })}
                            />
                        </div>
                        <div>
                            <Label htmlFor="dns-provider">{text("提供商", "Provider")}</Label>
                            <select
                                id="dns-provider"
                                className="h-10 w-full rounded-md border bg-background px-3"
                                value={account.provider}
                                onChange={(e) =>
                                    setAccount({ ...account, provider: e.target.value })
                                }
                            >
                                <option value="cloudflare">Cloudflare</option>
                                <option value="dnspod">DNSPod (Tencent Cloud)</option>
                                <option value="edgeone">EdgeOne DNS (NS)</option>
                            </select>
                        </div>
                        <div>
                            <Label htmlFor="dns-credential">
                                {text("凭据 ID", "Credential ID")}
                            </Label>
                            <Input
                                id="dns-credential"
                                required
                                value={account.credential_id}
                                onChange={(e) =>
                                    setAccount({ ...account, credential_id: e.target.value })
                                }
                            />
                        </div>
                        <div>
                            <Label htmlFor="dns-version">
                                {text("凭据版本", "Credential version")}
                            </Label>
                            <Input
                                id="dns-version"
                                required
                                type="number"
                                min="1"
                                value={account.credential_version}
                                onChange={(e) =>
                                    setAccount({
                                        ...account,
                                        credential_version: Number(e.target.value),
                                    })
                                }
                            />
                        </div>
                        <p className="text-muted-foreground sm:col-span-2">
                            {text(
                                "Cloudflare 使用 cloudflare 类型凭据；DNSPod 与 EdgeOne 使用 tencentcloud 类型的 secretId / secretKey，不接受传统 DNSPod Token。",
                                "Cloudflare uses a cloudflare credential. DNSPod and EdgeOne use tencentcloud secretId / secretKey credentials; legacy DNSPod tokens are not interchangeable.",
                            )}
                        </p>
                        <Button disabled={busy} className="sm:col-span-2" type="submit">
                            {text("绑定账户", "Bind account")}
                        </Button>
                    </form>
                    <ul className="mt-4 divide-y">
                        {accounts?.map((a) => (
                            <li
                                key={a.id}
                                className="flex flex-wrap items-center justify-between gap-2 py-3"
                            >
                                <span>
                                    {a.name} · {a.provider} · v{a.credential_version}
                                </span>
                                <div className="flex flex-wrap gap-2">
                                    <Button
                                        disabled={busy}
                                        variant="outline"
                                        onClick={() =>
                                            void action(async () => {
                                                await fetcher(
                                                    FetcherMethod.PATCH,
                                                    `/api/v1/unified/dns/accounts/${a.id}`,
                                                    { ...a, enabled: !a.enabled },
                                                )
                                                await refreshAccounts()
                                            })
                                        }
                                    >
                                        {a.enabled
                                            ? text("停用", "Disable")
                                            : text("启用", "Enable")}
                                    </Button>
                                    <Button
                                        disabled={busy}
                                        variant="outline"
                                        onClick={() => {
                                            const version = window.prompt(
                                                text(
                                                    "绑定新的凭据版本",
                                                    "Bind a new credential version",
                                                ),
                                                String(a.credential_version),
                                            )
                                            if (version === null) return
                                            void action(async () => {
                                                if (!/^\d+$/.test(version) || Number(version) < 1)
                                                    throw new Error(
                                                        text("无效版本", "Invalid version"),
                                                    )
                                                await fetcher(
                                                    FetcherMethod.PATCH,
                                                    `/api/v1/unified/dns/accounts/${a.id}`,
                                                    { ...a, credential_version: Number(version) },
                                                )
                                                await refreshAccounts()
                                            })
                                        }}
                                    >
                                        {text("更新凭据版本", "Update credential version")}
                                    </Button>
                                    <Button
                                        disabled={busy || !a.enabled}
                                        variant="outline"
                                        onClick={() =>
                                            void action(async () => {
                                                await fetcher(
                                                    FetcherMethod.POST,
                                                    `/api/v1/unified/dns/accounts/${a.id}/refresh`,
                                                    {},
                                                )
                                                await refreshZones()
                                            })
                                        }
                                    >
                                        {text("发现 / 刷新 Zone", "Discover / refresh zones")}
                                    </Button>
                                </div>
                            </li>
                        ))}
                    </ul>
                </details>
            )}
            <div className="flex flex-wrap items-end gap-3">
                <div className="min-w-60 flex-1">
                    <Label htmlFor="dns-zone">Zone</Label>
                    <select
                        id="dns-zone"
                        className="h-11 w-full rounded-md border bg-background px-3"
                        value={zone}
                        onChange={(e) => {
                            setZone(e.target.value)
                            setPage(1)
                            setEditing(false)
                        }}
                    >
                        <option value="">{text("选择已纳管 Zone", "Choose a managed zone")}</option>
                        {zones?.map((z) => (
                            <option key={z.id} value={z.id}>
                                {z.name} ·{" "}
                                {accounts?.find((a) => a.id === z.account_id)?.name || z.account_id}
                            </option>
                        ))}
                    </select>
                </div>
                <Button disabled={!zone || busy} onClick={() => beginEdit(initialRecord)}>
                    {text("新建记录", "Create record")}
                </Button>
                <Button
                    variant="outline"
                    disabled={!zone || busy}
                    onClick={() => void action(exportRecords)}
                >
                    {text("导出获准记录", "Export authorized records")}
                </Button>
            </div>
            {options?.warnings?.map((warning) => (
                <p
                    key={warning}
                    role="status"
                    className="rounded-md border p-3 text-muted-foreground"
                >
                    {warning}
                </p>
            ))}
            {zone && (
                <Check
                    label={text(
                        "修改记录时暂停关联的 DDNS 绑定",
                        "Pause linked DDNS bindings when changing a record",
                    )}
                    checked={pauseDDNS}
                    onChange={setPauseDDNS}
                />
            )}
            {editing && (
                <form
                    onSubmit={saveRecord}
                    className="grid gap-4 rounded-lg border p-4 sm:grid-cols-2"
                >
                    <h2 className="font-semibold sm:col-span-2">
                        {text(
                            record.id ? "编辑记录" : "新建记录",
                            record.id ? "Edit record" : "Create record",
                        )}
                    </h2>
                    <div>
                        <Label htmlFor="record-name">
                            {text("完整记录名称", "Fully qualified name")}
                        </Label>
                        <Input
                            id="record-name"
                            required
                            value={record.name}
                            onChange={(e) => setRecord({ ...record, name: e.target.value })}
                        />
                    </div>
                    <div>
                        <Label htmlFor="record-type">{text("记录类型", "Record type")}</Label>
                        <select
                            id="record-type"
                            className="h-10 w-full rounded-md border bg-background px-3"
                            value={record.type}
                            onChange={(e) => setRecord({ ...record, type: e.target.value })}
                        >
                            {(data?.capabilities.types || ["A", "AAAA", "TXT"]).map((type) => (
                                <option key={type}>{type}</option>
                            ))}
                        </select>
                    </div>
                    <div>
                        <Label htmlFor="record-value">{text("记录值", "Value")}</Label>
                        <Input
                            id="record-value"
                            value={record.value}
                            onChange={(e) => setRecord({ ...record, value: e.target.value })}
                        />
                    </div>
                    <div>
                        <Label htmlFor="record-ttl">TTL</Label>
                        <Input
                            id="record-ttl"
                            required
                            type="number"
                            min="1"
                            value={record.ttl}
                            onChange={(e) => setRecord({ ...record, ttl: Number(e.target.value) })}
                        />
                    </div>
                    <div className="sm:col-span-2">
                        <Label htmlFor="record-extras">
                            {text("提供商高级字段（JSON）", "Provider-specific fields (JSON)")}
                        </Label>
                        <Textarea
                            id="record-extras"
                            rows={5}
                            spellCheck={false}
                            value={advanced}
                            onChange={(e) => setAdvanced(e.target.value)}
                        />
                        <p className="mt-2 text-muted-foreground">
                            {text("支持字段", "Supported fields")}:{" "}
                            {data?.capabilities.fields.join(", ") || "—"}
                            {!!options?.lines?.length && (
                                <span className="mt-1 block">
                                    {text("线路", "Lines")}:{" "}
                                    {options.lines
                                        .map((line) => `${line.name} (${line.id})`)
                                        .join(", ")}
                                </span>
                            )}
                        </p>
                    </div>
                    <div className="flex gap-3 sm:col-span-2">
                        <Button disabled={busy} type="submit">
                            {text("保存记录", "Save record")}
                        </Button>
                        <Button
                            disabled={busy}
                            variant="outline"
                            type="button"
                            onClick={() => setEditing(false)}
                        >
                            {text("取消", "Cancel")}
                        </Button>
                    </div>
                </form>
            )}
            {isLoading && (
                <p role="status">{text("正在读取提供商记录…", "Loading provider records…")}</p>
            )}
            {data && (
                <>
                    <div className="overflow-x-auto rounded-lg border">
                        <table className="w-full text-left">
                            <thead className="bg-muted">
                                <tr>
                                    {[
                                        text("名称", "Name"),
                                        text("类型", "Type"),
                                        text("值", "Value"),
                                        "TTL",
                                        text("操作", "Actions"),
                                    ].map((h) => (
                                        <th key={h} className="p-3" scope="col">
                                            {h}
                                        </th>
                                    ))}
                                </tr>
                            </thead>
                            <tbody>
                                {data.items.map((r) => (
                                    <tr key={r.id} className="border-t">
                                        <td className="max-w-60 break-words p-3">
                                            {r.name}
                                            <span className="block text-xs text-muted-foreground">
                                                {r.id}
                                            </span>
                                        </td>
                                        <td className="p-3">{r.type}</td>
                                        <td className="max-w-72 break-all p-3">
                                            {r.value || JSON.stringify(r.data)}
                                        </td>
                                        <td className="p-3">{r.ttl}</td>
                                        <td className="p-3">
                                            <div className="flex gap-2">
                                                <Button
                                                    variant="outline"
                                                    disabled={busy}
                                                    onClick={() => beginEdit(r)}
                                                >
                                                    {text("编辑", "Edit")}
                                                </Button>
                                                <Button
                                                    variant="outline"
                                                    disabled={busy}
                                                    onClick={() => {
                                                        if (
                                                            window.confirm(
                                                                text(
                                                                    `删除 ${r.name} 的 ${r.type} 记录？`,
                                                                    `Delete the ${r.type} record for ${r.name}?`,
                                                                ),
                                                            )
                                                        )
                                                            void action(() => change("delete", r))
                                                    }}
                                                >
                                                    {text("删除", "Delete")}
                                                </Button>
                                            </div>
                                        </td>
                                    </tr>
                                ))}
                            </tbody>
                        </table>
                        {data.items.length === 0 && (
                            <p className="p-8 text-center text-muted-foreground">
                                {text("暂无获准的记录", "No authorized records")}
                            </p>
                        )}
                    </div>
                    <div className="flex items-center justify-between">
                        <span>
                            {data.total} {text("条获准记录", "authorized records")}
                        </span>
                        <div className="flex gap-2">
                            <Button
                                variant="outline"
                                disabled={page === 1 || busy}
                                onClick={() => setPage(page - 1)}
                            >
                                {text("上一页", "Previous")}
                            </Button>
                            <Button
                                variant="outline"
                                disabled={page * 50 >= data.total || busy}
                                onClick={() => setPage(page + 1)}
                            >
                                {text("下一页", "Next")}
                            </Button>
                        </div>
                    </div>
                </>
            )}
            {zone && (
                <details className="rounded-lg border p-4">
                    <summary className="cursor-pointer font-medium">
                        {text("批量导入 JSON", "Import JSON records")}
                    </summary>
                    <Label htmlFor="dns-import">
                        {text("记录数组（不含 id）", "Record array (without IDs)")}
                    </Label>
                    <Textarea
                        id="dns-import"
                        className="my-3"
                        rows={6}
                        value={importJSON}
                        onChange={(e) => setImportJSON(e.target.value)}
                    />
                    <Button
                        disabled={busy}
                        onClick={() =>
                            void action(async () => {
                                const imported = JSON.parse(importJSON)
                                const batch = Array.isArray(imported) ? imported : imported.records
                                if (!Array.isArray(batch) || batch.length > 100)
                                    throw new Error(
                                        text(
                                            "每批最多 100 条记录",
                                            "A batch supports up to 100 records",
                                        ),
                                    )
                                const result = await fetcher<
                                    { operation?: Operation; error?: string }[]
                                >(FetcherMethod.POST, `/api/v1/unified/dns/zones/${zone}/import`, {
                                    operation_key: crypto.randomUUID(),
                                    records: batch.map(
                                        ({ id: _id, ...record }: RecordValue) => record,
                                    ),
                                })
                                await refreshRecords()
                                await refreshOperations()
                                if (
                                    result.some(
                                        (item) =>
                                            item.error || item.operation?.state !== "succeeded",
                                    )
                                )
                                    throw new Error(
                                        text(
                                            "部分记录未成功，请核验操作历史。",
                                            "Some records did not succeed. Inspect operation history.",
                                        ),
                                    )
                                setImportJSON("")
                            })
                        }
                    >
                        {text("逐项校验并导入", "Validate and import each record")}
                    </Button>
                    <p className="mt-2 text-muted-foreground">
                        {text(
                            "各记录分别校验，保留成功项；请核验每项操作结果。",
                            "Each record is checked independently. Completed records are preserved; inspect every operation result.",
                        )}
                    </p>
                </details>
            )}
            {admin && zone && (
                <details className="rounded-lg border p-4">
                    <summary className="cursor-pointer font-medium">
                        {text("名称子树授权", "Delegate a name subtree")}
                    </summary>
                    <form
                        className="mt-4 grid gap-3 sm:grid-cols-2"
                        onSubmit={(e) => {
                            e.preventDefault()
                            void action(async () => {
                                await fetcher(FetcherMethod.POST, "/api/v1/unified/dns/policies", {
                                    ...grant,
                                    user_id: Number(grant.user_id),
                                    service_id: Number(grant.service_id),
                                    zone_id: Number(zone),
                                    types: grant.types.split(",").map((v) => v.trim()),
                                    actions: grant.actions.split(",").map((v) => v.trim()),
                                })
                                await refreshGrants()
                            })
                        }}
                    >
                        <div>
                            <Label htmlFor="dns-user">{text("用户 ID", "User ID")}</Label>
                            <Input
                                id="dns-user"
                                required={!grant.service_id}
                                type="number"
                                min="1"
                                value={grant.user_id}
                                onChange={(e) =>
                                    setGrant({ ...grant, user_id: e.target.value, service_id: "" })
                                }
                            />
                        </div>
                        <div>
                            <Label htmlFor="dns-service">
                                {text(
                                    "执行身份 ID（与用户二选一）",
                                    "Execution identity ID (instead of a user)",
                                )}
                            </Label>
                            <Input
                                id="dns-service"
                                type="number"
                                min="1"
                                value={grant.service_id}
                                onChange={(e) =>
                                    setGrant({ ...grant, service_id: e.target.value, user_id: "" })
                                }
                            />
                        </div>
                        <div>
                            <Label htmlFor="dns-root">{text("子树根名称", "Subtree root")}</Label>
                            <Input
                                id="dns-root"
                                required
                                value={grant.root}
                                onChange={(e) => setGrant({ ...grant, root: e.target.value })}
                            />
                        </div>
                        <div>
                            <Label htmlFor="dns-types">
                                {text("获准类型（逗号分隔）", "Allowed types (comma separated)")}
                            </Label>
                            <Input
                                id="dns-types"
                                required
                                value={grant.types}
                                onChange={(e) => setGrant({ ...grant, types: e.target.value })}
                            />
                        </div>
                        <div>
                            <Label htmlFor="dns-actions">
                                {text(
                                    "动作（read/create/update/delete/ddns）",
                                    "Actions (read/create/update/delete/ddns)",
                                )}
                            </Label>
                            <Input
                                id="dns-actions"
                                required
                                value={grant.actions}
                                onChange={(e) => setGrant({ ...grant, actions: e.target.value })}
                            />
                        </div>
                        <div className="flex flex-wrap gap-4 sm:col-span-2">
                            {(
                                [
                                    ["include_root", text("包含根名称", "Include root")],
                                    ["descendants", text("包含后代", "Descendants")],
                                    ["allow_wildcard", text("允许通配记录", "Allow wildcards")],
                                ] as const
                            ).map(([key, label]) => (
                                <label key={key} className="flex min-h-11 items-center gap-2">
                                    <input
                                        type="checkbox"
                                        checked={grant[key]}
                                        onChange={(e) =>
                                            setGrant({ ...grant, [key]: e.target.checked })
                                        }
                                    />
                                    {label}
                                </label>
                            ))}
                        </div>
                        <Button disabled={busy} type="submit" className="sm:col-span-2">
                            {text("授予 DNS 权限", "Grant DNS access")}
                        </Button>
                    </form>
                    <ul className="mt-4 divide-y">
                        {grants
                            ?.filter((g) => g.zone_id === Number(zone))
                            .map((g) => (
                                <li
                                    key={g.id}
                                    className="flex flex-wrap items-center justify-between gap-3 py-3"
                                >
                                    <span>
                                        {g.service_id
                                            ? `Service #${g.service_id}`
                                            : `User #${g.user_id}`}{" "}
                                        · {g.root} · {g.types.join(", ")} · {g.actions.join(", ")}
                                    </span>
                                    <Button
                                        variant="outline"
                                        disabled={busy}
                                        onClick={() =>
                                            void action(async () => {
                                                await fetcher(
                                                    FetcherMethod.DELETE,
                                                    `/api/v1/unified/dns/policies/${g.id}`,
                                                )
                                                await refreshGrants()
                                            })
                                        }
                                    >
                                        {text("撤销", "Revoke")}
                                    </Button>
                                </li>
                            ))}
                    </ul>
                </details>
            )}
            <section>
                <h2 className="mb-3 font-semibold">{text("最近操作", "Recent operations")}</h2>
                <ul className="divide-y rounded-lg border">
                    {operations?.map((op) => (
                        <li key={op.id} className="flex flex-wrap justify-between gap-2 p-3">
                            <span>
                                #{op.id} · {op.action} · {op.name}
                            </span>
                            {["unknown", "applying"].includes(op.state) && (
                                <Button
                                    variant="outline"
                                    disabled={busy}
                                    onClick={() => {
                                        const recordID = window.prompt(
                                            text(
                                                "输入已核验的提供商记录 ID（已绑定时可留空）",
                                                "Verified provider record ID (leave blank if already linked)",
                                            ),
                                            "",
                                        )
                                        if (recordID === null) return
                                        void action(async () => {
                                            await fetcher(
                                                FetcherMethod.POST,
                                                `/api/v1/unified/dns/operations/${op.id}/reconcile`,
                                                { record_id: recordID },
                                            )
                                            await refreshOperations()
                                            await refreshRecords()
                                        })
                                    }}
                                >
                                    {text("核验当前结果", "Reconcile current result")}
                                </Button>
                            )}
                            <span>
                                {op.state === "applying"
                                    ? text(
                                          "执行中；重启后需核验",
                                          "Applying; reconcile after restart",
                                      )
                                    : op.state}
                                {op.error ? ` · ${op.error}` : ""}
                            </span>
                        </li>
                    ))}
                </ul>
            </section>
        </main>
    )
}
