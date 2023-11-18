<script lang="ts" setup>
// Core
import { onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";

const route = useRoute();
const router = useRouter();

onMounted(() => {
    if (!route.params?.id) {
        router.push("/");
    }
});
</script>

<template>
    {{ route.params?.id }}
</template>
