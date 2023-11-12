export interface ZodGeneric<T> {
    parse: (value: T) => T;
}

export type Rule<T = string> = {
    validator?: ZodGeneric<T>;

    message: string;
} & {
    validation?: (value: T) => boolean;

    message: string;
};

export enum ValidationError {
    Success = "--success--",
    NotError = "--not-error--",
    Invalid = "--invalid--",
}

export type Rules<T = any> = Record<keyof T, Array<Rule>>;
