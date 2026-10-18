import react from "@vitejs/plugin-react"
import path from "path"
import { defineConfig } from "vite"

export default defineConfig({
    base: "/dashboard",
    plugins: [react()],
    server: {
        // Bind the dev server to loopback so an in-browser CSRF or LAN
        // scan cannot reach a developer's dashboard token. Override with
        // `bun run dev -- --host 0.0.0.0` for Docker / remote dev only.
        host: "127.0.0.1",
        proxy: {
            "^/api/v1/ws/.*": {
                target: "ws://127.0.0.1:8008",
                changeOrigin: true,
                ws: true,
            },
            "/api": {
                target: "http://127.0.0.1:8008",
                changeOrigin: true,
            },
            "/mcp": {
                target: "http://127.0.0.1:8008",
                changeOrigin: true,
            },
        },
    },
    resolve: {
        alias: {
            "@": path.resolve(import.meta.dirname, "./src"),
            "@cert": path.resolve(import.meta.dirname, "./src/certimate"),
        },
    },
    build: {
        cssCodeSplit: true,
        sourcemap: false,
        chunkSizeWarningLimit: 1500,
        // Let the bundler preserve lazy route boundaries; a catch-all vendor
        // chunk would pull the complete certificate designer into every host page.
    },
})
