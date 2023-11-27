import { VueRenderer } from "@tiptap/vue-3";
import tippy from "tippy.js";

import CommandsList from "@/components/editor/CommandList.vue";
import { actions } from "../actions";

export default {
    items: ({ query }: { query: string }) => {
        return actions.filter((item) =>
            item.label.toLowerCase().startsWith(query.toLowerCase())
        );
    },

    render: () => {
        let component: any;
        let popup: any;

        return {
            onStart: (props: any) => {
                component = new VueRenderer(CommandsList, {
                    // using vue 2:
                    // parent: this,
                    // propsData: props,
                    props,
                    editor: props.editor,
                });

                if (!props.clientRect) {
                    return;
                }

                popup = tippy("body", {
                    getReferenceClientRect: props.clientRect,
                    appendTo: () => document.body,
                    content: component.element,
                    showOnCreate: true,
                    interactive: true,
                    trigger: "manual",
                    placement: "bottom-start",
                });
            },

            onUpdate(props: any) {
                component.updateProps(props);

                if (!props.clientRect) {
                    return;
                }

                popup[0].setProps({
                    getReferenceClientRect: props.clientRect,
                });
            },

            onKeyDown(props: any) {
                if (props.event.key === "Escape") {
                    popup[0].hide();

                    return true;
                } else if (
                    props.event.key === "ArrowDown" ||
                    props.event.key === "ArrowUp" ||
                    props.event.key === "Enter"
                ) {
                    return true;
                }
                return false;
            },

            onExit() {
                popup[0].destroy();
                component.destroy();
            },
        };
    },
};
