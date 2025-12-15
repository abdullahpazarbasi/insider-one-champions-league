import { describe, expect, it } from "vitest";
import { isWeekComplete } from "./isWeekComplete";

describe("isWeekComplete", () => {
    it("returns false for missing week", () => {
        expect(isWeekComplete()).toBe(false);
    });

    it("returns true only when all scores are present", () => {
        const week = {
            weekNumber: 1,
            matches: [
                { id: "m1", homeTeamId: "h1", awayTeamId: "a1", homeScore: 1, awayScore: 0 },
                { id: "m2", homeTeamId: "h2", awayTeamId: "a2", homeScore: 2, awayScore: 2 },
            ],
        };

        expect(isWeekComplete(week)).toBe(true);
    });
});
