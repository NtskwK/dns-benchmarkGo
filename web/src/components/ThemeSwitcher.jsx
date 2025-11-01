import { useEffect, useState } from "react";
import { Switch, Tooltip } from "@nextui-org/react";
import { FaSun as SunIcon, FaMoon as MoonIcon } from "react-icons/fa";
import { useTranslation } from "react-i18next";

export default function ThemeSwitcher() {
  const { t } = useTranslation();
  const [isSelected, setIsSelected] = useState(false);

  useEffect(() => {
    // 检查系统主题
    const darkModeMediaQuery = window.matchMedia("(prefers-color-scheme: dark)");
    /**
     * 更新主题
     * @param {MediaQueryListEvent | MediaQueryList} e
    */
    const updateTheme = (e) => {
      setIsSelected(e.matches);
      document.documentElement.classList.toggle("dark", e.matches);
    };

    // 初始化主题
    updateTheme(darkModeMediaQuery);

    // 监听系统主题变化
    darkModeMediaQuery.addEventListener("change", updateTheme);

    return () => darkModeMediaQuery.removeEventListener("change", updateTheme);
  }, []);

  const toggleTheme = () => {
    setIsSelected(!isSelected);
    document.documentElement.classList.toggle("dark");
  };

  return (
    <Tooltip content={t("tip.theme")}>
      <Switch
        defaultSelected={isSelected}
        size="lg"
        color="secondary"
        onChange={toggleTheme}
        startContent={<SunIcon />}
        endContent={<MoonIcon />}
        aria-label={t("tip.theme")}
      />
    </Tooltip>
  );
}
