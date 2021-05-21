import Axios from 'axios';
import { ElMessage } from 'element-plus';

const baseURL = 'http://192.168.10.240:5050';

const axios = Axios.create({
  baseURL: baseURL,
  timeout: 3000 // 请求超时 3s
});

// 前置拦截器（发起请求之前的拦截）
axios.interceptors.request.use(
  (requestConfig) => {
    // 模拟指定请求令牌
    requestConfig.headers["Authorization"] = "Bearer Token...";
    return requestConfig;
  },
  (error) => {
    // 请求错误的统一处理
    console.log('[Axios Error]', error);
    return Promise.reject(error);
  }
);

// 后置拦截器（获取到响应时的拦截）
axios.interceptors.response.use(
  (response) => {
    console.log('[Axios Response]', response.data);
    return response;
  },
  (error) => {
    if (error.response && error.response.data) {
      const code = error.response.status;
      const msg = error.response.data.message;
      ElMessage.error(`Code: ${code}, Message: ${msg}`);
      console.log(`[Axios Error]`, error.response);
    } else {
      ElMessage.error(`${error}`);
      console.log('[Axios Error]', error);
    }
    return Promise.reject(error);
  }
);

export default axios;