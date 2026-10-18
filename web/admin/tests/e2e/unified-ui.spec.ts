import { type Page, type Route, expect, test } from "@playwright/test"

// Synthetic application transport; this suite never calls a CA or DNS provider.
async function nativeFixture(page: Page, baseURL: string, role = 0) {
    const writes: { path: string; body: any; method: string }[] = []
    const requests: string[] = []
    const workflow: any = {
        id: "fixtureflow0001",
        name: "Synthetic workflow",
        enabled: false,
        hasDraft: false,
        hasContent: false,
        created: "2026-10-18 00:00:00Z",
        updated: "2026-10-18 00:00:00Z",
    }
    let run: any
    let denied = false
    await page.addInitScript(() => {
        localStorage.setItem("language", "en-US")
        ;(window as any).DisableAnimatedMan = true
    })
    await page
        .context()
        .addCookies([{ name: "nz-csrf", value: "synthetic-csrf", url: baseURL }])
    await page.route("**/api/v1/**", async (route: Route) => {
        const request = route.request(),
            path = new URL(request.url()).pathname
        requests.push(path)
        const method = request.method()
        const body = request.postDataJSON()
        if (method !== "GET") writes.push({ path, body, method })
        let data: any = []
        let status = 200
        if (path === "/api/v1/profile") data = { id: 1, username: "Synthetic user", role }
        else if (path === "/api/v1/setting")
            data = { version: "test", config: { site_name: "Nezha Unified", language: "en-US" } }
        else if (path.endsWith("/resources/workflow") && method === "POST") {
            Object.assign(workflow, body)
            data = workflow
        } else if (path.endsWith("/resources/workflow/fixtureflow0001")) {
            if (method === "PATCH") Object.assign(workflow, body)
            data = workflow
        } else if (path.endsWith("/workflows/fixtureflow0001/runs")) {
            run = {
                id: "fixturerun00001",
                workflowRef: workflow.id,
                status: "processing",
                trigger: "manual",
                startedAt: "2026-10-18 00:00:00Z",
                created: "2026-10-18 00:00:00Z",
                graph: workflow.graphContent,
            }
            workflow.lastRunStatus = "processing"
            workflow.lastRunRef = run.id
            data = run
        } else if (path.endsWith("/cancel")) {
            run.status = "canceled"
            workflow.lastRunStatus = "canceled"
            data = {}
        } else if (path.endsWith("/resources/workflow_run/fixturerun00001")) data = run
        else if (path.includes("/resources/")) {
            const collection = path.split("/resources/")[1]
            const items =
                collection === "workflow_run" && run
                    ? [run]
                    : collection === "workflow"
                      ? [workflow]
                      : []
            data = { page: 1, perPage: 15, totalItems: items.length, totalPages: 1, items }
            if (collection === "certificate" && denied) status = 403
        } else if (path.endsWith("/workflows/stats"))
            data = { concurrency: 1, pendingRunIds: [], processingRunIds: [] }
        else if (path.endsWith("/system/logs")) data = { items: [], totalItems: 0, totalPages: 1 }
        else if (path.endsWith("/statistics")) data = { certificates: {}, workflows: {} }
        await route.fulfill({
            status,
            contentType: "application/json",
            body: JSON.stringify(
                status === 403
                    ? { success: false, error: "Resource access denied" }
                    : { success: true, data },
            ),
        })
    })
    return {
        writes,
        requests,
        workflow,
        deny: () => {
            denied = true
        },
    }
}

