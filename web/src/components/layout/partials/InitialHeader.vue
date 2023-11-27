<script lang="ts" setup>
// Core
import { RouterLink, useRoute } from "vue-router";
import { computed } from "vue";

// Components
import Button from "@/components/core/Button.vue";
import Logo from "@/components/Logo.vue";
import ThemeSwitch from "@/components/ThemeSwitch.vue";

// Assets
import signIconUrl from "@/assets/icons/sign.svg?url";
import signinIconUrl from "@/assets/icons/signin.svg?url";

type HeaderContent = {
    buttonText: string;
    link: string;
    icon: string;
};

const route = useRoute();

const headerContent: Record<string, HeaderContent> = {
    "/signup": {
        buttonText: "Login",
        link: "/",
        icon: signinIconUrl,
    },
    "/": {
        buttonText: "Sign up",
        link: "/signup",
        icon: signIconUrl,
    },
};

const content = computed(() => {
    if (!route.path) {
        return {
            buttonText: "Sign up",
            link: "/signup",
            icon: signIconUrl,
        };
    }
    return headerContent[route?.path as string];
});
</script>

<template>
    <header class="app-header">
        <div class="app-header__content">
            <Logo />

            <div class="app-header__content__right">
                <ThemeSwitch />
                <RouterLink :to="content?.link || ''">
                    <Button
                        :icon="content?.icon"
                        :text="content?.buttonText"
                        btn-type="outlined"
                        size="md"
                    />
                </RouterLink>
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

    .app-header__content {
        width: 100%;
        max-width: $header_content_width;
        @include flex(row, center, space-between);
        padding-inline: $spacing_13;

        .app-header__content__right {
            @include flex(row, center, center);
            gap: $spacing_16;
        }
    }
}
</style>
