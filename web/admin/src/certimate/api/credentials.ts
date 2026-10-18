import { type AccessModel } from "@cert/domain/access"

import { post } from "./_api"

export const duplicate = (id: string, name: string) =>
    post<AccessModel>({
        url: `/api/credentials/${encodeURIComponent(id)}/duplicate`,
        body: { name },
    })
export const exportMaterial = (id: string) =>
    post<{ config: Record<string, unknown> }>({
        url: `/api/credentials/${encodeURIComponent(id)}/export`,
    })
