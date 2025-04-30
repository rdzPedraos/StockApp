<script setup lang="ts">
import { RouterLink } from 'vue-router';
</script>

<template>
    <div>
        <header class="flex items-center gap-4 border-b-2 border-gray-100 bg-white p-4">
            <RouterLink to="/">
                <AppLogo />
            </RouterLink>
        </header>

        <main class="max-w-[1800px] mx-auto px-4 py-8">
            <slot></slot>
        </main>
    </div>
</template>
