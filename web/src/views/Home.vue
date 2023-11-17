<script lang="ts" setup>
import { useUserStore } from "@/store/user.ts";

const userStore = useUserStore();
</script>

<template>
    <div class="home">
        <p>USER: {{ userStore.user?.email }}</p>
    </div>
</template>

<style lang="scss" scoped>
.home p {
    color: $neutral_11 !important;
}
</style>
