import { boot } from 'quasar/wrappers';
import axios, { AxiosInstance } from 'axios';
import {Loading} from 'quasar';
import {QSpinnerBars} from 'quasar';
import {NotifyNegative} from 'src/utils/notify';
import {useInitStore} from 'stores/init'

declare module '@vue/runtime-core' {
  interface ComponentCustomProperties {
    $axios: AxiosInstance;
    $api: AxiosInstance;
  }
}

// Be careful when using SSR for cross-request state pollution
// due to creating a Singleton instance here;
// If any client changes this (global) instance, it might be a
// good idea to move this instance creation inside of the
// "export default () => {}" function below (which runs individually
// for each client)
const api = axios.create({ baseURL: process.env.baseURL });

// 请求数据拦截
api.interceptors.request.use((config: any) => {
  const initStore = useInitStore();

  if (!config.hasOwnProperty('showLoading') || config.showLoading) {
    Loading.show({
      spinner: QSpinnerBars,
      spinnerColor: 'secondary',
      spinnerSize: 50,
      message: '一些重要的过程正在进行中, 请等待...',
    });
  }

  // 如果存在Token，那么请求带上Token
  if (
    initStore.userToken !== '' &&
    !config.headers.hasOwnProperty('Authorization')
  ) {
    config.headers['Authorization'] = 'Bearer ' + initStore.userToken;
  }
  config.headers['Accept-Language'] = 'zh-CN';

  return config;
});

// 响应数据拦截
api.interceptors.response.use(
  (response) => {
    Loading.hide();
    const res = response.data;

    if (res.hasOwnProperty('code')) {
      if (res.code === 0) {
        return res.data;
      }
      NotifyNegative(res.msg);
      return Promise.reject(res.msg);
    } else {
      return res;
    }
  },
  (err) => {
    Loading.hide();
    if (err.response) {
      switch (err.response.status) {
        case 401:
          NotifyNegative('当前请求没有权限, 请退出重新登录...');
          break;
        case 500:
          NotifyNegative('服务器运行错误, 请联系系统管理员....');
          break;
        default:
          NotifyNegative('ERR_UNHANDLED_REJECTION');
      }
    }
    //  返回错误
    return Promise.reject('ERR_UNHANDLED_REJECTION');
  }
);

export default boot(({ app }) => {
  // for use inside Vue files (Options API) through this.$axios and this.$api

  app.config.globalProperties.$axios = axios;
  // ^ ^ ^ this will allow you to use this.$axios (for Vue Options API form)
  //       so you won't necessarily have to import axios in each vue file

  app.config.globalProperties.$api = api;
  // ^ ^ ^ this will allow you to use this.$api (for Vue Options API form)
  //       so you can easily perform requests against your app's API
});

export { api };
