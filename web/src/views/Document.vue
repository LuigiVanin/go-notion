<script lang="ts" setup>
// Core
import { onMounted, provide, ref } from "vue";
import { useRoute, useRouter } from "vue-router";

// Libraries
import { EditorContent } from "@tiptap/vue-3";

// Composables
import { useWrittableEditor } from "@/composables/editor/useWrittableEditor.ts";
import { useFetchDocument } from "@/composables/api/useFetchDocument.ts";
import { useUpdateDocument } from "@/composables/api/useUpdateDocument.ts";

// Components
import EditorActions from "@/components/editor/EditorActions.vue";
import TextInput from "@/components/core/TextInput.vue";
import StatusSelect from "@/components/editor/StatusSelect.vue";
import Button from "@/components/core/Button.vue";

const route = useRoute();
const router = useRouter();

const { editor } = useWrittableEditor({
    onBlur: ({ editor }) => {
        console.log("EDITOR: ", editor.getHTML());
        saveDocument();
    },
});
const { fetchDocument, document, fetchError } = useFetchDocument();
const { updateDocument } = useUpdateDocument();

const title = ref("");

provide("editor", {
    editor,
});

onMounted(async () => {
    if (!route.params?.id) {
        router.push("/");
    }

    await fetchDocument(route.params?.id as string);

    if (fetchError.value || !document.value) {
        router.push("/");
        return;
    }

    console.log("EDITOR 2: ", document.value);

    editor.value?.commands.setContent(
        document.value?.text.replaceAll("\n", "<p></p>")
    );
    title.value = document.value?.title;
});

const saveDocument = async () => {
    if (!document.value) {
        return;
    }
    console.log("EVINADO EDITOR: ", document.value);

    await updateDocument(document.value.id, {
        title: title.value,
        text: editor.value?.getHTML() || "",
    });
};
</script>

<template>
    <div class="editor-container">
        <header>
            <Button size="lg" btn-type="simple">Voltar</Button>
            <Button size="lg">Exportar para PDF </Button>
        </header>
        <TextInput v-model="title" />
        <div class="editor-container__actions">
            <EditorActions />
            <StatusSelect />
        </div>
        <div class="editor-content__wrapper">
            <EditorContent :editor="editor" />
        </div>
    </div>
</template>

<style lang="scss" scoped>
.editor-container {
    width: 100%;
    @include flex(column, center, center);
    gap: $spacing_13;

    header {
        width: 100%;
        @include flex(row, center, space-between);
    }

    .editor-container__actions {
        width: 100%;
        @include flex(row, start, space-between);
    }

    :deep(.editor-content__wrapper) {
        width: 100%;
    }

    :deep(.input-main) {
        height: 40px;

        input {
            font-weight: 600;
            font-size: $font_8;
        }
    }
}
</style>

<style lang="scss">
.tiptap p.is-empty::before {
    color: #adb5bd;
    content: attr(data-placeholder);
    float: left;
    height: 0;
    pointer-events: none;
}

.tiptap.ProseMirror {
    padding-block: $spacing_8;
    min-height: 500px;

    strong {
        font-weight: 700;
    }

    p {
        margin-bottom: $spacing_6;
        color: $neutral_11;
    }
}
</style>
