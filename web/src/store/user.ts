// Core
import { defineStore } from "pinia";

// Types
import { User } from "@/types/user";

const USER_STORAGE_KEY = "writtable-user";
const AUTH_STORAGE_KEY = "writtable-auth";

export const useUserStore = defineStore({
    id: "user",
    state: () => {
        let user: User | null = null;
        let token: { jwt: string } | null = null;
        const userSerialized = localStorage.getItem(USER_STORAGE_KEY);
        const tokenSerialized = localStorage.getItem(AUTH_STORAGE_KEY);

        if (userSerialized) {
            try {
                user = JSON.parse(userSerialized);
            } catch {
                user = null;
            }
        }
        if (tokenSerialized) {
            try {
                token = JSON.parse(tokenSerialized);
            } catch {
                token = null;
            }
        }

        return {
            token,
            user,
        };
    },
    getters: {},
    actions: {
        setUser(user: User) {
            this.user = user;
            console.log("SET USER");

            localStorage.setItem(USER_STORAGE_KEY, JSON.stringify(user));
        },
        setToken(token: string) {
            this.token = { jwt: token };
            localStorage.setItem(
                AUTH_STORAGE_KEY,
                JSON.stringify({ jwt: token })
            );
        },
        setAuth(auth: { user: User; token: string }) {
            console.log("SET AUTH");
            this.setUser(auth.user);
            this.setToken(auth.token);
        },
        clearUser() {
            this.user = null;
            localStorage.removeItem(USER_STORAGE_KEY);
        },
    },
});
