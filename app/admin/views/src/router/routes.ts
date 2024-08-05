import { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'Layouts',
    component: () => import('layouts/main.vue'),
    children: [
      { path: '/', component: () => import('pages/index.vue'), meta: { requireAuth: true, keepAlive: true }}
    ],
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/login',
    name: 'Login',
    component: () => import('pages/login.vue'),
  },
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/404.vue'),
  },
];

export default routes;
