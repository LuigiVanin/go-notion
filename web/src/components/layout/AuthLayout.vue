<script lang="ts" setup>
// Core
import { RouterView } from "vue-router";

// Components
import AuthHeader from "./partials/AuthHeader.vue";
</script>

<template>
    <div class="layout-container">
        <AuthHeader />
        <RouterView />
    </div>
</template>

<style scoped lang="scss">
.layout-container {
    @include flex-center();
    height: 100vh;
    position: relative;
    background: $neutral_2;
}
</style>
