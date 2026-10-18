import { FetcherMethod, fetcher, swrFetcher } from "@/api/api"
import { Button } from "@/components/ui/button"
import {
    Dialog,
    DialogContent,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "@/components/ui/dialog"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Textarea } from "@/components/ui/textarea"
import { useMainStore } from "@/hooks/useMainStore"
import { Download, KeyRound, RefreshCw, ShieldCheck } from "lucide-react"
import { type FormEvent, useState } from "react"
import { useTranslation } from "react-i18next"
import { NavLink } from "react-router-dom"
import useSWR from "swr"

type Resource = {
    id: string
    kind: string
    name: string
    provider?: string
    names?: string
    expires?: string
    enabled: boolean
}
type ResourcePage = { items: Resource[]; total: number; page: number; per_page: number }
type Grant = { id: number; user_id: number; kind: string; resource_id: string; action: string }
type ResourceKind = "credentials" | "certificates" | "workflows"

const links = [
    ["/dashboard/certificates", "Certificates"],
    ["/dashboard/credentials", "Credentials"],
    ["/dashboard/certificates/workflows", "Workflows"],
    ["/dashboard/permissions", "Permissions"],
]

function ResourceNav() {
    const { t } = useTranslation()
    return (
        <nav aria-label={t("Unified.Navigation")} className="flex flex-wrap gap-2 border-b pb-4">
            {links.map(([path, label]) => (
                <NavLink
                    key={path}
                    to={path}
                    end
                    className={({ isActive }) =>
                        `inline-flex min-h-11 items-center rounded-md px-3 font-medium transition-colors hover:bg-muted focus-visible:outline-2 focus-visible:outline-ring ${isActive ? "bg-primary text-primary-foreground hover:bg-primary/90" : "text-muted-foreground"}`
                    }
                >
                    {t(`Unified.${label}`)}
                </NavLink>
            ))}
        </nav>
    )
}

function ResourceForm({
    kind,
    done,
}: {
    kind: Exclude<ResourceKind, "workflows">
    done: () => void
}) {
    const { t } = useTranslation()
    const [open, setOpen] = useState(false)
    const [busy, setBusy] = useState(false)
    const [error, setError] = useState("")
    const [name, setName] = useState("")
    const [provider, setProvider] = useState("cloudflare")
    const [config, setConfig] = useState("")
    const [certificate, setCertificate] = useState("")
    const [privateKey, setPrivateKey] = useState("")
    const credential = kind === "credentials"
    const title = t(credential ? "Unified.CreateCredential" : "Unified.ImportCertificate")
    const clearSecrets = () => {
        setConfig("")
        setPrivateKey("")
        setCertificate("")
        setError("")
    }
    async function submit(event: FormEvent) {
        event.preventDefault()
        setBusy(true)
        setError("")
        try {
            const input = credential
                ? { name, provider, config: JSON.parse(config) }
                : { certificate, private_key: privateKey }
            await fetcher(
                FetcherMethod.POST,
                `/api/v1/unified/${credential ? "credentials" : "certificates/import"}`,
                input,
            )
            clearSecrets()
            setName("")
            setOpen(false)
            done()
        } catch (err) {
            setError(
                err instanceof SyntaxError
                    ? t("Unified.InvalidJSON")
                    : err instanceof Error
                      ? err.message
                      : t("Unified.Failed"),
            )
        } finally {
            setBusy(false)
        }
    }
    return (
        <Dialog
            open={open}
            onOpenChange={(next) => {
                if (!busy) {
                    setOpen(next)
                    if (!next) clearSecrets()
                }
            }}
        >
            <DialogTrigger asChild>
                <Button className="min-h-11">{title}</Button>
            </DialogTrigger>
            <DialogContent className="max-h-[90dvh] overflow-y-auto sm:max-w-xl">
                <DialogHeader>
                    <DialogTitle>{title}</DialogTitle>
                </DialogHeader>
                <form className="space-y-4" onSubmit={submit}>
                    {credential ? (
                        <>
                            <div className="space-y-2">
                                <Label htmlFor="unified-name">{t("Unified.Name")}</Label>
                                <Input
                                    id="unified-name"
                                    required
                                    maxLength={120}
                                    value={name}
                                    onChange={(e) => setName(e.target.value)}
                                />
                            </div>
                            <div className="space-y-2">
                                <Label htmlFor="unified-provider">{t("Unified.Provider")}</Label>
                                <Input
                                    id="unified-provider"
                                    required
                                    maxLength={100}
                                    value={provider}
                                    onChange={(e) => setProvider(e.target.value)}
                                    list="unified-providers"
                                />
                                <datalist id="unified-providers">
                                    <option value="cloudflare" />
                                    <option value="dnspod" />
                                    <option value="tencentcloud" />
                                    <option value="ssh" />
                                    <option value="baotapanel" />
                                </datalist>
                            </div>
                            <div className="space-y-2">
                                <Label htmlFor="unified-config">{t("Unified.Config")}</Label>
                                <Textarea
                                    id="unified-config"
                                    required
                                    rows={6}
                                    maxLength={65536}
                                    autoComplete="off"
                                    spellCheck={false}
                                    value={config}
                                    onChange={(e) => setConfig(e.target.value)}
                                    placeholder={'{"apiToken": "..."}'}
                                    aria-describedby="unified-secret-help"
                                />
                            </div>
                        </>
                    ) : (
                        <>
                            <div className="space-y-2">
                                <Label htmlFor="unified-cert">{t("Unified.CertificatePEM")}</Label>
                                <Textarea
                                    id="unified-cert"
                                    required
                                    rows={6}
                                    maxLength={131072}
                                    spellCheck={false}
                                    value={certificate}
                                    onChange={(e) => setCertificate(e.target.value)}
                                />
                            </div>
                            <div className="space-y-2">
                                <Label htmlFor="unified-key">{t("Unified.PrivateKeyPEM")}</Label>
                                <Textarea
                                    id="unified-key"
                                    required
                                    rows={5}
                                    maxLength={32768}
                                    spellCheck={false}
                                    autoComplete="off"
                                    value={privateKey}
                                    onChange={(e) => setPrivateKey(e.target.value)}
                                    aria-describedby="unified-secret-help"
                                />
                            </div>
                        </>
                    )}
                    <p id="unified-secret-help" className="text-sm text-muted-foreground">
                        {t(credential ? "Unified.CredentialHelp" : "Unified.CertificateHelp")}
                    </p>
                    {error && (
                        <p role="alert" className="text-destructive">
                            {error}
                        </p>
                    )}
                    <Button type="submit" disabled={busy} className="min-h-11 w-full">
                        {busy ? t("Unified.Saving") : title}
                    </Button>
                </form>
            </DialogContent>
        </Dialog>
    )
}

