import { computed, onMounted, reactive } from "vue";
import { z } from "zod";

const createNullObject = <T extends {}>(obj: T) => {
    return Object.fromEntries(
        Object.keys(obj).map((key) => [key, null])
    ) as Record<keyof T, string | null>;
};

export const useValidation = <T>(
    data: Record<keyof T, string | number>,
    rules: z.ZodObject<z.ZodRawShape>,
    initialValues?: Record<keyof T, string | null>
) => {
    const fieldErrors: Record<keyof T, string | null> = reactive(
        initialValues || createNullObject(data)
    );

    const resetFieldErrors = (field?: string) => {
        for (const key in fieldErrors) {
            if (field && key === field) {
                fieldErrors[key as keyof T] = null;
            }
        }
    };

    const validate = (field?: string) => {
        resetFieldErrors(field);
        try {
            rules.parse(data);
            return true;
        } catch (err) {
            if (err instanceof z.ZodError) {
                err.issues.forEach((issue) =>
                    issue.path.forEach((path) => {
                        if (path in fieldErrors && (!field || field === path)) {
                            fieldErrors[path as keyof T] = issue.message;
                        }
                    })
                );
            }
            return false;
        }
    };

    const isValid = computed(() => {
        return Object.values(fieldErrors).every((error) => !error);
    });

    onMounted(() => {
        resetFieldErrors();
    });

    return {
        isValid,
        validate,
        fieldErrors,
        resetFieldErrors,
    };
};
