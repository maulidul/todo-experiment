import './assets/main.css' // Mengimpor file CSS utama untuk styling aplikasi.

import { createApp } from 'vue' // Mengimpor fungsi createApp dari Vue.js untuk membuat instance aplikasi.
import { createPinia } from 'pinia' // Mengimpor fungsi createPinia dari Pinia untuk manajemen state.
import App from './App.vue' // Mengimpor komponen root App.vue.

const app = createApp(App) // Membuat instance aplikasi Vue dengan komponen root App.vue.
app.use(createPinia()) // Menggunakan Pinia sebagai manajemen state untuk aplikasi.
app.mount('#app') // Memasang aplikasi ke elemen HTML dengan ID 'app'.
