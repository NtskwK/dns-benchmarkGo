import { Navbar, NavbarBrand, NavbarContent, NavbarItem, Button, Link, Tooltip } from "@nextui-org/react";
import { FaGithub as GithubIcon } from "react-icons/fa6";
import { MdDns as DnsIcon } from "react-icons/md";
import { useTranslation } from "react-i18next";

import ThemeSwitcher from "./ThemeSwitcher";
import LangSwitcher from "./LangSwitcher";
import Upload from "./Upload";

export default function NavBar() {
  const { t } = useTranslation();
  return (
    <div id="navbar">
      <Navbar isBordered isBlurred shouldHideOnScroll>
        <NavbarBrand>
          <Link href="/" color="foreground">
            <DnsIcon className="w-6 h-6 mr-2" />
            <p className="font-bold text-inherit">{t("title")}</p>
          </Link>
        </NavbarBrand>

        <NavbarContent justify="end">
          <NavbarItem>
            <Tooltip content={t("tip.github")}>
              <Link href="https://github.com/NtskwK/dns-benchmarkGo" target="_blank">
                <Button variant="ghost" aria-label={t("tip.github")}>
                  <GithubIcon />
                  <span className="ml-2">{t("tip.github")}</span>
                </Button>
              </Link>
            </Tooltip>
          </NavbarItem>
          <NavbarItem>
            <LangSwitcher />
          </NavbarItem>
          <NavbarItem>
            <ThemeSwitcher />
          </NavbarItem>
          <NavbarItem>
            <Upload />
          </NavbarItem>
        </NavbarContent>
      </Navbar>
    </div>
  );
}
