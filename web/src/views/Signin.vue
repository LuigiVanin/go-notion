<script lang="ts" setup>
// Core
import { computed, reactive } from "vue";
import { useRouter } from "vue-router";

// Libraries
import { RouterLink } from "vue-router";
import { z } from "zod";

// Composables
import { useValidation } from "@/composables/useValidation.ts";
import { useSignin } from "@/composables/api/useSignin.ts";

// Components
import Card from "@/components/core/Card.vue";
import TextInput from "@/components/core/TextInput.vue";
import PasswordInput from "@/components/PasswordInput.vue";
import SpecialButton from "@/components/SpecialButton.vue";
import PageTitle from "@/components/layout/partials/PageTitle.vue";
import AlertBox from "@/components/core/AlertBox.vue";

// Assets
import emaiIconUrl from "@/assets/icons/email.svg?url";
import keyIconUrl from "@/assets/icons/key.svg?url";

const loginFormRule = z.object({
    email: z.string().email("Email inválido!"),
    password: z.string().min(6, "Senha deve ter no mínimo 6 caracteres!"),
});

const loginForm = reactive({
    email: "luisfvanin9com",
    password: "senha123",
});

const router = useRouter();
const { fieldErrors, validate } = useValidation(loginForm, loginFormRule);
const { loginLoading, loginError, signin } = useSignin();

const handleSignin = async () => {
    const isValidForm = validate();

    if (!isValidForm) return;
    await signin(loginForm);
    if (loginError.value) return;

    router.push("/home");
};

const errorMessage = computed(() => {
    const data = loginError.value?.response?.data;
    if (!data) return "Erro na requisição";

    if (data.fields?.length) {
        return "Um campo do formulário está incorreto!";
    }
    if ([404, 400].includes(data.code)) {
        return "Email inexistente ou senha incorretos!";
    }

    if (data.message) return data.message;
    return "Erro na requisição";
});
</script>

<template>
    <main class="page-main">
        <PageTitle
            title="Faça login na sua conta!"
            text="entre na sua conta usando seu email e senha. Caso não tenha conta, vá para a tela de signup"
        />
        <Card :loading="loginLoading">
            <form class="page-main__form" @submit.prevent="handleSignin">
                <TextInput
                    v-model="loginForm.email"
                    :icon="emaiIconUrl"
                    :error-message="fieldErrors.email"
                    placeholder="Insira seu email..."
                    label="Email"
                    @blur="validate('email')"
                />
                <PasswordInput
                    v-model="loginForm.password"
                    :icon="keyIconUrl"
                    :error-message="fieldErrors.password"
                    label="Senha"
                    placeholder="Insira a sua senha..."
                    @blur="validate('password')"
                />
                <AlertBox :show="!!loginError" :text="errorMessage ?? ''" />
                <footer>
                    <SpecialButton :disabled="loginLoading" text="Log In!" />
                </footer>
            </form>
        </Card>
        <p>
            Faça Login! caso não tenha conta,
            <RouterLink to="signup"> CRIE UMA CONTA AGORA! </RouterLink>
        </p>
    </main>
</template>

<style lang="scss" scoped>
main.page-main {
    width: 100%;
    max-width: 400px;
    z-index: 10;

    @include flex-center(column);
    gap: $spacing_24;

    p {
        color: $neutral_8;
        font-size: $font_3;
        text-align: center;
        padding-inline: $spacing_32;
        line-height: 1.25rem;

        :deep(a) {
            color: $primary_4;
            font-weight: 500;
            transition: all 0.2s;

            &:hover {
                color: $primary_5;
            }
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
