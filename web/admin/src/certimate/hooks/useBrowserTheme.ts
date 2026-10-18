import { useTheme } from "@/components/theme-provider";
import { useMediaQuery } from "@/hooks/useMediaQuery";
const useBrowserTheme=()=>{const host=useTheme();const systemDark=useMediaQuery("(prefers-color-scheme: dark)");return{theme:host.theme==="system"?(systemDark?"dark":"light"):host.theme,themeMode:host.theme,setThemeMode:host.setTheme};};
export type UseBrowserThemeReturns=ReturnType<typeof useBrowserTheme>;
export default useBrowserTheme;
