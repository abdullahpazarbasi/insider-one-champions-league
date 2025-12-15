import type { SimulationState } from "./SimulationState";

export function createEmptyState(): SimulationState {
    return {
        teams: [],
        weeks: [],
        standings: [],
        championChances: [],
        currentWeekIndex: 0,
    };
}
