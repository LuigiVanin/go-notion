<script lang="ts" setup>
// Libraries
import InlineSvg from "vue-inline-svg";

type EditorButtonIconProps = {
    icon: string;
    action: () => void;
};

const props = defineProps<EditorButtonIconProps>();
</script>

<template>
    <button @click="props.action">
        <InlineSvg :src="props.icon" />
    </button>
</template>

<style lang="scss" scoped>
button {
    @include reset-button;
    @include flex-center;
    width: 32px;
    height: 32px;

    border-radius: $border_r_mdx;
    border: 1px solid $neutral_5;
    transition: all 0.2s ease-in-out;

    &:hover {
        background: $neutral_3;
    }

    :deep(svg) {
        width: 13px;
        height: 13px;

        path,
        rect,
        circle {
            fill: $primary_6;
        }
    }
}
</style>
