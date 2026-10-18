import { Suspense, lazy } from "react"
import type { RouteObject } from "react-router-dom"

const Shell = lazy(() => import("./Shell"))
const Overview = lazy(() => import("./pages/dashboard/Dashboard"))
const Certificates = lazy(() => import("./pages/certificates/CertificateList"))
const Accesses = lazy(() => import("./pages/accesses/AccessList"))
const AccessNew = lazy(() => import("./pages/accesses/AccessNew"))
const Workflows = lazy(() => import("./pages/workflows/WorkflowList"))
const WorkflowNew = lazy(() => import("./pages/workflows/WorkflowNew"))
const WorkflowDetail = lazy(() => import("./pages/workflows/WorkflowDetail"))
const WorkflowDesign = lazy(() => import("./pages/workflows/WorkflowDetailDesign"))
const WorkflowRuns = lazy(() => import("./pages/workflows/WorkflowDetailRuns"))
const Presets = lazy(() => import("./pages/presets/PresetList"))
const Settings = lazy(() => import("./pages/settings/Settings"))
const SSLProvider = lazy(() => import("./pages/settings/SettingsSSLProvider"))
const Persistence = lazy(() => import("./pages/settings/SettingsPersistence"))
const Diagnostics = lazy(() => import("./pages/settings/SettingsDiagnostics"))
const About = lazy(() => import("./pages/settings/SettingsAbout"))
const shell = (
    <Suspense
        fallback={
            <p role="status" className="p-8">
                Loading certificate resources…
            </p>
        }
    >
        <Shell />
    </Suspense>
)
export const certificateRoutes: RouteObject[] = [
    {
        path: "/dashboard/certificates",
        element: shell,
        children: [
            { index: true, element: <Certificates /> },
            { path: "overview", element: <Overview /> },
            { path: "workflows", element: <Workflows /> },
            { path: "workflows/new", element: <WorkflowNew /> },
            {
                path: "workflows/:id",
                element: <WorkflowDetail />,
                children: [
                    { path: "design", element: <WorkflowDesign /> },
                    { path: "runs", element: <WorkflowRuns /> },
                ],
            },
            { path: "presets", element: <Presets /> },
        ],
    },
    {
        path: "/dashboard/credentials",
        element: shell,
        children: [
            { index: true, element: <Accesses /> },
            { path: "new", element: <AccessNew /> },
        ],
    },
    {
        path: "/dashboard/settings/certificates",
        element: shell,
        children: [
            {
                element: <Settings />,
                children: [
                    { path: "ssl-provider", element: <SSLProvider /> },
                    { path: "persistence", element: <Persistence /> },
                    { path: "diagnostics", element: <Diagnostics /> },
                    { path: "about", element: <About /> },
                    { index: true, element: <SSLProvider /> },
                ],
            },
        ],
    },
]
