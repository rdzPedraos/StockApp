import { ref, watch } from 'vue';
import type { Meta, Ticker } from '~/shared/types/models';
import api from '~/utils/axios';

export function useLoadTickerData() {
    const loading = ref<boolean>(true);
    const error = ref<string | null>(null);
    const tickers = ref<Ticker[]>([]);
    const meta = ref<Meta>({
        page: 1,
        total: 0,
        last_page: 0,
        per_page: 10,
    });

    const loadData = async () => {
        loading.value = true;
        error.value = null;

        const { page, per_page } = meta.value;

        try {
            const response = await api.get('/tickers', { params: { page, per_page } });
            const { data, meta: responseMeta } = response.data;

            tickers.value = data;
            meta.value.total = responseMeta.total;
            meta.value.last_page = responseMeta.last_page;
        } catch (err) {
            console.error('Error al cargar los tickers:', err);
            error.value = 'No se pudieron cargar los datos de los tickers';
        } finally {
            loading.value = false;
        }
    };

    watch(() => [meta.value.page, meta.value.per_page], loadData);
    loadData();

    return {
        tickers,
        loading,
        error,
        meta,
    };
}
