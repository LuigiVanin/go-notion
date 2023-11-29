// Types
import { Editor } from "@tiptap/vue-3";

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

export const actions = [
    {
        icon: undoIconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.undo();
        },
        label: "Undo",
    },
    {
        icon: redoIconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.redo();
        },
        label: "Redo",
    },
    {
        icon: boldIconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.toggleBold();
        },
        label: "Bold",
    },
    {
        icon: italicIconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.toggleItalic();
        },
        label: "Italic",
    },
    {
        icon: underlineIconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.toggleUnderline();
        },
        label: "Underline",
    },
    {
        icon: strikeIconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.toggleStrike();
        },
        label: "Strikethrough",
    },
    {
        icon: heading1IconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.toggleHeading({ level: 1 });
        },
        label: "Heading 1",
    },
    {
        icon: heading2IconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.toggleHeading({ level: 2 });
        },
        label: "Heading 2",
    },
    {
        icon: heading3IconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.toggleHeading({ level: 3 });
        },
        label: "Heading 3",
    },
    {
        icon: blockquoteIconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.toggleBlockquote();
        },
        label: "Blockquote",
    },
    {
        icon: bulletListIconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.toggleBulletList();
        },
        label: "Bullet List",
    },
    {
        icon: orderListIconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.toggleOrderedList();
        },
        label: "Order List",
    },
    {
        icon: alignleftIconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.setTextAlign("left");
        },
        label: "Align Left",
    },
    {
        icon: aligncenterIconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.setTextAlign("center");
        },
        label: "Align Center",
    },
    {
        icon: alignrightIconUrl,
        action: (editor?: Editor) => {
            editor && editor.commands.setTextAlign("right");
        },
        label: "Align Right",
    },
];
