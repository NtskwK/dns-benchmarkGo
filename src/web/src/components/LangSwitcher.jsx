import { useTranslation } from "react-i18next";
import { Button, Dropdown, DropdownTrigger, DropdownMenu, DropdownItem } from "@nextui-org/react";
import { FaLanguage as LanguageIcon } from "react-icons/fa6";

import { resources } from "../locales/resources";

export default function LangSwitcher() {
  const { i18n, t } = useTranslation();

  const toggleLanguage = (lang: string) => {
    i18n.changeLanguage(lang);
  };

  // 获取当前语言，如果不存在则默认使用 'en'
  const currentLang = i18n.language?.split("-")[0] || "en";
  const displayName = resources[currentLang as keyof typeof resources]?.name || resources.en.name;

  return (
    <Dropdown>
      <DropdownTrigger>
        <Button variant="ghost" aria-label={t("tip.language")}>
          <LanguageIcon />
          <span className="ml-2">{displayName}</span>
        </Button>
      </DropdownTrigger>
      <DropdownMenu aria-label={t("tip.language")}>
        {Object.keys(resources).map((lang) => (
          <DropdownItem key={lang} onClick={() => toggleLanguage(lang)}>
            {resources[lang as keyof typeof resources].name}
          </DropdownItem>
        ))}
      </DropdownMenu>
    </Dropdown>
  );
}
