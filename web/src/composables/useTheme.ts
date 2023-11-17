import { useDark, useToggle } from "@vueuse/core";

export const useTheme = () => {
    const isDark = useDark({
        selector: "body",
        attribute: "data-theme",
        valueDark: "dark",
        valueLight: "light",
    });
    const toggleTheme = useToggle(isDark);

    return {
        isDark,
        toggleTheme,
    };
};
