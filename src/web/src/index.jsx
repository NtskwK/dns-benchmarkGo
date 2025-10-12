import React from 'react';
import ReactDOM from 'react-dom/client';
import { useTranslation } from 'react-i18next';
import { NextUIProvider } from '@nextui-org/react';

import App from './App';
import './components/i18n';

const rootElement = document.getElementById('root');
if (!rootElement) throw new Error('Failed to find the root element');

const root = ReactDOM.createRoot(rootElement);

// 将 useTranslation 移到组件内部使用
function Main() {
  const { t } = useTranslation();
  // 设置页面标题
  React.useEffect(() => {
    document.title = t("title");
  }, [t]);

  return (
    <React.StrictMode>
      <NextUIProvider>
        <App />
      </NextUIProvider>
    </React.StrictMode>
  );
}

root.render(<Main />);
