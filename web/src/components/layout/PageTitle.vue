<script lang="ts" setup>
type PageTitleProps = {
    title: string;
    text?: string;
};

const props = defineProps<PageTitleProps>();
</script>

<template>
    <div class="page-title">
        <h2>{{ props.title }}</h2>
        <p v-if="props.text">{{ props.text }}</p>
    </div>
</template>

<style lang="scss">
.page-title {
    @include flex-center(column);
    gap: $spacing_13;
    margin-bottom: $spacing_10;

    h2 {
        font-size: $font_14;
        font-weight: 500;
        text-align: center;
        color: $primary_4;
        line-height: 2.5rem;

        @include gradient-text();
    }

    p {
        color: $neutral_8;
        font-size: $font_3;
        line-height: 1.25rem;
        text-align: center;
        padding-inline: $spacing_3;
    }
}
</style>
