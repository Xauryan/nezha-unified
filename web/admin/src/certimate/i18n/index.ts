import i18n from "@/lib/i18n";
import { localeNames } from "./locales";
import resources from "./resources";
for (const language of ["en", "zh", ...Object.keys(i18n.options.resources || {})]) {
 i18n.addResourceBundle(language,"translation",resources[language.startsWith("zh") ? "zh" : "en"].translation,true,false);
}

export { localeNames };
export const localeResources = resources;

export { useAntdLocale } from "./vendors/antd";
export { useDayjsLocale } from "./vendors/dayjs";
export { useZodLocale } from "./vendors/zod";

export default i18n;
