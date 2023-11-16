import { Api } from "@/api/api.ts";

export const useApi = () => {
    const api = new Api();
    return api;
};
