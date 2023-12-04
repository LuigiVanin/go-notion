<script lang="ts" setup>
// Libraries
import Popper from "vue3-popper";

// Components
import Button from "./core/Button.vue";

// Types
import { Document } from "@/types/document";

// Assets
import threeDotsIconUrl from "@/assets/icons/three-dots.svg?url";
import trashIconUrl from "@/assets/icons/trash.svg?url";

const props = defineProps<{
    document: Document;
}>();
const emit = defineEmits(["delete-document", "edit-document"]);
</script>

<template>
    <Popper placement="bottom-end">
        <Button :icon="threeDotsIconUrl" />

        <template #content>
            <div class="document-table__cell__popper">
                <p class="title">Info</p>
                <p>
                    Quantidade palavras:
                    <span>
                        {{ props.document.text.split(/[ \n]+/).length }}
                    </span>
                </p>

                <div class="btn-group">
                    <Button @click="emit(`edit-document`)"> Editar </Button>
                    <Button
                        :icon="trashIconUrl"
                        color="red"
                        btn-type="soft"
                        @click="emit(`delete-document`)"
                    />
                </div>
            </div>
        </template>
    </Popper>
</template>

<style lang="scss" scoped>
.document-table__cell__popper {
    background: $neutral_2;
    border-radius: $border_r_md;
    border: $spacing_0 solid $neutral_3;

    width: 100%;
    box-shadow: 0px 0px 15px -1px $neutral_7;
    padding: $spacing_10 $spacing_10;
    @include flex(column, start, start);
    gap: $spacing_6;

    .btn-group {
        width: 100%;
        @include flex(row, center, end);
        gap: $spacing_6;
        margin-top: $spacing_8;

        :deep(button.button-main--soft) {
            padding: $spacing_5 $spacing_5;
        }
    }

    p.title {
        font-size: $font_6;
        font-weight: 500;
        text-decoration: none;
    }

    p:not(.title) {
        font-size: $font_5;
        font-weight: 400;
        color: $neutral_7;
        text-decoration: none;

        span {
            font-weight: 500;
            color: $neutral_10;
        }
    }
}
</style>
