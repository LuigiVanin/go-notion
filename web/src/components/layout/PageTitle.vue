<script lang="ts" setup>
type PageTitleProps = {
    title: string;
    text?: string;
};

const props = defineProps<PageTitleProps>();
</script>

<template>
    <div class="page-title">
        <h1>{{ props.title }}</h1>
        <p v-if="props.text">{{ props.text }}</p>
    </div>
</template>