export default function UnifiedResources({ kind }: { kind: ResourceKind }) {
    const { t } = useTranslation()
    const admin = useMainStore((s) => s.profile?.role === 0)
    const [page, setPage] = useState(1)
    const [actionError, setActionError] = useState("")
    const [downloading, setDownloading] = useState("")
    const { data, error, isLoading, isValidating, mutate } = useSWR<ResourcePage>(
        `/api/v1/unified/${kind}?page=${page}&per_page=20`,
        swrFetcher,
    )
    const label =
        kind === "credentials"
            ? "Credentials"
            : kind === "certificates"
              ? "Certificates"
              : "Workflows"
    async function download(item: Resource) {
        setActionError("")
        setDownloading(item.id)
        try {
            const response = await fetch(
                `/api/v1/unified/certificates/${encodeURIComponent(item.id)}/pem`,
            )
            if (!response.ok)
                throw new Error(
                    response.status === 403 ? t("Unified.DownloadDenied") : t("Unified.Failed"),
                )
            const url = URL.createObjectURL(await response.blob())
            const link = document.createElement("a")
            link.href = url
            link.download = "certificate.pem"
            link.click()
            URL.revokeObjectURL(url)
        } catch (err) {
            setActionError(err instanceof Error ? err.message : t("Unified.Failed"))
        } finally {
            setDownloading("")
        }
    }
    return (
        <main className="space-y-6 px-4 py-8">
            <ResourceNav />
            <div className="flex flex-wrap items-start justify-between gap-4">
                <div>
                    <div className="mb-2 flex items-center gap-2 text-muted-foreground">
                        {kind === "credentials" ? (
                            <KeyRound size={18} />
                        ) : (
                            <ShieldCheck size={18} />
                        )}
                        <span>Nezha Unified</span>
                    </div>
                    <h1 className="text-2xl font-semibold tracking-tight">
                        {t(`Unified.${label}`)}
                    </h1>
                    <p className="mt-2 max-w-2xl text-muted-foreground">
                        {t(`Unified.${label}Description`)}
                    </p>
                </div>
                <div className="flex gap-2">
                    <Button
                        variant="outline"
                        className="min-h-11"
                        disabled={isValidating}
                        onClick={() => void mutate()}
                        aria-label={t("Unified.Refresh")}
                    >
                        <RefreshCw size={16} />
                        {t("Unified.Refresh")}
                    </Button>
                    {admin && kind !== "workflows" && (
                        <ResourceForm
                            kind={kind}
                            done={() => {
                                setPage(1)
                                void mutate()
                            }}
                        />
                    )}
                </div>
            </div>
            {(error || actionError) && (
                <div
                    role="alert"
                    className="rounded-md border border-destructive p-4 text-destructive"
                >
                    {actionError || error.message}
                </div>
            )}
            {isLoading ? (
                <p role="status" className="py-12 text-center text-muted-foreground">
                    {t("Unified.Loading")}
                </p>
            ) : data && data.items.length === 0 ? (
                <div className="rounded-lg border border-dashed p-10 text-center">
                    <ShieldCheck className="mx-auto mb-4 text-muted-foreground" size={28} />
                    <p className="font-medium">{t("Unified.Empty")}</p>
                    <p className="mt-2 text-muted-foreground">
                        {t(kind === "workflows" ? "Unified.WorkflowPending" : "Unified.EmptyHelp")}
                    </p>
                </div>
            ) : (
                data && (
                    <div className="overflow-x-auto rounded-lg border">
                        <table className="w-full text-left text-sm">
                            <caption className="sr-only">{t(`Unified.${label}`)}</caption>
                            <thead className="bg-muted">
                                <tr>
                                    <th scope="col" className="p-4">
                                        {t("Unified.Name")}
                                    </th>
                                    <th scope="col" className="p-4">
                                        {kind === "certificates"
                                            ? t("Unified.Names")
                                            : t("Unified.Provider")}
                                    </th>
                                    <th scope="col" className="p-4">
                                        {kind === "certificates"
                                            ? t("Unified.Expires")
                                            : t("Unified.Status")}
                                    </th>
                                    {kind === "certificates" && (
                                        <th scope="col" className="p-4">
                                            {t("Unified.Actions")}
                                        </th>
                                    )}
                                </tr>
                            </thead>
                            <tbody>
                                {data.items.map((item) => (
                                    <tr key={item.id} className="border-t hover:bg-muted/50">
                                        <th
                                            scope="row"
                                            className="max-w-64 break-words p-4 font-medium"
                                        >
                                            {item.name || item.id}
                                            <span className="mt-1 block font-mono text-xs font-normal text-muted-foreground">
                                                {item.id}
                                            </span>
                                        </th>
                                        <td className="max-w-64 break-words p-4">
                                            {item.names || item.provider || "—"}
                                        </td>
                                        <td className="whitespace-nowrap p-4">
                                            {item.expires ||
                                                t(
                                                    item.enabled
                                                        ? "Unified.Enabled"
                                                        : "Unified.Disabled",
                                                )}
                                        </td>
                                        {kind === "certificates" && (
                                            <td className="p-4">
                                                <Button
                                                    variant="outline"
                                                    className="min-h-11"
                                                    disabled={downloading === item.id}
                                                    onClick={() => void download(item)}
                                                >
                                                    <Download size={16} />
                                                    {t("Unified.Download")}
                                                </Button>
                                            </td>
                                        )}
                                    </tr>
                                ))}
                            </tbody>
                        </table>
                    </div>
                )
            )}
            {data && data.total > 0 && (
                <div className="flex flex-wrap items-center justify-between gap-3">
                    <p className="text-muted-foreground">
                        {t("Unified.Pagination", {
                            total: data.total,
                            page,
                            pages: Math.ceil(data.total / data.per_page),
                        })}
                    </p>
                    <div className="flex gap-2">
                        <Button
                            variant="outline"
                            disabled={page === 1 || isLoading}
                            onClick={() => setPage(page - 1)}
                        >
                            {t("Unified.Previous")}
                        </Button>
                        <Button
                            variant="outline"
                            disabled={page * data.per_page >= data.total || isLoading}
                            onClick={() => setPage(page + 1)}
                        >
                            {t("Unified.Next")}
                        </Button>
                    </div>
                </div>
            )}
        </main>
    )
}

