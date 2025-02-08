import { createRouter, createWebHistory, type RouteRecordRaw } from "vue-router";
import Login from "../pages/Login.vue";
import Welcome from "../components/Welcome.vue";
import Signup from "../pages/Signup.vue";

const routes: RouteRecordRaw[] = [
  { path: "/login", component: Login },
  { path: "/", component: Login },
  { path: "/welcome", component: Welcome },
  { path: "/signup", component: Signup },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;