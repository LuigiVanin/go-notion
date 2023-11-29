// Core
import { Ref } from "vue";

// Composables
import { useApi } from "@/composables/api/useApi.ts";
import { useAsync } from "../useAsync.ts";

// Stores
import { useUserStore } from "@/store/user.ts";

// Types
import { ApiError } from "@/types/api.ts";
import { ISigninForm } from "@/types/user.ts";

export const useSignin = () => {
    const api = useApi();
    const userStore = useUserStore();

    const signin = async (payload: { email: string; password: string }) => {
        const { data, error } = await api.auth.signin(payload);
        if (error) {
            throw error;
        }
        if (data) {
            userStore.setAuth({
                user: data.user,
                token: data.token,
            });
        }
        return data;
    };

    const { action, data, error, loading } = useAsync(signin, false);

    return {
        signin: action as (payload: ISigninForm) => Promise<void>,
        userData: data,
        loginError: error as Ref<null | ApiError>,
        loginLoading: loading,
    };
};
