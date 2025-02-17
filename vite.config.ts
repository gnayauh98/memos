import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import UnoCSS from 'unocss/vite'

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue(), UnoCSS()],

  server: {
    port: 80,
    proxy: {
      "/api": {
        // target: 'http://14.103.242.223',
        target: "http://localhost:3000",
        changeOrigin: true
      }
    }
  }
})
