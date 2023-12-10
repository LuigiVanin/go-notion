<script lang="ts" setup>
import { ref } from "vue";
import CommandItem from "@/components/editor/partials/CommandItem.vue";
import { onKeyStroke } from "@vueuse/core";

const props = defineProps<{
    command: any;
    items: { icon: string; label: string; action: Function }[];
}>();

const selectedIndex = ref(0);

const handleActionClick = (item: any) => {
    props.command && props.command(item); // props.command(item);
};

const moveSelectedIndex = (steps: number) => {
    selectedIndex.value = (selectedIndex.value + steps) % props.items.length;

    if (selectedIndex.value < 0) {
        selectedIndex.value = props.items.length - 1;
    }
};

onKeyStroke(["ArrowDown", "ArrowUp"], (event) => {
    if (event.key === "ArrowDown") {
        moveSelectedIndex(1);
    } else {
        moveSelectedIndex(-1);
    }
});

onKeyStroke("Enter", () => {
    const item = props.items[selectedIndex.value];
    item && handleActionClick(item);
});
</script>

<template>
    <div class="command-list">
        <h3>Ações</h3>
        <p v-if="!items.length">Nenhuma ação foi encontrada...</p>
        <ul v-else>
            <CommandItem
                v-for="(item, index) in props.items"
                :key="item.label"
                :icon-url="item.icon"
                :label="item.label"
                :selected="index === selectedIndex"
                @click="handleActionClick(item)"
            />
        </ul>
    </div>
</template>

<style lang="scss" scoped>
.command-list {
    // max-height: 200px;
    min-width: 250px;

    background-color: red;

    padding: $spacing_6 $spacing_6;
    border: $spacing_0 solid $neutral_3;
    background: $neutral_2;
    border-radius: $border_r_md;
    box-shadow: 0px 0px 6px -2px $neutral_7;

    h3 {
        font-size: $font_6;
        color: $neutral_8;
        margin-bottom: $spacing_6;
    }

    ul {
        max-height: 200px;
        overflow-y: auto;
        @include set-scrollbar;
    }
}
</style>
