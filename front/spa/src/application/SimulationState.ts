import type { ChampionChance } from "../domain/entities/ChampionChance";
import type { Standing } from "../domain/entities/Standing";
import type { Team } from "../domain/entities/Team";
import type { Week } from "../domain/entities/Week";

export interface SimulationState {
    teams: Team[];
    weeks: Week[];
    standings: Standing[];
    championChances: ChampionChance[];
    currentWeekIndex: number;
}
