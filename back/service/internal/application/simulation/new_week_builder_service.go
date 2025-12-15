package simulation

func NewWeekBuilderService(generator ScheduleGenerator) WeekBuilderService {
	return WeekBuilderService{scheduleGenerator: generator}
}
