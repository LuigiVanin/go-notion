<script lang="ts" setup>
// Core
import { useRouter } from "vue-router";

// Libraries
import InlineSvg from "vue-inline-svg";
import Popper from "vue3-popper";

// Copmposables
import { useUpdateDocument } from "@/composables/api/useUpdateDocument.ts";

// Components
import StatusSelect from "./editor/StatusSelect.vue";
import Button from "./core/Button.vue";

// Assets
import paperIconUrl from "@/assets/icons/paper.svg?url";
import threeDotsIconUrl from "@/assets/icons/three-dots.svg?url";
import trashIconUrl from "@/assets/icons/trash.svg?url";

// Types
import { Document } from "@/types/document";

type DocumentTableProps = {
    documents: Document[];
};

const props = defineProps<DocumentTableProps>();

const router = useRouter();

const { updateDocument } = useUpdateDocument();

const formatDateString = (date: string) => {
    const dateObj = new Date(date);
    return `${dateObj.getDate()}/${
        dateObj.getMonth() + 1
    }/${dateObj.getFullYear()}`;
};
</script>

<template>
    <div class="document-table">
        <header>
            <div class="document-table__cell">Título</div>
            <div class="document-table__cell">Status</div>
            <div class="document-table__cell">Data</div>
            <div class="document-table__cell">Ações</div>
        </header>
        <main>
            <div
                v-for="document in props.documents"
                :key="document.id"
                class="document-table__line"
                @click="router.push(`/document/${document.id}`)"
            >
                <div class="document-table__cell document-table__cell--title">
                    <InlineSvg
                        :src="paperIconUrl"
                        width="16px"
                        min-width="16px"
                    />

                    <p>
                        {{ document.title }}
                    </p>
                </div>
                <div class="document-table__cell" @click.stop>
                    <StatusSelect
                        v-model="document.status"
                        @select="
                            updateDocument(document.id, { status: $event })
                        "
                    />
                </div>
                <div class="document-table__cell">
                    {{ formatDateString(document.createdAt) }}
                </div>
                <div class="document-table__cell" @click.stop>
                    <Popper placement="bottom-end">
                        <Button :icon="threeDotsIconUrl" />

                        <template #content>
                            <div class="document-table__cell__popper">
                                <p class="title">Info</p>
                                <p>
                                    Quantidade palavras:
                                    <span>
                                        {{
                                            document.text.split(/[ \n]+/).length
                                        }}
                                    </span>
                                </p>

                                <div class="button-group">
                                    <Button
                                        @click="
                                            router.push(
                                                `/document/${document.id}`
                                            )
                                        "
                                    >
                                        Editar
                                    </Button>
                                    <Button
                                        :icon="trashIconUrl"
                                        color="red"
                                        btn-type="soft"
                                    />
                                </div>
                            </div>
                        </template>
                    </Popper>
                </div>
            </div>
        </main>
    </div>
</template>

<style lang="scss" scoped>
.document-table {
    width: 100%;
    @include flex(column, start, start);
    overflow-x: auto;

    main {
        width: 100%;
        min-width: 600px;
        overflow-x: auto;
    }

    header {
        width: 100%;
        min-width: 600px;

        background: $neutral_4;
        border-radius: $border_r_lg;

        .document-table__cell {
            padding-block: $spacing_13;
            font-size: $font_6;
            color: $primary_6;
            font-weight: 500;
        }
    }

    .document-table__cell {
        @include flex(row, center, start);
        padding-inline: $spacing_13;
        padding-block: $spacing_6;
        color: $neutral_10;
        gap: $spacing_4;
        font-size: $font_5;
        position: relative;

        &.document-table__cell--title :deep(svg path) {
            stroke: $primary_6;
        }

        .document-table__cell__popper {
            background: $neutral_2;
            border-radius: $border_r_md;
            border: $spacing_0 solid $neutral_3;

            // padding: $spacing_6 $spacing_6;
            width: 100%;
            box-shadow: 0px 0px 4px -1px $neutral_7;
            padding: $spacing_10 $spacing_10;
            @include flex(column, start, start);
            gap: $spacing_6;

            .button-group {
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
    }

    header,
    main > .document-table__line {
        display: grid;

        grid-template-columns:
            minmax(200px, 1.4fr)
            minmax(145px, 145px)
            minmax(100px, 0.4fr)
            minmax(100px, 0.3fr);
    }

    main .document-table__line {
        // padding: $spacing_6 0;
        // background: red;
        &:not(:last-child) {
            border-bottom: $spacing_0 solid $neutral_4;
        }

        &:hover {
            background: $neutral_3;

            p {
                color: $neutral_11;
                text-decoration: underline;
            }
        }

        .document-table__cell {
            cursor: pointer;

            p {
                @include line-clamp(1);
            }

            :deep(.input-main) {
                width: 125px !important;

                .inline-block {
                    min-width: 100px !important;
                    //         width: 100px !important;
                }
            }
        }
    }
}
</style>

<style lang="scss">
.document-table .document-table__cell .popper {
    min-width: 250px;
}
</style>
