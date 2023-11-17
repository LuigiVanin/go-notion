<script lang="ts" setup>
// Libraries
import InlineSvg from "vue-inline-svg";

// Assets
import userIconUrl from "@/assets/icons/user.svg?url";
</script>

<template>
    <div class="user-avatar">
        <InlineSvg :src="userIconUrl" />
    </div>
</template>

<style lang="scss" scoped>
.user-avatar {
    :deep(svg) {
        width: 32px;
        height: 32px;

        path,
        circle,
        rect {
            stroke: $primary_7;
        }
    }
}
</style>
