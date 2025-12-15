import type { Week } from "../../domain/entities/Week";

export function isWeekComplete(week?: Week): boolean {
    if (!week) return false;

    return week.matches.every(
        (match) =>
            match.homeScore !== undefined &&
            match.homeScore !== null &&
            match.awayScore !== undefined &&
            match.awayScore !== null,
    );
}
