import { AxiosError } from "axios";

export type ApiErrorData = {
    message: string;
    code: number;
    fields: { tag: string; field: string }[];
};

export type ApiError = AxiosError<ApiErrorData>;

export type ApiResponse<T, Err = ApiError> = {
    data: T | null;
    error: Err | null;
};

export type Query = Record<string, string | number>;
