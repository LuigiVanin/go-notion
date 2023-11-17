// Libraries
import axios from "axios";

// Helpers
import { authStorage } from "@/helpers/storage.ts";

// Types
import { ISigninForm, SigninResponse, SignupForm } from "@/types/user.ts";
import { ApiError, ApiResponse, Query } from "@/types/api.ts";

const baseUrl = import.meta.env.VITE_API_URL;
const axiosInstance = axios.create({
    baseURL: baseUrl,
});

axiosInstance.interceptors.request.use((config) => {
    const token = authStorage.getItem();
    if (token) {
        config.headers.Authorization = `${token.jwt}`;
    }
    return config;
});

async function fetch<T>(
    basePath: string,
    query?: Query
): Promise<ApiResponse<T>> {
    try {
        const queryString =
            query &&
            Object.entries(query || {})
                .map(([key, value]) => `${key}=${value}`)
                .join("&");
        const url = queryString ? `${basePath}?${queryString}` : basePath;
        const res = await axiosInstance.get(url);
        return { data: res.data, error: null };
    } catch (err) {
        return { data: null, error: err as ApiError };
    }
}

async function create<T, R>(
    basePath: string,
    body?: T
): Promise<ApiResponse<R>> {
    try {
        const res = await axiosInstance.post(basePath, body);
        return { data: res.data, error: null };
    } catch (err) {
        return { data: null, error: err as ApiError };
    }
}

const registerCreateStrategy =
    <T, R>(basePath: string) =>
    (body?: T) =>
        create<T, R>(basePath, body);

const registerFetchStrategy =
    <T>(basePath: string) =>
    (query?: Query) =>
        fetch<T>(basePath, query);

interface IUser {}

export class Api {
    auth = {
        signin: registerCreateStrategy<ISigninForm, SigninResponse>(
            "/auth/signin"
        ),
        signup: registerCreateStrategy<SignupForm, SigninResponse>(
            "/auth/signup"
        ),
    };
    user = {
        fetch: registerFetchStrategy<IUser>("/user"),
    };
}
