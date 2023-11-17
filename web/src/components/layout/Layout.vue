<script lang="ts" setup>
// Core
import { RouterView } from "vue-router";

// Composables
import { useTheme } from "@/composables/useTheme.ts";

// Components
import InlineSvg from "vue-inline-svg";
import InitialHeader from "@/components/layout/partials/InitialHeader.vue";

// Assets
import gradientBackgroundUrl from "@/assets/illustration/background.svg?url";

useTheme();
</script>

<template>
    <div class="layout-container">
        <InitialHeader />
        <InlineSvg
            :src="gradientBackgroundUrl"
            class="layout-container__background"
        />
        <RouterView />
    </div>
</template>

<style scoped lang="scss">
.layout-container {
    @include flex-center();
    height: 100vh;
    position: relative;
    background: $neutral_1;

    .layout-container__background {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        // z-index: -1;
    }
}
</style>
