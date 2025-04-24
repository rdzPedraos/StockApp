<template>
    <div v-if="isLoading" class="flex items-center gap-4">
        <USkeleton class="h-12 w-12 rounded-full" />

        <div class="flex flex-col gap-2">
            <USkeleton class="h-4 w-[250px]" />
            <USkeleton class="h-4 w-[200px]" />
        </div>
    </div>

    <div v-else>
        <TickerInfo :ticker="ticker.ticker" :company="ticker.company" :logo="ticker.logo" />
        <p class="text-5xl mt-2 font-bold font-mono">{{ formatPrice(ticker.price) }}</p>
    </div>
</template>

<script setup lang="ts">
const store = useTickerStore();

const ticker = computed(() => store.data);
const isLoading = computed(() => store.isLoading);
</script>
