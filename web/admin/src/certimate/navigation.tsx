import { useTranslation } from "react-i18next"
import { NavLink } from "react-router-dom"

const entries = [
    ["/dashboard/certificates/overview", "dashboard.page.heading"],
    ["/dashboard/certificates", "certificate.page.heading"],
    ["/dashboard/certificates/workflows", "workflow.page.heading"],
    ["/dashboard/credentials", "access.page.heading"],
    ["/dashboard/certificates/presets", "preset.page.heading"],
    ["/dashboard/settings/certificates", "settings.page.heading"],
    ["/dashboard/runs", "Unified.RunHistory"],
    ["/dashboard/targets", "Unified.DeploymentTargets"],
    ["/dashboard/permissions", "Unified.Permissions"],
    ["/dashboard/identities", "Unified.Identities"],
]
export default function CertificateNav() {
    const { t } = useTranslation()
    return (
        <nav
            aria-label="Certificate management"
            className="flex flex-wrap gap-2 border-b px-4 py-4"
        >
            {entries.map(([path, key]) => (
                <NavLink
                    key={path}
                    to={path}
                    end={path === "/dashboard/certificates"}
                    className={({ isActive }) =>
                        `inline-flex min-h-11 items-center rounded-md px-3 transition-colors hover:bg-muted focus-visible:outline-2 focus-visible:outline-ring ${isActive ? "bg-primary text-primary-foreground" : "text-muted-foreground"}`
                    }
                >
                    {t(key)}
                </NavLink>
            ))}
        </nav>
    )
}
