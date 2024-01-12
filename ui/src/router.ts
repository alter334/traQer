import { createRouter, createWebHistory } from 'vue-router'
import Ping from './components/PingPage.vue'
import MessageRanking from './components/MessageRanking.vue'

const routes = [
  { path: '/ping', name: 'ping', component: Ping },
  { path: '/messages', name: 'message', component: MessageRanking },
  { path: '/messages/:groupid', name: 'messagegroup', component: MessageRanking},
  

]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
