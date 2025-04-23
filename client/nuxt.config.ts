import tailwindcss from '@tailwindcss/vite';

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    compatibilityDate: '2024-11-01',
    devtools: { enabled: true },

    modules: ['@nuxt/eslint', '@nuxt/icon', '@nuxt/image', '@nuxt/test-utils', '@nuxt/ui'],

    css: ['~/assets/main.css'],

    routeRules: {
        '/': {
            redirect: '/home',
        },
    },

    vite: {
        plugins: [tailwindcss()],
        server: {
            proxy: {
                '/back': {
                    target: process.env.API_URL || 'http://localhost:3001',
                    changeOrigin: true,
                    rewrite: (path) => path.replace(/^\/back/, '/api'),
                },
            },
        },
    },
});
