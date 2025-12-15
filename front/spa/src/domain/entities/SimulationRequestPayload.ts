import type { Team } from "./Team";
import type { Week } from "./Week";

export interface SimulationRequestPayload {
    teams: Team[];
    weeks: Week[];
    targetWeekIndex?: number;
}
