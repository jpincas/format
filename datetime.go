package format

import "time"

const (
	weekDayMonthDayTimeNoSecondsYear = "Mon Jan 2 15:04 2006"
)

//DateToWeekDayMonthDayTimeNoSecondsYear formats as follows: "Mon Jan 2 15:04 2006"
func DateToWeekDayMonthDayTimeNoSecondsYear(t time.Time) string {

	return t.Format(weekDayMonthDayTimeNoSecondsYear)
}
