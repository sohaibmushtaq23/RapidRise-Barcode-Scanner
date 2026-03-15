import { fileURLToPath, URL } from 'node:url'
import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import fs from 'fs'

export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd())
  return {
    plugins: [vue()],
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url))
      }
    },
    server: {
      host: true,
      port: 5173,
      https: {
        key: fs.readFileSync('../certs/192.168.100.13+1-key.pem'),
        cert: fs.readFileSync('../certs/192.168.100.13+1.pem'),
      },
      proxy: {
        '/api': {
          target: env.VITE_API_URL || 'https://192.168.100.13:8080',
          changeOrigin: true,
          secure: false
        }
      },
      // In production, allowedHosts should be restricted
      allowedHosts: mode === 'development' ? ['all'] : ['yourdomain.com']
    }
  }
})