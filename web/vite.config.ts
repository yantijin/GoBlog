import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import eslintPlugin from "vite-plugin-eslint";
import { resolve } from "path";
import * as dotenv from "dotenv";
import { readFileSync } from "fs";

// https://vitejs.dev/config/
export default defineConfig(({ command, mode }) => {
  const file = `.env.${mode}`;
  const envConfig = dotenv.parse(readFileSync(file));
  for (const k in envConfig) {
    process.env[k] = envConfig[k];
  }

  return {
    plugins: [vue(), eslintPlugin()],
    server: {
      host: "localhost",
      port: parseInt(process.env.VITE_CLI_PORT),
      open: true,
      proxy: {
        [process.env.VITE_BASE_API]: {
          target: `${process.env.VITE_BASE_PATH}:${process.env.VITE_SERVER_PORT}`,
          changeOrigin: true, // changes the origin of the host header to the target url
          rewrite: path =>
            path.replace(new RegExp("^" + process.env.VITE_BASE_API), ""),
        },
      },
    },
    resolve: {
      alias: {
        "@": resolve(__dirname, "./src"), // 设置@代表路径 ./src
      },
    },
  };
});
