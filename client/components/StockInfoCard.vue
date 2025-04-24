<template>
    <div
        class="flex flex-col items-center justify-center border-1 border-gray-200 shadow p-2 rounded-sm"
    >
        <div class="flex items-center gap-2 text-sm text-gray-500">
            <small>{{ props.title }}</small>
            <UTooltip :text="props.info">
                <Icon name="lucide:info" class="w-4 h-4" />
            </UTooltip>
        </div>

        <p class="text-sm font-mono font-bold">
            {{ typeof props.value === 'number' ? formatPrice(props.value) : props.value }}
        </p>
    </div>
</template>

<script setup lang="ts">
const props = defineProps<{
    title: string;
    info: string;
    value: number | string;
}>();
</script>
