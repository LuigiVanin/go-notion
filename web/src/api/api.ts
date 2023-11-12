import axios from "axios";

type ApiResponse<T, Err = unknown> = {
    data: T | null;
    error: Err | null;
};

interface Strategy {
    execute: (...args: any[]) => Promise<ApiResponse<unknown>>;
}

export class FetchStrategy<T> implements Strategy {
    protected basePath: string;

    constructor(basePath: string) {
        this.basePath = basePath;
    }

    async execute(
        query?: Record<string, string | number>
    ): Promise<ApiResponse<T>> {
        try {
            const queryString =
                query &&
                Object.entries(query || {})
                    .map(([key, value]) => `${key}=${value}`)
                    .join("&");
            const url = queryString
                ? `${this.basePath}?${queryString}`
                : this.basePath;
            const res = await axios.get(url);
            return { data: res.data, error: null };
        } catch (err) {
            return { data: null, error: err };
        }
    }
}

export class CreateStrategy<R, T> implements Strategy {
    protected basePath: string;

    constructor(basePath: string) {
        this.basePath = basePath;
    }

    async execute(payload: R): Promise<ApiResponse<T>> {
        try {
            const res = await axios.post(this.basePath, payload);
            return { data: res.data, error: null };
        } catch (err) {
            return { data: null, error: err };
        }
    }
}

class Api {
    user = {
        fetch: new FetchStrategy<{ email: string }>("/api/user").execute,
        create: new CreateStrategy<undefined, { email: string }>("/api/user")
            .execute,
    };
}

export const api = new Api();
