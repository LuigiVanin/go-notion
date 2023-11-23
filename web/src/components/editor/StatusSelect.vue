<script lang="ts" setup>
import Popper from "vue3-popper";
</script>

<template>
    <div class="input-main">
        <Popper content="Teste">
            <p>Rascunho</p>
        </Popper>
    </div>
</template>

<style lang="scss" scoped>
.input-main {
    :deep(.inline-block) {
        width: 100%;
        margin: 0px !important;
        border: none !important;
        min-width: 120px;

        @include flex-center(row);
    }
}
</style>
