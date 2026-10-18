import { useMainStore } from "@/hooks/useMainStore"

// Metadata caches belong to the active Nezha identity, never to the browser tab.
let generation = 0
const resetters = new Set<() => void>()
useMainStore.subscribe((state, previous) => {
    if (state.profile?.id === previous.profile?.id) return
    generation++
    for (const reset of resetters) reset()
})
export const sessionGeneration = () => generation
export const registerSessionReset = (reset: () => void) => {
    resetters.add(reset)
}
