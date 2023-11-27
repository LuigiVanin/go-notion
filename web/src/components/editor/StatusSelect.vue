<script lang="ts" setup>
// Core
import { computed, ref } from "vue";

// Libraries
import InlineSvg from "vue-inline-svg";

// Components
import Popper from "vue3-popper";

// Assets
import chevronDownIconUrl from "@/assets/icons/chevron-down.svg?url";

// Types
import { DocumentStatus } from "@/types/document";
import { onKeyStroke } from "@vueuse/core";

type StatusSelectProps = {
    modelValue: DocumentStatus;
};

const props = defineProps<StatusSelectProps>();
const emit = defineEmits(["update:modelValue", "select"]);

const openDropdown = ref(false);
const optionEl = ref<HTMLElement[] | null>(null);
const optionSelectedIndex = ref(0);

const moveSelectedIndex = (steps: number) => {
    optionSelectedIndex.value =
        (optionSelectedIndex.value + steps) % optionEl.value!.length;

    if (optionSelectedIndex.value < 0) {
        optionSelectedIndex.value = optionEl.value!.length - 1;
    }
};

onKeyStroke(["ArrowDown", "ArrowUp"], (event) => {
    if (openDropdown.value) {
        if (event.key === "ArrowDown") {
            moveSelectedIndex(1);
        } else {
            moveSelectedIndex(-1);
        }
    }
});

onKeyStroke("Enter", () => {
    if (openDropdown.value) {
        optionEl.value![optionSelectedIndex.value].click();
    }
});

const options: {
    label: string;
    value: DocumentStatus;
}[] = [
    {
        label: "Rascunho",
        value: "draft",
    },
    {
        label: "Publicado",
        value: "completed",
    },
    {
        label: "Rejeitado",
        value: "rejected",
    },
    {
        label: "Revisado",
        value: "review",
    },
];

const handleSelectStatus = (status: DocumentStatus) => {
    emit("update:modelValue", status);
    emit("select", status);
    openDropdown.value = false;
};

const selectedStatusObject = computed(() => {
    return options.find((option) => option.value === props.modelValue);
});
</script>

<template>
    <div
        :class="[
            'input-main',
            openDropdown && 'input-main--focus',
            props.modelValue,
        ]"
    >
        <Popper
            @open:popper="openDropdown = true"
            @close:popper="openDropdown = false"
        >
            <div class="input-main__inner">
                <span class="input-main__inner__text">
                    {{ selectedStatusObject?.label || "" }}
                </span>
                <span class="input-main__inner__icon">
                    <InlineSvg
                        :src="chevronDownIconUrl"
                        width="16"
                        class="input-main__icon"
                    />
                </span>
            </div>

            <template #content="{ close }">
                <div class="input-main__popper">
                    <ul>
                        <li
                            v-for="(option, index) in options"
                            ref="optionEl"
                            :key="option.value"
                            :class="[
                                option.value === props.modelValue && 'selected',
                                optionSelectedIndex === index && 'hovered',
                            ]"
                            @click="
                                handleSelectStatus(option.value);
                                close();
                            "
                        >
                            {{ option.label }}
                        </li>
                    </ul>
                </div>
            </template>
        </Popper>
    </div>
</template>

<style lang="scss" scoped>
.input-main {
    @include flex-center(row);
    // background: red;
    border-radius: $border_r_md;
    height: $select_input_height !important;
    border: $spacing_0 solid var(--strong-color);
    background: var(--soft-color);
    position: relative;

    --strong-color: var(--neutral_11);
    --soft-color: var(--neutral_4);

    &.draft {
        --strong-color: var(--neutral_7);
        --soft-color: var(--neutral_4);
    }

    &.completed {
        --strong-color: var(--green_4);
        --soft-color: var(--green_1);
    }

    &.rejected {
        --strong-color: var(--red_4);
        --soft-color: var(--red_1);
    }

    &.review {
        --strong-color: var(--blue_4);
        --soft-color: var(--blue_1);
    }

    &:hover {
        border-color: $primary_4;
    }

    &--focus {
        box-shadow: 0px 0px 4px -1px $primary_5;
        border-color: $primary_4;

        .input-main__inner span.input-main__inner__icon {
            transform: rotate(180deg) !important;
        }
    }

    :deep(.inline-block) {
        width: 100%;
        min-width: 150px;
        height: $select_input_height !important;
        margin: 0px !important;
        border: none !important;
        cursor: pointer;

        @include flex-center(row);

        > div {
            padding: $spacing_6 $spacing_6;
            width: 100%;
        }
    }

    .input-main__inner {
        color: $neutral_8;
        font-size: $font_4;
        font-weight: 400;
        width: 100%;
        @include flex(row, center, start);

        span.input-main__inner__text {
            width: 100%;
            text-align: start;
            white-space: nowrap;
            overflow: hidden;
            text-overflow: ellipsis;
            color: var(--strong-color);
        }

        span.input-main__inner__icon {
            margin-left: auto;
            transition: $transition;
            transform: rotate(0deg);
            @include flex-center;
            color: var(--strong-color);
        }
    }

    .input-main__popper {
        background: $neutral_2;
        border-radius: $border_r_md;
        border: $spacing_0 solid $neutral_3;

        // padding: $spacing_6 $spacing_6;
        width: 100%;
        box-shadow: 0px 0px 4px -1px $neutral_7;

        ul {
            li {
                padding: $spacing_8 $spacing_6;
                transition: $transition;
                font-size: $font_4;
                color: $neutral_9;
                cursor: pointer;

                &.selected {
                    background: $neutral_3;
                    background: var(--soft-color);
                    color: var(--strong-color);
                }

                &:hover:not(.selected),
                &.hovered:not(.selected) {
                    background: $neutral_4;
                }
            }
        }
    }
}
</style>
