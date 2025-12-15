<template>
    <div class="standings">
        <div class="panel-header">
            <h2>Standings</h2>
            <p class="hint">Edit team names and strengths inline by clicking a team.</p>
        </div>
        <table>
            <thead>
                <tr>
                    <th>Team</th>
                    <th>PTS</th>
                    <th>P</th>
                    <th>W</th>
                    <th>D</th>
                    <th>L</th>
                    <th>GD</th>
                </tr>
            </thead>
            <tbody>
                <tr
                    v-for="row in standings"
                    :key="row.teamId"
                    :class="{ editing: editingTeamId === row.teamId }"
                >
                    <td class="team-cell">
                        <div v-if="editingTeamId === row.teamId" class="team-edit-form">
                            <input
                                type="text"
                                name="team-name"
                                autocomplete="off"
                                :value="editingDraft(row.teamId)?.name ?? ''"
                                :disabled="loading"
                                @input="
                                    $emit(
                                        'edit-field',
                                        'name',
                                        ($event.target as HTMLInputElement).value,
                                    )
                                "
                            />
                            <input
                                type="number"
                                min="0"
                                max="100"
                                name="team-strength"
                                :value="editingDraft(row.teamId)?.strength ?? ''"
                                :disabled="loading"
                                @input="
                                    $emit(
                                        'edit-field',
                                        'strength',
                                        ($event.target as HTMLInputElement).value,
                                    )
                                "
                            />
                            <button
                                type="button"
                                class="primary ok-button"
                                :disabled="loading"
                                @click="$emit('confirm-edit')"
                            >
                                OK
                            </button>
                        </div>
                        <button
                            v-else
                            type="button"
                            class="link-button team-name"
                            :disabled="loading"
                            @click="$emit('start-edit', row.teamId)"
                        >
                            {{ row.name }}
                        </button>
                    </td>
                    <td>{{ row.points }}</td>
                    <td>{{ row.played }}</td>
                    <td>{{ row.wins }}</td>
                    <td>{{ row.draws }}</td>
                    <td>{{ row.losses }}</td>
                    <td>{{ row.goalDifference }}</td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script setup lang="ts">
import type { Standing, Team } from "../domain/entities";

const props = defineProps<{
    standings: Standing[];
    teamLookup: Record<string, Team>;
    editingTeamId: string | null;
    editingTeamDraft: { name: string; strength: string } | null;
    loading: boolean;
}>();

defineEmits<{
    (e: "start-edit", teamId: string): void;
    (e: "edit-field", field: "name" | "strength", value: string): void;
    (e: "confirm-edit"): void;
}>();

function editingDraft(teamId: string) {
    if (props.editingTeamId !== teamId) return null;
    return (
        props.editingTeamDraft ?? {
            name: props.teamLookup[teamId]?.name ?? "",
            strength: String(props.teamLookup[teamId]?.strength ?? ""),
        }
    );
}
</script>
