<template>
    <div>
        <div class="flex justify-between items-center mb-4">
            <UInput
                v-model="search"
                class="mb-4"
                icon="i-lucide-search"
                placeholder="Search ticker"
                size="xl"
                :loading="loading"
                :disabled="loading"
            />

            <USelect
                v-model="meta.per_page"
                size="xl"
                :items="[8, 10, 20, 30]"
                class="w-24"
                :disabled="loading"
            />
        </div>

        <LoadingCells v-if="loading" />

        <div v-else>
            <UTable
                :loading="loading"
                sticky
                :data="tickers"
                :columns="columns"
                :meta="meta"
                @select="openTickerModal"
            />

            <div v-if="error" class="p-4 text-red-500 text-center">
                {{ error }}
            </div>

            <div class="flex justify-center mt-4">
                <UPagination
                    size="lg"
                    :disabled="loading"
                    :default-page="meta.page"
                    :items-per-page="meta.per_page"
                    :total="meta.total"
                    @update:page="handlePageChange"
                />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import type { TableRow } from '@nuxt/ui';
import type { Ticker } from '~/shared/types/models';
import { getColumns } from './columns';
import LoadingCells from './LoadingCells.vue';
const tickerData = useLoadTickerData();
const { tickers, loading, search, meta, error, loadData } = tickerData;

const columns = getColumns();

const handlePageChange = (newPage: number) => {
    meta.value.page = newPage;
};

const openTickerModal = (row: TableRow<Ticker>) => {
    navigateTo(`/ticker/${row.original.ticker}`);
};

onMounted(loadData);
</script>
