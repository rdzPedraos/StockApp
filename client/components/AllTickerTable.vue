<script setup lang="ts">
import type { TableColumn } from '@nuxt/ui';
import type { Ticker } from '~/shared/types/models';

const items = useLoadTickerData();
const tickerCell = resolveComponent('TickerCell');
const UBadge = resolveComponent('UBadge');

const columns: TableColumn<Ticker>[] = [
    {
        header: '#',
        cell: ({ row }) => {
            const index: number = row.index + 1;

            return h('p', {}, index);
        },
    },
    {
        accessorKey: 'ticker',
        header: 'Ticker / Company',
        cell: ({ row }) => {
            const { company, logo, ticker } = row.original;

            return h(tickerCell, {
                ticker,
                company,
                logo,
            });
        },
    },
    {
        accessorKey: 'sentiment',
        header: 'Sentiment',
        cell: ({ row }) => {
            const sentiment: string = row.getValue('sentiment');

            const color = {
                buy: 'success',
                sell: 'error',
                hold: 'neutral',
            }[sentiment];

            return h(UBadge, { class: 'capitalize', variant: 'outline', color }, () => sentiment);
        },
    },
    {
        accessorKey: 'price',
        header: 'Price',
        cell: ({ row }) => {
            const price: number = row.getValue('price');

            const formattedPrice = new Intl.NumberFormat('en-US', {
                style: 'currency',
                currency: 'USD',
            }).format(price);

            return h('p', { class: 'font-bold' }, formattedPrice);
        },
    },
    {
        accessorKey: 'market_cap',
        header: 'Market Cap',
        cell: ({ row }) => {
            const marketCap: number = row.getValue('market_cap');

            const formattedMarketCap = new Intl.NumberFormat('en-US', {
                style: 'currency',
                currency: 'USD',
            }).format(marketCap);

            return h('p', { class: 'font-bold' }, formattedMarketCap);
        },
    },
    {
        accessorKey: 'last_review',
        header: 'Last Review',
        cell: ({ row }) => {
            const lastReview: string = row.getValue('last_review');

            return h('p', {}, lastReview);
        },
    },
];
</script>

<template>
    <UTable sticky :data="items" :columns="columns" />
</template>
