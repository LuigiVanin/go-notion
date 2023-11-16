<script lang="ts" setup>
// Libraries
import { ref } from "vue";
import InlineSvg from "vue-inline-svg";

type ButtonsProps = {
    text?: string;
    btnType?: "soft" | "no-border" | "filled" | "outlined";
    color?: "primary" | "blue" | "orange";
    size?: "sm" | "md" | "lg";
    icon?: string;

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
            width: 15px;
            height: 15px;
        }
    }

    &--size-md {
        padding: $spacing_5 $spacing_10;
        font-size: $font_4;

        :deep(svg.button-main__icon) {
            width: 17px;
            height: 17px;
        }
    }

    &--size-lg {
        padding: $spacing_6 $spacing_13;
        font-size: $font_5;

        :deep(svg.button-main__icon) {
            width: 20px;
            height: 20px;
        }

        &:hover {
            font-weight: 500;
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
}
</style>
