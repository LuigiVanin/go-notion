<script lang="ts" setup>
import EditorButtonIcon from "./EditorButtonIcon.vue";

// Assets
import boldIconUrl from "@/assets/icons/bold.svg?url";
import italicIconUrl from "@/assets/icons/italic.svg?url";
import underlineIconUrl from "@/assets/icons/underline.svg?url";
import strikeIconUrl from "@/assets/icons/strike.svg?url";
import heading1IconUrl from "@/assets/icons/heading1.svg?url";
import heading2IconUrl from "@/assets/icons/heading2.svg?url";
import heading3IconUrl from "@/assets/icons/heading3.svg?url";
import blockquoteIconUrl from "@/assets/icons/blockquote.svg?url";
import undoIconUrl from "@/assets/icons/undo.svg?url";
import redoIconUrl from "@/assets/icons/redo.svg?url";
import bulletListIconUrl from "@/assets/icons/bullet-list.svg?url";
import orderListIconUrl from "@/assets/icons/order-list.svg?url";
import alignleftIconUrl from "@/assets/icons/align-left.svg?url";
import aligncenterIconUrl from "@/assets/icons/align-center.svg?url";
import alignrightIconUrl from "@/assets/icons/align-right.svg?url";
import { ShallowRef, inject, onMounted } from "vue";
import { Editor } from "@tiptap/vue-3";

const injected = inject<{
    editor: ShallowRef<Editor | undefined>;
}>("editor");

onMounted(() => {
    console.log("EDITOR: ", injected);
});

const actions = [
    {
        icon: undoIconUrl,
        action: (editor?: Editor) => {
            console.log(injected?.editor.value);
            console.log("EDITOR: ", editor);
            editor && editor.commands.setBold();
        },
        label: "Undo",
    },
    {
        icon: redoIconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.setBold();
        },
        label: "Redo",
    },
    {
        icon: boldIconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.setBold();
        },
        label: "Bold",
    },
    {
        icon: italicIconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.setBold();
        },
        label: "Italic",
    },
    {
        icon: heading1IconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.setBold();
        },
        label: "Heading 1",
    },
    {
        icon: heading2IconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.setBold();
        },
        label: "Heading 2",
    },
    {
        icon: heading3IconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.setBold();
        },
        label: "Heading 3",
    },
    {
        icon: blockquoteIconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.setBold();
        },
        label: "Blockquote",
    },
    {
        icon: bulletListIconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.setBold();
        },
        label: "Bullet List",
    },
    {
        icon: orderListIconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.setBold();
        },
        label: "Order List",
    },
    {
        icon: alignleftIconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.setBold();
        },
        label: "Align Left",
    },
    {
        icon: aligncenterIconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.setBold();
        },
        label: "Align Center",
    },
    {
        icon: alignrightIconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.setBold();
        },
        label: "Align Right",
    },
    {
        icon: underlineIconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.setBold();
        },
        label: "Underline",
    },
    {
        icon: strikeIconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.setBold();
        },
        label: "Strikethrough",
    },
];
</script>

<template>
    <div class="editor-action">
        <EditorButtonIcon
            v-for="act in actions"
            :icon="act.icon"
            :action="() => act.action(injected?.editor.value)"
        />
    </div>
</template>

<style lang="scss" scoped>
.editor-action {
    width: 100%;
    @include flex(row, start, start);
    gap: $spacing_5;
}
</style>
