import { createRouter, createWebHistory, type RouteRecordRaw } from "vue-router";
import Login from "../pages/Login.vue";
import Discover from "../pages/Discover.vue";
import Signup from "../pages/Signup.vue";
import ElHome from "../total-client-ui/Home.vue";

const routes: RouteRecordRaw[] = [
  { path: "/login", component: Login },
  { path: "/", component: Discover },
  { path: "/discover", component: Discover },
  { path: "/signup", component: Signup },
  { path: "/home-elarduspark", component: ElHome },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;