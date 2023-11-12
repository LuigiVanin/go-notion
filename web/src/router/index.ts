import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import NProgress from "nprogress";

const routes: RouteRecordRaw[] = [
    {
        path: "/",
        component: () => import("@/layouts/Layout.vue"),
        children: [
            {
                path: "/signup",
                name: "Sign Up",
                component: () => import("@/views/Signup.vue"),
            },
            {
                path: "/",
                name: "Sign In",
                component: () => import("@/views/Signin.vue"),
            },
        ],
    },
    {
        path: "/",
        component: () => import("@/layouts/Layout.vue"),
        children: [
            {
                path: "/home",
                name: "Home",
                component: () => import("@/views/Home.vue"),
            },
        ],
    },
];

export const router = createRouter({
    history: createWebHistory(),
    routes,
    scrollBehavior(_to, _from, savedPosition) {
        if (savedPosition) {
            return savedPosition;
        } else {
            return { top: 0 };
        }
    },
});

router.beforeEach(() => {
    NProgress.start();
});

router.afterEach(() => {
    NProgress.done();
});
