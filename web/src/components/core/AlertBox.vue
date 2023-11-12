<script lang="ts" setup>
// Core
import { Transition } from "vue";

// Libraries
import InlineSvg from "vue-inline-svg";

// Assets
import alertCircleIconUrl from "@/assets/icons/alert-circle-outline.svg?url";

type AlertBoxProps = {
    text: string;
    type?: "error" | "success" | "warning" | "info";
    icon?: string;
    show?: boolean;
};

const props = withDefaults(defineProps<AlertBoxProps>(), {
    type: "error",
    icon: alertCircleIconUrl,
    show: true,
});
</script>

<template>
    <Transition name="alert-box-fade">
        <div
            v-if="props.show"
            :class="['alert-box', `alert-box--${props.type}`]"
        >
            <InlineSvg :src="props.icon" />
            <p>{{ props.text }}</p>
        </div>
    </Transition>
</template>

<style lang="scss" scoped>
.alert-box {
    --alert-soft-color: var(--red_1);
    --alert-strong-color: var(--red_6);

    border-radius: $border_r_md;
    width: 100%;
    padding: $spacing_8 $spacing_8;
    @include flex(row, center, start);
    gap: $spacing_8;

    &--error {
        --alert-soft-color: var(--red_1);
        --alert-strong-color: var(--red_6);
    }

    &--success {
        --alert-soft-color: var(--green_1);
        --alert-strong-color: var(--green_6);
    }

    &--warning {
        --alert-soft-color: var(--orange_1);
        --alert-strong-color: var(--orange_6);
    }

    &--info {
        --alert-soft-color: var(--blue_1);
        --alert-strong-color: var(--blue_6);
    }

    background: var(--alert-soft-color);
    border: 1px solid var(--alert-strong-color);

    :deep(svg) {
        width: 25px;
        height: 25px;

        path {
            stroke: var(--alert-strong-color);
        }
    }

    p {
        color: var(--alert-strong-color);
        font-size: $font_3;
        text-align: start;
        line-height: 1.25rem;
    }
}

.alert-box-fade-enter-active,
.alert-box-fade-leave-active {
    transition: all 0.1s $transition_spring;
    height: 47px;
}

.alert-box-fade-enter-from {
    opacity: 0;
    height: 0px;
}
</style>
