import tailwindcss from '@tailwindcss/vite';

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    ssr: false,

    compatibilityDate: '2024-11-01',
    devtools: { enabled: true },

    modules: [
        '@nuxt/eslint',
        '@nuxt/icon',
        '@nuxt/image',
        '@nuxt/test-utils',
        '@nuxt/ui',
        '@pinia/nuxt',
    ],

    css: ['~/assets/main.css'],

    routeRules: {
        '/': {
            redirect: '/home',
        },
    },

    vite: {
        plugins: [tailwindcss()],
    },
});
