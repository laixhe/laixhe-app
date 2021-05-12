import {
    createRouter,
    createWebHashHistory,
    RouteRecordRaw
} from 'vue-router';

import Login from '@/views/login.vue'; // 必须添加.vue后缀

// 路由地址列表
const routes: Array<RouteRecordRaw> = [
{
    path: '/login',
    name: 'Login',
    component: Login
},
// {
//     path: '/todo-list',
//     name: 'TodoList',
//     component: () => import('@/views/todo-list.vue') // 懒加载组件 必须添加.vue后缀
// }
];

const router = createRouter({
    // 指定路由的模式
    history: createWebHashHistory(),
    // 路由地址
    routes
});

export default router;