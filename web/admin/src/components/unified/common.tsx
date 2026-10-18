import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Textarea } from "@/components/ui/textarea"
import { type ReactNode, useState } from "react"
import { useTranslation } from "react-i18next"
import { NavLink } from "react-router-dom"

export function useText() {
    const { i18n } = useTranslation()
    return (zh: string, en: string) => (i18n.language.startsWith("zh") ? zh : en)
}
export function useAction() {
    const [busy, setBusy] = useState(false),
        [error, setError] = useState("")
    async function run(work: () => Promise<unknown>) {
        setBusy(true)
        setError("")
        try {
            await work()
        } catch (e) {
            setError(e instanceof Error ? e.message : "Request failed")
        } finally {
            setBusy(false)
        }
    }
    return { busy, error, run }
}
export function Field({
    label,
    id,
    value,
    onChange,
    type = "text",
    required = false,
    min,
    help,
}: {
    label: string
    id: string
    value: string | number
    onChange: (v: string) => void
    type?: string
    required?: boolean
    min?: number
    help?: string
}) {
    return (
        <div className="space-y-2">
            <Label htmlFor={id}>{label}</Label>
            <Input
                id={id}
                type={type}
                value={value}
                min={min}
                required={required}
                onChange={(e) => onChange(e.target.value)}
            />
            {help && <p className="text-xs text-muted-foreground">{help}</p>}
        </div>
    )
}
export function JSONField({
    label,
    id,
    value,
    onChange,
    help,
}: {
    label: string
    id: string
    value: string
    onChange: (v: string) => void
    help?: string
}) {
    return (
        <div className="space-y-2">
            <Label htmlFor={id}>{label}</Label>
            <Textarea
                id={id}
                value={value}
                onChange={(e) => onChange(e.target.value)}
                rows={5}
                spellCheck={false}
                autoComplete="off"
                className="font-mono"
            />
            {help && <p className="text-xs text-muted-foreground">{help}</p>}
        </div>
    )
}
export function Check({
    label,
    checked,
    onChange,
}: {
    label: string
    checked: boolean
    onChange: (v: boolean) => void
}) {
    return (
        <label className="flex min-h-11 items-center gap-2">
            <input type="checkbox" checked={checked} onChange={(e) => onChange(e.target.checked)} />
            {label}
        </label>
    )
}
export function ErrorNotice({ error }: { error: unknown }) {
    return error ? (
        <p role="alert" className="rounded-md border border-destructive p-4 text-destructive">
            {error instanceof Error ? error.message : String(error)}
        </p>
    ) : null
}
export function UnifiedNav() {
    const t = useText()
    return (
        <nav
            aria-label={t("资源管理", "Resource management")}
            className="flex flex-wrap gap-2 border-b pb-4"
        >
            {[
                ["runs", t("运行记录", "Run history")],
                ["dns", "DNS"],
                ["credentials", t("凭据", "Credentials")],
                ["certificates", t("证书", "Certificates")],
                ["targets", t("部署目标", "Deployment targets")],
                ["permissions", t("资源授权", "Permissions")],
                ["identities", t("执行身份与令牌", "Identities & tokens")],
            ].map(([path, label]) => (
                <NavLink
                    key={path}
                    to={"/dashboard/" + path}
                    end
                    className={({ isActive }) =>
                        `inline-flex min-h-11 items-center rounded-md px-3 hover:bg-muted ${isActive ? "bg-primary text-primary-foreground" : "text-muted-foreground"}`
                    }
                >
                    {label}
                </NavLink>
            ))}
        </nav>
    )
}
export function Panel({
    title,
    children,
    open = false,
}: {
    title: string
    children: ReactNode
    open?: boolean
}) {
    return (
        <details className="rounded-lg border p-4" open={open || undefined}>
            <summary className="cursor-pointer font-medium">{title}</summary>
            <div className="mt-4 space-y-4">{children}</div>
        </details>
    )
}
export function parseArray<T>(value: string): T[] {
    const result = JSON.parse(value)
    if (!Array.isArray(result)) throw new Error("Expected a JSON array")
    return result
}
export function parseObject<T>(value: string): T {
    const result = JSON.parse(value)
    if (!result || typeof result !== "object" || Array.isArray(result))
        throw new Error("Expected a JSON object")
    return result
}
