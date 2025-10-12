import { defineConfig } from "@rsbuild/core";
import { pluginReact } from "@rsbuild/plugin-react";

export default defineConfig({
  plugins: [pluginReact()],
  html: {
    title: "DNS-BenchGo",
  },
  output: {
    filenameHash: false,
  },
  performance: {
    chunkSplit: {
      strategy: "all-in-one",
    },
  },
});

