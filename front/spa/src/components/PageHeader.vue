<template>
    <header class="page-header">
        <div>
            <!-- <p class="eyebrow">UEFA Champions League</p> -->
            <h1>{{ title }}</h1>
            <p class="lede">{{ description }}</p>
        </div>
    </header>
</template>

<script setup lang="ts">
defineProps<{ title: string; description: string; loading: boolean }>();
</script>
