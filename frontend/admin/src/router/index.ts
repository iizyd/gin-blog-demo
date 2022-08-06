import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";

const routes: RouteRecordRaw[] = [
  {
    path: "/",
    name: "board",
    redirect: '/article',
    component: () => import("@/view/Layout/Layout.vue"),
    children: [
      {
        path: "article",
        name: "article",
        component: () => import("@/view/Article/index.vue"),
      },
      {
        path: "tag",
        name: "tag",
        component: () => import("@/view/Tag/index.vue"),
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
