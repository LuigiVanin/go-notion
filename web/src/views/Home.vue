<script lang="ts" setup>
// Core
import { onMounted, ref } from "vue";
import { useRouter } from "vue-router";

// Libraries
import InlineSvg from "vue-inline-svg";

// Composables
import { useApi } from "@/composables/api/useApi.ts";
// Components
import CreateDocumentButton from "@/components/CreateDocumentButton.vue";
import AlertBox from "@/components/core/AlertBox.vue";

// Assets
import paperIconUrl from "@/assets/icons/paper.svg?url";

const api = useApi();
const router = useRouter();

const documentLoading = ref(false);
const documents = ref<Document[]>([]);

onMounted(async () => {
    const { data, error } = await api.document.fetch();
    if (error || !data) {
        console.log(error);
        return;
    }
    console.log(data);

    documents.value = data.docs;
});

const createDocument = async () => {
    console.log("Create Document");
    documentLoading.value = true;
    const { data, error } = await api.document.create({
        title: "Teste",
        text: "\n\n\n",
        status: "draft",
    });
    documentLoading.value = false;

    if (error || !data?.id) {
        console.log(error);
        return;
    }

    router.push(`/document/${data.id}`);

    console.log(data);
};
</script>

<template>
    <div class="home">
        <header class="home__title">
            <InlineSvg :src="paperIconUrl" />
            <h3>Home</h3>
            <span class="home__title__divider" />
            <p>A sua home page possui a listagem de documentos e ações.</p>
        </header>
        <CreateDocumentButton
            :loading="documentLoading"
            :action="createDocument"
        />
        <AlertBox
            type="info"
            text="Lembrando que você possui o direito de apenas 10 documentos por conta."
            closable
        />
        <p>USER: {{ documents }}</p>
    </div>
</template>

<style lang="scss" scoped>
.home {
    width: 100%;
    @include flex(column, start, start);
    gap: $spacing_20;

    header.home__title {
        width: 100%;
        @include flex(row, center, start);
        gap: $spacing_6;

        :deep(svg) {
            width: 20px;
            height: 20px;
            path {
                stroke: $primary_5;
            }
        }

        h3 {
            font-size: $font_7;
            font-weight: 600;
            @include gradient-text;
        }

        span.home__title__divider {
            height: 20px;
            width: 1px;
            background: $neutral_6;
        }

        p {
            color: $neutral_8 !important;
            font-size: $font_4;
            font-weight: 400;
        }
    }

    p {
        color: $neutral_11 !important;
    }
}
</style>
