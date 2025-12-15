<template>
    <div>
        <div class="panel-header">
            <h2>Teams and Form</h2>
            <p class="hint">Click a row to edit the name and strength.</p>
        </div>
        <div class="team-grid">
            <div
                v-for="team in teams"
                :key="team.id"
                class="team-row"
                :class="{ editing: editingTeamId === team.id }"
                @click="$emit('edit-team', team.id)"
            >
                <label class="field">
                    <span>Name</span>
                    <input
                        v-if="editingTeamId === team.id"
                        :value="team.name"
                        type="text"
                        @input="
                            $emit('update-name', team.id, ($event.target as HTMLInputElement).value)
                        "
                    />
                    <span v-else>{{ team.name }}</span>
                </label>
                <label class="field strength">
                    <span>Strength</span>
                    <input
                        v-if="editingTeamId === team.id"
                        :value="team.strength"
                        type="number"
                        min="0"
                        max="100"
                        @input="
                            $emit(
                                'update-strength',
                                team.id,
                                ($event.target as HTMLInputElement).value,
                            )
                        "
                    />
                    <span v-else>{{ team.strength }}</span>
                </label>
                <div class="chance">Championship: %{{ formatChance(team.id) }}</div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import type { ChampionChance, Team } from "../domain/entities";

const props = defineProps<{
    teams: Team[];
    editingTeamId: string | null;
    championChances: ChampionChance[];
}>();

function formatChance(teamId: string) {
    const value = props.championChances.find((chance) => chance.teamId === teamId)?.percentage ?? 0;
    return value.toFixed(2);
}
</script>
