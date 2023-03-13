import { createWebHashHistory, createRouter, RouteRecordRaw } from "vue-router";
import Home from "@/views/Home/index.vue";
import Index from "../views/Index/index.vue";
import type { App } from "vue";

const routes: RouteRecordRaw[] = [
  {
    path: "/",
    component: Home,
  },
  {
    path: "/home",
    component: Home,
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export const setupRouter = (app: App) => {
  app.use(router);
};
