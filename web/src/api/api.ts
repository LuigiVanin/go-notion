// Libraries
import axios from "axios";

// Helpers
import { authStorage } from "@/helpers/storage.ts";

// Types
import { ISigninForm, SigninResponse, SignupForm, User } from "@/types/user.ts";
import { ApiError, ApiResponse, Query } from "@/types/api.ts";
import { CreateDocument, Document } from "@/types/document";

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

async function update<T, R>(
    basePath: string,
    body?: T
): Promise<ApiResponse<R>> {
    try {
        const res = await axiosInstance.patch(basePath, body);
        return { data: res.data, error: null };
    } catch (err) {
        return { data: null, error: err as ApiError };
    }
}

const registerCreateStrategy =
    <T, R = void>(basePath: string) =>
    (body?: T) =>
        create<T, R>(basePath, body);

const registerPatchParamStrategy =
    <T, R = void>(basePath: string) =>
    (param: number | string, body?: T) =>
        update<T, R>(`${basePath}/${param}`, body);

const registerFetchStrategy =
    <T>(basePath: string) =>
    (query?: Query) =>
        fetch<T>(basePath, query);

const registerFetchParamStrategy =
    <T>(basePath: string) =>
    (param: number | string, query?: Query) =>
        fetch<T>(`${basePath}/${param}`, query);

export class Api {
    auth = {
        signin: registerCreateStrategy<ISigninForm, SigninResponse>(
            "/auth/signin"
        ),
        signup: registerCreateStrategy<SignupForm>("/auth/signup"),
    };
    user = {
        fetch: registerFetchStrategy<User>("/user"),
    };
    document = {
        fetch: registerFetchStrategy<{ docs: Document[] }>("/document"),
        fetchOne: registerFetchParamStrategy<Document>("/document"),
        create: registerCreateStrategy<CreateDocument, { id: number }>(
            "/document"
        ),
        update: registerPatchParamStrategy<Partial<CreateDocument>>(
            "/document"
        ),
    };
}
