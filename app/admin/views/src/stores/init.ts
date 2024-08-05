import { defineStore } from 'pinia';
import {LocalStorage} from 'quasar';

const UserTokenKey = '_UserToken';
const UserInfoKey = '_UserInfoKey';
const RouterList = '_RouterList';
const MenuList = '_MenuList';
const EndRouteKey = '_EndRoute';

export const InitStoreState = {
  //  管理Token
  userToken: '',

  //  管理信息
  userInfo: {
    avatar: '',
    nickname: '',
    email: '',
  } as any,

  // 管理路由列表
  routerList: [] as any,

  // 最后的路由
  endRoute: '',

  //  管理菜单路由
  menuList: [] as any,
};

// 初始化数据
export const useInitStore = defineStore('init', {
  state: () => {
    return {
      userToken: LocalStorage.getItem(UserTokenKey) ?? '',
      endRoute: LocalStorage.getItem(EndRouteKey) ?? '',
      userInfo: LocalStorage.getItem(UserInfoKey) ? JSON.parse(<string>LocalStorage.getItem(UserInfoKey)) : InitStoreState.userInfo,
      routerList: LocalStorage.getItem(RouterList) ? JSON.parse(<string>LocalStorage.getItem(RouterList)) : InitStoreState.routerList,
      menuList: LocalStorage.getItem(MenuList) ? JSON.parse(<string>LocalStorage.getItem(MenuList)) : InitStoreState.menuList,
    }
  },

  getters: {},

  actions: {
    //  更新用户Token
    updateUserToken(token: string) {
      this.userToken = token;
      LocalStorage.set(UserTokenKey, token)
    },

    // 更新最后的路由
    updateEndRoute(endRoute: string) {
        this.endRoute = endRoute
      LocalStorage.set(EndRouteKey, endRoute)
    },

    // 更新用户信息
    updateUserInfo(info: object) {
      this.userInfo = info
      LocalStorage.set(UserInfoKey, JSON.stringify(info))
    },

    // 更新路由列表
    updateRouterList(routerList: []) {
      this.routerList = routerList
      LocalStorage.set(RouterList, JSON.stringify(routerList))
    },

    // 更新用户菜单
    updateMenuList(menuList: []) {
      this.menuList = menuList
      LocalStorage.set(MenuList, JSON.stringify(menuList))
    },

    //  判断是否存在当前路由
    hasRoute(url: string) {
      if (url == '') {
        return true
      }

      const envBaseURL = <string>process.env.baseURL
      const baseURL = envBaseURL.toString().indexOf('//') == 0 ? new URL(document.location.protocol + envBaseURL) : new URL(envBaseURL)
      return this.routerList.indexOf(baseURL.pathname + url) > -1
    },
  },
});
