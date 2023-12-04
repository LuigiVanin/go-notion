<script lang="ts" setup>
// Components

// Libraries
import { useMagicKeys, whenever } from "@vueuse/core";
import InlineSvg from "vue-inline-svg";

// Assets
import closeIconUrl from "@/assets/icons/close.svg?url";

type ModalConfigurationProps = {
    show: boolean;
    disableAction: () => void;
    loading?: boolean;
    disabled?: boolean;

    appearence?: {
        mClass?: string;
        customClass?: string;
        enableGradient?: boolean;
        disableBorder?: boolean;
        disableTeleport?: boolean;
        disableCloseButton?: boolean;
        disableCloseOnBackdrop?: boolean;
        backdropAppearence?: "normal" | "dark" | "darker" | "transparent";
        primaryButtonStatus?: "success" | "danger";
    };
};

const props = defineProps<ModalConfigurationProps>();
const emit = defineEmits([
    "primary-button-clicked",
    "secondary-button-clicked",
    "backdrop-click",
    "close-button-click",
]);

const DISABLED_TELEPORT_CLASSES = [
    "negative-feedback",
    "video-learn-modal",
    "negative-feedback img",
];

const { current } = useMagicKeys();

const backdropClickHandler = () => {
    emit("backdrop-click");
    !props.appearence?.disableCloseOnBackdrop && props.disableAction();
};

whenever(
    () => current.has("escape"),
    () => {
        if (props.show) props.disableAction();
    }
);
</script>

<template>
    <Teleport
        to="body"
        :disabled="
            DISABLED_TELEPORT_CLASSES.includes(
                props.appearence?.mClass ?? ''
            ) || props.appearence?.disableTeleport
        "
    >
        <transition name="modal-down-fade">
            <div
                v-if="props.show"
                class="modal__backdrop"
                :class="[
                    `modal__backdrop--${
                        props.appearence?.backdropAppearence || 'normal'
                    }`,
                    props.appearence?.disableTeleport &&
                        'modal__backdrop--no-teleport',
                ]"
                @mousedown="backdropClickHandler"
            >
                <div
                    :class="[
                        'modal--core',
                        !props.appearence?.enableGradient &&
                            'modal--disable-gradient',
                        props.appearence?.disableBorder &&
                            'modal--disable-border',
                        props.appearence?.customClass,
                    ]"
                    :aria-labelledby="'modal'"
                    role="dialog"
                    @mousedown.stop
                >
                    <button
                        v-if="!props.appearence?.disableCloseButton"
                        class="modal__close-btn"
                    >
                        <InlineSvg
                            :src="closeIconUrl"
                            width="22"
                            height="22"
                            class="modal__close"
                            @click="props.disableAction"
                        />
                    </button>

                    <slot />
                </div>
            </div>
        </transition>
    </Teleport>
</template>

<style lang="scss" scoped>
@mixin base-index($index: $base_index) {
    z-index: $index;
    position: relative;
}

.modal__backdrop {
    position: fixed;
    inset: 0;
    background-color: rgba(78, 78, 78, 0.075);
    backdrop-filter: blur(1px);
    z-index: 99000;
    padding: $spacing_8;
    @include flex-center();

    &.modal__backdrop--no-teleport {
        position: absolute;
    }

    & > .modal--core {
        border-top: $spacing_8 solid $primary_6;
        max-width: 400px;
        width: 100%;
        border-radius: $border_r_lg;
        background: $neutral_2;
        box-shadow: 0px 0px 20px -1px $neutral_7;
        position: relative;
        @include flex-center(column);
        overflow: hidden;
        padding: $spacing_10;
        padding: $spacing_20;

        &::before {
            content: "";
            position: absolute;
            height: 100px;
            z-index: 1000;
            left: 0;
            top: 0;
            width: 100%;
            background: linear-gradient(1deg, #efe7f200 1.39%, #efe7f2 99.1%);
            pointer-events: none;
        }

        &.modal--disable-border {
            border-top: none;
        }

        &.modal--disable-gradient {
            &::after,
            &::before {
                display: none;
            }
        }

        button.modal__close-btn {
            position: absolute;
            top: $spacing_6;
            right: $spacing_6;
            @include reset-button();
            @include flex-center();
            padding: $spacing_2;
            border-radius: $border_r_md;
            transition: all $animation__duration ease-in-out;

            &:hover {
                background: $primary_1;
                :deep(svg path) {
                    stroke: $primary_6;
                }
            }

            :deep(svg path) {
                transition: all $animation__duration ease-in-out;
                stroke: $neutral_8;
            }
        }
    }
}

.modal-down-fade-enter-active,
.modal-down-fade-leave-active {
    transition: all $animation__duration $transition_spring;
    opacity: 1;

    > .modal--core {
        transition: all 0.3s $transition_spring;
        transform: translateY(0px);
    }
}

.modal-down-fade-enter-from,
.modal-down-fade-leave-to {
    opacity: 0;

    .modal--core {
        transform: translateY(-20px);
    }
}
</style>
