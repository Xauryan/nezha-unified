import { expect } from "@playwright/test"

import { csrfRequest, test } from "./fixtures"

test("file manager creation only accepts POST and reaches the connected Agent", async ({
    adminPage: page,
}) => {
    const serverID = process.env.E2E_AGENT_SERVER_ID
    expect(serverID, "E2E_AGENT_SERVER_ID must identify the connected fixture Agent").toBeTruthy()
    const fileURL = `/api/v1/file?id=${serverID}`
    const getResp = await page.request.get(fileURL, {
        failOnStatusCode: false,
    })
    expect(
        getResp.status() === 404 || getResp.status() === 405,
        `GET /api/v1/file must no longer be routable (got ${getResp.status()})`,
    ).toBeTruthy()

    const postResp = await csrfRequest(page, "post", fileURL, {
        failOnStatusCode: false,
    })
    expect(postResp.status()).toBe(200)
    const body = (await postResp.json()) as {
        data?: { session_id?: string }
        error?: string
        success?: boolean
    }
    expect(
        body.success,
        `POST must create a file-manager session through the real Agent: ${body.error ?? ""}`,
    ).toBe(true)
    expect(body.data?.session_id).toBeTruthy()
})
