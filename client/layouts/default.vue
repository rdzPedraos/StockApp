<script setup lang="ts">
import { RouterLink } from 'vue-router';

const searchTicker = ref('');
</script>

<template>
    <div>
        <header class="flex items-center gap-4 border-b-2 border-gray-100 bg-white p-4">
            <AppLogo />

            <nav class="flex items-center gap-4">
                <RouterLink to="/">Home</RouterLink>
            </nav>

            <UInput
                v-model="searchTicker"
                class="block ml-auto"
                icon="i-lucide-search"
                size="xl"
                placeholder="Search ticker"
            />
        </header>
        <main class="max-w-[1800px] mx-auto px-4 py-8">
            <slot ></slot>
        </main>
    </div>
</template>
