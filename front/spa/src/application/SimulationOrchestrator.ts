import type { SimulationRequestPayload } from "../domain/entities/SimulationRequestPayload";
import type { SimulationResponse } from "../domain/entities/SimulationResponse";
import type { SimulationGateway } from "../domain/ports/SimulationGateway";
import { isWeekComplete } from "./helpers/isWeekComplete";
import type { SimulationState } from "./SimulationState";

interface SimulationOptions {
    targetWeekIndex?: number;
    moveToTargetWeek?: boolean;
}

export class SimulationOrchestrator {
    constructor(private readonly gateway: SimulationGateway) {}

    async bootstrap(): Promise<SimulationState> {
        const payload = await this.gateway.fetchBootstrap();
        return this.applyResponse(payload);
    }

    async simulate(state: SimulationState, options?: SimulationOptions): Promise<SimulationState> {
        const request: SimulationRequestPayload = {
            teams: state.teams,
            weeks: state.weeks,
            ...(typeof options?.targetWeekIndex === "number" ? { targetWeekIndex: options.targetWeekIndex } : {}),
        };
        const payload = await this.gateway.simulate(request);
        const nextState = this.applyResponse(payload);
        if (options?.moveToTargetWeek && typeof options.targetWeekIndex === "number") {
            return this.moveToWeek(nextState, Math.min(options.targetWeekIndex, nextState.weeks.length - 1));
        }
        return nextState;
    }

    async recalculateStandings(state: SimulationState): Promise<SimulationState> {
        const payload = await this.gateway.simulate({
            teams: state.teams,
            weeks: state.weeks,
        });
        return {
            ...state,
            standings: payload.standings,
            championChances: payload.championChances ?? [],
        };
    }

    updateScore(
        state: SimulationState,
        weekNumber: number,
        matchId: string,
        side: "homeScore" | "awayScore",
        value: string,
    ): SimulationState {
        const nextWeeks = state.weeks.map((week) => {
            if (week.weekNumber !== weekNumber) return week;
            return {
                ...week,
                matches: week.matches.map((match) => {
                    if (match.id !== matchId) return match;
                    const parsed = value === "" ? null : Number(value);
                    return {
                        ...match,
                        [side]: Number.isNaN(parsed) ? undefined : parsed,
                    };
                }),
            };
        });
        return { ...state, weeks: nextWeeks };
    }

    clearWeekFromIndex(state: SimulationState, fromIndex: number): SimulationState {
        const targetIndex = fromIndex - 1;
        if (targetIndex < 0) return state;

        const clearedWeeks = state.weeks.map((week, index) => {
            if (index < fromIndex) return week;
            return {
                ...week,
                matches: week.matches.map((match) => ({
                    ...match,
                    homeScore: undefined,
                    awayScore: undefined,
                })),
            };
        });

        return { ...state, weeks: clearedWeeks };
    }

    applyTeamEdit(state: SimulationState, teamId: string, name: string, strength: number): SimulationState {
        const updatedTeams = state.teams.map((team) =>
            team.id === teamId ? { ...team, name: name.trim(), strength } : team,
        );
        return { ...state, teams: updatedTeams };
    }

    playNextWeek(state: SimulationState): SimulationOptions {
        if (state.weeks.length === 0) return {};

        const targetIndex = Math.min(state.currentWeekIndex + 1, state.weeks.length - 1);
        if (this.needsSimulation(state, targetIndex)) {
            return { targetWeekIndex: targetIndex, moveToTargetWeek: true };
        }

        return { targetWeekIndex: targetIndex, moveToTargetWeek: false };
    }

    playAllWeeks(state: SimulationState): SimulationOptions {
        if (state.weeks.length === 0) return {};

        const targetIndex = Math.min(state.weeks.length - 1, 5);
        if (this.needsSimulation(state, targetIndex)) {
            return { targetWeekIndex: targetIndex, moveToTargetWeek: true };
        }

        return { targetWeekIndex: targetIndex, moveToTargetWeek: false };
    }

    moveToWeek(state: SimulationState, index: number): SimulationState {
        if (index < 0 || index >= state.weeks.length) return state;
        return { ...state, currentWeekIndex: index };
    }

    withCurrentWeekAfterSimulation(state: SimulationState): SimulationState {
        return { ...state, currentWeekIndex: this.findLastPlayedWeekIndex(state) };
    }

    shouldDisplayChampionChances(state: SimulationState): boolean {
        return state.championChances.length > 0 && isWeekComplete(state.weeks[3]);
    }

    allWeeksCompleted(state: SimulationState): boolean {
        return state.weeks.length > 0 && state.weeks.every((week) => isWeekComplete(week));
    }

    private applyResponse(payload: SimulationResponse): SimulationState {
        const baseState: SimulationState = {
            teams: payload.teams,
            weeks: payload.weeks,
            standings: payload.standings,
            championChances: payload.championChances ?? [],
            currentWeekIndex: 0,
        };
        return this.withCurrentWeekAfterSimulation(baseState);
    }

    private needsSimulation(state: SimulationState, targetIndex: number): boolean {
        return state.weeks.some((week, index) => index <= targetIndex && !isWeekComplete(week));
    }

    private findLastPlayedWeekIndex(state: SimulationState): number {
        const lastPlayed = state.weeks.reduce((last, week, index) => {
            const hasScore = week.matches.some((match) =>
                [match.homeScore, match.awayScore].every(
                    (value) => value !== undefined && value !== null,
                ),
            );
            return hasScore ? index : last;
        }, -1);

        return lastPlayed >= 0 ? lastPlayed : 0;
    }
}
