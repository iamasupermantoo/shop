// 图片处理方法
export const imageSrc = (url: string) => {
  // 默认值
  if (url == '' || url == null)
    url = '/logo.png'

  // 如果是域名那么返回当前地址
  if (url.indexOf('//') == 0 || url.indexOf('//') == 5 || url.indexOf('//') == 6) {
    return url;
  }

  // 如果使用 // 自动获取的话, 那么自动加上头
  const envBaseURL = <string>process.env.baseURL
  const baseURL = envBaseURL.toString().indexOf('//') == 0 ? new URL(document.location.protocol + envBaseURL) : new URL(envBaseURL)
  return baseURL.origin + url;
};
