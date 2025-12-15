<template>
    <div v-if="week" class="matches">
        <div v-for="match in week.matches" :key="match.id" class="match-row">
            <div class="team-label">
                <span>{{ teamLookup[match.homeTeamId]?.name ?? "Unknown" }}</span>
                <input
                    type="number"
                    min="0"
                    max="99"
                    :value="match.homeScore ?? ''"
                    :disabled="loading"
                    @change="updateScore(match.id, 'homeScore', $event)"
                />
            </div>
            <span class="vs">-</span>
            <div class="team-label">
                <input
                    type="number"
                    min="0"
                    max="99"
                    :value="match.awayScore ?? ''"
                    :disabled="loading"
                    @change="updateScore(match.id, 'awayScore', $event)"
                />
                <span>{{ teamLookup[match.awayTeamId]?.name ?? "Unknown" }}</span>
            </div>
        </div>
    </div>
    <p v-else class="hint">Week not found.</p>
</template>

<script setup lang="ts">
import type { Team, Week } from "../domain/entities";

const props = defineProps<{
    week?: Week;
    teamLookup: Record<string, Team>;
    loading: boolean;
}>();

const emit = defineEmits<{
    (e: "change-score", matchId: string, side: "homeScore" | "awayScore", value: string): void;
}>();

function updateScore(matchId: string, side: "homeScore" | "awayScore", event: Event) {
    emit("change-score", matchId, side, (event.target as HTMLInputElement).value);
}
</script>
