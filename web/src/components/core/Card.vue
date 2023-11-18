<script lang="ts" setup>
type CardProps = {
    loading?: boolean;
};

const props = defineProps<CardProps>();
</script>

<template>
    <div class="card">
        <slot />
        <span v-if="props.loading" class="card__loading" />
    </div>
</template>

<style lang="scss" scoped>
.card {
    width: 100%;
    border: 1px solid $neutral_5;
    box-shadow: 0px 1px 12px 0px rgba(145, 145, 145, 0.2);
    padding: $spacing_16;
    border-radius: $border_r_lg;
    @include flex(column, start, start);
    background: $neutral_1;
    position: relative;
    overflow: hidden;

    span.card__loading {
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
