import { Ref } from "vue";
import { useAsync } from "../useAsync.ts";
import { useApi } from "./useApi.ts";
import { ApiError } from "@/types/api.ts";

export const useFetchDocument = () => {
    const api = useApi();

    const fetchDocument = async (id: string) => {
        const { data, error } = await api.document.fetchOne(id);
        if (error) {
            throw error;
        }
        return data;
    };

    const { action, data, error, loading } = useAsync(fetchDocument, false);

    return {
        fetchDocument: action as (id: string) => Promise<void>,
        document: data,
        fetchError: error as Ref<null | ApiError>,
        fetchLoading: loading,
    };
};
