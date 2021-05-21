import {
  createRouter,
  createWebHistory,
  RouteRecordRaw
} from 'vue-router';

import Login from '@/views/login.vue'; // 必须添加.vue后缀

// 路由地址列表
const routes: Array<RouteRecordRaw> = [
  {
      path: '/home',
      name: 'Home',
      component: () => import('@/views/home.vue') // 懒加载组件 必须添加.vue后缀
  },
  {
      path: '/login',
      name: 'Login',
      component: Login
  }
];

const router = createRouter({
  // 指定路由的模式
  history: createWebHistory(process.env.BASE_URL),
  routes
});

export default router;
