<script lang="ts" setup>
// Core
import { ref } from "vue";

// Libraries
import InlineSvg from "vue-inline-svg";

// Components
import TextInput from "@/components/core/TextInput.vue";

// Types
import type { InputProps } from "@/types/props";

// Assets
import keyIconUrl from "@/assets/icons/key.svg?url";
import lockIconUrl from "@/assets/icons/lock.svg?url";
import openLockIconUrl from "@/assets/icons/open-lock.svg?url";

type PasswordInputProps = {} & InputProps;

const props = defineProps<PasswordInputProps>();
const emit = defineEmits(["update:modelValue"]);

const passwordVisibility = ref(false);

const togglePasswordVisibility = () => {
    passwordVisibility.value = !passwordVisibility.value;
};
</script>

<template>
    <TextInput
        :type="!passwordVisibility ? 'password' : 'text'"
        :model-value="props.modelValue"
        :label="props.label"
        :max-width="props.maxWidth"
        :name="props.name"
        :placeholder="props.placeholder"
        :icon="keyIconUrl"
        @update:model-value="emit('update:modelValue', $event)"
    >
        <template #suffix>
            <button
                class="password-visibility"
                @click="togglePasswordVisibility"
            >
                <InlineSvg
                    v-if="!passwordVisibility"
                    :src="lockIconUrl"
                    class="password-visibility__icon password-visibility__icon--hide"
                />
                <InlineSvg
                    v-show="passwordVisibility"
                    :src="openLockIconUrl"
                    class="password-visibility__icon password-visibility__icon--show"
                />
            </button>
        </template>
    </TextInput>
</template>

<style lang="scss" scoped>
button.password-visibility {
    @include reset-button();
    @include flex-center;
    border-radius: $border_r_md;
    min-width: 25px;
    height: 25px;
    transition: all 0.2s ease-in-out;

    &:hover,
    &:focus {
        background-color: $neutral_4;
    }

    :deep(svg.password-visibility__icon) {
        width: 17px;
        height: 17px;

        path {
            transition: all 0.2s ease-in-out;
        }

        &.password-visibility__icon--hide {
            path,
            rect {
                stroke: $neutral_7;
            }
        }

        &.password-visibility__icon--show {
            path,
            rect {
                stroke: $green_6;
            }
        }
    }
}
</style>
