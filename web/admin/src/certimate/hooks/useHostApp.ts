import { App } from "antd"
import { type ReactNode, isValidElement } from "react"
import { toast } from "sonner"

type Message = ReactNode | { content: ReactNode; key?: string | number; duration?: number }
type Notice = {
    title?: ReactNode
    message?: ReactNode
    description?: ReactNode
    key?: string | number
    duration?: number
}
const duration = (seconds?: number) =>
    seconds === 0 ? Infinity : seconds == null ? undefined : seconds * 1000
const showMessage = (
    kind: "success" | "error" | "info" | "warning" | "loading",
    input: Message,
) => {
    const options =
        input && typeof input === "object" && !isValidElement(input) && "content" in input
            ? input
            : { content: input as ReactNode }
    return toast[kind](options.content, { id: options.key, duration: duration(options.duration) })
}
const showNotice = (kind: "success" | "error" | "info" | "warning", options: Notice) =>
    toast[kind](options.title || options.message || "", {
        id: options.key,
        description: options.description,
        duration: duration(options.duration),
    })
const message = {
    success: (input: Message) => showMessage("success", input),
    error: (input: Message) => showMessage("error", input),
    info: (input: Message) => showMessage("info", input),
    warning: (input: Message) => showMessage("warning", input),
    loading: (input: Message) => showMessage("loading", input),
    destroy: (key?: string | number) => toast.dismiss(key),
}
const notification = {
    success: (options: Notice) => showNotice("success", options),
    error: (options: Notice) => showNotice("error", options),
    info: (options: Notice) => showNotice("info", options),
    warning: (options: Notice) => showNotice("warning", options),
    destroy: (key?: string | number) => toast.dismiss(key),
}

// Keep AntD modal behavior and send all messages to the host's notification entry.
export default function useHostApp() {
    return { ...App.useApp(), message, notification }
}
