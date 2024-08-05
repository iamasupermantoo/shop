import { api } from 'src/boot/axios';

// 新的验证码
export const captchaCreateAPI = () => {
  return api.get('/captcha/create', { showLoading: false } as any);
};

// 管理登录
export const adminLoginAPI = (params: any) => {
  return api.post('/login', params);
};

// 首页信息
export const homeIndexAPI = () => {
  return api.post('/auth/index')
}
