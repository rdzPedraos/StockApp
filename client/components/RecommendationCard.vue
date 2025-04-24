<template>
    <div class="flex flex-col gap-4 py-2 px-4 rounded bg-slate-50">
        <small class="text-gray-500">{{ formatDate(r.time) }}</small>

        <div class="border-y-1 border-gray-200 py-4">
            <p class="uppercase font-bold mb-2" :class="getColor(r.rating_to)">
                {{ r.action }}
            </p>

            <div class="flex items-center gap-1">
                <div
                    :class="getBackground(r.rating_from)"
                    class="px-2 py-1 rounded min-w-24 text-center"
                >
                    <p>{{ r.rating_from_txt }}</p>
                    <p>{{ formatPrice(r.target_from) }}</p>
                </div>

                <Icon name="lucide:arrow-right" />

                <div
                    :class="getBackground(r.rating_to)"
                    class="px-2 py-1 rounded min-w-24 text-center"
                >
                    <p>{{ r.rating_to_txt }}</p>
                    <p>{{ formatPrice(r.target_to) }}</p>
                </div>
            </div>
        </div>

        <div class="flex items-center gap-1">
            <Icon name="lucide:shield-check" />
            {{ r.broker }}
        </div>
    </div>
</template>

<script setup lang="ts">
import type { Recommendation } from '~/shared/types/models';

const { recommendation: r } = defineProps<{
    recommendation: Recommendation;
}>();

const getColor = (rating: string) => {
    return {
        positive: 'text-green-500',
        negative: 'text-red-500',
        neutral: 'text-gray-500',
    }[rating];
};

const getBackground = (rating: string) => {
    return {
        positive: 'bg-green-200',
        negative: 'bg-red-200',
        neutral: 'bg-gray-200',
    }[rating];
};
</script>
