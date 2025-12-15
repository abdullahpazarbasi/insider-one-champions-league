import { mount, flushPromises } from "@vue/test-utils";
import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";
const bffBaseUrl = "http://example.test/api";

type MockResponse = {
    teams: { id: string; name: string; strength: number }[];
    weeks: {
        weekNumber: number;
        matches: {
            id: string;
            homeTeamId: string;
            awayTeamId: string;
            homeScore?: number;
            awayScore?: number;
        }[];
    }[];
    standings: {
        teamId: string;
        name: string;
        played: number;
        wins: number;
        draws: number;
        losses: number;
        goalsFor: number;
        goalsAgainst: number;
        goalDifference: number;
        points: number;
    }[];
    championChances?: { teamId: string; percentage: number }[];
};

const bootstrapPayload: MockResponse = {
    teams: [
        { id: "t1", name: "Real Madrid", strength: 95 },
        { id: "t2", name: "Bayern", strength: 92 },
        { id: "t3", name: "City", strength: 94 },
        { id: "t4", name: "Inter", strength: 89 },
    ],
    weeks: [
        {
            weekNumber: 1,
            matches: [
                { id: "m1", homeTeamId: "t1", awayTeamId: "t2", homeScore: 2, awayScore: 1 },
                { id: "m2", homeTeamId: "t3", awayTeamId: "t4", homeScore: 0, awayScore: 0 },
            ],
        },
        {
            weekNumber: 2,
            matches: [
                { id: "m3", homeTeamId: "t1", awayTeamId: "t3" },
                { id: "m4", homeTeamId: "t2", awayTeamId: "t4" },
            ],
        },
    ],
    standings: [
        {
            teamId: "t1",
            name: "Real Madrid",
            played: 1,
            wins: 1,
            draws: 0,
            losses: 0,
            goalsFor: 2,
            goalsAgainst: 1,
            goalDifference: 1,
            points: 3,
        },
        {
            teamId: "t2",
            name: "Bayern",
            played: 1,
            wins: 0,
            draws: 0,
            losses: 1,
            goalsFor: 1,
            goalsAgainst: 2,
            goalDifference: -1,
            points: 0,
        },
        {
            teamId: "t3",
            name: "City",
            played: 1,
            wins: 0,
            draws: 1,
            losses: 0,
            goalsFor: 0,
            goalsAgainst: 0,
            goalDifference: 0,
            points: 1,
        },
        {
            teamId: "t4",
            name: "Inter",
            played: 1,
            wins: 0,
            draws: 1,
            losses: 0,
            goalsFor: 0,
            goalsAgainst: 0,
            goalDifference: 0,
            points: 1,
        },
    ],
};

let fetchMock: ReturnType<typeof vi.fn>;

beforeEach(() => {
    vi.stubEnv("VITE_BFF_BASE_URL", bffBaseUrl);
    fetchMock = vi.fn((input: RequestInfo | URL, init?: RequestInit) => {
        const url = typeof input === "string" ? input : input.toString();
        if (url.includes("/api/bootstrap")) {
            return Promise.resolve(new Response(JSON.stringify(bootstrapPayload), { status: 200 }));
        }

        if (url.includes("/api/simulate")) {
            const body = init?.body ? JSON.parse(init.body.toString()) : {};
            const baseChances = bootstrapPayload.teams.map((team) => ({
                teamId: team.id,
                percentage: 25,
            }));
            const includeChances = typeof body.targetWeekIndex === "number" && body.targetWeekIndex >= 3;
            const updated: MockResponse = {
                ...bootstrapPayload,
                weeks: bootstrapPayload.weeks.map((week) => ({
                    ...week,
                    matches: week.matches.map((match, index) => ({
                        ...match,
                        homeScore: match.homeScore ?? index,
                        awayScore: match.awayScore ?? index + 1,
                    })),
                })),
                standings: bootstrapPayload.standings.map((row) => ({ ...row, played: 1 })),
                ...(includeChances
                    ? {
                          championChances: baseChances.map((chance) => ({
                              ...chance,
                              percentage:
                                  chance.percentage +
                                  (body.targetWeekIndex === bootstrapPayload.weeks.length - 1 ? 0.5 : 0),
                          })),
                      }
                    : {}),
            };
            return Promise.resolve(new Response(JSON.stringify(updated), { status: 200 }));
        }

        return Promise.reject(new Error("unknown endpoint"));
    });

    vi.stubGlobal("fetch", fetchMock);
});

afterEach(() => {
    vi.unstubAllEnvs();
    vi.restoreAllMocks();
    vi.resetModules();
});

