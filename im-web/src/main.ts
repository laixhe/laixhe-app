// 程序的主入口文件

// 引入 createApp 函数，创建对应的应用
import { createApp } from 'vue'
// 引入 App 组件（所有组件的父级组件）
import App from './App.vue'

// 浏览器样式兼容
import 'normalize.css'

// UI框架
import ElementPlus from 'element-plus'
import 'element-plus/lib/theme-chalk/index.css'
import locale from 'element-plus/lib/locale/lang/zh-cn';

// 路由
import router from './router/index'

// 创建 App 应用返回对应的实例对象，调用 mount 方法进行挂载
createApp(App)
.use(ElementPlus, { size: 'small', locale })
.use(router)
.mount('#app')
