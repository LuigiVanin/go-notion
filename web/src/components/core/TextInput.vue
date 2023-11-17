<script lang="ts" setup>
// Libraries
import InlineSvg from "vue-inline-svg";
import { vMaska } from "maska";

// Components
import Label from "./Label.vue";
import { InputProps } from "@/types/props";
import { computed } from "vue";

type TextInputProps = InputProps & {};

const props = defineProps<TextInputProps>();
const emit = defineEmits(["update:modelValue", "blur"]);

const handleInputEvent = (event: any) => {
    emit("update:modelValue", event.target.value);
};

const hasError = computed(() => {
    return props.errorMessage !== undefined && props.errorMessage !== null;
});
</script>

<template>
    <div class="input-main__wrapper">
        <Label
            v-if="props.label"
            :text="props.label"
            :for="props.name"
            :optional="props.optional"
        />
        <div :class="['input-main', hasError && 'input-main--error']">
            <slot name="prefix">
                <InlineSvg :src="props.icon" class="input-main__icon" />
            </slot>

            <input
                v-maska
                :data-maska="props.mask"
                :value="props.modelValue"
                :placeholder="props.placeholder"
                :name="props.name"
                :type="props.type || 'text'"
                :style="{ maxWidth: maxWidth }"
                @input="handleInputEvent"
                @blur="emit('blur')"
            />
            <slot name="suffix">
                <InlineSvg :src="props.rightIcon" class="input-main__icon" />
            </slot>
        </div>
        <Transition name="error-fade">
            <p v-if="props.errorMessage" class="input-main__warning">
                {{ props.errorMessage }}
            </p>
        </Transition>
    </div>
</template>

<style lang="scss" scoped>
.input-main__wrapper {
    width: 100%;
    @include flex(column, start, start);

    p.input-main__warning {
        @include line-clamp();
        width: 100%;
        color: $red_4;
        font-size: $font_1;
        padding-top: $spacing_4;
        padding-left: $spacing_2;
        height: 17px;
        transition: all 0.2s ease-in-out;
    }
}

.input-main {
    @include flex-center;
    width: 100%;
    height: $text_input_height;
    border: $spacing_0 solid $neutral_6;
    padding: $spacing_6 $spacing_6;
    border-radius: $border_r_md;
    gap: $spacing_4;
    transition: $transition;
    font-family: "Poppins", sans-serif;
    background: $neutral_2;

    &:hover {
        border-color: $neutral_7;
    }

    &:focus-within {
        box-shadow: 0px 0px 4px -1px $primary_5;
        border-color: $primary_4;

        :deep(svg.input-main__icon) {
            path,
            rect,
            circle {
                stroke: $primary_4;
            }
        }
    }

    :deep(svg.input-main__icon) {
        width: $text_input_icon_size;
        height: $text_input_icon_size;

        path,
        rect,
        circle {
            stroke: $neutral_7;
        }
    }

    &--error {
        border-color: $red_4 !important;

        :deep(svg.input-main__icon) {
            path,
            rect,
            circle {
                stroke: $red_4 !important;
            }
        }

        &:hover {
            border-color: $red_5;
        }
    }

    input {
        width: 100%;
        font-size: $font_4;
        color: $neutral_9;
        padding-bottom: $spacing_0 !important;
        @include reset-input();
    }
}

.error-fade-enter-active,
.error-fade-leave-active {
    height: $text_input_warn_height;
    opacity: 1;
    transform: translateY(0);
}

.error-fade-enter-from,
.error-fade-leave-from {
    opacity: 0;
    height: 0px;
    transform: translateY(-5px);
}
</style>
