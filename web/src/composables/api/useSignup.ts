// Core
import { Ref } from "vue";

// Composables
import { useApi } from "@/composables/api/useApi.ts";
import { useAsync } from "../useAsync.ts";

// Types
import { ApiError } from "@/types/api.ts";
import { SignupForm } from "@/types/user.ts";

export const useSignup = () => {
    const api = useApi();

    const signup = async (payload: SignupForm) => {
        const { data, error } = await api.auth.signup(payload);
        if (error) {
            throw error;
        }
        return data;
    };

    const { action, data, error, loading } = useAsync(signup, false);

    return {
        signup: action as typeof signup,
        data,
        signupError: error as Ref<null | ApiError>,
        signupLoading: loading,
    };
};
