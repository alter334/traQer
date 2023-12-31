import { createRouter, createWebHistory } from 'vue-router'
import Ping from './components/PingPage.vue'

const routes = [
  { path: '/ping', name: 'ping', component: Ping },

]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
