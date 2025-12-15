<script setup lang="ts">
import { onMounted, computed } from "vue";
import ChampionChances from "./components/ChampionChances.vue";
import ErrorBanner from "./components/ErrorBanner.vue";
import MatchList from "./components/MatchList.vue";
import PageHeader from "./components/PageHeader.vue";
import StandingsTable from "./components/StandingsTable.vue";
import WeekNavigator from "./components/WeekNavigator.vue";
import { useSimulation } from "./composables/useSimulation";
import "./assets/app.css";

const {
    weeks,
    standings,
    championChances,
    currentWeekIndex,
    loading,
    errorMessage,
    editingTeamId,
    editingTeamDraft,
    teamLookup,
    currentWeek,
    allWeeksCompleted,
    shouldShowChampionChances,
    loadBootstrap,
    selectWeek,
    startEditingTeam,
    updateEditingTeamField,
    confirmTeamEdit,
    updateScore,
    clearWeekFromIndex,
    playNextWeek,
    playAllWeeks,
} = useSimulation();

const hasPreviousWeek = computed(() => currentWeekIndex.value > 0);
const hasNextWeek = computed(() => currentWeekIndex.value < weeks.value.length - 1);
const isLastWeek = computed(() => currentWeekIndex.value >= weeks.value.length - 1);
const nextWeekDisabled = computed(() => loading.value || (allWeeksCompleted.value && isLastWeek.value));
const playAllDisabled = computed(() => loading.value || allWeeksCompleted.value);
const hasPreviousPlayedWeek = computed(
    () =>
        currentWeekIndex.value > 0 &&
        weeks.value
            .slice(0, currentWeekIndex.value)
            .some((week) =>
                week.matches.some(
                    (match) =>
                        match.homeScore !== undefined &&
                        match.homeScore !== null &&
                        match.awayScore !== undefined &&
                        match.awayScore !== null,
                ),
            ),
);
const clearWeekDisabled = computed(() => loading.value || !hasPreviousPlayedWeek.value);

onMounted(() => {
    loadBootstrap();
});

function handleScoreChange(matchId: string, side: "homeScore" | "awayScore", value: string) {
    if (!currentWeek.value) return;
    updateScore(currentWeek.value.weekNumber, matchId, side, value);
}
</script>

<template>
    <main class="app-shell">
        <PageHeader
            title="Insider One Champions League"
            description="You can update team names and strengths."
            :loading="loading"
        />

        <ErrorBanner v-if="errorMessage" :message="errorMessage" />

        <section class="layout">
            <div class="panel">
                <StandingsTable
                    :standings="standings"
                    :team-lookup="teamLookup"
                    :editing-team-id="editingTeamId"
                    :editing-team-draft="editingTeamDraft"
                    :loading="loading"
                    @start-edit="startEditingTeam"
                    @edit-field="updateEditingTeamField"
                    @confirm-edit="confirmTeamEdit"
                />
            </div>

            <div class="panel">
                <WeekNavigator
                    :current-week-number="currentWeek?.weekNumber ?? 0"
                    :has-previous="hasPreviousWeek"
                    :has-next="hasNextWeek"
                    :loading="loading"
                    :simulate-next-disabled="nextWeekDisabled"
                    :simulate-all-disabled="playAllDisabled"
                    :clear-disabled="clearWeekDisabled"
                    @previous="selectWeek(-1)"
                    @next="selectWeek(1)"
                    @simulate-next="playNextWeek()"
                    @simulate-all="playAllWeeks()"
                    @clear-week="clearWeekFromIndex(currentWeekIndex)"
                />

                <MatchList
                    :week="currentWeek"
                    :team-lookup="teamLookup"
                    :loading="loading"
                    @change-score="handleScoreChange"
                />
            </div>

            <div v-if="shouldShowChampionChances" class="panel chances">
                <ChampionChances :chances="championChances" :team-lookup="teamLookup" />
            </div>
        </section>
    </main>
</template>
