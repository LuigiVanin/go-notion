<script lang="ts" setup>
// Core
import { ref, watch } from "vue";

// Libraries
import InlineSvg from "vue-inline-svg";
import { useMouse, useMouseInElement } from "@vueuse/core";

// Assets
import magicStarsIconUrl from "@/assets/icons/magic-stars.svg?url";

type SpecialButtonProps = {
    text: string;
    loading?: boolean;
    type?: "button" | "submit" | "reset";
};

const props = defineProps<SpecialButtonProps>();
const emit = defineEmits(["click"]);

const specialButtonEl = ref<HTMLButtonElement | null>(null);

const { x, y } = useMouse();
const { isOutside } = useMouseInElement(specialButtonEl);

const mouseUpdate = () => {
    if (specialButtonEl.value && !isOutside.value) {
        const bounds = specialButtonEl.value?.getBoundingClientRect();

        specialButtonEl.value.style.setProperty(
            "--rx",
            `${(x.value - bounds.x) / bounds.width}`
        );
        specialButtonEl.value.style.setProperty(
            "--x",
            `${(x.value - bounds.x) / bounds.width}`
        );
        specialButtonEl.value.style.setProperty(
            "--y",
            `${(y.value - bounds.y) / bounds.height}`
        );
    }
};

watch(x, mouseUpdate);
watch(y, mouseUpdate);
</script>

<template>
    <button ref="specialButtonEl" @click="emit('click')" :type="props.type">
        <span class="backdrop" />
        <span>
            <InlineSvg :src="magicStarsIconUrl" />
            {{ props.text }}
        </span>
    </button>
</template>

<style lang="scss" scoped>
button {
    @include reset-button;
    --text-padding: 8px 16px;
    --border: 1px;
    --padding: 1px;
    cursor: pointer;
    border-radius: 8px;

    background: hsl(0 0% 90%);
    box-shadow: inset 0 1px 0px 0px hsl(0 0% 100% / 0.5),
        inset 0 -1px 0px 0px hsl(0 0% 0% / 0.5);

    font-size: $font_4;
    position: relative;
    display: grid;
    place-items: center;
    padding: $spacing_0;
    border: var(--border) solid hsl(0 0% 80%);
    transform: translate(
        calc(var(--active, 0) * -2px),
        calc(var(--active, 0) * 2px)
    );
    transition: transform 0.1s;
    color: hsl(0 0% 20%);

    &:is(:hover, :focus-visible) {
        --hover: 1;
        /*   border-color: transparent !important; */
    }

    &:active {
        --active: 1;
    }

    &:before {
        content: "";
        position: absolute;
        inset: 0px;
        border-radius: calc(8px - var(--border));
        background: hsl(0 0% 100% / calc(1 - var(--hover, 0) * 0.25));
        background: grey;
        background: conic-gradient(
            from calc(var(--rx, 0) * 180deg) at calc(var(--x, 0) * 100%)
                calc(var(--y, 0) * 100%),
            hsl(10 90% 70%),
            hsl(140 80% 70%),
            hsl(320 80% 70%),
            hsl(210 80% 70%),
            hsl(10 80% 70%)
        );
        filter: saturate(0.7);
        opacity: var(--hover, 0);
        transition: opacity 0.2s;
    }

    .backdrop {
        position: relative;
        width: 100%;
        height: 100%;
        background: hsl(0 0% 98% / 0.6);
        border-radius: calc(8px - $spacing_0);
        display: block;
        grid-column: 1;
        grid-row: 1;
        backdrop-filter: blur(20px) brightness(1.5);
    }

    > span {
        padding: var(--text-padding);
        grid-row: 1;
        grid-column: 1;
        z-index: 2;
        @include flex(inline-flex, center, center);
        gap: 0.5rem;

        :deep(svg) {
            width: 20px;
        }
    }
}
</style>
