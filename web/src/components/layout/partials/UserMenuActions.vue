<script lang="ts" setup>
// Libraries
import InlineSvg from "vue-inline-svg";

// Types
import { Action } from "@/types/props";

type UserMenuActionsProps = {
    actions: Action[];
};

const props = defineProps<UserMenuActionsProps>();
</script>

<template>
    <ul>
        <li
            v-for="act in props.actions"
            :key="act.label"
            :disabled="act.disabled"
            :class="[act.disabled && 'user-menu--disabled']"
            @click="!act.disabled && act.onClick()"
        >
            <InlineSvg
                :src="act.icon"
                :class="['user-menu__icon', act.class]"
            />
            <span>
                {{ act.label }}
            </span>
        </li>
    </ul>
</template>

<style lang="scss" scoped>
ul {
    @include reset-list;
    @include flex(column, start, start);

    li {
        @include flex(row, center, start);
        width: 100%;
        gap: $spacing_6;
        padding-block: $spacing_5;
        padding-inline: $spacing_3;
        border-radius: $border_r_sm;
        cursor: pointer;

        &:disabled,
        &.user-menu--disabled {
            opacity: 0.5;
            cursor: not-allowed;
        }

        &:hover {
            background: $neutral_3;
        }

        :deep(svg) {
            width: 17px;
            height: 17px;

            path,
            rect,
            circle {
                stroke: $neutral_9;
            }
        }

        span {
            color: $neutral_9;
            font-size: $font_4;
        }
    }
}
</style>
