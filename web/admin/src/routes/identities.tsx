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

type Scope = { kind: string; resource_id: string; action: string }
type Token = { id: number; name: string; scopes: Scope[]; expires_at: string; revoked_at?: string }
type Identity = {
    id: number
    name: string
    owner_id: number
    active: boolean
    version: number
    scopes: Scope[]
}
export default function IdentitiesPage() {
    const t = useText(),
        admin = useMainStore((s) => s.profile?.role === 0),
        action = useAction()
    const {
        data: tokens,
        mutate: refreshTokens,
        error,
    } = useSWR<Token[]>("/api/v1/unified/tokens", swrFetcher)
    const {
        data: identities,
        mutate: refreshIdentities,
        error: identityError,
    } = useSWR<Identity[]>("/api/v1/unified/identities", swrFetcher)
    const [tokenName, setTokenName] = useState(""),
        [expires, setExpires] = useState(""),
        [tokenScopes, setTokenScopes] = useState("[]"),
        [secret, setSecret] = useState("")
    const [name, setName] = useState(""),
        [owner, setOwner] = useState(""),
        [scopes, setScopes] = useState("[]"),
        [active, setActive] = useState(true),
        [edit, setEdit] = useState<Identity>()
    const help = t(
        '精确授权数组，例如 [{"kind":"workflow","resource_id":"工作流 ID","action":"execute"}]；不支持通配符。',
        'Exact scopes, e.g. [{"kind":"workflow","resource_id":"workflow ID","action":"execute"}]. Wildcards are not supported.',
    )
    return (
        <main className="space-y-6 px-4 py-8">
            <UnifiedNav />
            <h1 className="text-2xl font-semibold">
                {t("执行身份与受限令牌", "Execution identities & scoped tokens")}
            </h1>
            <ErrorNotice error={action.error || error || identityError} />
            <Panel title={t("创建受限令牌", "Create scoped token")}>
                <form
                    className="space-y-4"
                    onSubmit={(e) => {
                        e.preventDefault()
                        void action.run(async () => {
                            const result = await fetcher<{ secret: string }>(
                                M.POST,
                                "/api/v1/unified/tokens",
                                {
                                    name: tokenName,
                                    expires_at: new Date(expires).toISOString(),
                                    scopes: parseArray<Scope>(tokenScopes),
                                },
                            )
                            setSecret(result.secret)
                            setTokenName("")
                            await refreshTokens()
                        })
                    }}
                >
                    <div className="grid gap-4 sm:grid-cols-2">
                        <Field
                            id="token-name"
                            label={t("名称", "Name")}
                            value={tokenName}
                            onChange={setTokenName}
                            required
                        />
                        <Field
                            id="token-expires"
                            label={t("过期时间（本地时区）", "Expires at (local time)")}
                            type="datetime-local"
                            value={expires}
                            onChange={setExpires}
                            required
                        />
                    </div>
                    <JSONField
                        id="token-scopes"
                        label={t("资源与操作范围", "Resources and actions")}
                        value={tokenScopes}
                        onChange={setTokenScopes}
                        help={help}
                    />
                    <Button disabled={action.busy}>{t("创建令牌", "Create token")}</Button>
                </form>
            </Panel>
            {secret && (
                <section className="space-y-3 rounded-lg border p-4">
                    <p className="font-medium">
                        {t(
                            "令牌仅在本次显示，请立即保存。",
                            "This token is shown once. Save it now.",
                        )}
                    </p>
                    <input
                        aria-label={t("新令牌", "New token")}
                        type="password"
                        readOnly
                        value={secret}
                        className="w-full rounded-md border bg-background p-3 font-mono"
                        autoComplete="off"
                    />
                    <div className="flex gap-2">
                        <Button
                            variant="outline"
                            onClick={() =>
                                void action.run(() => navigator.clipboard.writeText(secret))
                            }
                        >
                            {t("复制", "Copy")}
                        </Button>
                        <Button variant="outline" onClick={() => setSecret("")}>
                            {t("已保存，关闭", "Saved, dismiss")}
                        </Button>
                    </div>
                </section>
            )}
            <section className="space-y-3">
                <h2 className="text-lg font-semibold">{t("我的令牌", "My tokens")}</h2>
                {tokens?.length === 0 && (
                    <p className="text-muted-foreground">{t("暂无令牌。", "No tokens yet.")}</p>
                )}
                {tokens?.map((token) => (
                    <article key={token.id} className="rounded-lg border p-4">
                        <div className="flex items-center justify-between gap-4">
                            <div>
                                <h3 className="font-medium">{token.name}</h3>
                                <p className="text-xs text-muted-foreground">
                                    #{token.id} · {token.expires_at} ·{" "}
                                    {token.revoked_at
                                        ? t("已撤销", "Revoked")
                                        : t("有效", "Active")}
                                </p>
                            </div>
                            <Button
                                variant="outline"
                                disabled={action.busy || !!token.revoked_at}
                                onClick={() =>
                                    void action.run(async () => {
                                        await fetcher(
                                            M.DELETE,
                                            `/api/v1/unified/tokens/${token.id}`,
                                        )
                                        await refreshTokens()
                                    })
                                }
                            >
                                {t("撤销", "Revoke")}
                            </Button>
                        </div>
                        <ul className="mt-3 space-y-1 font-mono text-xs">
                            {token.scopes.map((s, i) => (
                                <li key={i}>
                                    {s.kind} / {s.resource_id} / {s.action}
                                </li>
                            ))}
                        </ul>
                    </article>
                ))}
            </section>
            {admin && (
                <Panel
                    title={
                        edit
                            ? t("编辑执行身份", "Edit execution identity")
                            : t("创建执行身份", "Create execution identity")
                    }
                    open={!!edit}
                >
                    <form
                        className="space-y-4"
                        onSubmit={(e) => {
                            e.preventDefault()
                            void action.run(async () => {
                                await fetcher(
                                    edit ? M.PATCH : M.POST,
                                    `/api/v1/unified/identities${edit ? "/" + edit.id : ""}`,
                                    {
                                        name,
                                        version: edit?.version,
                                        owner_id: Number(owner),
                                        active,
                                        scopes: parseArray<Scope>(scopes),
                                    },
                                )
                                setEdit(undefined)
                                setName("")
                                setOwner("")
                                setScopes("[]")
                                await refreshIdentities()
                            })
                        }}
                    >
                        <div className="grid gap-4 sm:grid-cols-2">
                            <Field
                                id="identity-name"
                                label={t("名称", "Name")}
                                value={name}
                                onChange={setName}
                                required
                            />
                            <Field
                                id="identity-owner"
                                label={t("所属用户 ID", "Owner user ID")}
                                type="number"
                                min={1}
                                value={owner}
                                onChange={setOwner}
                                required
                            />
                        </div>
                        <JSONField
                            id="identity-scopes"
                            label={t("资源与操作范围", "Resources and actions")}
                            value={scopes}
                            onChange={setScopes}
                            help={help}
                        />
                        <Check
                            label={t("启用身份", "Identity active")}
                            checked={active}
                            onChange={setActive}
                        />
                        <div className="flex gap-2">
                            <Button disabled={action.busy}>{t("保存", "Save")}</Button>
                            {edit && (
                                <Button
                                    type="button"
                                    variant="outline"
                                    onClick={() => {
                                        setEdit(undefined)
                                        setName("")
                                        setOwner("")
                                        setScopes("[]")
                                    }}
                                >
                                    {t("取消", "Cancel")}
                                </Button>
                            )}
                        </div>
                    </form>
                </Panel>
            )}
            <section className="space-y-3">
                <h2 className="text-lg font-semibold">{t("执行身份", "Execution identities")}</h2>
                {identities?.map((identity) => (
                    <article key={identity.id} className="rounded-lg border p-4">
                        <div className="flex flex-wrap justify-between gap-4">
                            <div>
                                <h3 className="font-medium">
                                    {identity.name} #{identity.id}
                                </h3>
                                <p className="text-xs text-muted-foreground">
                                    {t("所属用户", "Owner")} #{identity.owner_id} · v
                                    {identity.version} ·{" "}
                                    {identity.active ? t("启用", "Active") : t("停用", "Inactive")}
                                </p>
                            </div>
                            {admin && (
                                <Button
                                    variant="outline"
                                    onClick={() => {
                                        setEdit(identity)
                                        setName(identity.name)
                                        setOwner(String(identity.owner_id))
                                        setScopes(JSON.stringify(identity.scopes, null, 2))
                                        setActive(identity.active)
                                    }}
                                >
                                    {t("编辑", "Edit")}
                                </Button>
                            )}
                        </div>
                        <ul className="mt-3 space-y-1 font-mono text-xs">
                            {identity.scopes.map((s, i) => (
                                <li key={i}>
                                    {s.kind} / {s.resource_id} / {s.action}
                                </li>
                            ))}
                        </ul>
                    </article>
                ))}
            </section>
        </main>
    )
}
