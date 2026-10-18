import { swrFetcher } from "@/api/api"
import useSWR from "swr"

import { ErrorNotice, useText } from "./common"

type Capabilities = {
    version: string
    certimate: string
    storage: string
    workflow_execution: boolean
    dns_management: boolean
    agent_deployment: boolean
}
export const useUnifiedCapabilities = () =>
    useSWR<Capabilities>("/api/v1/unified/status", swrFetcher)
export default function UnifiedCapabilities() {
    const t = useText(),
        { data, error } = useUnifiedCapabilities()
    return (
        <section className="space-y-3 rounded-lg border p-4">
            <h2 className="font-semibold">{t("运行时能力", "Runtime capabilities")}</h2>
            <ErrorNotice error={error} />
            {data && (
                <>
                    <p className="text-sm text-muted-foreground">
                        Nezha Unified {data.version} · Certimate {data.certimate} · {data.storage}
                    </p>
                    <dl className="grid gap-3 sm:grid-cols-3">
                        {(
                            [
                                ["workflow_execution", t("工作流执行", "Workflow execution")],
                                ["dns_management", t("DNS 管理", "DNS management")],
                                ["agent_deployment", t("Agent 部署", "Agent deployment")],
                            ] as const
                        ).map(([key, label]) => (
                            <div key={key}>
                                <dt className="text-sm text-muted-foreground">{label}</dt>
                                <dd className="font-medium">
                                    {data[key]
                                        ? t("可用", "Available")
                                        : t("未启用", "Unavailable")}
                                </dd>
                            </div>
                        ))}
                    </dl>
                </>
            )}
        </section>
    )
}
