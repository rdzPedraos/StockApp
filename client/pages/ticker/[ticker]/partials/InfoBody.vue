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
            :value="ticker.company.marketCap"
        />

        <StockInfoCard title="Volume" info="Volume of the company" :value="ticker.company.volume" />

        <StockInfoCard
            title="Beta"
            info="Beta works as a measure of the stock's volatility relative to the market, if the beta is 1, the stock is as volatile as the market. If the beta is greater than 1, the stock is more volatile than the market. If the beta is less than 1, the stock is less volatile than the market."
            :value="ticker.company.beta"
        />

        <StockInfoCard title="Sector" info="Sector of the company" :value="ticker.company.sector" />
    </div>
</template>

<script setup lang="ts">
const store = useTickerStore();

const ticker = computed(() => store.data);
const isLoading = computed(() => store.isLoading);
</script>
