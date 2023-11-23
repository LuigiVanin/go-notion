// Core
import { Ref } from "vue";

// Composables
import { useApi } from "@/composables/api/useApi.ts";
import { useAsync } from "../useAsync.ts";

// Types
import { ApiError } from "@/types/api.ts";
import { CreateDocument } from "@/types/document.ts";

export const useUpdateDocument = () => {
    const api = useApi();

    const update = async (
        id: string | number,
        payload: Partial<CreateDocument>
    ) => {
        const { data, error } = await api.document.update(id, payload);
        if (error) {
            throw error;
        }
        return data;
    };

    const { action, data, error, loading } = useAsync(update, false);

    return {
        updateDocument: action as typeof update,
        data,
        udpateError: error as Ref<null | ApiError>,
        updateLoading: loading,
    };
};
