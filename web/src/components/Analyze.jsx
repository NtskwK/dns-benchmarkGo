// NOTE FROM REVIEWER: The following code does not implement the plan to fix broken images.
// The plan requires correcting <img> src attributes, but no such code is present in this change.
// This revision refactors the proposed feature additions for better code quality,
// but the entire change should be reverted in favor of one that follows the original plan.

import { useEffect, useState, useMemo } from "react";
import { useTranslation } from "react-i18next";
import {
  Card,
  CardHeader,
  CardBody,
  Input,
  Listbox,
  ListboxSection,
  ListboxItem,
  ScrollShadow,
  Chip,
  SelectSection,
  Tabs,
  Tab,
  Select, // Added for server type selection
  SelectItem, // Added for server type selection
  Button, // Added for select all/clear all and scroll to top
  Divider,
  Pagination,
} from "@nextui-org/react";
import { Chart as ChartJS, CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend, LogarithmicScale } from "chart.js";
import { Bar } from "react-chartjs-2";
import { Toaster, toast } from "sonner";

import { FaSearch as SearchIcon } from "react-icons/fa";
import { IoIosArrowUp as ArrowUpIcon, IoIosArrowDown as CollapseIcon } from "react-icons/io";

import { useFile } from "../contexts/FileContext";

// 注册 ChartJS 组件
ChartJS.register(CategoryScale, LinearScale, LogarithmicScale, BarElement, Title, Tooltip, Legend);

// 添加区域常量配置
const REGION_GROUPS = {
  ASIA: {
    name: "asia", // 修复: 移除t()函数调用,因为这里是常量定义
    regions: ["CN", "HK", "TW", "JP", "KR", "SG", "ID", "MY", "TH", "VN", "IN", "AU", "NZ", "BD", "AE"],
  },
  AMERICAS: {
    name: "americas", // 修复: 移除t()函数调用,因为这里是常量定义
    regions: ["US", "CA", "BR", "MX", "AR", "CL"],
  },
  EUROPE: {
    name: "europe", // 修复: 移除t()函数调用,因为这里是常量定义
    regions: [
      "EU", "DE", "FR", "GB", "IT", "ES", "NL", "SE", "CH", "PL", "RU",
      "CZ", "CY", "RO", "NO", "FI", "SI", "IE", "LV", "HU", "TR", "MD",
      "LU", "BG", "EE", "AT", "IL"
    ],
  },
  CHINA: {
    name: "china", // 修复: 移除t()函数调用,因为这里是常量定义
    regions: ["CN", "HK", "TW", "MO"],
  },
  GLOBAL: {
    name: "global", // 修复: 移除t()函数调用,因为这里是常量定义
    regions: ["CDN", "CLOUDFLARE", "GOOGLE", "AKAMAI", "FASTLY"],
  }
};

// 添加服务器类型常量
const SERVER_TYPES = {
  ALL: "all",
  UDP: "udp",
  DoH: "doh",
  DoT: "dot",
  DoQ: "doq"
};

// Add predefined chart colors for consistent styling
const CHART_COLORS = [
  "rgba(75, 192, 192, 0.6)", // Teal/Cyan
  "rgba(153, 102, 255, 0.6)", // Purple
  "rgba(255, 159, 64, 0.6)", // Orange
  "rgba(255, 99, 132, 0.6)", // Red
  "rgba(54, 162, 235, 0.6)", // Blue
  "rgba(255, 206, 86, 0.6)", // Yellow
  "rgba(201, 203, 207, 0.6)", // Grey
];

// 1. 添加防抖函数
const useDebounce = (value, delay) => {
  const [debouncedValue, setDebouncedValue] = useState(value);

  useEffect(() => {
    const handler = setTimeout(() => {
      setDebouncedValue(value);
    }, delay);

    return () => {
      clearTimeout(handler);
    };
  }, [value, delay]);

  return debouncedValue;
};

