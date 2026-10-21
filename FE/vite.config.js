import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  server: {
    port: 8082,
    host: '0.0.0.0',
    allowedHosts: ['my-react-app-im'],
    proxy: {
      '/api': {
        target: 'http://go-backend:8080',
        changeOrigin: true,
      },
    },
  },
})