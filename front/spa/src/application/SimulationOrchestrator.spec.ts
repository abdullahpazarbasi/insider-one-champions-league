import { describe, expect, it } from "vitest";
import { SimulationOrchestrator } from "./SimulationOrchestrator";
import { createEmptyState } from "./createEmptyState";
import type { SimulationGateway } from "../domain/ports/SimulationGateway";
import type { SimulationResponse } from "../domain/entities";

const bootstrapPayload: SimulationResponse = {
    teams: [
        { id: "t1", name: "Team 1", strength: 90 },
        { id: "t2", name: "Team 2", strength: 80 },
    ],
    weeks: [
        {
            weekNumber: 1,
            matches: [
                { id: "m1", homeTeamId: "t1", awayTeamId: "t2", homeScore: 1, awayScore: 0 },
            ],
        },
        {
            weekNumber: 2,
            matches: [{ id: "m2", homeTeamId: "t2", awayTeamId: "t1" }],
        },
    ],
    standings: [
        {
            teamId: "t1",
            name: "Team 1",
            played: 1,
            wins: 1,
            draws: 0,
            losses: 0,
            goalsFor: 1,
            goalsAgainst: 0,
            goalDifference: 1,
            points: 3,
        },
        {
            teamId: "t2",
            name: "Team 2",
            played: 1,
            wins: 0,
            draws: 0,
            losses: 1,
            goalsFor: 0,
            goalsAgainst: 1,
            goalDifference: -1,
            points: 0,
        },
    ],
    championChances: [
        { teamId: "t1", percentage: 60 },
        { teamId: "t2", percentage: 40 },
    ],
};

function createGateway(overrides: Partial<SimulationGateway> = {}): SimulationGateway {
    return {
        fetchBootstrap: async () => bootstrapPayload,
        simulate: async () => bootstrapPayload,
        ...overrides,
    };
}

describe("SimulationOrchestrator", () => {
    it("bootstraps state and moves to last played week", async () => {
        const gateway = createGateway();
        const orchestrator = new SimulationOrchestrator(gateway);

        const state = await orchestrator.bootstrap();

        expect(state.currentWeekIndex).toBe(0);
        expect(state.standings[0].points).toBe(3);
        expect(orchestrator.shouldDisplayChampionChances(state)).toBe(false);
    });

    it("simulates and moves to requested week", async () => {
        const gateway = createGateway({ simulate: async () => bootstrapPayload });
        const orchestrator = new SimulationOrchestrator(gateway);

        const next = await orchestrator.simulate(createEmptyState(), {
            targetWeekIndex: 1,
            moveToTargetWeek: true,
        });

        expect(next.currentWeekIndex).toBe(1);
        expect(next.weeks).toHaveLength(2);
    });

    it("applies edits and requests recalculation options", () => {
        const gateway = createGateway();
        const orchestrator = new SimulationOrchestrator(gateway);

        const base = {
            ...bootstrapPayload,
            championChances: bootstrapPayload.championChances ?? [],
            currentWeekIndex: 0,
        };
        const edited = orchestrator.applyTeamEdit(base, "t1", "  New Name  ", 75);

        expect(edited.teams[0].name).toBe("New Name");
        expect(orchestrator.playNextWeek(base).targetWeekIndex).toBe(1);
        expect(orchestrator.playAllWeeks(base).targetWeekIndex).toBe(1);
    });

    it("clears weeks and recalculates derived flags", () => {
        const gateway = createGateway();
        const orchestrator = new SimulationOrchestrator(gateway);

        const cleared = orchestrator.clearWeekFromIndex(bootstrapPayload, 1);

        expect(cleared.weeks[1].matches[0].homeScore).toBeUndefined();
        expect(orchestrator.allWeeksCompleted(cleared)).toBe(false);
    });
});