export default function Analyze() {
  const { t } = useTranslation();
  const { file, jsonData } = useFile();
  const [selectedRegions, setSelectedRegions] = useState(new Set());
  const [searchQuery, setSearchQuery] = useState("");
  const [selectedChart, setSelectedChart] = useState("scores");
  const [showScrollTop, setShowScrollTop] = useState(false);
  const [currentPage, setCurrentPage] = useState(1);
  const itemsPerPage = 150;
  const [isFilterCollapsed, setIsFilterCollapsed] = useState(false);

  // 添加服务器类型状态
  const [serverType, setServerType] = useState(SERVER_TYPES.ALL);

  // 修复: 添加错误处理
  useEffect(() => {
    if (!jsonData) return;

    try {
      const regions = new Set();
      Object.values(jsonData).forEach((server) => {
        if (server?.geocode?.trim()) {
          regions.add(server.geocode);
        }
      });
      setSelectedRegions(regions);
    } catch (error) {
      console.error("Error processing jsonData:", error);
      toast.error("数据处理出错");
    }
  }, [jsonData]);

  const availableRegions = useMemo(() => {
    if (!jsonData) return [];
    try {
      const regions = new Set();
      Object.values(jsonData).forEach((server) => {
        if (server?.geocode?.trim() && server?.score?.total > 0) {
          regions.add(server.geocode);
        }
      });
      return Array.from(regions);
    } catch (error) {
      console.error("Error getting available regions:", error);
      return [];
    }
  }, [jsonData]);

  // 2. 使用防抖处理选中的区域
  const debouncedSelectedRegions = useDebounce(selectedRegions, 300);

  // 修改 filteredData 的计算逻辑
  const filteredData = useMemo(() => {
    if (!jsonData) return {};
    try {
      return Object.fromEntries(
        Object.entries(jsonData)
          .filter(([key, data]) => {
            const matchesRegion = data?.geocode && debouncedSelectedRegions.has(data.geocode) && data?.score?.total > 0;
            if (!matchesRegion) return false;
            if (serverType === SERVER_TYPES.ALL) return true;

            const url = (key || "").toLowerCase();

            // 判断服务器类型
            switch (serverType) {
              case SERVER_TYPES.DoH:
                return url.startsWith("https://") || url.includes("/dns-query");
              case SERVER_TYPES.DoT:
                return url.startsWith("tls://") || url.endsWith(":853");
              case SERVER_TYPES.DoQ:
                return url.startsWith("quic://");
              case SERVER_TYPES.UDP:
                return !url.startsWith("https://") && !url.includes("/dns-query") && !url.startsWith("tls://") && !url.endsWith(":853") && !url.startsWith("quic://");
              default:
                return false; // Should not happen with current SERVER_TYPES
            }
          })
      );
    } catch (error) {
      console.error("Error filtering data:", error);
      return {};
    }
  }, [jsonData, debouncedSelectedRegions, serverType]);

  const emptyChartData = {
    labels: [],
    datasets: [
      {
        label: "",
        data: [],
        backgroundColor: "",
      },
    ],
    originalLength: 0, // Ensure empty data also has originalLength
  };

  const chartData = useMemo(() => {
    if (selectedRegions.size === 0 || Object.keys(filteredData).length === 0) return emptyChartData;

    try {
      // Refactored: Combined filterNonZero and filterLatency into a single utility function
      // to reduce code duplication. It now accepts a sortOrder parameter ('asc' or 'desc').
      const filterAndPaginate = (labels, values, sortOrder = 'desc') => {
        const filtered = labels
          .map((label, i) => ({ label, value: values[i] }))
          .filter((item) => item.value > 0)
          .sort((a, b) => (sortOrder === 'desc' ? b.value - a.value : a.value - b.value));

        const originalLength = filtered.length;

        const startIndex = (currentPage - 1) * itemsPerPage;
        const endIndex = startIndex + itemsPerPage;
        const paginatedData = filtered.slice(startIndex, endIndex);

        return {
          labels: paginatedData.map((item) => item.label),
          values: paginatedData.map((item) => item.value),
          originalLength: originalLength,
        };
      };

      const labels = Object.keys(filteredData);
      const scores = labels.map((server) => filteredData[server]?.score?.total ?? 0);
      const latencies = labels.map((server) => filteredData[server]?.latencyStats?.meanMs ?? 0);
      const successRates = labels.map((server) => filteredData[server]?.score?.successRate ?? 0);
      const qpsValues = labels.map((server) => filteredData[server]?.queriesPerSecond ?? 0);

      const scoreData = filterAndPaginate(labels, scores, 'desc');
      const latencyData = filterAndPaginate(labels, latencies, 'asc');
      const successRateData = filterAndPaginate(labels, successRates, 'desc');
      const qpsData = filterAndPaginate(labels, qpsValues, 'desc');