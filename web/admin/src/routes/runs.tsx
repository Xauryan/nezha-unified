import { swrFetcher } from "@/api/api"
import UnifiedCapabilities from "@/components/unified/capabilities"
import { ErrorNotice, UnifiedNav, useText } from "@/components/unified/common"
import { Link } from "react-router-dom"
import useSWR from "swr"

type WorkflowRun = {
    id: string
    workflowRef: string
    status: string
    created: string
    error?: string
}
type DNSRun = {
    id: number
    name: string
    action: string
    state: string
    created_at: string
    error?: string
}
type DeploymentRun = {
    id: number
    target_id: number
    state: string
    created_at: string
    code?: string
}
type Runs = {
    workflows: { items: WorkflowRun[]; totalItems: number }
    dns: DNSRun[]
    deployments: DeploymentRun[]
}
export default function RunsPage() {
    const t = useText(),
        { data, error, isLoading } = useSWR<Runs>("/api/v1/unified/runs", swrFetcher, {
            refreshInterval: 5000,
        })
    const items = [
        ...(data?.workflows?.items || []).map((r) => ({
            id: "workflow:" + r.id,
            label: t("证书工作流", "Certificate workflow"),
            resource: r.workflowRef,
            status: r.status,
            time: r.created,
            error: r.error,
            path: `/dashboard/certificates/workflows/${encodeURIComponent(r.workflowRef)}/runs`,
        })),
        ...(data?.dns || []).map((r) => ({
            id: "dns:" + r.id,
            label: "DNS · " + r.action,
            resource: r.name,
            status: r.state,
            time: r.created_at,
            error: r.error,
            path: "/dashboard/dns",
        })),
        ...(data?.deployments || []).map((r) => ({
            id: "deployment:" + r.id,
            label: t("证书部署", "Certificate deployment"),
            resource: String(r.target_id),
            status: r.state,
            time: r.created_at,
            error: r.code,
            path: "/dashboard/targets",
        })),
    ].sort((a, b) => b.time.localeCompare(a.time))
    return (
        <main className="space-y-6 px-4 py-8">
            <UnifiedNav />
            <h1 className="text-2xl font-semibold">{t("运行记录", "Run history")}</h1>
            <UnifiedCapabilities />
            <ErrorNotice error={error} />
            {isLoading && <p role="status">{t("正在加载…", "Loading…")}</p>}
            <div className="overflow-x-auto rounded-lg border">
                <table className="w-full text-left text-sm">
                    <thead className="bg-muted">
                        <tr>
                            {[
                                t("类型", "Type"),
                                t("资源", "Resource"),
                                t("状态", "Status"),
                                t("时间", "Time"),
                            ].map((label) => (
                                <th className="p-3" key={label}>
                                    {label}
                                </th>
                            ))}
                        </tr>
                    </thead>
                    <tbody>
                        {items.map((item) => (
                            <tr key={item.id} className="border-t">
                                <td className="p-3">
                                    {item.label}
                                    <p className="font-mono text-xs text-muted-foreground">
                                        {item.id}
                                    </p>
                                </td>
                                <td className="p-3">
                                    <Link className="underline underline-offset-4" to={item.path}>
                                        {item.resource}
                                    </Link>
                                </td>
                                <td className="p-3">
                                    {item.status}
                                    {item.error && <p className="text-destructive">{item.error}</p>}
                                </td>
                                <td className="whitespace-nowrap p-3">{item.time}</td>
                            </tr>
                        ))}
                    </tbody>
                </table>
                {data && items.length === 0 && (
                    <p className="p-8 text-center text-muted-foreground">
                        {t("暂无获准的运行记录。", "No authorized run records.")}
                    </p>
                )}
            </div>
            <p className="text-sm text-muted-foreground">
                {t(
                    "展示最近的获准记录；在相应模块中查看完整历史和执行详情。",
                    "Shows recent authorized records. Open the corresponding module for complete history and execution details.",
                )}
            </p>
        </main>
    )
}
