// @ts-check
import withNuxt from './.nuxt/eslint.config.mjs';

export default withNuxt({
    rules: {
        'vue/html-self-closing': [
            'error',
            {
                html: {
                    void: 'any', // Permitir cualquier estilo para elementos vacíos HTML
                    normal: 'never',
                    component: 'always',
                },
                svg: 'always',
                math: 'always',
            },
        ],
    },
});