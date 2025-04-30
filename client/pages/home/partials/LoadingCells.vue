<template>
    <div>
        <div v-for="i in 8" :key="i" class="flex gap-4">
            <USkeleton v-for="col in randomCols" :key="col" :class="'h-10 ' + col + ' mb-4'" />
        </div>
    </div>
</template>

<script setup lang="ts">
const cols = ['flex-1', 'flex-2', 'flex-3', 'flex-4', 'flex-5', 'flex-6'];
const randomCols = cols.sort(() => Math.random() - 0.5);
</script>
