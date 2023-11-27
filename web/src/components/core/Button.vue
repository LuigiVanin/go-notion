<script lang="ts" setup>
// Libraries
import { ref } from "vue";
import InlineSvg from "vue-inline-svg";

type ButtonsProps = {
    text?: string;
    btnType?: "soft" | "no-border" | "filled" | "outlined" | "simple";
    color?: "primary" | "blue" | "orange" | "red";
    size?: "sm" | "md" | "lg";
    icon?: string;
    loading?: boolean;
    disabled?: boolean;

    suffixIcon?: string;
};

const props = defineProps<ButtonsProps>();
const emit = defineEmits(["click"]);

const activeButton = ref(false);

const handleClick = () => {
    activeButton.value = !activeButton.value;
    emit("click");
    setTimeout(() => {
        activeButton.value = !activeButton.value;
    }, 200);
};
</script>

<template>
    <button
        :class="[
            'button-main',
            `button-main--${props.btnType || 'filled'}`,
            `button-main--color-${props.color || 'primary'}`,
            `button-main--size-${props.size || 'md'}`,
            `button-main--${activeButton ? 'active' : ''}`,
        ]"
        :disabled="props.disabled"
        @click="handleClick"
    >
        <span>
            <InlineSvg
                v-if="props.icon"
                :src="props.icon"
                class="button-main__icon"
            />
            <slot>
                {{ props.text }}
            </slot>
            <InlineSvg
                v-if="props.suffixIcon"
                :src="props.suffixIcon"
                class="button-main__icon"
            />
        </span>
    </button>
</template>

<style lang="scss" scoped>
button.button-main {
    @include reset-button;
    padding: $spacing_3 $spacing_8;
    border-radius: $border_r_md;
    transition: $transition, transform 0.05s $transition_spring;

    --button-color-strong: var(--primary_6);
    --button-color-medium: var(--primary_5);
    --button-color-soft: var(--primary_2);

    span {
        @include flex-center;
        gap: $spacing_3;
    }

    :deep(svg.button-main__icon) {
        path,
        rect,
        circle {
            stroke: currentColor;
        }
    }

    &--active,
    &:active {
        transform: translate(calc(-2px), calc(2px));
    }

    &--size-sm {
        padding: $spacing_3 $spacing_8;
        font-size: $font_3;

        :deep(svg.button-main__icon) {
            width: $button_icon_small;
            height: $button_icon_small;
        }
    }

    &--size-md {
        padding: $spacing_5 $spacing_10;
        font-size: $font_4;

        :deep(svg.button-main__icon) {
            width: $button_icon_medium;
            height: $button_icon_medium;
        }
    }

    &--size-lg {
        padding: $spacing_6 $spacing_13;
        font-size: $font_5;
        border-radius: $border_r_mdx;

        :deep(svg.button-main__icon) {
            width: $button_icon_large;
            height: $button_icon_large;
        }

        &:hover {
            // font-weight: 500;
        }

        span {
            gap: $spacing_4;
        }
    }

    &--color-primary {
        --button-color-strong: var(--primary_6);
        --button-color-medium: var(--primary_5);
        --button-color-soft: var(--primary_2);
    }

    &--color-blue {
        --button-color-strong: var(--blue_6);
        --button-color-medium: var(--blue_5);
        --button-color-soft: var(--blue_2);
    }

    &--color-red {
        --button-color-strong: var(--red_6);
        --button-color-medium: var(--red_5);
        --button-color-soft: var(--red_2);
    }

    &--filled {
        background-color: var(--button-color-strong);
        border: 1px solid var(--button-color-strong);
        color: $neutral_1;

        &:hover {
            background-color: var(--button-color-medium);
            border: 1px solid var(--button-color-medium);
            color: $neutral_1;
        }
    }

    &--soft {
        background-color: var(--button-color-soft);
        border: 1px solid var(--button-color-soft);
        color: var(--button-color-medium);

        &:hover {
            background-color: var(--button-color-medium);
            border: 1px solid var(--button-color-medium);
            color: $neutral_1;
        }
    }

    &--outlined {
        background-color: transparent;
        border: 1px solid var(--button-color-strong);
        color: var(--button-color-strong);

        &:hover {
            background-color: var(--button-color-medium);
            border: 1px solid var(--button-color-medium);
            color: $neutral_1;
        }
    }

    &--no-border {
        background-color: transparent;
        border: 1px solid transparent;
        color: $neutral_10;

        &:hover {
            background-color: $neutral_3;
            text-decoration: underline;
            border: 1px solid transparent;
            color: var(--button-color-strong);
        }
    }

    &--simple {
        background-color: transparent;
        border: 1px solid $neutral_6;
        color: $neutral_9;

        &:hover {
            background-color: $neutral_3;
            border: 1px solid $neutral_6;
            color: $neutral_12;
        }
    }

    &:disabled,
    &:hover:disabled {
        --button-color-strong: var(--neutral_6);
        --button-color-medium: var(--neutral_5);
        --button-color-soft: var(--neutral_2);

        cursor: not-allowed !important;
    }
}
</style>
