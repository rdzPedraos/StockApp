import type { TableColumn } from '@nuxt/ui';
import type { Ticker } from '~/shared/types/models';

import { h } from 'vue';
import { UBadge } from '#components';
import TickerCell from '~/components/TickerInfo.vue';

export const getColumns = (): TableColumn<Ticker>[] => {
    return [
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
                const { company, ticker } = row.original;

                return h(TickerCell, {
                    ticker,
                    company: company.companyName,
                    logo: company.image,
                });
            },
        },
        {
            accessorKey: 'sentiment',
            header: 'Sentiment',
            cell: ({ row }) => {
                const sentiment: string = row.getValue('sentiment');

                const color = {
                    positive: 'success',
                    negative: 'error',
                    neutral: 'neutral',
                }[sentiment];

                return h(
                    UBadge,
                    {
                        class: 'capitalize',
                        variant: 'outline',
                        color: color as 'success' | 'error' | 'neutral',
                    },
                    () => sentiment
                );
            },
        },
        {
            accessorKey: 'price',
            header: 'Price',
            cell: ({ row }) => {
                const price: number = row.original.company.price;

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
                const marketCap: number = row.original.company.marketCap;

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

                const formattedLastReview = new Date(lastReview).toLocaleDateString('en-US', {
                    day: '2-digit',
                    month: '2-digit',
                    year: 'numeric',
                    hour: '2-digit',
                    minute: '2-digit',
                    second: '2-digit',
                });

                return h('p', {}, formattedLastReview);
            },
        },
    ];
};
