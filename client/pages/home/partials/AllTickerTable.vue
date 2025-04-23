<script setup lang="ts">
import type { TableRow } from '@nuxt/ui';
import type { Ticker } from '~/shared/types/models';
import { getColumns } from './columns';

const tickerData = useLoadTickerData();
const { tickers, loading, error, meta } = tickerData;

const columns = getColumns();

const handlePageChange = (newPage: number) => {
    meta.value.page = newPage;
};

const openTickerModal = (row: TableRow<Ticker>) => {
    navigateTo(`/ticker/${row.original.ticker}`);
};
</script>

<template>
    <div>
        <div v-if="error" class="p-4 text-red-500 text-center">
            {{ error }}
        </div>

        <div v-else>
            <UTable
                :loading="loading"
                sticky
                :data="tickers"
                :columns="columns"
                :meta="meta"
                @select="openTickerModal"
            />

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
