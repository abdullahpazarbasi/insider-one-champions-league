package simulation

import "github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"

type ScheduleGenerator interface {
	GenerateSchedule(teams []domain.Team) []domain.Week
}
