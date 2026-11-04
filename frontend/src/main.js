import 'virtual:uno.css'
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import VueApexCharts from 'vue3-apexcharts'
import App from './App.vue'
import './style.css'

const app = createApp(App)

app.use(createPinia())
app.use(VueApexCharts)
app.mount('#app')