test("native workflow create, node configuration, publish, execute and cancel", async ({
    page,
    baseURL,
}, testInfo) => {
    const fixture = await nativeFixture(page, baseURL!)
    const errors: string[] = []
    page.on("pageerror", (error) => errors.push(error.message))
    await page.goto("/dashboard/certificates/workflows/new")
    await page.getByRole("button", { name: "Create from blank", exact: true }).click()
    await expect(page).toHaveURL(/fixtureflow0001\/design$/)
    await expect(page.getByRole("button", { name: "Publish", exact: true })).toBeVisible()
    await expect(page.getByPlaceholder("Please enter workflow name")).toBeHidden()
    await expect(
        page
            .getByRole("navigation", { name: "Certificate management" })
            .getByRole("link", { name: "Workflows", exact: true }),
    ).toBeVisible()
    await expect(page.locator(".workflow-design-surface")).toHaveCSS("min-height", /\d+px/)
    await page.getByText("Start", { exact: true }).first().click()
    await expect(page.locator("#certimate-module .ant-drawer")).toBeVisible()
    await page.locator(".ant-drawer").getByRole("radio", { name: "Scheduled", exact: true }).check()
    await page
        .locator(".ant-drawer")
        .getByPlaceholder("Please enter CRON expression")
        .fill("17 3 * * *")
    await page
        .locator(".ant-drawer")
        .getByRole("button", { name: "Save changes", exact: true })
        .click()
    await expect(page.locator(".ant-drawer")).toBeHidden()
    await expect
        .poll(() =>
            fixture.writes.some(
                (w) => w.body?.graphDraft?.nodes?.[0]?.data?.config?.triggerCron === "17 3 * * *",
            ),
        )
        .toBe(true)
    await page.getByRole("button", { name: "Publish", exact: true }).click()
    await expect(page.getByRole("dialog")).toBeVisible()
    await expect(page.locator("#certimate-module .ant-modal")).toBeVisible()
    await page
        .getByRole("dialog")
        .getByRole("button", { name: /OK|Confirm/ })
        .click()
    await expect.poll(() => fixture.workflow.hasContent).toBe(true)
    await expect(page.getByRole("dialog")).toBeHidden()
    await page.screenshot({ path: testInfo.outputPath("workflow-design.png"), fullPage: true })
    await page.getByRole("button", { name: "Execute", exact: true }).click()
    await expect.poll(() => fixture.writes.some((w) => w.path.endsWith("/runs"))).toBe(true)
    await page.goto("/dashboard/certificates/workflows/fixtureflow0001/runs")
    await expect(page.getByText("fixturerun00001", { exact: true })).toBeVisible()
    await page.getByRole("row").filter({ hasText: "fixturerun00001" }).getByRole("button").click()
    await page.getByRole("menuitem", { name: "Cancel", exact: true }).click()
    await page
        .getByRole("dialog")
        .getByRole("button", { name: /OK|Confirm/ })
        .click()
    await expect.poll(() => fixture.writes.some((w) => w.path.endsWith("/cancel"))).toBe(true)
    expect(fixture.requests.every((path) => path.startsWith("/api/v1/"))).toBe(true)
    expect(errors).toEqual([])
})

