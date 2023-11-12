<script lang="ts" setup>
import { RouterView } from "vue-router";
</script>

<template>
    <div class="layout-container">
        <RouterView />
    </div>
</template>

<style scoped lang="scss">
.layout-container {
    @include flex-center();
    height: 100vh;
}
</style>
