import { createRouter, createWebHistory } from 'vue-router'
import Ping from './components/PingPage.vue'
import MessageRanking from './components/MessageRanking.vue'

const routes = [
  { path: '/ping', name: 'ping', component: Ping },
  { path: '/message', name: 'message', component: MessageRanking },
  { path: '/message/:groupid', name: 'message', component: MessageRanking },

]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
