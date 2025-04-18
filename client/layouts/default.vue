<script setup lang="ts">
import { RouterLink } from 'vue-router';
</script>

<template>
    <div>
        <header class="flex items-center gap-4 border-b-2 border-gray-100 bg-white px-3 py-4">
            <AppLogo />

            <nav class="flex items-center gap-4">
                <RouterLink to="/">Home</RouterLink>
                <RouterLink to="/about">About</RouterLink>
            </nav>
        </header>

        <main>
            <slot></slot>
        </main>
    </div>
</template>
