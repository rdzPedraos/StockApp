<template>
    <div class="grid grid-cols-[1fr_3fr] gap-4">
        <div class="relative h-[calc(100vh-10rem)] pr-2">
            <div
                ref="scrollContainer"
                class="overflow-y-scroll h-full beautify-scroll"
                @scroll="handleScroll"
            >
                <InfoHeader />
                <InfoBody class="mt-4" />
                <Recommendations class="mt-4" />
            </div>

            <div
                v-show="showScrollIndicator"
                class="h-10 absolute bottom-0 w-full bg-gradient-to-t from-white to-transparent pointer-events-none"
            ></div>
        </div>

        <ChartStock v-if="symbol" :symbol="symbol" />
    </div>
</template>

<script setup lang="ts">
import InfoHeader from './partials/InfoHeader.vue';
import InfoBody from './partials/InfoBody.vue';
import Recommendations from './partials/Recommendations.vue';

const showScrollIndicator = ref(true);
const scrollContainer = ref<HTMLElement | null>(null);

const route = useRoute();
const store = useTickerStore();
const symbol = computed(() => store.data.ticker);

// Función para manejar el evento de scroll
const handleScroll = () => {
    if (!scrollContainer.value) return;

    const { scrollTop, scrollHeight, clientHeight } = scrollContainer.value;

    // Calculamos si estamos cerca del final del scroll
    // Si la distancia al final es menor que un umbral (ej: 20px), ocultamos el indicador
    const isAtBottom = scrollHeight - scrollTop - clientHeight < 20;
    showScrollIndicator.value = !isAtBottom;
};

watchEffect(() => {
    const tickerParam = route.params.ticker as string;

    if (tickerParam) {
        store.fetchTickerData(tickerParam);
    }
});
</script>
