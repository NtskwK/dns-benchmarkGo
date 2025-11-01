import { Tooltip, Button } from "@nextui-org/react";
import { useTranslation } from "react-i18next";
import { useFile } from "../contexts/FileContext";

export default function Upload() {
  const { t } = useTranslation();
  const { setFile, jsonData } = useFile();

  const handleFileChange = (event) => {
    const file = event.target.files?.[0];
    if (file) {
      setFile(file);
    }
  };

  return (
    <div className="flex gap-2">
      <Tooltip content={t("tip.upload")}>
        <Button color="primary" variant="flat" as="label" className="cursor-pointer">
          <input type="file" className="hidden" accept=".json" onChange={handleFileChange} />
          {t("button.upload")}
        </Button>
      </Tooltip>
    </div>
  );
}
