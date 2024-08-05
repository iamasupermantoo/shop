export const InputTypeList = {
  Text: 1, //  文本
  TextArea: 2, //  文本域
  Editor: 3, //  富文本
  Number: 4, //  数字
  Password: 5, //  密码
  Select: 10, //  选择图片
  Radio: 11, //  单选框
  Checkbox: 12, //  多选框
  Toggle: 13, //  开关
  File: 21, //  文件
  Image: 22, //  图片
  Images: 23, //  图片组
  Icon: 24, //  图标
  DatePicker: 31, //  时间格式
  RangeDatePicker: 32, //  时间范围
  InputJson: 41, //  Json对象
  InputChildren: 42, //  子对象
  InputTranslate: 61, //  翻译
};

export const ContentTypeList = {
  Text: 1, //  文本类型
  Translate: 61, //  翻译类型
  Select: 10, //  选择图片
  Image: 22, //  图片
  Images: 23, //  图片组
  DatePicker: 31, //  时间格式
  InputEditText: 51, //  编辑文本
  InputEditNumber: 52, //  编辑数字
  InputEditTextArea: 53, //  编辑富文本
  InputEditToggle: 54, //  编辑开关
};

// 颜色集合
export const quasarColorsObject = {
  primary: 'primary',
  secondary: 'secondary',
  accent: 'accent',
  dark: 'dark',
  positive: 'positive',
  negative: 'negative',
  info: 'info',
  warning: 'warning',
  red: 'red',
  pink: 'pink',
  purple: 'purple',
  deepPurple: 'deep-purple',
  indigo: 'indigo',
  blue: 'blue',
  lightBlue: 'light-blue',
  cyan: 'cyan',
  teal: 'teal',
  green: 'green',
  lightGreen: 'light-green',
  lime: 'lime',
  yellow: 'yellow',
  amber: 'amber',
  orange: 'orange',
  deepOrange: 'deep-orange',
  brown: 'brown',
  grey: 'grey',
  blueGrey: 'blue-grey',
};

// cookieOptions cookies其他参数
export const cookieOptions = () => {
  return { expires: '30d 3h 5m' };
};
