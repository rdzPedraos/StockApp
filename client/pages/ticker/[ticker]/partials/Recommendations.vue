<template>
    <div v-if="isLoading" class="flex flex-col gap-4">
        <USkeleton class="h-8 w-24" />
        <USkeleton class="h-44 w-full" />
        <USkeleton class="h-44 w-full" />
    </div>

    <div v-else class="flex flex-col gap-4">
        <h3 class="text-lg font-bold">Recommendations</h3>

        <RecommendationCard
            v-for="recommendation in recommendations"
            :key="recommendation.time"
            class="border-1 border-gray-200"
            :recommendation="recommendation"
        />
    </div>
</template>

<script setup lang="ts">
const store = useTickerStore();

const recommendations = computed(() => store.data.recommendations);
const isLoading = computed(() => store.isLoading);
</script>
