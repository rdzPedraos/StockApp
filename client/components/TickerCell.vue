<script setup lang="ts">
const props = defineProps<{
    ticker: string;
    company: string;
    logo: string;
}>();
</script>

<template>
    <div class="flex items-center gap-2">
        <img :src="props.logo" class="w-6 h-6 rounded-full" />
        <div class="flex items-center gap-1">
            <span class="capitalize font-bold">{{ props.company }}</span>
            <span>{{ props.ticker }}</span>
        </div>
    </div>
</template>
