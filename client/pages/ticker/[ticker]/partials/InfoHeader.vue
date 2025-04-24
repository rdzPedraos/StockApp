<template>
    <div v-if="isLoading" class="flex items-center gap-4">
        <USkeleton class="h-12 w-12 rounded-full" />

        <div class="flex flex-col gap-2">
            <USkeleton class="h-4 w-[250px]" />
            <USkeleton class="h-4 w-[200px]" />
        </div>
    </div>

    <div v-else>
        <TickerInfo
            :ticker="ticker.ticker"
            :company="ticker.company.companyName"
            :logo="ticker.company.image"
        />

        <div class="flex items-center gap-2">
            <p class="text-5xl mt-2 font-bold font-mono">{{ formatPrice(ticker.company.price) }}</p>

            <div class="flex items-center gap-2" :class="grow ? 'text-green-500' : 'text-red-500'">
                <Icon :name="grow ? 'lucide:arrow-up' : 'lucide:arrow-down'" class="w-4 h-4" />
                {{ formatPercentage(ticker.company.changePercentage) }}
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
const store = useTickerStore();

const ticker = computed(() => store.data);
const isLoading = computed(() => store.isLoading);

const grow = computed(() => ticker.value.company.changePercentage > 0);
</script>
