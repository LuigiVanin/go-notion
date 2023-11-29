<script lang="ts" setup>
// Core
import { ShallowRef, computed, inject } from "vue";

// Services & Helpers
import { actions } from "@/helpers/actions.ts";

// Components
import EditorButtonIcon from "./EditorButtonIcon.vue";
import { Editor } from "@tiptap/vue-3";

const injected = inject<{
    editor: ShallowRef<Editor | undefined>;
}>("editor");

type EditorActionsProps = {
    actions?: string[];
};

const props = defineProps<EditorActionsProps>();

const filteredActions = computed(() => {
    if (props.actions) {
        return actions.filter((act) => props.actions?.includes(act.label));
    }
    return actions;
});
</script>

<template>
    <div class="editor-action">
        <EditorButtonIcon
            v-for="act in filteredActions"
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
