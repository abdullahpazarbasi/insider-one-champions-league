<template>
    <section class="error-banner" role="alert">{{ message }}</section>
</template>

<script setup lang="ts">
defineProps<{ message: string }>();
</script>
