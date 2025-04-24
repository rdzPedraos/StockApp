import MarkdownIt from 'markdown-it';

export default defineNuxtPlugin(() => {
    const md = new MarkdownIt({ html: true });
    return {
        provide: {
            md: (txt: string) => md.render(txt),
        },
    };
});