export function UnifiedPermissions() {
    const { t } = useTranslation()
    const admin = useMainStore((s) => s.profile?.role === 0)
    const { data, error, mutate, isLoading } = useSWR<Grant[]>("/api/v1/unified/grants", swrFetcher)
    const [userID, setUserID] = useState("")
    const [kind, setKind] = useState("certificate")
    const [resourceID, setResourceID] = useState("")
    const [action, setAction] = useState("read")
    const [busy, setBusy] = useState(false)
    const [message, setMessage] = useState("")
    async function submit(event: FormEvent) {
        event.preventDefault()
        setBusy(true)
        setMessage("")
        try {
            await fetcher(FetcherMethod.POST, "/api/v1/unified/grants", {
                user_id: Number(userID),
                kind,
                resource_id: resourceID,
                action,
            })
            setResourceID("")
            await mutate()
        } catch (err) {
            setMessage(err instanceof Error ? err.message : t("Unified.Failed"))
        } finally {
            setBusy(false)
        }
    }
    async function revoke(id: number) {
        setBusy(true)
        setMessage("")
        try {
            await fetcher(FetcherMethod.DELETE, `/api/v1/unified/grants/${id}`)
            await mutate()
        } catch (err) {
            setMessage(err instanceof Error ? err.message : t("Unified.Failed"))
        } finally {
            setBusy(false)
        }
    }
    return (
        <main className="space-y-6 px-4 py-8">
            <ResourceNav />
            <div>
                <h1 className="text-2xl font-semibold">{t("Unified.Permissions")}</h1>
                <p className="mt-2 text-muted-foreground">{t("Unified.PermissionsDescription")}</p>
            </div>
            {(error || message) && (
                <p role="alert" className="text-destructive">
                    {message || error.message}
                </p>
            )}
            {admin && (
                <form onSubmit={submit} className="grid gap-4 rounded-lg border p-5 sm:grid-cols-2">
                    <div className="space-y-2">
                        <Label htmlFor="grant-user">{t("Unified.UserID")}</Label>
                        <Input
                            id="grant-user"
                            required
                            type="number"
                            min="1"
                            step="1"
                            value={userID}
                            onChange={(e) => setUserID(e.target.value)}
                        />
                    </div>
                    <div className="space-y-2">
                        <Label htmlFor="grant-kind">{t("Unified.ResourceType")}</Label>
                        <select
                            id="grant-kind"
                            className="h-11 w-full rounded-md border bg-background px-3"
                            value={kind}
                            onChange={(e) => {
                                setKind(e.target.value)
                                setAction("read")
                            }}
                        >
                            <option value="certificate">{t("Unified.Certificates")}</option>
                            <option value="credential">{t("Unified.Credentials")}</option>
                            <option value="workflow">{t("Unified.Workflows")}</option>
                        </select>
                    </div>
                    <div className="space-y-2">
                        <Label htmlFor="grant-resource">{t("Unified.ResourceID")}</Label>
                        <Input
                            id="grant-resource"
                            required
                            value={resourceID}
                            onChange={(e) => setResourceID(e.target.value)}
                        />
                    </div>
                    <div className="space-y-2">
                        <Label htmlFor="grant-action">{t("Unified.Action")}</Label>
                        <select
                            id="grant-action"
                            className="h-11 w-full rounded-md border bg-background px-3"
                            value={action}
                            onChange={(e) => setAction(e.target.value)}
                        >
                            <option value="read">{t("Unified.ReadMetadata")}</option>
                            {kind === "certificate" && (
                                <option value="download">{t("Unified.Download")}</option>
                            )}
                        </select>
                    </div>
                    <Button className="min-h-11 sm:col-span-2" disabled={busy} type="submit">
                        {busy ? t("Unified.Saving") : t("Unified.CreateGrant")}
                    </Button>
                </form>
            )}
            {isLoading && <p role="status">{t("Unified.Loading")}</p>}
            {data && data.length === 0 && (
                <p className="rounded-lg border border-dashed p-8 text-center text-muted-foreground">
                    {t("Unified.NoGrants")}
                </p>
            )}
            {data && data.length > 0 && (
                <ul className="divide-y rounded-lg border">
                    {data.map((g) => (
                        <li
                            key={g.id}
                            className="flex flex-wrap items-center justify-between gap-3 p-4"
                        >
                            <div>
                                <span className="font-medium">
                                    {t("Unified.UserID")} {g.user_id} · {g.kind} · {g.action}
                                </span>
                                <p className="mt-1 break-all font-mono text-xs text-muted-foreground">
                                    {g.resource_id}
                                </p>
                            </div>
                            {admin && (
                                <Button
                                    variant="outline"
                                    className="min-h-11"
                                    disabled={busy}
                                    onClick={() => void revoke(g.id)}
                                >
                                    {t("Unified.Revoke")}
                                </Button>
                            )}
                        </li>
                    ))}
                </ul>
            )}
        </main>
    )
}
