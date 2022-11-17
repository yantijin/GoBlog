import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import eslintPlugin from 'vite-plugin-eslint';
import { resolve } from 'path';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(), eslintPlugin()],
  server: {
    host: 'localhost',
    port: 8080,
    open: true,
  },
  resolve: {
    alias: {
      '@': resolve(__dirname, './src'), // 设置@代表路径 ./src
    },
  },
});
