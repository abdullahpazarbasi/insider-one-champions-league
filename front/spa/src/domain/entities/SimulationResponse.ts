import type { ChampionChance } from "./ChampionChance";
import type { Standing } from "./Standing";
import type { Team } from "./Team";
import type { Week } from "./Week";

export interface SimulationResponse {
    teams: Team[];
    weeks: Week[];
    standings: Standing[];
    championChances?: ChampionChance[];
}
