import jsPDF, { HTMLOptions } from "jspdf";

interface Exporter<T> {
    export(source: T, callback?: () => void): void;
}

export type DocumentSource = {
    title: string;
    content: string;
};

const htmlExportOptions: HTMLOptions = {
    margin: [10, 10, 10, 10],
    autoPaging: "text",
    x: 0,
    y: 0,
    width: 190,
    windowWidth: 675,
};

export class Html2PdfExporter implements Exporter<DocumentSource> {
    export(source: DocumentSource, callback?: () => void) {
        const pdf = new jsPDF();
        const name = source.title;
        const content =
            `<h1 style="line-height: 2rem; font-weight: 700; font-size: 1.538rem">${name}</h1>` +
            source.content;

        const htmlWrapper = (text: string) => `
        <div width="675px" style="width: 675px; line-height: 1.5rem; font-size: 13px">
        ${text}
        </div>`;

        const htmlContent = htmlWrapper(
            `
            ${content
                .replaceAll(/<a(.*?)>/g, "")
                .replaceAll(
                    "<ul>",
                    '<ul style="list-style-type: disc !important; padding-left: 2rem">'
                )
                .replaceAll(
                    "<p>",
                    '<p style="line-height: 1.4rem; font-size: 1.077rem; margin-bottom: 0.5rem">'
                )
                .replaceAll("<strong>", '<strong style="font-weight: 700">')
                .replace(
                    "<h3>",
                    '<br><h3 style="line-height: 1.7rem; font-weight: 700">'
                )
                .replaceAll(
                    "<h2>",
                    '<br><h2 style="line-height: 2.5rem; font-weight: 700; margin-bottom: 1rem; font-size: 1.385rem">'
                )
                .replace(
                    "<h1>",
                    '<br><h1 style="line-height: 2rem; font-weight: 700; font-size: 1.538rem">'
                )}`
        );

        pdf.html(htmlContent, {
            ...htmlExportOptions,
            callback: function (doc) {
                callback && callback();
                console.log(htmlContent);
                doc.save(`${name}.pdf`);
            },
        });
    }
}
