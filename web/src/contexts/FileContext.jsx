// @ts-check
import { createContext, useContext, useState, useEffect, useRef, useCallback } from "react";
import { useTranslation } from "react-i18next";
import { toast } from "sonner";

/** @type {React.Context<any>} */
const FileContext = createContext(null);

const LOCAL_STORAGE_KEY = "dnsAnalyzerData";

/**
 * @param {{ children: React.ReactNode }} props
 */
export function FileProvider({ children }) {
  const { t } = useTranslation();
  const [file, setFile] = useState(/** @type {File | null} */(null));
  const hasShownInitialToast = useRef(false);

  const [jsonData, setJsonData] = useState(/** @type {any} */(null))

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

  const showToast = useCallback(
    /**
     * @param {'success' | 'error'} type
     * @param {string} title
     * @param {string} desc
     */
    (type, title, desc) => {
      if (type === 'success') {
        toast.success(t(title), {
          description: t(desc),
          duration: 5000,
          className: "dark:text-neutral-200",
          dismissible: true,
        });
      } else if (type === 'error') {
        toast.error(t(title), {
          description: t(desc),
          duration: 6000,
          className: "dark:text-neutral-200",
          dismissible: true,
        });
      }
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
        const result = e.target?.result;
        if (typeof result !== 'string') {
          throw new Error('File content is not a string');
        }
        const data = JSON.parse(result);
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
