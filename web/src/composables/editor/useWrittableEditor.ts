// Core

// Libraries
import { useEditor } from "@tiptap/vue-3";
import StarterKit from "@tiptap/starter-kit";
import Placeholder from "@tiptap/extension-placeholder";

type EditorProps = Parameters<typeof useEditor>[0];

export function useWrittableEditor(options?: Partial<EditorProps>) {
    const editor = useEditor({
        ...options,
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
