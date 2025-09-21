import '@/assets/global.css'
import '@/assets/dark.css'
import '@/assets/css/form.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

console.log(`

--- %cHealty Information System%c ---

Develop in : Sunday, 01 June 2025
Developer  : @mn.syp

---         End line          ---

`, `font-weight: bold`, ``)

const app = createApp(App)

app.use(router)

app.mount('#app')
