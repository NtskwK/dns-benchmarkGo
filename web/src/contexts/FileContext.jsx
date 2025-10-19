import React, { createContext, useContext, useState, useEffect, useRef, useCallback } from "react";
import { useTranslation } from "react-i18next";
import { toast } from "sonner";

const FileContext = createContext();

const LOCAL_STORAGE_KEY = "dnsAnalyzerData";

export function FileProvider({ children }) {
  const { t } = useTranslation();
  const [file, setFile] = useState(null);
  const hasShownInitialToast = useRef(false);

  const [jsonData, setJsonData] = useState(null)

  useEffect(() => {
    const savedData = localStorage.getItem(LOCAL_STORAGE_KEY);
    if (!savedData) return;
    try {
      const data = JSON.parse(savedData);
      if (!hasShownInitialToast.current) {
        setTimeout(() => {
          toast.success(t("tip.data_loaded"), {
            description: t("tip.data_loaded_desc"),
            duration: 5000,
            className: "dark:text-neutral-200",
            dismissible: true,
          });
        }, 0);
        hasShownInitialToast.current = true;
      }
      setJsonData(data);
    } catch (error) {
      console.error("解析保存的JSON时出错:", error);
    }
  }, []);

  const showToast = useCallback((type, title, desc) => {
    toast[type](t(title), {
      description: t(desc),
      duration: type === 'error' ? 6000 : 5000,
      className: "dark:text-neutral-200",
      dismissible: true,
    });
  }, [t]);


  useEffect(() => {
    if (!file) return;

    if (!file.name.toLowerCase().endsWith('.json')) {
      showToast('error', 'tip.invalid_file_type', 'tip.only_json_allowed');
      setFile(null);
      return;
    }

    const reader = new FileReader();
    reader.onload = (e) => {
      try {
        const data = JSON.parse(e.target?.result);
        setJsonData(data);
        showToast('success', 'tip.data_loaded', 'tip.data_loaded_desc');
        localStorage.setItem(LOCAL_STORAGE_KEY, JSON.stringify(data));
      } catch (error) {
        console.error("解析JSON时出错:", error);
        showToast('error', 'tip.data_load_failed', 'tip.data_load_failed_desc');
      }
    };
    reader.readAsText(file);
  }, [file, showToast]);

  const value = {
    file,
    setFile,
    jsonData,
    setJsonData
  };

  return <FileContext.Provider value={value}>{children}</FileContext.Provider>;
}

export function useFile() {
  const context = useContext(FileContext);
  if (context === undefined) {
    throw new Error("useFile必须在FileProvider中使用");
  }
  return context;
}
