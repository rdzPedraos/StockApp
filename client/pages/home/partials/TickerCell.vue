<script setup lang="ts">
const props = defineProps<{
    ticker: string;
    company: string;
    logo: string;
}>();

// Importar la imagen de respaldo desde assets
const fallbackImage = new URL('~/assets/img/logo.png', import.meta.url).href;

// Función para manejar el error de carga de imagen
function handleImageError(event: Event) {
    const imgElement = event.target as HTMLImageElement;

    if (imgElement.src !== fallbackImage) {
        imgElement.src = fallbackImage;
    }
}
</script>

<template>
    <div class="flex items-center gap-2">
        <img
            :src="props.logo"
            class="w-8 h-8 rounded-full drop-shadow-md"
            alt="Logo de la empresa"
            @error="handleImageError"
        />

        <div class="flex items-center gap-1">
            <span class="capitalize font-bold">{{ props.company }}</span>
            <span>{{ props.ticker }}</span>
        </div>
    </div>
</template>
