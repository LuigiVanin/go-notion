import { Extension, Range } from "@tiptap/core";
import Suggestion from "@tiptap/suggestion";
import { Editor } from "@tiptap/vue-3";

type CommandProps = {
    editor: Editor;
    range: Range;
    props: any;
};

export default Extension.create({
    name: "command",

    addOptions() {
        return {
            suggestion: {
                char: "/",
                command: ({ editor, range, props }: CommandProps) => {
                    console.log("command!!", props);
                    editor.commands.deleteRange(range);
                    props?.action && props.action(editor);
                },
            },
        };
    },

    addProseMirrorPlugins() {
        return [
            Suggestion({
                editor: this.editor,
                startOfLine: true,
                ...this.options.suggestion,
            }),
        ];
    },
});
