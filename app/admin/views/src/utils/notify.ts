import { Notify } from 'quasar';

// 错误提示
export const NotifyNegative = (msg: string) => {
  Notify.create({
    type: 'negative',
    position: 'top',
    timeout: 3000,
    message: msg,
  });
};

// 成功提示
export const NotifyPositive = (msg: string) => {
  Notify.create({
    type: 'positive',
    position: 'top',
    timeout: 3000,
    message: msg,
  });
};

// 系统通知
export const WarningNotify = (msg: string) => {
  Notify.create({
    type: 'warning',
    position: 'top-right',
    timeout: 3000,
    message: msg,
  });
};

export const ShowNotify = (config = {message: '', avatar: '', icon: '', color: 'primary', position: 'right', timeout: 3000} as any) => {
  Notify.create({color: config.color, icon: config.icon, avatar: config.avatar, position: config.position, message: config.message, timeout: config.timeout})
}
