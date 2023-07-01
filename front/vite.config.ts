import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  build: {
    sourcemap: true,
  },
  server: {
    proxy: {
      '/api': {
        target: 'http://127.25.23.1:8080',
        // rewrite: path => path.replace(/^\/api/, ''),
      },
    },
  },
  resolve: {
    alias:{
      '@': path.resolve(__dirname, './src')
    }
  }
})
