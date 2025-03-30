import { createApp } from 'vue' // 引入 createApp
import App from './App.vue' // 引入 App 根组件
import router from '@/router/index' // 引入路由器

const app = createApp(App) // 创建一个应用

app.use(router) // 使用路由器

app.mount('#app') // 挂载整个应用到 App 容器中
