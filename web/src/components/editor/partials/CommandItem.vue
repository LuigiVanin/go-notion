<script lang="ts" setup>
// Libraries
import InlineSvg from "vue-inline-svg";
import { watch, ref } from "vue";

type CommandItemProps = {
    label: string;
    iconUrl: string;
    selected: boolean;
};

const props = defineProps<CommandItemProps>();
const emit = defineEmits(["click"]);

const itemEl = ref<HTMLElement | null>(null);

const handleActionClick = () => {
    emit("click");
};

watch(
    () => props.selected,
    (selected) => {
        if (selected) {
            itemEl.value?.scrollIntoView({
                behavior: "smooth",
                block: "nearest",
                inline: "start",
            });
        }
    }
);
</script>

<template>
    <li
        ref="itemEl"
        :class="['action', props.selected && 'action--selected']"
        @click="handleActionClick"
    >
        <div class="action__icon">
            <InlineSvg :src="props.iconUrl" width="24" />
        </div>
        <span class="action__label">
            {{ props.label }}
        </span>
    </li>
</template>

<style lang="scss" scoped>
.action {
    height: 45px;
    @include flex(row, center, start);
    gap: $spacing_8;
    border-radius: $border_r_md;
    padding-inline: $spacing_3;
    transition: all 0.1s ease-in-out;
    width: 100%;

    &:not(:last-child) {
        margin-bottom: $spacing_2;
    }

    &.action--selected {
        background: $neutral_3;

        .action__icon {
            background: $neutral_2;

            :deep(svg path) {
                fill: $primary_5;
            }
        }

        span.action__label {
            color: $neutral_11;
        }
    }

    .action__icon {
        height: 35px;
        min-width: 35px;
        background: $neutral_4;
        border-radius: $border_r_md;
        @include flex-center;
        transition: all 0.1s ease-in-out;

        :deep(svg path) {
            fill: $neutral_11;
        }
    }

    span.action__label {
        font-weight: 400;
        font-size: $font_5;
        color: $neutral_9;
    }
}
</style>
