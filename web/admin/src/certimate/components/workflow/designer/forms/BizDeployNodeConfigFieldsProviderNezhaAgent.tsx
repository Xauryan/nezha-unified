import { swrFetcher } from "@/api/api"
import { Alert, Form, InputNumber, Select } from "antd"
import { useTranslation } from "react-i18next"
import useSWR from "swr"

type Target = {
    id: number
    name: string
    server_id: number
    version: number
    enabled: boolean
    allowed_names: string[]
}
export default function NezhaAgentTargetFields() {
    const { i18n } = useTranslation()
    const zh = i18n.language.startsWith("zh")
    const form = Form.useFormInstance()
    const selected = Form.useWatch("providerAccessId", form)
    const { data: targets, error } = useSWR<Target[]>("/api/v1/unified/targets", swrFetcher)
    const { data: capability, error: capabilityError } = useSWR<{
        online: boolean
        certificate_deploy: boolean
    }>(
        selected ? `/api/v1/unified/targets/${encodeURIComponent(selected)}/capabilities` : null,
        swrFetcher,
    )
    const target = targets?.find((item) => String(item.id) === selected)
    return (
        <>
            <Form.Item
                name="providerAccessId"
                label={zh ? "哪吒部署目标" : "Nezha deployment target"}
                rules={[{ required: true }]}
            >
                <Select
                    showSearch
                    optionFilterProp="label"
                    loading={!targets && !error}
                    options={targets?.map((item) => ({
                        value: String(item.id),
                        label: `${item.name} · Host #${item.server_id} · v${item.version}`,
                        disabled: !item.enabled,
                    }))}
                    onChange={(id) =>
                        form.setFieldValue(
                            ["providerConfig", "targetVersion"],
                            targets?.find((item) => String(item.id) === id)?.version,
                        )
                    }
                />
            </Form.Item>
            <Form.Item
                name={["providerConfig", "targetVersion"]}
                label={zh ? "固定的目标版本" : "Pinned target version"}
                rules={[{ required: true, type: "number", min: 1 }]}
            >
                <InputNumber min={1} precision={0} readOnly className="w-full" />
            </Form.Item>
            {target && (
                <p className="mb-4 text-muted-foreground">{target.allowed_names.join(", ")}</p>
            )}
            {(error || capabilityError) && (
                <Alert type="warning" title={String(error?.message || capabilityError?.message)} />
            )}
            {capability && (
                <Alert
                    className="mb-4"
                    type={capability.online && capability.certificate_deploy ? "info" : "warning"}
                    title={
                        capability.online && capability.certificate_deploy
                            ? zh
                                ? "Agent 支持原子部署与结果核验。"
                                : "Agent supports atomic deployment and result verification."
                            : zh
                              ? "Agent 离线或不支持证书部署，执行前需要就绪。"
                              : "Agent is offline or lacks certificate deployment; it must be ready before execution."
                    }
                />
            )}
        </>
    )
}
