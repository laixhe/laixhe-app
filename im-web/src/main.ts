import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';

// UI框架
import ElementPlus from 'element-plus';
import 'element-plus/lib/theme-chalk/index.css';
import locale from 'element-plus/lib/locale/lang/zh-cn';

// 创建 App 应用返回对应的实例对象，调用 mount 方法进行挂载
const app = createApp(App);
app.use(ElementPlus, { size: 'small', locale });
app.use(store);
app.use(router);
app.mount('#app');
