import { defineStore } from 'pinia';
import type { CompanyData, Recommendation } from '~/shared/types/models';
import api from '~/utils/axios';

type TickerData = {
    ticker: string;
    company: CompanyData;
    recommendations: Recommendation[];
    advice: string;
};

export const useTickerStore = defineStore('ticker', () => {
    const data = ref<TickerData>({} as TickerData);
    const isLoading = ref(true);
    const error = ref<string | null>(null);

    // Acción para cargar los datos
    async function fetchTickerData(ticker: string) {
        isLoading.value = true;
        error.value = null;

        try {
            const response = await api.get(`/tickers/${ticker}`);
            data.value = response.data;
        } catch (err) {
            error.value = err instanceof Error ? err.message : 'Error al cargar los datos';
            console.error('Error al cargar los datos del ticker:', err);
        }

        isLoading.value = false;
    }

    return {
        data,
        isLoading,
        error,
        fetchTickerData,
    };
});
