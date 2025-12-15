<template>
    <div class="panel-header">
        <h2>Championship Odds</h2>
        <!-- <p class="hint">Server prediction</p> -->
    </div>
    <ul class="chance-list">
        <li v-for="chance in chances" :key="chance.teamId">
            <div class="chance-team">{{ teamLookup[chance.teamId]?.name ?? chance.teamId }}</div>
            <div class="chance-value">%{{ chance.percentage.toFixed(2) }}</div>
        </li>
    </ul>
</template>

<script setup lang="ts">
import type { ChampionChance, Team } from "../domain/entities";

defineProps<{
    chances: ChampionChance[];
    teamLookup: Record<string, Team>;
}>();
</script>
