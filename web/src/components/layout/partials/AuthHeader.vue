<script lang="ts" setup>
// Components
import UserHeaderItem from "./UserHeaderItem.vue";
import ThemeSwitch from "@/components/ThemeSwitch.vue";
import Logo from "@/components/Logo.vue";
</script>

<template>
    <header class="app-header">
        <div class="app-header__content">
            <Logo />
            <div class="app-header__content__actions">
                <ThemeSwitch />
                <UserHeaderItem />
            </div>
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

    height: $header_height;
    border-bottom: $spacing_0 solid $neutral_4;
    backdrop-filter: blur(20px);

    .app-header__content {
        width: 100%;
        max-width: $header_auth_content_width;
        @include flex(row, center, space-between);
        padding-inline: $spacing_13;

        .app-header__content__actions {
            height: 100%;
            @include flex(row, center, center);
            gap: $spacing_8;
        }
    }
}
</style>
