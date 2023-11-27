<script lang="ts" setup>
// Core
import { onMounted, ref } from "vue";
import { useRouter } from "vue-router";

// Libraries
import InlineSvg from "vue-inline-svg";

// Composables
import { useApi } from "@/composables/api/useApi.ts";
import { useFetchDocuments } from "@/composables/api/useFetchDocuments.ts";

// Components
import CreateDocumentButton from "@/components/CreateDocumentButton.vue";
import AlertBox from "@/components/core/AlertBox.vue";
import DocumentTable from "@/components/DocumentTable.vue";

// Assets
import paperIconUrl from "@/assets/icons/paper.svg?url";

const api = useApi();
const router = useRouter();
const { documents, fetchDocuments, fetchError } = useFetchDocuments();

const documentLoading = ref(false);

onMounted(async () => {
    await fetchDocuments();
    if (fetchError.value || !fetchError.value) {
        console.log(fetchError.value);
        return;
    }
    console.log(documents.value);
});

const createDocument = async () => {
    documentLoading.value = true;
    const { data, error } = await api.document.create({
        title: "Teste",
        text: "\n",
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
        <h3>Documentos:</h3>
        <DocumentTable v-if="documents?.docs" :documents="documents.docs" />
    </div>
</template>

<style lang="scss" scoped>
.home {
    width: 100%;
    @include flex(column, start, start);
    gap: $spacing_20;

    > h3 {
        padding-top: $spacing_6;
        font-size: $font_7;
        color: $neutral_8;
        font-weight: 600;
    }

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
