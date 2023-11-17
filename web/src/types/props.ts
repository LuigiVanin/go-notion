export type InputProps = {
    modelValue: string;
    placeholder?: string;
    type?: string;
    name?: string;
    rightIcon?: string;
    icon?: string;
    maxWidth?: string;
    label?: string;
    errorMessage?: string | null | undefined;
    mask?: string;
    optional?: boolean;
};

export type Action = {
    icon?: string;
    disabled?: boolean;
    class?: string;
    label: string;
    onClick: () => void;
};
