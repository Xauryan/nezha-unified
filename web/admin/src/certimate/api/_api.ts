import { ClientResponseError, type SendOptions } from "pocketbase";

import { getNativeClient } from "@cert/repository/_native";

const pb = getNativeClient();

type Options = Pick<SendOptions, "body" | "headers"> & {
  url: string;
};

export async function get<T = any>({ url, ...opts }: Options) {
  const resp = await pb.send<BaseResponse<T>>(url, {
    method: "GET",
    ...opts,
  });

  if (resp.code != 0) {
    throw new ClientResponseError({ status: resp.code, response: resp, data: {} });
  }

  return resp;
}

export async function post<T = any>({ url, headers, ...opts }: Options) {
  const resp = await pb.send<BaseResponse<T>>(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      ...headers,
    },
    ...opts,
  });

  if (resp.code != 0) {
    throw new ClientResponseError({ status: resp.code, response: resp, data: {} });
  }

  return resp;
}
