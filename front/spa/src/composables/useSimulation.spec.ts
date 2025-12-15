import { describe, expect, it, vi } from "vitest";
import { useSimulation } from "./useSimulation";
import type { SimulationGateway } from "../domain/ports/SimulationGateway";
import type { SimulationResponse } from "../domain/entities";

const bootstrapResponse: SimulationResponse = {
    teams: [
        { id: "t1", name: "Team 1", strength: 90 },
        { id: "t2", name: "Team 2", strength: 80 },
    ],
    weeks: [
        {
            weekNumber: 1,
            matches: [
                { id: "m1", homeTeamId: "t1", awayTeamId: "t2" },
            ],
        },
    ],
    standings: [
        {
            teamId: "t1",
            name: "Team 1",
            played: 0,
            wins: 0,
            draws: 0,
            losses: 0,
            goalsFor: 0,
            goalsAgainst: 0,
            goalDifference: 0,
            points: 0,
        },
        {
            teamId: "t2",
            name: "Team 2",
            played: 0,
            wins: 0,
            draws: 0,
            losses: 0,
            goalsFor: 0,
            goalsAgainst: 0,
            goalDifference: 0,
            points: 0,
        },
    ],
    championChances: [
        { teamId: "t1", percentage: 50 },
        { teamId: "t2", percentage: 50 },
    ],
};

function createApi(overrides: Partial<SimulationGateway> = {}): SimulationGateway {
    const simulate = vi
        .fn()
        .mockResolvedValue({ ...bootstrapResponse, standings: [{ ...bootstrapResponse.standings[0], points: 3 }] });

    return {
        fetchBootstrap: vi.fn().mockResolvedValue(bootstrapResponse),
        simulate,
        ...overrides,
    };
}

describe("useSimulation", () => {
    it("loads bootstrap data and exposes lookups", async () => {
        const api = createApi();
        const simulation = useSimulation(api);

        await simulation.loadBootstrap();

        expect(api.fetchBootstrap).toHaveBeenCalled();
        expect(simulation.teams.value).toHaveLength(2);
        expect(simulation.teamLookup.value.t1.name).toBe("Team 1");
        expect(simulation.currentWeek.value?.weekNumber).toBe(1);
    });

    it("updates scores and recalculates standings", async () => {
        const api = createApi({
            simulate: vi.fn(async ({ weeks }) => ({
                ...bootstrapResponse,
                weeks,
                standings: [
                    { ...bootstrapResponse.standings[0], points: 3 },
                    bootstrapResponse.standings[1],
                ],
                championChances: bootstrapResponse.championChances,
            })),
        });
        const simulation = useSimulation(api);
        await simulation.loadBootstrap();

        await simulation.updateScore(1, "m1", "homeScore", "2");

        expect(api.simulate).toHaveBeenCalledWith(expect.objectContaining({ weeks: expect.any(Array) }));
        expect(simulation.standings.value[0].points).toBe(3);
        expect(simulation.weeks.value[0].matches[0].homeScore).toBe(2);
    });

    it("clears weeks from an index", async () => {
        const api = createApi({
            simulate: vi.fn().mockResolvedValue({
                ...bootstrapResponse,
                weeks: [
                    {
                        weekNumber: 1,
                        matches: [
                            {
                                id: "m1",
                                homeTeamId: "t1",
                                awayTeamId: "t2",
                                homeScore: undefined,
                                awayScore: undefined,
                            },
                        ],
                    },
                    {
                        weekNumber: 2,
                        matches: [
                            {
                                id: "m2",
                                homeTeamId: "t1",
                                awayTeamId: "t2",
                                homeScore: undefined,
                                awayScore: undefined,
                            },
                        ],
                    },
                ],
            }),
        });
        const simulation = useSimulation(api);
        await simulation.loadBootstrap();

        simulation.weeks.value[0].matches[0].homeScore = 1;
        simulation.weeks.value[0].matches[0].awayScore = 1;

        simulation.weeks.value.push({
            weekNumber: 2,
            matches: [
                { id: "m2", homeTeamId: "t1", awayTeamId: "t2", homeScore: 1, awayScore: 0 },
            ],
        });

        await simulation.clearWeekFromIndex(1);

        expect(simulation.weeks.value[1].matches[0].homeScore).toBeUndefined();
        expect(api.simulate).toHaveBeenCalledWith(expect.objectContaining({ targetWeekIndex: 0 }));
    });

    it("simulates remaining weeks up to target and moves to the latest", async () => {
        const api = createApi({
            simulate: vi.fn().mockResolvedValue({
                ...bootstrapResponse,
                weeks: [
                    bootstrapResponse.weeks[0],
                    { weekNumber: 2, matches: [] },
                ],
            }),
        });
        const simulation = useSimulation(api);
        await simulation.loadBootstrap();

        await simulation.simulate({ targetWeekIndex: 1, moveToTargetWeek: true });

        expect(api.simulate).toHaveBeenCalledWith(
            expect.objectContaining({ targetWeekIndex: 1, weeks: bootstrapResponse.weeks }),
        );
        expect(simulation.currentWeekIndex.value).toBe(1);
    });

    it("validates team edits and avoids unnecessary simulations", async () => {
        const simulate = vi.fn();
        const api = createApi({ simulate });
        const simulation = useSimulation(api);
        await simulation.loadBootstrap();

        simulation.startEditingTeam("t1");
        simulation.updateEditingTeamField("name", " ");
        await simulation.confirmTeamEdit();
        expect(simulation.errorMessage.value).toBe("Team name cannot be empty");

        simulation.startEditingTeam("t1");
        simulation.updateEditingTeamField("name", "Team 1");
        simulation.updateEditingTeamField("strength", "200");
        await simulation.confirmTeamEdit();
        expect(simulation.errorMessage.value).toBe("Strength must be between 0 and 100");

        simulation.startEditingTeam("t1");
        simulation.updateEditingTeamField("strength", String(bootstrapResponse.teams[0].strength));
        await simulation.confirmTeamEdit();
        expect(simulate).not.toHaveBeenCalled();
    });

    it("moves weeks without simulation when matches are complete", async () => {
        const api = createApi({ simulate: vi.fn() });
        const simulation = useSimulation(api);
        await simulation.loadBootstrap();

        await simulation.playAllWeeks();

        expect(simulation.currentWeekIndex.value).toBe(1);
        expect(api.simulate).not.toHaveBeenCalled();
    });
});
