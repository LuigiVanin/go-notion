<script lang="ts" setup>
// Components
import Button from "./core/Button.vue";

// Assets
import circlePlusIconUrl from "@/assets/icons/plus-circle.svg?url";

type CreateDocumentProps = {
    action: () => void;
    loading?: boolean;
};

const props = defineProps<CreateDocumentProps>();
</script>

<template>
    <div class="create-document">
        <Button
            :icon="circlePlusIconUrl"
            :disabled="props.loading"
            text="Criar novo Documento"
            btn-type="outlined"
            size="lg"
            @click="props.action"
        />
        <span v-if="loading" class="create-document__loading" />
    </div>
</template>

<style lang="scss" scoped>
.create-document {
    border-radius: $border_r_lg;
    position: relative;

    &:hover {
        // box-shadow: 0px 0px 10px -4px $primary_5;
    }

    :deep(button.button-main.button-main--size-lg) {
        padding-inline: $spacing_13;
        padding-block: $spacing_10;
        font-size: $font_7;
        border-radius: $border_r_lg;
        font-weight: 500;

        svg.button-main__icon {
            width: 25px;
            height: 25px;
        }

        &:hover:not(:disabled) {
            background: $primary_1;
            color: $primary_6;
            box-shadow: 0px 0px 10px -4px $primary_5;
        }
    }

    span.create-document__loading {
        position: absolute;
        background: $primary_4;
        bottom: 0px;
        left: 0px;
        width: 100px;
        height: 5px;
        z-index: 10;
        border-radius: 3px;

        animation: bar-loading 2.5s infinite $transition_spring;
    }
}
</style>
