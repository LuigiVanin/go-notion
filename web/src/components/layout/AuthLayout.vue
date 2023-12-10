<script lang="ts" setup>
// Core
import { onMounted } from "vue";
import { RouterView, useRouter } from "vue-router";

// Components
import AuthHeader from "./partials/AuthHeader.vue";
import { useApi } from "@/composables/api/useApi.ts";
import { useUserStore } from "@/store/user.ts";

const router = useRouter();

const api = useApi();
const userStore = useUserStore();

onMounted(async () => {
    const { data, error } = await api.user.fetch();

    if (error || !data) {
        router.push("/");
        return;
    }
    userStore.setUser(data);
});
</script>

<template>
    <div class="layout-container">
        <AuthHeader />
        <main>
            <RouterView />
        </main>
    </div>
</template>

<style scoped lang="scss">
.layout-container {
    @include flex(column, center, start);
    height: 100vh;
    position: relative;
    background: $neutral_2;

    padding-top: $header_height;

    main {
        padding-block: $spacing_20;
        padding-inline: $spacing_13;
        width: 100%;
        max-width: $header_content_width;
        background: $neutral_2;
        flex: 1;
    }
}
</style>
