import { route } from 'quasar/wrappers';
import {
  createMemoryHistory,
  createRouter,
  createWebHashHistory,
  createWebHistory,
} from 'vue-router';

export interface MenuListDataInterface {
  icon: string;
  tmp: string;
  confURL: string;
}

export interface MenuListInterface {
  name: string;
  route: string;
  children: MenuListInterface[];
  data: MenuListDataInterface;
}

const routePageList = import.meta.glob('../pages/**/*.vue');
import routes from 'src/router/routes';
import {useInitStore} from 'stores/init';

/*
 * If not building with SSR mode, you can
 * directly export the Router instantiation;
 *
 * The function below can be async too; either use
 * async/await or return a Promise which resolves
 * with the Router instance.
 */

export default route(function (/* { store, ssrContext } */) {
  const createHistory = process.env.SERVER
    ? createMemoryHistory
    : (process.env.VUE_ROUTER_MODE === 'history' ? createWebHistory : createWebHashHistory);

  const $initStore = useInitStore()

  const Router = createRouter({
    scrollBehavior: () => ({ left: 0, top: 0 }),
    routes,

    // Leave this as is and make changes in quasar.conf.js instead!
    // quasar.conf.js -> build -> vueRouterMode
    // quasar.conf.js -> build -> publicPath
    history: createHistory(process.env.VUE_ROUTER_BASE),
  });


  // 路由前置守卫
  Router.beforeEach((to, form, next) => {
    if (
      (to.name === 'Login' || to.name === 'Register') &&
      $initStore.userToken != null &&
      $initStore.userToken.length > 0
    ) {
      next('/');
    } else {
      // 验证是否跳转到登录页面
      if (
        to.matched.some((record) => record.meta.requireAuth) &&
        ($initStore.userToken == null ||
          $initStore.userToken.length === 0)
      ) {
        next('/login');
      } else {
        $initStore.updateEndRoute(to.fullPath)
        next();
      }
    }
  });

  return Router;
});

// 递归动态添加路由
export const dynamicRouter = async (router: any, menuList: MenuListInterface[]) => {
  const $initStore = useInitStore()
  dynamicRouterFunc(router, menuList)
  router.replace($initStore.endRoute);
};

const dynamicRouterFunc = (router: any, menuList: MenuListInterface[]) => {
  for (let index = 0; index < menuList.length; index++) {
    const element = menuList[index];
    if (element.route !== '' && !router.hasRoute(element.route)) {
      router.addRoute('Layouts', {
        path: element.route,
        component: routePageList['../pages' + element.data.tmp + '.vue'],
        meta: {
          requireAuth: true,
          keepAlive: true,
          views: element.data.confURL,
        },
      });
    }
    if (
      element.hasOwnProperty('children') &&
      element.children !== null &&
      element.children.length > 0
    ) {
      dynamicRouterFunc(router, element.children);
    }
  }
}
