// Native Nezha transport: no PocketBase client, independent session or authStore.
import { sessionGeneration } from "@cert/session"
import { ClientResponseError, type RecordModel, type RecordSubscription } from "pocketbase"

type Options = {
    filter?: string
    sort?: string
    fields?: string
    expand?: string
    requestKey?: string | null
    batch?: number
    skipTotal?: boolean
    page?: number
    perPage?: number
}
type Page<T> = { page: number; perPage: number; totalItems: number; totalPages: number; items: T[] }
const base = "/api/v1/unified/certimate"
async function request<T>(
    path: string,
    method = "GET",
    body?: unknown,
    options?: Options,
): Promise<T> {
    const generation = sessionGeneration()
    const url = new URL(base + path, window.location.origin)
    for (const [key, value] of Object.entries(options || {}))
        if (value != null && key !== "requestKey" && key !== "batch")
            url.searchParams.set(key, String(value))
    const headers: Record<string, string> = { "Content-Type": "application/json" }
    if (method !== "GET") {
        const csrf = document.cookie
            .split(";")
            .map((v) => v.trim())
            .find((v) => v.startsWith("nz-csrf="))
        if (csrf) headers["X-CSRF-Token"] = csrf.slice("nz-csrf=".length)
    }
    const response = await fetch(url, {
        method,
        headers,
        credentials: "same-origin",
        body: body == null ? undefined : JSON.stringify(body),
    })
    const payload = await response.json()
    if (generation !== sessionGeneration())
        throw new ClientResponseError({
            status: 0,
            isAbort: true,
            response: { message: "Session changed" },
        })
    if (!response.ok || !payload.success)
        throw new ClientResponseError({
            status: response.ok ? 403 : response.status,
            response: { message: payload.error || response.statusText, data: payload.data || {} },
        })
    return payload.data
}
class Collection {
    private subscriptions = new Map<string, Set<() => void>>()
    constructor(private name: string) {}
    private path(id?: string) {
        return `/resources/${encodeURIComponent(this.name)}${id ? "/" + encodeURIComponent(id) : ""}`
    }
    getList<T = RecordModel>(page = 1, perPage = 20, options: Options = {}) {
        return request<Page<T>>(this.path(), "GET", undefined, { ...options, page, perPage })
    }
    async getFullList<T = RecordModel>(options: Options = {}) {
        const items: T[] = []
        for (let page = 1; page <= 10000; page++) {
            const result = await this.getList<T>(page, 200, options)
            items.push(...result.items)
            if (page >= result.totalPages || !result.items.length) return items
        }
        throw new Error("Resource pagination limit exceeded")
    }
    getOne<T = RecordModel>(id: string, options: Options = {}) {
        return request<T>(this.path(id), "GET", undefined, options)
    }
    async getFirstListItem<T = RecordModel>(filter: string, options: Options = {}) {
        const result = await this.getList<T>(1, 1, { ...options, filter })
        if (!result.items.length)
            throw new ClientResponseError({
                status: 404,
                response: { message: "Record not found" },
            })
        return result.items[0]
    }
    create<T = RecordModel>(body: unknown) {
        return request<T>(this.path(), "POST", body)
    }
    update<T = RecordModel>(id: string, body: unknown) {
        return request<T>(this.path(id), "PATCH", body)
    }
    delete(id: string) {
        return request<boolean>(this.path(id), "DELETE")
    }
    async subscribe<T = RecordModel>(id: string, callback: (event: RecordSubscription<T>) => void) {
        let active = true,
            busy = false,
            previous = ""
        const tick = async () => {
            if (!active || busy) return
            busy = true
            try {
                const record = await this.getOne<T>(id)
                const current = JSON.stringify(record)
                if (active && current !== previous) {
                    previous = current
                    callback({ action: "update", record } as RecordSubscription<T>)
                }
            } catch {
                /* Authority is checked on every poll; a denial never clears Nezha login. */
            } finally {
                busy = false
            }
        }
        const timer = setInterval(() => void tick(), 2000)
        const stop = () => {
            active = false
            clearInterval(timer)
            this.subscriptions.get(id)?.delete(stop)
        }
        if (!this.subscriptions.has(id)) this.subscriptions.set(id, new Set())
        this.subscriptions.get(id)!.add(stop)
        void tick()
        return stop
    }
    async unsubscribe(id?: string) {
        for (const [key, stops] of this.subscriptions) {
            if (id === undefined || id === key) {
                for (const stop of stops) stop()
                this.subscriptions.delete(key)
            }
        }
    }
}
const collections = new Map<string, Collection>()
const client = {
    collection(name: string) {
        if (!collections.has(name)) collections.set(name, new Collection(name))
        return collections.get(name)!
    },
    filter(expression: string, params: Record<string, unknown> = {}) {
        return expression.replace(/\{:(\w+)\}/g, (_, key) => {
            const value = params[key]
            if (value == null) return "null"
            if (typeof value === "number" || typeof value === "boolean") return String(value)
            const text =
                value instanceof Date ? value.toISOString().replace("T", " ") : String(value)
            return "'" + text.replace(/\\/g, "\\\\").replace(/'/g, "\\'") + "'"
        })
    },
    createBatch() {
        const operations: { collection: string; method: string; id?: string; body?: unknown }[] = []
        return {
            collection: (collection: string) => ({
                update: (id: string, body: unknown) =>
                    operations.push({ collection, method: "PATCH", id, body }),
                create: (body: unknown) => operations.push({ collection, method: "POST", body }),
                delete: (id: string) => operations.push({ collection, method: "DELETE", id }),
            }),
            send: () => request<{ status: number }[]>("/batch", "POST", { operations }),
        }
    },
    async send<T>(
        url: string,
        options: { method?: string; body?: unknown; headers?: unknown } = {},
    ): Promise<T> {
        if (!url.startsWith("/api/")) throw new Error("Unsupported native business endpoint")
        const data = await request(url.slice(4), options.method || "GET", options.body)
        return { code: 0, data } as T
    },
    crons: {
        getFullList: (_options?: Options) =>
            request<{ id: string; expression: string }[]>("/system/crons"),
    },
    logs: {
        getList: (page: number, perPage: number, _options?: Options) =>
            request<Page<RecordModel>>(`/system/logs?page=${page}&perPage=${perPage}`),
    },
}
export const getNativeClient = () => client
export const COLLECTION_NAME_ACCESS = "access"
export const COLLECTION_NAME_CERTIFICATE = "certificate"
export const COLLECTION_NAME_SETTINGS = "settings"
export const COLLECTION_NAME_WORKFLOW = "workflow"
export const COLLECTION_NAME_WORKFLOW_RUN = "workflow_run"
export const COLLECTION_NAME_WORKFLOW_OUTPUT = "workflow_output"
export const COLLECTION_NAME_WORKFLOW_LOG = "workflow_logs"
