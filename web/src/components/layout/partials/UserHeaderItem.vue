<script lang="ts" setup>
// Core
import { ref } from "vue";

// Libraries
import { onClickOutside } from "@vueuse/core";

// Components
import Button from "@/components/core/Button.vue";
import UserMenu from "@/components/layout/partials/UserMenu.vue";

// Assets
import userIconUrl from "@/assets/icons/user.svg?url";

const showUserMenu = ref(false);
const userMenuEl = ref<HTMLElement | null>(null);

onClickOutside(userMenuEl, () => {
    showUserMenu.value = false;
});
</script>

<template>
    <div
        ref="userMenuEl"
        :class="['header-item', showUserMenu && 'header-item--active']"
    >
        <Button
            :icon="userIconUrl"
            btn-type="no-border"
            size="lg"
            @click="showUserMenu = !showUserMenu"
        />
        <UserMenu :show="showUserMenu" @close="showUserMenu = false" />
    </div>
</template>

<style lang="scss" scoped>
.header-item {
    position: relative;

    &.header-item--active {
        :deep(> button) {
            color: $primary_6;
            background: $neutral_3;
        }
    }

    :deep(> button) {
        padding-block: $spacing_3;
        padding-inline: $spacing_3;
        color: $neutral_7;

        &:hover {
            color: $primary_6;
        }

        svg {
            width: 22px !important;
            height: 22px !important;
        }
    }

    :deep(.user-menu) {
        .fade-down-enter-active,
        .fade-down-leave-active {
            height: $text_input_warn_height;
            opacity: 1;
            transform: translateY(0);
        }

        .fade-down-enter-from,
        .fade-down-leave-from {
            opacity: 0;
            height: 0px;
            transform: translateY(-5px);
        }
    }
}
</style>
