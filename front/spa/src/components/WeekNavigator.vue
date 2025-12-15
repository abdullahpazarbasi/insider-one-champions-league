<template>
    <div>
        <div class="panel-header">
            <h2>Matchweeks</h2>
            <div class="controls">
                <button type="button" :disabled="!hasPrevious" @click="$emit('previous')">⟵</button>
                <span>Week {{ currentWeekNumber }}</span>
                <button type="button" :disabled="!hasNext" @click="$emit('next')">⟶</button>
            </div>
        </div>

        <div class="week-actions">
            <button class="secondary" :disabled="simulateNextDisabled" @click="$emit('simulate-next')">
                Next Week
            </button>
            <button class="primary" :disabled="simulateAllDisabled" @click="$emit('simulate-all')">
                Play All
            </button>
            <button class="danger" :disabled="clearDisabled || loading" @click="$emit('clear-week')">
                Clear Week
            </button>
        </div>
    </div>
</template>

<script setup lang="ts">
defineProps<{
    currentWeekNumber: number;
    hasPrevious: boolean;
    hasNext: boolean;
    loading: boolean;
    simulateNextDisabled: boolean;
    simulateAllDisabled: boolean;
    clearDisabled: boolean;
}>();
</script>
