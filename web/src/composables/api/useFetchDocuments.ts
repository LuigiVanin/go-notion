import { useAsync } from "../useAsync.ts";
import { useApi } from "./useApi.ts";

export const useFetchDocuments = () => {
    const api = useApi();

    const {
        action: fetchDocuments,
        data: documents,
        error: fetchError,
        loading: fetchLoading,
    } = useAsync(async () => {
        const { data, error } = await api.document.fetch();
        if (error) {
            throw error;
        }
        return data;
    });

    return {
        fetchDocuments,
        documents,
        fetchError,
        fetchLoading,
    };
};
