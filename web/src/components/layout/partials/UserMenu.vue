<script lang="ts" setup>
// Core
import { Transition } from "vue";
import { useRouter } from "vue-router";

// Stores
import { useUserStore } from "@/store/user.ts";

// Components
import UserAvatar from "./UserAvatar.vue";
import UserMenuActions from "./UserMenuActions.vue";
import Button from "@/components/core/Button.vue";

// Assets
import githubIconUrl from "@/assets/icons/logo-github.svg?url";
import paperIconUrl from "@/assets/icons/paper.svg?url";
import userIconUrl from "@/assets/icons/user.svg?url";
import logoutIconUrl from "@/assets/icons/log-out-outline.svg?url";

// Types
import { Action } from "@/types/props";

type UserMenuProps = {
    show: boolean;
};

const props = defineProps<UserMenuProps>();
const emit = defineEmits(["close"]);

const router = useRouter();

const actionList: Action[] = [
    {
        icon: paperIconUrl,
        label: "Home",
        onClick: () => {
            console.log("Logout");
            router.push("/home");
        },
    },
    {
        icon: githubIconUrl,
        label: "Repositório",
        class: "user-menu__icon--github",
        onClick: () => {
            console.log("Logout");
        },
    },
    {
        icon: userIconUrl,
        label: "Usuário",
        disabled: true,
        onClick: () => {
            console.log("Logout");
        },
    },
];

const userStore = useUserStore();

const logout = () => {
    router.push("/");
    userStore.clearUser();
};
</script>

<template>
    <Transition name="fade-down">
        <div v-if="props.show" class="user-menu">
            <div class="user-menu__auth">
                <UserAvatar />

                <div class="user-menu__auth__content">
                    <h3>{{ userStore.user?.name }}</h3>
                    <p>{{ userStore.user?.email }}</p>
                </div>
            </div>

            <span class="user-menu__divider" />

            <UserMenuActions :actions="actionList" />

            <span class="user-menu__divider" />

            <Button
                :icon="logoutIconUrl"
                text="Sair"
                btn-type="no-border"
                @click="logout"
            />
        </div>
    </Transition>
</template>

<style lang="scss" scoped>
.user-menu {
    position: absolute;
    right: 0;
    width: 250px;
    padding: $spacing_10;
    border-radius: $border_r_lg;
    background: $neutral_2;
    border: 1px solid $neutral_3;
    box-shadow: 0px 1px 12px 0px rgba(145, 145, 145, 0.2);
    transition: all 0.2s ease-in-out;
    margin-top: $spacing_3;

    :deep(svg.user-menu__icon--github) {
        path {
            fill: $neutral_9;
        }
    }

    .user-menu__auth {
        width: 100%;
        @include flex(row, center, start);
        gap: $spacing_6;

        .user-menu__auth__content {
            @include flex(column, start, start);
            gap: $spacing_2;

            h3 {
                font-size: $font_4;
                font-weight: 500;
                color: $neutral_11;
            }

            p {
                font-size: $font_2;
                color: $neutral_8;
            }
        }
    }

    span.user-menu__divider {
        @include flex(row, center, start);
        width: 100%;
        height: $spacing_0;
        background: $neutral_4;
        margin-block: $spacing_10;
    }
}

.fade-down-enter-active,
.fade-down-leave-active {
    transition: all 0.2s ease-in-out;
}

.fade-down-enter-from,
.fade-down-leave-to {
    opacity: 0;
    transform: translateY(10%);
}
</style>