describe("App", () => {
    it("renders bootstrap data", async () => {
        const App = (await import("./App.vue")).default;
        const wrapper = mount(App);
        await flushPromises();

        expect(fetchMock).toHaveBeenCalledWith(`${bffBaseUrl}/bootstrap`, undefined);
        expect(wrapper.text()).toContain("Insider One Champions League");
        expect(wrapper.text()).toContain("Real Madrid");
        expect(wrapper.text()).toContain("Standings");
    });

    it("hides champion chances when they are not provided", async () => {
        const App = (await import("./App.vue")).default;
        const wrapper = mount(App);
        await flushPromises();

        expect(wrapper.find(".panel.chances").exists()).toBe(false);
    });

    it("requests simulation when Next Week is clicked", async () => {
        const App = (await import("./App.vue")).default;
        const wrapper = mount(App);
        await flushPromises();

        const button = wrapper.get("button.secondary");
        await button.trigger("click");
        await flushPromises();

        const [, simulateCall] = fetchMock.mock.calls;
        const [simulateUrl, simulateInit] = simulateCall;
        expect(simulateUrl).toBe(`${bffBaseUrl}/simulate`);
        expect(JSON.parse((simulateInit as RequestInit).body as string)).toMatchObject({
            targetWeekIndex: 1,
        });
    });

    it("advances to the next week after simulating", async () => {
        const App = (await import("./App.vue")).default;
        const wrapper = mount(App);
        await flushPromises();

        const weekLabel = () => wrapper.get(".controls span").text();
        expect(weekLabel()).toContain("Week 1");

        const button = wrapper.get("button.secondary");
        await button.trigger("click");
        await flushPromises();

        expect(weekLabel()).toContain("Week 2");
    });

    it("skips simulation when current and next weeks are already played", async () => {
        const fullyPlayedPayload: MockResponse = {
            ...bootstrapPayload,
            weeks: bootstrapPayload.weeks.map((week) => ({
                ...week,
                matches: week.matches.map((match, index) => ({
                    ...match,
                    homeScore: match.homeScore ?? index,
                    awayScore: match.awayScore ?? index + 1,
                })),
            })),
        };

        fetchMock.mockImplementation((input: RequestInfo | URL) => {
            const url = typeof input === "string" ? input : input.toString();
            if (url.includes("/api/bootstrap")) {
                return Promise.resolve(new Response(JSON.stringify(fullyPlayedPayload), { status: 200 }));
            }

            if (url.includes("/api/simulate")) {
                return Promise.reject(new Error("simulate should not be called"));
            }

            return Promise.reject(new Error("unknown endpoint"));
        });

        const App = (await import("./App.vue")).default;
        const wrapper = mount(App);
        await flushPromises();

        const [previousButton] = wrapper.findAll(".controls button");
        await previousButton.trigger("click");
        expect(wrapper.get(".controls span").text()).toContain("Week 1");

        const nextWeekButton = wrapper.get("button.secondary");
        await nextWeekButton.trigger("click");
        await flushPromises();

        expect(wrapper.get(".controls span").text()).toContain("Week 2");
        expect(fetchMock).toHaveBeenCalledTimes(1);
    });

    it("disables next and play-all actions when all weeks are completed and on the last week", async () => {
        const fullyPlayedPayload: MockResponse = {
            ...bootstrapPayload,
            weeks: bootstrapPayload.weeks.map((week) => ({
                ...week,
                matches: week.matches.map((match, index) => ({
                    ...match,
                    homeScore: match.homeScore ?? index,
                    awayScore: match.awayScore ?? index + 1,
                })),
            })),
        };

        fetchMock.mockImplementation((input: RequestInfo | URL) => {
            const url = typeof input === "string" ? input : input.toString();
            if (url.includes("/api/bootstrap")) {
                return Promise.resolve(new Response(JSON.stringify(fullyPlayedPayload), { status: 200 }));
            }

            if (url.includes("/api/simulate")) {
                return Promise.reject(new Error("simulate should not be called"));
            }

            return Promise.reject(new Error("unknown endpoint"));
        });

        const App = (await import("./App.vue")).default;
        const wrapper = mount(App);
        await flushPromises();

        const playAllButton = wrapper.get("button.primary");
        const nextWeekButton = wrapper.get("button.secondary");

        expect(playAllButton.attributes("disabled")).toBeDefined();
        expect(nextWeekButton.attributes("disabled")).toBeDefined();
    });

    it("disables clear when no previous week is available and clears to the previous week when possible", async () => {
        const App = (await import("./App.vue")).default;
        const wrapper = mount(App);
        await flushPromises();

        const actionButtons = wrapper.findAll(".week-actions button");
        const clearButton = actionButtons[2];
        expect(clearButton.attributes("disabled")).toBeDefined();

        const nextWeekButton = actionButtons[0];
        await nextWeekButton.trigger("click");
        await flushPromises();

        expect(clearButton.attributes("disabled")).toBeUndefined();

        await clearButton.trigger("click");
        await flushPromises();

        const weekLabel = wrapper.get(".controls span").text();
        expect(weekLabel).toContain("Week 1");

        const simulateCalls = fetchMock.mock.calls.filter(([url]) =>
            url.toString().includes("/api/simulate"),
        );
        const lastCallBody = JSON.parse((simulateCalls.at(-1)?.[1] as RequestInit).body as string);
        expect(lastCallBody.targetWeekIndex).toBe(0);
    });
});
