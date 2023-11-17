<script lang="ts" setup>
// Composables
import { useTheme } from "@/composables/useTheme.ts";

// Libraries
import InlineSvg from "vue-inline-svg";

// Assets
import sunnyIconUrl from "@/assets/icons/sunny-outline.svg?url";
import moonIconUrl from "@/assets/icons/moon-outline.svg?url";

const { toggleTheme, isDark } = useTheme();

const handleThemeToggle = () => {
    toggleTheme();
};
</script>

<template>
    <button
        :class="['theme-switch', isDark && 'theme-switch--dark']"
        @click="handleThemeToggle"
    >
        <InlineSvg :src="moonIconUrl" />
        <span class="theme-switch__core" />
        <InlineSvg :src="sunnyIconUrl" />
    </button>
</template>

<style lang="scss" scoped>
.theme-switch {
    @include reset-button();
    @include flex(row, center, space-between);
    transition: all 0.2s ease-in-out;
    height: $theme_switch_height;
    width: $theme_switch_width;
    background: transparent;
    border-radius: calc($theme_switch_height / 2);
    border: 1px solid $neutral_7;
    position: relative;
    background: $neutral_2;
    cursor: pointer;
    padding-inline: $spacing_3;

    :deep(svg) {
        width: 16px;
        height: 16px;
        transition: all 0.3s ease-in-out;
        margin-left: 1px;

        &:last-child {
            opacity: 1;
        }

        &:first-child {
            opacity: 0;
        }

        path,
        circle {
            transition: all 0.3s ease-in-out;
            stroke: $primary_7;
        }
    }

    .theme-switch__core {
        height: calc($theme_switch_height - 8px);
        width: calc($theme_switch_height - 8px);
        border-radius: calc(($theme_switch_height - 8px) / 2);
        background: $primary_6;
        transition: left 0.2s $transition_spring !important;
        position: absolute;
        left: 3px;
    }

    &.theme-switch--dark {
        :deep(svg) {
            &:last-child {
                opacity: 0;
            }

            &:first-child {
                opacity: 1;
            }
        }

        opacity: 1 !important;

        .theme-switch__core {
            left: calc($theme_switch_width - $theme_switch_height + 3px);
        }
    }

    &:hover {
        border-color: $primary_6;
        background: $neutral_3;
    }
}
</style>
