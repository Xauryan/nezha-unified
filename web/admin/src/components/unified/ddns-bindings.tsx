import { FetcherMethod as M, fetcher, swrFetcher } from "@/api/api"
import { Button } from "@/components/ui/button"
import { useMainStore } from "@/hooks/useMainStore"
import { useState } from "react"
import useSWR from "swr"

import { Check, ErrorNotice, Field, Panel, useAction, useText } from "./common"

type Binding = {
    id: number
    version: number
    zone_id: number
    record_id: string
    name: string
    record_type: string
    source_kind: string
    host_id: number
    interface: string
    address: string
    stability_seconds: number
    max_age_seconds: number
    enabled: boolean
    network_enabled: boolean
    owner_id: number
    service_id?: number
    identity_version?: number
    last_observed: string
    last_success: string
    last_observed_at: string
    last_updated_at: string
    state: string
    error: string
}
type Device = { id: number; binding_id: number; revoked_at?: string }
const initial = {
    zone_id: "",
    record_id: "",
    name: "",
    record_type: "A",
    source_kind: "host",
    host_id: "",
    interface: "",
    address: "ipv4",
    stability_seconds: "30",
    max_age_seconds: "300",
    enabled: false,
    network_enabled: false,
}
export default function DDNSBindings() {
    const admin = useMainStore((s) => s.profile?.role === 0)
    const t = useText(),
        action = useAction(),
        [form, setForm] = useState(initial),
        [edit, setEdit] = useState<Binding>(),
        [token, setToken] = useState("")
    const { data, mutate, error } = useSWR<Binding[]>("/api/v1/unified/ddns/bindings", swrFetcher, {
        refreshInterval: 10000,
    })
    const { data: devices, mutate: refreshDevices } = useSWR<Device[]>(
        "/api/v1/unified/ddns/devices",
        swrFetcher,
    )
    const { data: addresses } = useSWR<unknown>(
        form.source_kind === "host" && Number(form.host_id) > 0
            ? `/api/v1/unified/ddns/hosts/${form.host_id}/addresses`
            : null,
        swrFetcher,
    )
    const field = (key: keyof typeof initial, value: string | boolean) =>
        setForm((p) => ({ ...p, [key]: value }))
    function startEdit(binding: Binding) {
        setEdit(binding)
        setForm({
            zone_id: String(binding.zone_id),
            record_id: binding.record_id,
            name: binding.name,
            record_type: binding.record_type,
            source_kind: binding.source_kind,
            host_id: String(binding.host_id || ""),
            interface: binding.interface || "",
            address: binding.address,
            stability_seconds: String(binding.stability_seconds),
            max_age_seconds: String(binding.max_age_seconds),
            enabled: binding.enabled,
            network_enabled: binding.network_enabled,
        })
    }
    return (
        <section className="my-8 space-y-4 border-t pt-6">
            <h2 className="text-xl font-semibold">
                {t("统一 DNS 记录绑定", "Unified DNS record bindings")}
            </h2>
            <p className="text-sm text-muted-foreground">
                {t(
                    "复用 DNS 账户和名称授权，使用主机观测地址或设备上报更新指定记录。",
                    "Reuse DNS accounts and name permissions to update a record from host observations or device reports.",
                )}
            </p>
            <ErrorNotice error={action.error || error} />
            <Panel
                title={edit ? t("编辑绑定", "Edit binding") : t("新增绑定", "Create binding")}
                open={!!edit}
            >
                <form
                    className="grid gap-4 sm:grid-cols-2"
                    onSubmit={(e) => {
                        e.preventDefault()
                        void action.run(async () => {
                            await fetcher(
                                edit ? M.PATCH : M.POST,
                                `/api/v1/unified/ddns/bindings${edit ? "/" + edit.id : ""}`,
                                {
                                    ...form,
                                    version: edit?.version,
                                    zone_id: Number(form.zone_id),
                                    host_id: Number(form.host_id),
                                    stability_seconds: Number(form.stability_seconds),
                                    max_age_seconds: Number(form.max_age_seconds),
                                },
                            )
                            setEdit(undefined)
                            setForm(initial)
                            await mutate()
                        })
                    }}
                >
                    <Field
                        id="binding-zone"
                        label={t("区域 ID", "Zone ID")}
                        type="number"
                        min={1}
                        value={form.zone_id}
                        onChange={(v) => field("zone_id", v)}
                        required
                    />
                    <Field
                        id="binding-record"
                        label={t("记录 ID", "Record ID")}
                        value={form.record_id}
                        onChange={(v) => field("record_id", v)}
                        required
                    />
                    <Field
                        id="binding-name"
                        label={t("完整记录名称", "Full record name")}
                        value={form.name}
                        onChange={(v) => field("name", v)}
                        required
                    />
                    <label className="space-y-2">
                        <span className="block">{t("记录类型", "Record type")}</span>
                        <select
                            className="h-11 w-full rounded-md border bg-background px-3"
                            value={form.record_type}
                            onChange={(e) => {
                                field("record_type", e.target.value)
                                field("address", e.target.value === "A" ? "ipv4" : "ipv6")
                            }}
                        >
                            <option>A</option>
                            <option>AAAA</option>
                        </select>
                    </label>
                    <label className="space-y-2">
                        <span className="block">{t("地址来源", "Address source")}</span>
                        <select
                            className="h-11 w-full rounded-md border bg-background px-3"
                            value={form.source_kind}
                            onChange={(e) => field("source_kind", e.target.value)}
                        >
                            <option value="host">{t("哪吒主机", "Nezha host")}</option>
                            <option value="device">{t("设备令牌", "Device token")}</option>
                        </select>
                    </label>
                    {form.source_kind === "host" && (
                        <>
                            <Field
                                id="binding-host"
                                label={t("主机 ID", "Host ID")}
                                type="number"
                                min={1}
                                value={form.host_id}
                                onChange={(v) => field("host_id", v)}
                                required
                            />
                            <Field
                                id="binding-interface"
                                label={t("网卡（可选）", "Interface (optional)")}
                                value={form.interface}
                                onChange={(v) => field("interface", v)}
                            />
                            <Field
                                id="binding-address"
                                label={t("地址选择", "Address selection")}
                                value={form.address}
                                onChange={(v) => field("address", v)}
                                help={t(
                                    "ipv4 / ipv6 或 Agent 上报的具体地址。",
                                    "ipv4 / ipv6 or a specific Agent-observed address.",
                                )}
                            />
                            {addresses != null && (
                                <pre className="max-h-40 overflow-auto rounded-md bg-muted p-3 text-xs sm:col-span-2">
                                    {JSON.stringify(addresses, null, 2)}
                                </pre>
                            )}
                        </>
                    )}
                    <Field
                        id="binding-stability"
                        label={t("稳定窗口（秒）", "Stability window (seconds)")}
                        type="number"
                        min={0}
                        value={form.stability_seconds}
                        onChange={(v) => field("stability_seconds", v)}
                        required
                    />
                    <Field
                        id="binding-age"
                        label={t("观测有效期（秒）", "Observation max age (seconds)")}
                        type="number"
                        min={1}
                        value={form.max_age_seconds}
                        onChange={(v) => field("max_age_seconds", v)}
                        required
                    />
                    <Check
                        label={t("启用绑定", "Enable binding")}
                        checked={form.enabled}
                        onChange={(v) => field("enabled", v)}
                    />
                    <Check
                        label={t("允许向 DNS 提供商发送更新", "Allow updates to the DNS provider")}
                        checked={form.network_enabled}
                        onChange={(v) => field("network_enabled", v)}
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
                                {t("取消", "Cancel")}
                            </Button>
                        )}
                    </div>
                </form>
            </Panel>
            {token && (
                <div className="space-y-3 rounded-lg border p-4">
                    <p>{t("设备令牌仅显示一次。", "The device token is shown once.")}</p>
                    <input
                        type="password"
                        readOnly
                        value={token}
                        aria-label={t("设备令牌", "Device token")}
                        className="w-full rounded-md border bg-background p-3"
                        autoComplete="off"
                    />
                    <div className="flex gap-2">
                        <Button
                            variant="outline"
                            onClick={() =>
                                void action.run(() => navigator.clipboard.writeText(token))
                            }
                        >
                            {t("复制", "Copy")}
                        </Button>
                        <Button variant="outline" onClick={() => setToken("")}>
                            {t("关闭", "Dismiss")}
                        </Button>
                    </div>
                </div>
            )}
            {data?.map((binding) => (
                <article className="space-y-3 rounded-lg border p-4" key={binding.id}>
                    <div>
                        <h3 className="font-medium">
                            {binding.name} · {binding.record_type}{" "}
                            <span className="text-muted-foreground">#{binding.id}</span>
                        </h3>
                        <p className="mt-1 text-sm">
                            {binding.state || t("尚未运行", "Not run")} · {t("观测", "Observed")}:{" "}
                            {binding.last_observed || "—"} · {t("已应用", "Applied")}:{" "}
                            {binding.last_success || "—"}
                        </p>
                        <p className="text-xs text-muted-foreground">
                            {t("最近观测", "Last observation")}: {binding.last_observed_at || "—"} ·{" "}
                            {t("最近更新", "Last update")}: {binding.last_updated_at || "—"}
                        </p>
                        {binding.error && <p className="text-destructive">{binding.error}</p>}
                    </div>
                    {binding.service_id ? (
                        <p className="text-xs text-muted-foreground">
                            {t("执行身份", "Execution identity")} #{binding.service_id} · v
                            {binding.identity_version}
                        </p>
                    ) : null}
                    <div className="flex flex-wrap gap-2">
                        {admin && (
                            <Button
                                variant="outline"
                                disabled={action.busy}
                                onClick={() => {
                                    const identity = window.prompt(
                                        t(
                                            "新的执行身份 ID；交接会重置观测并撤销旧设备令牌。",
                                            "New execution identity ID; handover resets observations and revokes prior device tokens.",
                                        ),
                                    )
                                    if (identity === null) return
                                    void action.run(async () => {
                                        if (!/^\d+$/.test(identity) || Number(identity) < 1)
                                            throw new Error(
                                                t("请输入有效身份 ID", "Enter a valid identity ID"),
                                            )
                                        await fetcher(
                                            M.POST,
                                            `/api/v1/unified/ddns/bindings/${binding.id}/service`,
                                            {
                                                service_id: Number(identity),
                                                version: binding.version,
                                            },
                                        )
                                        await mutate()
                                        await refreshDevices()
                                    })
                                }}
                            >
                                {t("交接执行身份", "Assign execution identity")}
                            </Button>
                        )}
                        <Button variant="outline" onClick={() => startEdit(binding)}>
                            {t("编辑", "Edit")}
                        </Button>
                        <Button
                            variant="outline"
                            disabled={action.busy || !binding.enabled || !binding.network_enabled}
                            onClick={() =>
                                void action.run(async () => {
                                    await fetcher(
                                        M.POST,
                                        `/api/v1/unified/ddns/bindings/${binding.id}/run`,
                                    )
                                    await mutate()
                                })
                            }
                        >
                            {t("立即运行", "Run now")}
                        </Button>
                        {binding.source_kind === "device" && (
                            <Button
                                variant="outline"
                                disabled={action.busy}
                                onClick={() =>
                                    void action.run(async () => {
                                        const result = await fetcher<{ token: string }>(
                                            M.POST,
                                            "/api/v1/unified/ddns/devices",
                                            { binding_id: binding.id },
                                        )
                                        setToken(result.token)
                                        await refreshDevices()
                                    })
                                }
                            >
                                {t("创建设备令牌", "Create device token")}
                            </Button>
                        )}
                        <Button
                            variant="outline"
                            disabled={action.busy}
                            onClick={() =>
                                void action.run(async () => {
                                    await fetcher(
                                        M.DELETE,
                                        `/api/v1/unified/ddns/bindings/${binding.id}`,
                                    )
                                    await mutate()
                                })
                            }
                        >
                            {t("删除绑定", "Delete binding")}
                        </Button>
                    </div>
                    {devices
                        ?.filter((d) => d.binding_id === binding.id)
                        .map((d) => (
                            <div
                                key={d.id}
                                className="flex items-center justify-between gap-2 border-t pt-3 text-sm"
                            >
                                <span>
                                    {t("设备", "Device")} #{d.id} ·{" "}
                                    {d.revoked_at ? t("已撤销", "Revoked") : t("有效", "Active")}
                                </span>
                                <Button
                                    variant="outline"
                                    disabled={action.busy || !!d.revoked_at}
                                    onClick={() =>
                                        void action.run(async () => {
                                            await fetcher(
                                                M.DELETE,
                                                `/api/v1/unified/ddns/devices/${d.id}`,
                                            )
                                            await refreshDevices()
                                        })
                                    }
                                >
                                    {t("撤销令牌", "Revoke token")}
                                </Button>
                            </div>
                        ))}
                </article>
            ))}
            {data?.length === 0 && (
                <p className="text-muted-foreground">
                    {t("暂无统一记录绑定。", "No unified record bindings yet.")}
                </p>
            )}
        </section>
    )
}
