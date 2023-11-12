import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";
import Signin from "@/views/Signin.vue";

const routes: RouteRecordRaw[] = [
    {
        path: "/",
        component: () => import("@/layouts/Layout.vue"),
        children: [
            {
                path: "/",
                component: Signin,
            },
        ],
    },
];

export const router = createRouter({
    history: createWebHashHistory(),
    routes,
});
