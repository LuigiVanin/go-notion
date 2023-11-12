<script lang="ts" setup>
// Libraries
import InlineSvg from "vue-inline-svg";

// Components
import Label from "./Label.vue";

type TextInput = {
    modelValue: string;
    placeholder?: string;
    type?: string;
    name?: string;
    rightIcon?: string;
    icon?: string;
    maxWidth?: string;
    label?: string;
};

const props = defineProps<TextInput>();
const emit = defineEmits(["update:modelValue"]);

const handleInputEvent = (event: any) => {
    emit("update:modelValue", event.target.value);
};
</script>

<template>
    <div class="input-main__wrapper">
        <Label v-if="props.label" :text="props.label" :for="props.name" />
        <div class="input-main">
            <slot name="prefix">
                <InlineSvg :src="props.icon" class="input-main__icon" />
            </slot>

            <input
                :value="props.modelValue"
                :placeholder="props.placeholder"
                :name="props.name"
                :type="props.type || 'text'"
                :style="{ maxWidth: maxWidth }"
                @input="handleInputEvent"
            />
            <slot name="suffix">
                <InlineSvg :src="props.rightIcon" class="input-main__icon" />
            </slot>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.input-main__wrapper {
    width: 100%;
    @include flex(column, start, start);
}

.input-main {
    @include flex-center;
    width: 100%;
    height: 38px;
    border: $spacing_0 solid $neutral_5;
    padding: $spacing_6 $spacing_6;
    border-radius: $border_r_md;
    gap: $spacing_4;
    transition: $transition;
    font-family: "Poppins", sans-serif;

    &:hover {
        border-color: $neutral_7;
    }

    &:focus-within {
        box-shadow: 0px 0px 4px -1px $primary_5;
        border-color: $primary_3;

        :deep(svg.input-main__icon) {
            path,
            rect {
                stroke: $primary_3;
            }
        }
    }

    input {
        width: 100%;
        font-size: $font_4;
        color: $neutral_9;
        padding-bottom: $spacing_1 !important;
        @include reset-input();
    }
    :deep(svg.input-main__icon) {
        width: 20px;
        height: 20px;

        path,
        rect {
            stroke: $neutral_7;
        }
    }
}
</style>
