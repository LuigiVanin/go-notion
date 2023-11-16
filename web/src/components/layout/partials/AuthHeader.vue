<script lang="ts" setup>
// Core

// Components
import Button from "@/components/core/Button.vue";

// Assets
import userIconUrl from "@/assets/icons/user.svg?url";
</script>

<template>
    <header class="app-header">
        <div class="app-header__content">
            <span class="app-header__content__login">Writabble</span>

            <Button :icon="userIconUrl" btn-type="no-border" size="lg" />
        </div>
    </header>
</template>

<style lang="scss" scoped>
header.app-header {
    @include flex(row, center, center);
    z-index: 5;
    width: 100%;
    top: 0;
    position: fixed;

    height: 50px;
    border-bottom: $spacing_1 solid $neutral_4;
    backdrop-filter: blur(20px);

    .app-header__content {
        width: 100%;
        max-width: 1080px;
        @include flex(row, center, space-between);
        padding-inline: $spacing_13;

        &__login {
            font-weight: 700;
            color: $primary_4;
            @include gradient-text();
        }

        :deep(button) {
            padding-block: $spacing_4;
            padding-inline: $spacing_4;
        }
    }
}
</style>
