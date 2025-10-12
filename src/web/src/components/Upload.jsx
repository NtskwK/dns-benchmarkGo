import { Tooltip, Button } from "@nextui-org/react";
import { useTranslation } from "react-i18next";
import { useFile } from "../contexts/FileContext";
import SampleData from "./SampleData";

export default function Upload() {
  const { t } = useTranslation();
  const { setFile, jsonData, setJsonData } = useFile();

  const handleFileChange = (event) => {
    const file = event.target.files?.[0];
    if (file) {
      setFile(file);
    }
  };

  const handleClearData = () => {
    setFile(null);
    setJsonData(null);
    localStorage.removeItem("dnsAnalyzerData");
  };

  const handleLoadSample = () => {
    setJsonData(SampleData);
    localStorage.setItem("dnsAnalyzerData", JSON.stringify(SampleData));
  };

  return (
    <div className="flex gap-2">
      <Tooltip content={t("tip.upload")}>
        <Button color="primary" variant="flat" as="label" className="cursor-pointer">
          <input type="file" className="hidden" accept=".json" onChange={handleFileChange} />
          {t("button.upload")}
        </Button>
      </Tooltip>

      <Tooltip content={t("tip.sample")}>
        <Button color="secondary" variant="flat" onClick={handleLoadSample}>
          {t("button.sample")}
        </Button>
      </Tooltip>

      {jsonData && (
        <Tooltip content={t("tip.clear")}>
          <Button color="danger" variant="flat" onClick={handleClearData}>
            {t("button.clear")}
          </Button>
        </Tooltip>
      )}
    </div>
  );
}
