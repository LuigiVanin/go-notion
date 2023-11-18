export type CreateDocument = {
    title: string;
    text: string;
    status: "completed" | "draft" | "rejected" | "review";
};

export type Document = {
    id: number;
    title: string;
    text: string;
    status: "completed" | "draft" | "rejected" | "review";
    createdAt: string;
    updatedAt: string;
};
