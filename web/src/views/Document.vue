<script lang="ts" setup>
// Core
import { onMounted, provide, ref } from "vue";
import { useRoute, useRouter } from "vue-router";

// Libraries
import { EditorContent } from "@tiptap/vue-3";

// Composables
import { useWrittableEditor } from "@/composables/editor/useWrittableEditor.ts";
import { useFetchDocument } from "@/composables/api/useFetchDocument.ts";

// Components
import EditorActions from "@/components/editor/EditorActions.vue";
import TextInput from "@/components/core/TextInput.vue";
import StatusSelect from "@/components/editor/StatusSelect.vue";

const route = useRoute();
const router = useRouter();

const { editor } = useWrittableEditor();
const { fetchDocument, document, fetchError } = useFetchDocument();

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

    console.log("EDITOR 2: ", document.value?.text);

    editor.value?.commands.setContent(
        document.value?.text.replaceAll("\n", "<p></p>")
    );
    title.value = document.value?.title;
});
</script>

<template>
    <div class="editor-container">
        <TextInput v-model="title" />
        <EditorActions />
        <StatusSelect />
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
