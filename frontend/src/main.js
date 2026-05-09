import 'virtual:uno.css'
import {createApp} from 'vue'
import {createPinia} from 'pinia'
import App from './App.vue'
import VueApexCharts from "vue3-apexcharts"; 
import './style.css';

const app = createApp(App)
const pinia = createPinia()

app.use(pinia)
app.use(VueApexCharts)
app.component('apexchart', VueApexCharts)
app.mount('#app')