<script lang="ts" setup>
type LabelProps = {
    text: string;
    for?: string;
    optional?: boolean;
};

const props = defineProps<LabelProps>();
</script>

<template>
    <label :for="props.for">
        <span>
            {{ props.text }}
        </span>
        <span v-if="props.optional" class="label__optional">(optional)</span>
    </label>
</template>

<style lang="scss" scoped>
label {
    font-size: $font_4;
    color: $neutral_10;
    margin-bottom: $spacing_4;
    width: 100%;
    display: block;
    font-weight: 500;
    padding-left: $spacing_2;

    span.label__optional {
        color: $neutral_8;
        font-size: $font_1;
        font-weight: 500;
        padding-left: $spacing_1;
    }
}
</style>
