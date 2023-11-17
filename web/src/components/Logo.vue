<template>
    <span class="app-logo">Writabble</span>
</template>

<style lang="scss" scoped>
.app-logo {
    font-weight: 700;
    color: $primary_4;
    @include gradient-text();
    @include not-selectable;
}
</style>
