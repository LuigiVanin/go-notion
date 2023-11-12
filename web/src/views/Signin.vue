<script lang="ts" setup>
// Core
import { reactive } from "vue";

// Libraries
import { RouterLink } from "vue-router";

// Components
import Card from "@/components/core/Card.vue";
import TextInput from "@/components/core/TextInput.vue";
import PasswordInput from "@/components/PasswordInput.vue";
import SpecialButton from "@/components/SpecialButton.vue";
import PageTitle from "@/components/layout/PageTitle.vue";

// Assets
import emaiIconUrl from "@/assets/icons/email.svg?url";
import keyIconUrl from "@/assets/icons/key.svg?url";

const loginForm = reactive({
    email: "",
    password: "",
});
</script>

<template>
    <main class="page-main">
        <PageTitle
            title="Faça login na sua conta!"
            text="entre na sua conta usando seu email e senha. Caso não tenha conta, vá para a tela de signup"
        />
        <Card>
            <form class="page-main__form" @submit.prevent="">
                <TextInput
                    v-model="loginForm.email"
                    :icon="emaiIconUrl"
                    placeholder="Insira seu email..."
                    label="Email"
                />
                <PasswordInput
                    v-model="loginForm.password"
                    :icon="keyIconUrl"
                    label="Senha"
                    placeholder="Insira a sua senha..."
                />
                <footer>
                    <SpecialButton text="Log In!" />
                </footer>
            </form>
        </Card>
        <p>
            Faça Login! caso não tenha conta,
            <RouterLink to="/signup"> CRIE UMA CONTA AGORA! </RouterLink>
        </p>
    </main>
</template>

<style lang="scss" scoped>
main.page-main {
    width: 100%;
    max-width: 400px;

    @include flex-center(column);
    gap: $spacing_24;

    p {
        color: $neutral_9;
        font-size: $font_3;
        text-align: center;
        padding-inline: $spacing_32;
        line-height: 1.25rem;

        :deep(a) {
            color: $primary_4;
            font-weight: 500;
        }
    }

    .page-main__form {
        @include flex-center(column);
        width: 100%;
        padding-block: $spacing_8;
        gap: $spacing_13;

        footer {
            width: 100%;
            padding-top: $spacing_10;
            @include flex-center;
            :deep(button) {
                width: 75%;
                transition: all 0.2s $transition_spring;

                &:hover {
                    width: 100%;
                }
            }
        }
    }
}
</style>
