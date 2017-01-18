package cron

import "time"

type MonthDelaySchedule struct {
	Delay int
}

func MonthDelay(delay int) MonthDelaySchedule {
	if delay <= 0 {
		delay = 0
	}
	return MonthDelaySchedule{
		Delay: delay,
	}
}

// Next returns the next time this should be run.
// This rounds so that the next activation time will be on the second.
func (schedule MonthDelaySchedule) Next(t time.Time) time.Time {
	return t.AddDate(0, schedule.Delay, 0)
}
