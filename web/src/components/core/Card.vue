<script lang="ts" setup></script>

<template>
    <div class="card">
        <slot />
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
}
</style>
