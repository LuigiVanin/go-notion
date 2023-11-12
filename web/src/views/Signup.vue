<script lang="ts" setup>
// Core
import { computed, reactive, ref } from "vue";
import { RouterLink } from "vue-router";

// Helpers & Services
import { Api, type ApiError } from "@/api/api.ts";

// Components
import Card from "@/components/core/Card.vue";
import TextInput from "@/components/core/TextInput.vue";
import PasswordInput from "@/components/PasswordInput.vue";
import SpecialButton from "@/components/SpecialButton.vue";
import PageTitle from "@/components/layout/PageTitle.vue";
import AlertBox from "@/components/core/AlertBox.vue";

// Assets
import emaiIconUrl from "@/assets/icons/email.svg?url";
import keyIconUrl from "@/assets/icons/key.svg?url";
import { z } from "zod";

const api = new Api();

const validation = z
    .object({
        name: z
            .string()
            .min(3, { message: "O nome deve ter no mínimo 3 caracteres" }),
        email: z
            .string()
            .min(1, { message: "O email é obrigatório" })
            .email({ message: "O email deve ser válido" }),
        password: z
            .string()
            .min(6, { message: "A senha deve ter no mínimo 6 caracteres" }),
        confirmPassword: z
            .string()
            .min(6, { message: "A senha deve ter no mínimo 6 caracteres" }),
    })
    .refine((data) => data.password === data.confirmPassword, {
        message: "As senhas devem ser iguais",
        path: ["confirmPassword"],
    });

const loginForm = reactive({
    name: "Luis",
    email: "luisfvanin9com",
    password: "senha123",
    confirmPassword: "senha123",
});

const loginError = ref<null | ApiError>(null);

const signup = async () => {
    try {
        validation.parse(loginForm);
    } catch (err) {
        console.log(err);
    }
    loginError.value = null;
    const { error } = await api.auth.signin(loginForm);
    if (!error) {
        return;
    }

    loginError.value = error;
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
        <Card>
            <form class="page-main__form" @submit.prevent="signup">
                <TextInput
                    v-model="loginForm.name"
                    :icon="emaiIconUrl"
                    placeholder="Insira seu nome..."
                    label="Nome"
                />
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
                <PasswordInput
                    v-model="loginForm.confirmPassword"
                    :icon="keyIconUrl"
                    label="Confirmar Senha"
                    placeholder="Repita a sua senha..."
                />
                <AlertBox :show="!!loginError" :text="errorMessage ?? ''" />
                <footer>
                    <SpecialButton text="Criar conta" />
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
