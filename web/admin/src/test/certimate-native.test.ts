import { getNativeClient } from "@cert/repository/_native"
import { remove as removeCredential } from "@cert/repository/access"
import { useMainStore } from "@/hooks/useMainStore"
import { type ModelProfile } from "@/types"
import { afterEach, beforeEach, describe, expect, it, vi } from "vitest"

describe("Nezha certificate transport", () => {
    const client = getNativeClient()
    beforeEach(() => {
        document.cookie = "nz-csrf=synthetic-csrf; path=/"
    })
    afterEach(async () => {
        await client.collection("workflow").unsubscribe()
        vi.useRealTimers()
        vi.unstubAllGlobals()
        useMainStore.getState().setProfile(undefined)
    })
    it("uses the current Nezha session and adds CSRF only to mutations", async () => {
        const fetchMock = vi
            .fn()
            .mockResolvedValue({
                ok: true,
                json: async () => ({ success: true, data: { id: "fixture" } }),
            })
        vi.stubGlobal("fetch", fetchMock)
        await client.collection("workflow").getOne("fixture")
        await client.collection("workflow").update("fixture", { name: "Synthetic workflow" })
        const [getURL, getOptions] = fetchMock.mock.calls[0]
        expect(getURL.pathname).toBe("/api/v1/unified/certimate/resources/workflow/fixture")
        expect(getOptions.credentials).toBe("same-origin")
        expect(getOptions.headers["X-CSRF-Token"]).toBeUndefined()
        expect(fetchMock.mock.calls[1][1].headers["X-CSRF-Token"]).toBe("synthetic-csrf")
        expect(fetchMock.mock.calls[1][1].headers.Authorization).toBeUndefined()
        expect(client).not.toHaveProperty("authStore")
    })
    it("retains permission errors without logging out the host", async () => {
        vi.stubGlobal(
            "fetch",
            vi
                .fn()
                .mockResolvedValue({
                    ok: false,
                    status: 403,
                    json: async () => ({ success: false, error: "Resource access denied" }),
                }),
        )
        await expect(client.collection("certificate").getOne("denied")).rejects.toMatchObject({
            status: 403,
            response: { message: "Resource access denied" },
        })
        expect(document.cookie).toContain("nz-csrf=synthetic-csrf")
    })
    it("uses DELETE for credential deletion instead of mutating metadata", async () => {
        const fetchMock = vi.fn().mockResolvedValue({ ok: true, json: async () => ({ success: true, data: true }) })
        vi.stubGlobal("fetch", fetchMock)
        await removeCredential({ id: "fixture" })
        expect(fetchMock.mock.calls[0][1].method).toBe("DELETE")
        expect(fetchMock.mock.calls[0][1].body).toBeUndefined()
    })
    it("discards an old identity's pending response after login changes", async () => {
        useMainStore.getState().setProfile({ id: 1, username: "first", role: 0 } as ModelProfile)
        let complete!: (response: unknown) => void
        vi.stubGlobal("fetch", vi.fn().mockImplementation(() => new Promise(resolve => { complete = resolve })))
        const pending = client.collection("access").getOne("fixture")
        const rejected = expect(pending).rejects.toMatchObject({ isAbort: true })
        useMainStore.getState().setProfile({ id: 2, username: "second", role: 1 } as ModelProfile)
        complete({ ok: true, json: async () => ({ success: true, data: { id: "fixture" } }) })
        await rejected
    })
    it("stops delivery from an outstanding poll after route disposal", async () => {
        vi.useFakeTimers()
        let complete!: (value: unknown) => void
        const fetchMock = vi.fn().mockImplementation(
            () =>
                new Promise((resolve) => {
                    complete = resolve
                }),
        )
        vi.stubGlobal("fetch", fetchMock)
        const callback = vi.fn()
        const stop = await client.collection("workflow").subscribe("fixture", callback)
        stop()
        complete({ ok: true, json: async () => ({ success: true, data: { id: "fixture" } }) })
        await vi.advanceTimersByTimeAsync(10000)
        expect(callback).not.toHaveBeenCalled()
        expect(fetchMock).toHaveBeenCalledTimes(1)
    })
    it("escapes filter parameters and never permits alternate business origins", async () => {
        expect(client.filter("name={:value}", { value: "a'b\\c" })).toBe("name='a\\'b\\\\c'")
        await expect(client.send("https://example.invalid/api/workflows")).rejects.toThrow(
            "Unsupported native business endpoint",
        )
    })
})
