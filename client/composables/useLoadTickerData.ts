import { ref, watch } from 'vue';
import type { Meta, Ticker } from '~/shared/types/models';
import api from '~/utils/axios';

export function useLoadTickerData() {
    const tickers = ref<Ticker[]>([]);
    const loading = ref<boolean>(true);
    const error = ref<string | null>(null);
    const search = ref<string>('');
    const meta = ref<Meta>({
        page: 1,
        total: 0,
        last_page: 0,
        per_page: 8,
    });

    const loadData = debounce(() => {
        loading.value = true;
        error.value = null;

        const { page, per_page } = meta.value;

        api.get('/tickers', { params: { page, per_page, search: search.value } })
            .then((response) => {
                const { data, meta: responseMeta } = response.data;

                tickers.value = data;
                meta.value.total = responseMeta.total;
                meta.value.last_page = responseMeta.last_page;
            })
            .catch(() => {
                error.value = 'No se pudieron cargar los datos de los tickers';
            })
            .finally(() => {
                loading.value = false;
            });
    }, 600);

    watch(() => [meta.value.page, meta.value.per_page, search.value], loadData);

    return {
        tickers,
        loading,
        error,
        meta,
        search,
        loadData,
    };
}
