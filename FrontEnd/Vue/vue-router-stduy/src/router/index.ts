// 创建一个路由器，并暴露出去
// 第一步: 引入 createRouter
import { createRouter, createWebHistory } from 'vue-router'

// 第二步: 引入可能要呈现的组件
import Home from '@/view/MainHome.vue'
import Topic from '@/view/MainTopic.vue'
import About from '@/view/MainAbout.vue'

// 第三步: 创建路由器
const router = createRouter({
  history: createWebHistory(),
  routes: [
    // home
    {
      path: '/home',
      component: Home,
    },
    // topic
    {
      path: '/topic',
      component: Topic,
    },
    // about
    {
      path: '/about',
      component: About,
    },
  ],
})

// 第四步: 暴露路由
export default router
