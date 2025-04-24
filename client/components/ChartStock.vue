<template>
    <div ref="container" class="tradingview-widget-container" style="height: 100%; width: 100%">
        <div
            class="tradingview-widget-container__widget"
            style="height: calc(100% - 32px); width: 100%"
        ></div>
    </div>
</template>

<script setup lang="ts">
const props = defineProps<{
    symbol: string;
}>();

const container = ref<HTMLElement | null>(null);

onMounted(() => {
    const script = document.createElement('script');
    script.src = 'https://s3.tradingview.com/external-embedding/embed-widget-advanced-chart.js';
    script.type = 'text/javascript';
    script.async = true;
    script.innerHTML = `
        {
          "autosize": true,
          "symbol": "${props.symbol}USD",
          "interval": "D",
          "timezone": "Etc/UTC",
          "theme": "light",
          "style": "3",
          "locale": "en",
          "allow_symbol_change": true,
          "support_host": "https://www.tradingview.com"
        }`;

    if (container.value) {
        container.value.appendChild(script);
    }
});
</script>
