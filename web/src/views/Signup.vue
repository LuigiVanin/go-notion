<script lang="ts" setup>
// Core
import { computed, reactive, ref } from "vue";
import { RouterLink, useRouter } from "vue-router";

// Libraries
import { useToast } from "vue-toastification";
import { z } from "zod";

// Composables
import { useSignup } from "@/composables/api/useSignup.ts";
import { useValidation } from "@/composables/useValidation.ts";

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
import keypadIconUrl from "@/assets/icons/keypad-outline.svg?url";
import userIconUrl from "@/assets/icons/user.svg?url";

const signupForm = reactive({
    name: "",
    email: "",
    password: "",
    passwordConfirm: "",
});
const phone = ref("");

const validation = z.object({
    name: z
        .string()
        .min(5, { message: "O nome deve ter no mínimo 3 caracteres" })
        .refine((data) => data.split(" ").length >= 2, {
            message: "O nome deve ter nome e sobrenome",
        }),
    email: z
        .string()
        .min(1, { message: "O email é obrigatório" })
        .email({ message: "O email deve ser válido" }),
    password: z
        .string()
        .min(6, { message: "A senha deve ter no mínimo 6 caracteres" }),
    passwordConfirm: z
        .string()
        .min(6, { message: "A senha deve ter no mínimo 6 caracteres" })
        .refine((data) => data === signupForm.password, {
            message: "As senhas devem ser iguais",
        }),
});

const { fieldErrors, validate } = useValidation(signupForm, validation);
const { signup, signupError, signupLoading } = useSignup();
const toast = useToast();
const router = useRouter();

const handleSignup = async () => {
    const isValid = validate();
    if (!isValid) return;

    await signup(signupForm);

    if (signupError.value) {
        return;
    }
    toast.success("Conta criada com sucesso!");
    router.push("/");
};

const errorMessage = computed(() => {
    const data = signupError.value?.response?.data;
    if (!data) return "Erro na requisição";

    if (data.fields?.length) {
        return "Um campo do formulário está incorreto!";
    }
    if ([400].includes(data.code)) {
        return "Email utilizado já existe na plataforma!";
    }

    if (data.message) return data.message;
    return "Erro na requisição";
});
</script>

<template>
    <main class="page-main">
        <PageTitle
            title="Crie a sua conta!"
            text="Crie a sua conta usando seus dados pessoais. Caso já tenha conta, vá para a tela de login"
        />
        <Card :loading="signupLoading">
            <form class="page-main__form" @submit.prevent="handleSignup">
                <TextInput
                    v-model="signupForm.name"
                    :icon="userIconUrl"
                    :error-message="fieldErrors.name"
                    placeholder="Insira seu nome..."
                    label="Nome"
                    @blur="validate('name')"
                />
                <TextInput
                    v-model="signupForm.email"
                    :icon="emaiIconUrl"
                    :error-message="fieldErrors.email"
                    placeholder="Insira seu email..."
                    label="Email"
                    @blur="validate('email')"
                />
                <TextInput
                    v-model="phone"
                    :icon="keypadIconUrl"
                    mask="+## (##) #####-####"
                    placeholder="+55 (11) 99999-9999"
                    label="Telefone"
                    optional
                />
                <PasswordInput
                    v-model="signupForm.password"
                    :icon="keyIconUrl"
                    :error-message="fieldErrors.password"
                    label="Senha"
                    placeholder="Insira a sua senha..."
                    @blur="validate('password')"
                />
                <PasswordInput
                    v-model="signupForm.passwordConfirm"
                    :icon="keyIconUrl"
                    :error-message="fieldErrors.passwordConfirm"
                    label="Confirmar Senha"
                    placeholder="Repita a sua senha..."
                    @blur="validate('passwordConfirm')"
                />
                <AlertBox :show="!!signupError" :text="errorMessage ?? ''" />
                <footer>
                    <SpecialButton
                        text="Criar conta"
                        type="submit"
                        :disabled="signupLoading"
                    />
                </footer>
            </form>
        </Card>
        <p>
            Crie a sua conta! caso já tenha conta,
            <RouterLink to="/"> VÁ PARA A TELA DE LOGIN! </RouterLink>
        </p>
    </main>
</template>

<style lang="scss" scoped>
main.page-main {
    width: 100%;
    max-width: $sign_main_width;
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

                &:not(:disabled):hover {
                    width: 100%;
                }

                &:disabled {
                    cursor: progress;
                }
            }
        }
    }
}
</style>
