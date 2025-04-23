<script setup lang="ts">
const { ticker } = useRoute().params;
</script>

<template>
    <div>
        <h1>{{ ticker }}</h1>
    </div>
</template>
