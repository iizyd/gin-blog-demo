import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";

const routes:RouteRecordRaw[] = [
  {
    path: "/",
    name: 'board',
    component: () => import('@/view/Board/index.vue')
  },
  {
    path: "/article",
    name: 'article',
    component: () => import('@/view/Article/index.vue')
  },
  {
    path: "/tag",
    name: 'tag',
    component: () => import('@/view/Tag/index.vue')
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
