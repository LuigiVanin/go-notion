// Core

// Helpers & Services
import Command from "@/helpers/tiptap/commands.ts";
import suggestion from "@/helpers/tiptap/suggestion.ts";

// Libraries
import { useEditor } from "@tiptap/vue-3";
import BubbleMenuExt from "@tiptap/extension-bubble-menu";
import StarterKit from "@tiptap/starter-kit";
import Placeholder from "@tiptap/extension-placeholder";
import TextAlign from "@tiptap/extension-text-align";
import Underline from "@tiptap/extension-underline";

type EditorProps = Parameters<typeof useEditor>[0];

export function useWrittableEditor(options?: Partial<EditorProps>) {
    const editor = useEditor({
        extensions: [
            StarterKit,
            Placeholder.configure({ placeholder: "Write something..." }),
            TextAlign.configure({
                types: ["heading", "paragraph"],
                alignments: ["left", "center", "right"],
            }),
            Underline,
            BubbleMenuExt,
            Command.configure({
                suggestion,
            }),
        ],
        content: "",
        editorProps: {
            attributes: {
                class: "prose",
            },
        },
        ...options,
    });

    return {
        editor,
    };
}
