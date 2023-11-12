import { Ref, computed, ref } from "vue";
import { type Rules, ValidationError } from "@/types/utils.ts";

export const useValidation = <T>(
    input: Record<keyof T, any>,
    rules: Rules<T>
) => {
    const errors: Ref<Record<keyof T, ValidationError | string>> = ref(
        {} as any
    );

    const valid = computed(() => {
        for (const key in input) {
            if (errors.value[key] !== ValidationError.Success) {
                return false;
            }
        }
        return true;
    });

    const validate = () => {
        errors.value = {} as Record<keyof T, ValidationError>;

        for (const field in input) {
            validateField(field);
        }
    };

    const validateField = (field: keyof T) => {
        const value = input[field];
        const fieldRules = rules[field];
        errors.value[field] = ValidationError.Success;

        for (const rule of fieldRules) {
            if (rule.validator) {
                try {
                    rule.validator.parse(value);
                } catch (err) {
                    errors.value[field] = rule.message;
                }
            } else if (rule.validation && !rule?.validation(value)) {
                errors.value[field] = rule.message;
            }
        }

        return errors.value[field] || ValidationError.Success;
    };

    const fieldStatus = (field: keyof T) => {
        if (!errors.value[field]) {
            return ValidationError.NotError;
        }
        if (errors.value[field] !== ValidationError.Success) {
            return ValidationError.Invalid;
        }
        return errors.value[field];
    };

    return { errors, valid, validate, fieldStatus, validateField };
};
