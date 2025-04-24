<template>
    <div v-if="isLoading" class="grid grid-cols-2 gap-4">
        <USkeleton class="h-12 w-full" />
        <USkeleton class="h-12 w-full" />
        <USkeleton class="h-12 w-full" />
        <USkeleton class="h-12 w-full" />
    </div>

    <div v-else class="grid grid-cols-2 gap-4">
        <StockInfoCard
            title="Marketcap"
            info="Capitalization of the company"
            :value="ticker.market_cap"
        />

        <StockInfoCard title="Volume" info="Volume of the company" :value="40818" />
        <StockInfoCard title="Total Supply" info="Total supply of the company" :value="40818" />
        <StockInfoCard
            title="Circulating Supply"
            info="Circulating supply of the company"
            :value="40818"
        />
    </div>
</template>

<script setup lang="ts">
const store = useTickerStore();

const ticker = computed(() => store.data);
const isLoading = computed(() => store.isLoading);
</script>
