<template>
    <div v-if="isLoading" class="flex flex-col gap-4">
        <USkeleton class="h-8 w-24" />
        <USkeleton class="h-44 w-full" />
        <USkeleton class="h-44 w-full" />
    </div>

    <div v-else class="flex flex-col gap-4">
        <h3 class="text-lg font-bold">Recommendations</h3>

        <UCollapsible class="bg-emerald-100 p-4 rounded-lg">
            <div class="flex items-center gap-2 cursor-pointer">
                <Icon name="lucide:sprout" />
                <span>Show AI resume</span>
            </div>

            <template #content>
                <div class="mt-2" v-html="advice"></div>
            </template>
        </UCollapsible>

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
const { $md } = useNuxtApp();

const recommendations = computed(() => store.data.recommendations);
const advice = computed(() => $md(store.data.advice));
const isLoading = computed(() => store.isLoading);
</script>
