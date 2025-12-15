import { computed, ref } from "vue";
import { createEmptyState } from "../application/createEmptyState";
import { SimulationOrchestrator } from "../application/SimulationOrchestrator";
import type { SimulationState } from "../application/SimulationState";
import type { SimulationGateway } from "../domain/ports/SimulationGateway";
import { HttpSimulationGateway } from "../infrastructure/gateways/HttpSimulationGateway";
import type { ChampionChance } from "../domain/entities/ChampionChance";
import type { Team } from "../domain/entities/Team";
import type { Week } from "../domain/entities/Week";

export interface UseSimulationState {
    teams: ReturnType<typeof computed<Team[]>>;
    weeks: ReturnType<typeof computed<Week[]>>;
    standings: ReturnType<typeof computed<SimulationState["standings"]>>;
    championChances: ReturnType<typeof computed<ChampionChance[]>>;
    currentWeekIndex: ReturnType<typeof computed<number>>;
    loading: ReturnType<typeof ref<boolean>>;
    errorMessage: ReturnType<typeof ref<string>>;
    editingTeamId: ReturnType<typeof ref<string | null>>;
    editingTeamDraft: ReturnType<typeof ref<{ name: string; strength: string } | null>>;
    teamLookup: ReturnType<typeof computed<Record<string, Team>>>;
    currentWeek: ReturnType<typeof computed<Week | undefined>>;
    allWeeksCompleted: ReturnType<typeof computed<boolean>>;
    shouldShowChampionChances: ReturnType<typeof computed<boolean>>;
    loadBootstrap: () => Promise<void>;
    selectWeek: (direction: -1 | 1) => void;
    startEditingTeam: (teamId: string) => void;
    updateEditingTeamField: (field: "name" | "strength", value: string) => void;
    confirmTeamEdit: () => Promise<void>;
    updateScore: (
        weekNumber: number,
        matchId: string,
        side: "homeScore" | "awayScore",
        value: string,
    ) => Promise<void>;
    clearWeekFromIndex: (fromIndex: number) => Promise<void>;
    simulate: (options?: { targetWeekIndex?: number; moveToTargetWeek?: boolean }) => Promise<void>;
    playNextWeek: () => Promise<void>;
    playAllWeeks: () => Promise<void>;
    recalculateStandings: () => Promise<void>;
}

export function useSimulation(gateway: SimulationGateway = new HttpSimulationGateway()): UseSimulationState {
    const orchestrator = new SimulationOrchestrator(gateway);
    const state = ref<SimulationState>(createEmptyState());
    const loading = ref(false);
    const errorMessage = ref("");
    const editingTeamId = ref<string | null>(null);
    const editingTeamDraft = ref<{ name: string; strength: string } | null>(null);

    const teams = computed(() => state.value.teams);
    const weeks = computed(() => state.value.weeks);
    const standings = computed(() => state.value.standings);
    const championChances = computed(() => state.value.championChances);
    const currentWeekIndex = computed(() => state.value.currentWeekIndex);

    const teamLookup = computed<Record<string, Team>>(() =>
        teams.value.reduce((acc, team) => ({ ...acc, [team.id]: team }), {} as Record<string, Team>),
    );

    const currentWeek = computed(() => weeks.value[currentWeekIndex.value]);
    const allWeeksCompleted = computed(() => orchestrator.allWeeksCompleted(state.value));
    const shouldShowChampionChances = computed(() => orchestrator.shouldDisplayChampionChances(state.value));

    async function loadBootstrap() {
        await runSafely(async () => {
            state.value = await orchestrator.bootstrap();
        });
    }

    function selectWeek(direction: -1 | 1) {
        state.value = orchestrator.moveToWeek(state.value, state.value.currentWeekIndex + direction);
    }

    function startEditingTeam(teamId: string) {
        const team = teamLookup.value[teamId];
        if (!team) return;
        editingTeamId.value = teamId;
        editingTeamDraft.value = { name: team.name, strength: String(team.strength) };
    }

    function updateEditingTeamField(field: "name" | "strength", value: string) {
        if (!editingTeamDraft.value) return;
        editingTeamDraft.value = { ...editingTeamDraft.value, [field]: value };
    }

    function resetEditingState() {
        editingTeamId.value = null;
        editingTeamDraft.value = null;
    }

    async function confirmTeamEdit() {
        if (!editingTeamId.value || !editingTeamDraft.value) return;

        const current = teamLookup.value[editingTeamId.value];
        if (!current) {
            resetEditingState();
            return;
        }

        const nextName = editingTeamDraft.value.name.trim();
        const parsedStrength = Number(editingTeamDraft.value.strength);

        if (!nextName) {
            errorMessage.value = "Team name cannot be empty";
            return;
        }

        if (!Number.isFinite(parsedStrength) || parsedStrength < 0 || parsedStrength > 100) {
            errorMessage.value = "Strength must be between 0 and 100";
            return;
        }

        const updatedState = orchestrator.applyTeamEdit(state.value, editingTeamId.value, nextName, parsedStrength);
        const hasChanged =
            nextName !== current.name || parsedStrength !== current.strength;

        state.value = updatedState;
        resetEditingState();

        if (!hasChanged) return;

        await simulate({ targetWeekIndex: 0, moveToTargetWeek: true });
    }

    async function updateScore(
        weekNumber: number,
        matchId: string,
        side: "homeScore" | "awayScore",
        value: string,
    ) {
        const updated = orchestrator.updateScore(state.value, weekNumber, matchId, side, value);
        state.value = updated;
        await recalculateStandings();
    }

    async function clearWeekFromIndex(fromIndex: number) {
        const cleared = orchestrator.clearWeekFromIndex(state.value, fromIndex);
        if (cleared === state.value) return;
        state.value = cleared;
        await simulate({ targetWeekIndex: fromIndex - 1, moveToTargetWeek: true });
    }

    async function simulate(options?: { targetWeekIndex?: number; moveToTargetWeek?: boolean }) {
        await runSafely(async () => {
            state.value = await orchestrator.simulate(state.value, options);
        });
    }

    async function playNextWeek() {
        const options = orchestrator.playNextWeek(state.value);
        if (typeof options.targetWeekIndex !== "number") return;
        if (options.moveToTargetWeek) {
            await simulate(options);
            return;
        }
        state.value = orchestrator.moveToWeek(state.value, options.targetWeekIndex);
    }

    async function playAllWeeks() {
        const options = orchestrator.playAllWeeks(state.value);
        if (typeof options.targetWeekIndex !== "number") return;
        if (options.moveToTargetWeek) {
            await simulate(options);
            return;
        }
        state.value = orchestrator.moveToWeek(state.value, options.targetWeekIndex);
    }

    async function recalculateStandings() {
        await runSafely(async () => {
            state.value = await orchestrator.recalculateStandings(state.value);
        });
    }

    async function runSafely(operation: () => Promise<void>) {
        loading.value = true;
        errorMessage.value = "";
        try {
            await operation();
        } catch (err) {
            errorMessage.value = (err as Error).message;
        } finally {
            loading.value = false;
        }
    }

    return {
        teams,
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
        simulate,
        playNextWeek,
        playAllWeeks,
        recalculateStandings,
    };
}
