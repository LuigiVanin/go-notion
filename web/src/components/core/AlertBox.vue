<script lang="ts" setup>
// Core
import { Transition, computed, ref } from "vue";

// Libraries
import InlineSvg from "vue-inline-svg";

// Assets
import alertCircleIconUrl from "@/assets/icons/alert-circle-outline.svg?url";
import infoCircleIconUrl from "@/assets/icons/info-circle.svg?url";
import closeIconUrl from "@/assets/icons/close.svg?url";

type AlertBoxType = "error" | "success" | "warning" | "info";

type AlertBoxProps = {
    text: string;
    type?: AlertBoxType;
    icon?: string;
    show?: boolean;
    closable?: boolean;
};

const iconsLUT: Record<AlertBoxType, string> = {
    error: alertCircleIconUrl,
    success: alertCircleIconUrl,
    warning: alertCircleIconUrl,
    info: infoCircleIconUrl,
};

const props = withDefaults(defineProps<AlertBoxProps>(), {
    type: "error",
    show: true,
});

const closedAlert = ref(false);

const boxIcon = computed(() => {
    return props.icon || iconsLUT[props.type] || alertCircleIconUrl;
});
</script>

<template>
    <Transition name="alert-box-fade">
        <div
            v-if="props.show && !closedAlert"
            :class="['alert-box', `alert-box--${props.type}`]"
        >
            <InlineSvg :src="boxIcon" />
            <p>{{ props.text }}</p>
            <InlineSvg
                v-if="props.closable"
                :src="closeIconUrl"
                class="alert-box__close"
                @click="closedAlert = true"
            />
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
    opacity: 1;

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
        --alert-strong-color: var(--blue_5);
    }

    background: var(--alert-soft-color);
    border: 1px solid var(--alert-strong-color);

    :deep(svg) {
        @include set-fixed-width(25px);

        path {
            stroke: var(--alert-strong-color);
        }
    }

    :deep(svg.alert-box__close) {
        @include set-fixed-width(20px);
        height: 20px;
        margin-left: auto;

        &:hover {
            cursor: pointer;
            opacity: 0.8;
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
}

.alert-box-fade-enter-from,
.alert-box-fade-leave-to {
    opacity: 0;
    height: 0px;
}
</style>
