export type DocumentStatus = "completed" | "draft" | "rejected" | "review";

export type CreateDocument = {
    title: string;
    text: string;
    status: DocumentStatus;
};

export type Document = {
    id: number;
    title: string;
    text: string;
    status: DocumentStatus;
    createdAt: string;
    updatedAt: string;
};
