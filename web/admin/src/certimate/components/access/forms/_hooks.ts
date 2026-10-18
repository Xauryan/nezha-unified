import { useMemo } from "react";

import { ACCESS_USAGES, type AccessProvider } from "@cert/domain/provider";

export const useProviderFilterByUsage = (usage?: "dns" | "hosting" | "dns-hosting" | "ca" | "notification") => {
  return useMemo(() => {
    if (usage == null) return;

    switch (usage) {
      case "dns":
        return (_props: string, option: AccessProvider) => option.usages.includes(ACCESS_USAGES.DNS);
      case "hosting":
        return (_props: string, option: AccessProvider) => option.usages.includes(ACCESS_USAGES.HOSTING);
      case "dns-hosting":
        return (_props: string, option: AccessProvider) => option.usages.includes(ACCESS_USAGES.DNS) || option.usages.includes(ACCESS_USAGES.HOSTING);
      case "ca":
        return (_props: string, option: AccessProvider) => option.usages.includes(ACCESS_USAGES.CA);
      case "notification":
        return (_props: string, option: AccessProvider) => option.usages.includes(ACCESS_USAGES.NOTIFICATION);
      default:
        console.warn(`[certimate] unsupported provider usage: '${usage}'`);
    }
  }, [usage]);
};