test("workflow node deletion preserves saved graph without delayed hover errors", async ({
    page,
    baseURL,
}) => {
    const fixture = await nativeFixture(page, baseURL!)
    const errors: string[] = []
    page.on("pageerror", (error) => errors.push(error.message))
    await page.setViewportSize({ width: 1440, height: 1080 })
    await page.goto("/dashboard/certificates/workflows/new")
    await page.getByRole("button", { name: "Create from blank", exact: true }).click()
    await expect(page).toHaveURL(/fixtureflow0001\/design$/)

    const adder = page.locator(".flow-canvas-adder").first()
    await adder.hover()
    await adder.getByRole("button").click()
    await page.getByRole("menuitem", { name: "Delay", exact: true }).click()
    await expect.poll(() => fixture.workflow.graphDraft.nodes.length).toBe(3)
    const delayId = fixture.workflow.graphDraft.nodes[1].id as string
    const delay = page.locator(`[data-node-id="${delayId}"]`)
    await delay.getByText("Delay", { exact: true }).click()
    await page.locator(".ant-drawer").getByRole("spinbutton").fill("7")
    await page.locator(".ant-drawer").getByRole("button", { name: "Save changes", exact: true }).click()
    await expect(page.locator(".ant-drawer")).toBeHidden()
    await expect.poll(() => fixture.workflow.graphDraft.nodes[1].data.config.wait).toBe(7)

    await delay.hover()
    await page.locator(".ant-popover:visible button").click()
    await page.getByRole("menuitem", { name: "Duplicate node", exact: true }).click()
    await expect.poll(() => fixture.workflow.graphDraft.nodes.length).toBe(4)
    const copyId = fixture.workflow.graphDraft.nodes[2].id as string
    expect(copyId).not.toBe(delayId)
    const copy = page.locator(`[data-node-id="${copyId}"]`)
    await expect(copy.getByText("Delay-copy", { exact: true })).toBeVisible()
    expect(fixture.workflow.graphDraft.nodes[2].data.config.wait).toBe(7)

    await copy.hover()
    const grip = await copy.locator(".tabler-icon-grip-vertical").boundingBox()
    const destination = await page.locator(".flow-canvas-adder").first().boundingBox()
    expect(grip).not.toBeNull()
    expect(destination).not.toBeNull()
    await page.mouse.move(grip!.x + grip!.width / 2, grip!.y + grip!.height / 2)
    await page.mouse.down()
    await page.mouse.move(
        destination!.x + destination!.width / 2,
        destination!.y + destination!.height / 2,
        { steps: 30 },
    )
    await page.mouse.up()
    await expect.poll(() => fixture.workflow.graphDraft.nodes[1].id).toBe(copyId)

    await delay.hover()
    await page.locator(".ant-popover:visible button").click()
    await page.getByRole("menuitem", { name: "Delete node", exact: true }).click()
    await expect(delay).toHaveCount(0)
    await expect.poll(() => fixture.workflow.graphDraft.nodes.length).toBe(3)
    // The canvas schedules mouseleave cleanup for 200ms; deleting a hovered node must
    // remain safe after that callback runs, as well as after the draft autosave.
    await page.waitForTimeout(350)
    expect(errors).toEqual([])
    const savedGraph = structuredClone(fixture.workflow.graphDraft)
    expect(savedGraph.nodes.map((node: { id: string }) => node.id)).not.toContain(delayId)
    expect(savedGraph.nodes[1].id).toBe(copyId)

    await page.reload()
    await expect(copy.getByText("Delay-copy", { exact: true })).toBeVisible()
    await expect(delay).toHaveCount(0)
    expect(fixture.workflow.graphDraft).toEqual(savedGraph)
    await copy.getByText("Delay-copy", { exact: true }).click()
    await expect(page.locator(".ant-drawer").getByRole("spinbutton")).toHaveValue("7")
    await page.locator(".ant-drawer").getByRole("button", { name: "Cancel", exact: true }).click()
    await expect(page.locator(".ant-drawer")).toBeHidden()
    await page.getByText("History runs", { exact: true }).click()
    await expect(page).toHaveURL(/fixtureflow0001\/runs$/)
    await page.goBack()
    await expect(page).toHaveURL(/fixtureflow0001\/design$/)
    await expect(copy).toBeVisible()
    await expect(delay).toHaveCount(0)
    await page.goForward()
    await expect(page).toHaveURL(/fixtureflow0001\/runs$/)
    await page.goBack()
    await expect(copy).toBeVisible()
    await expect(delay).toHaveCount(0)
    expect(fixture.workflow.graphDraft).toEqual(savedGraph)
    expect(errors).toEqual([])
})

test("certificate permission denial retains Nezha session and scoped dark theme", async ({
    page,
    baseURL,
}) => {
    const fixture = await nativeFixture(page, baseURL!, 1)
    fixture.deny()
    await page.addInitScript(() => localStorage.setItem("vite-ui-theme", "dark"))
    await page.goto("/dashboard/certificates")
    await expect(page.getByText("Resource access denied").first()).toBeVisible()
    await expect(page.locator("html")).toHaveClass(/dark/)
    await expect(page.locator("#certimate-module")).toBeVisible()
    await page.goto("/dashboard/permissions")
    await expect(
        page.getByRole("heading", { name: "Resource permissions", exact: true }),
    ).toBeVisible()
    await expect(page.locator("#grant-user")).toHaveCount(0)
    await expect(page).not.toHaveURL(/login/)
})

test("native management deep links render on mobile", async ({ page, baseURL }) => {
    await nativeFixture(page, baseURL!)
    await page.setViewportSize({ width: 390, height: 844 })
    for (const [route, title] of [
        ["dns", "DNS management"],
        ["targets", "Deployment targets"],
        ["identities", "Execution identities & scoped tokens"],
    ]) {
        await page.goto(`/dashboard/${route}`)
        await expect(page.getByRole("heading", { name: title, exact: true })).toBeVisible()
        expect(await page.evaluate(() => document.documentElement.scrollWidth <= innerWidth)).toBe(
            true,
        )
    }
})
