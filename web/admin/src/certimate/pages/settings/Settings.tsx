import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import { Outlet, useLocation, useNavigate } from "react-router-dom";
import { IconDatabaseCog, IconHeartRateMonitor, IconInfoCircle, IconPlugConnected } from "@tabler/icons-react";
import { Menu } from "antd";

const Settings = () => {
  const location = useLocation();
  const navigate = useNavigate();

  const { t } = useTranslation();

  const menus = [
    ["ssl-provider", "settings.sslprovider.tab", <IconPlugConnected size="1em" />],
    ["persistence", "settings.persistence.tab", <IconDatabaseCog size="1em" />],
    ["diagnostics", "settings.diagnostics.tab", <IconHeartRateMonitor size="1em" />],
    ["about", "settings.about.tab", <IconInfoCircle size="1em" />],
  ] satisfies [string, string, React.ReactElement][];
  const [menuKey, setMenuKey] = useState<string>(() => location.pathname.replace("/dashboard/settings/certificates", "").split("/")[1]);
  useEffect(() => {
    const subpath = location.pathname.replace("/dashboard/settings/certificates", "").split("/")[1];
    if (!subpath) {
      navigate("/dashboard/settings/certificates/ssl-provider");
      return;
    }

    setMenuKey(subpath);
  }, [location.pathname]);

  const handleMenuClick = ({ key }: { key: string }) => {
    setMenuKey(key);
    navigate(`/dashboard/settings/certificates/${key}`);
  };

  return (
    <div className="px-6 py-4">
      <div className="container">
        <h1>{t("settings.page.heading")}</h1>
      </div>

      <div className="container">
        <div className="hidden select-none max-lg:block">
          <Menu
            style={{ background: "transparent", borderInlineEnd: "none" }}
            mode="horizontal"
            selectedKeys={[menuKey]}
            items={menus.map(([key, label, icon]) => ({
              key,
              label: t(label),
              icon: (
                <span className="anticon scale-125" role="img">
                  {icon}
                </span>
              ),
            }))}
            onClick={handleMenuClick}
          />
        </div>

        <div className="flex h-full justify-stretch gap-x-4 overflow-hidden">
          <div className="w-[256px] select-none max-lg:hidden">
            <Menu
              style={{ background: "transparent", borderInlineEnd: "none" }}
              mode="vertical"
              selectedKeys={[menuKey]}
              items={menus.map(([key, label, icon]) => ({
                key,
                label: t(label),
                icon: (
                  <span className="anticon scale-125" role="img">
                    {icon}
                  </span>
                ),
              }))}
              onClick={handleMenuClick}
            />
          </div>

          <div className="w-full flex-1">
            <div className="px-4 max-lg:px-0 max-lg:py-6">
              <Outlet />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Settings;
