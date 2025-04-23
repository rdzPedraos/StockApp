<script setup lang="ts">
import AllTickerTable from './partials/AllTickerTable.vue';
</script>

<template>
    <div>
        <section class="mb-4">
            <h1 class="text-2xl font-bold">Today Ticker Prices</h1>
            <span class="text-sm text-gray-500">Last updated: 10 minutes ago</span>
        </section>

        <AllTickerTable class="sticky top-0 overflow-visible" />
    </div>
</template>
