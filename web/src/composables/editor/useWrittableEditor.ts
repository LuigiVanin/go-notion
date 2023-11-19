// Core

// Libraries
import { useEditor } from "@tiptap/vue-3";
import StarterKit from "@tiptap/starter-kit";
import Placeholder from "@tiptap/extension-placeholder";

export function useWrittableEditor() {
    const editor = useEditor({
        extensions: [
            StarterKit,
            Placeholder.configure({ placeholder: "Write something..." }),
        ],
        content: "\n\n\n",
        editorProps: {
            attributes: {
                class: "prose",
            },
        },
    });

    return {
        editor,
    };
}
