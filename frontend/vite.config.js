import { resolve } from 'path'
import { defineConfig } from 'vite'

export default defineConfig({
  build: {
    rollupOptions: {
      input: {
        main: resolve(__dirname, 'index.html'),
        downloads: resolve(__dirname, 'src/html-navegation/downloads.html'),
        folders: resolve(__dirname, 'src/html-navegation/folders-menu.html'),
      }
    }
  }
})