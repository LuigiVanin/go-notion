<script lang="ts" setup>
// Core
import { RouterLink } from "vue-router";

// Components
import Button from "../core/Button.vue";

// Assets
import signIconUrl from "@/assets/icons/sign.svg?url";
</script>

<template>
    <header class="app-header">
        <div class="app-heaeder__content">
            <span class="app-heaeder__content__login">Writabble</span>

            <RouterLink to="/signup">
                <Button
                    :icon="signIconUrl"
                    text="Sign up"
                    btn-type="outlined"
                    size="md"
                />
            </RouterLink>
        </div>
    </header>
</template>

<style lang="scss" scoped>
header.app-header {
    @include flex(row, center, center);
    z-index: 5;
    width: 100%;
    position: fixed;
    top: 0;

    height: 50px;

    .app-heaeder__content {
        width: 100%;
        max-width: 868px;
        @include flex(row, center, space-between);
        padding-inline: $spacing_13;

        &__login {
            font-weight: 700;
            color: $primary_4;
            @include gradient-text();
        }
    }
}
</style>
