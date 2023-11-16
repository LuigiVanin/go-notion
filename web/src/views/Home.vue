<script lang="ts" setup>
import { useUserStore } from "@/store/user.ts";

const userStore = useUserStore();
</script>

<template>
    <div class="">
        <p>USER: {{ userStore.user?.email }}</p>
    </div>
</template>
