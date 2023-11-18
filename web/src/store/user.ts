// Core
import { defineStore } from "pinia";

// Helpers & Services
import { authStorage, userStorage } from "@/helpers/storage.ts";

// Types
import { User } from "@/types/user";

export const useUserStore = defineStore({
    id: "user",
    state: () => {
        let user: User | null = userStorage.getItem();
        let token: { jwt: string } | null = authStorage.getItem();

        return {
            token,
            user,
        };
    },
    getters: {},
    actions: {
        setUser(user: User) {
            console.log("AQUIIII");

            this.user = user;

            console.log("THIS.USER", this.user);

            userStorage.setItem(user);
        },
        setToken(token: string) {
            this.token = { jwt: token };
            authStorage.setItem({ jwt: token });
        },
        setAuth(auth: { user: User; token: string }) {
            console.log("SET AUTH");
            this.setUser(auth.user);
            this.setToken(auth.token);
        },
        clearUser() {
            this.user = null;
            userStorage.removeItem();
            authStorage.removeItem();
        },
    },
});
