package holidays

import "time"

func init() {
	registerHolidayDataSource("us", holidaysUSNational{})
}

type holidaysUSNational struct{}

func (h holidaysUSNational) GetIncludes() []string { return []string{} }
func (h holidaysUSNational) GetHolidays(year int) []Holiday {
	return []Holiday{
		newHoliday("New Year's Day", map[string]string{}, h.fixWeekend(dateFromNumbers(year, 1, 1))),                                         // 1st January
		newHoliday("Birthday of Martin Luther King, Jr.", map[string]string{}, h.findNextWeekday(dateFromNumbers(year, 1, 15), time.Monday)), // Third Monday in January
		newHoliday("Washington's Birthday", map[string]string{}, h.findNextWeekday(dateFromNumbers(year, 2, 15), time.Monday)),               // Third Monday in February
		newHoliday("Memorial Day", map[string]string{}, h.findNextWeekday(dateFromNumbers(year, 5, 25), time.Monday)),                        // Last Monday in May
		newHoliday("Independence Day", map[string]string{}, h.fixWeekend(dateFromNumbers(year, 7, 4))),                                       // 4 July
		newHoliday("Labor Day", map[string]string{}, h.findNextWeekday(dateFromNumbers(year, 9, 1), time.Monday)),                            // First Monday in September
		newHoliday("Columbus Day", map[string]string{}, h.findNextWeekday(dateFromNumbers(year, 10, 8), time.Monday)),                        // Second Monday in October
		newHoliday("Veterans Day", map[string]string{}, h.fixWeekend(dateFromNumbers(year, 11, 11))),                                         // 11 November
		newHoliday("Thanksgiving Day", map[string]string{}, h.findNextWeekday(dateFromNumbers(year, 11, 23), time.Thursday)),                 // Fourth Thursday in November
		newHoliday("Christmas Day", map[string]string{}, h.fixWeekend(dateFromNumbers(year, 12, 25))),                                        // 25 December
	}
}

func (h holidaysUSNational) findNextWeekday(in time.Time, wd time.Weekday) time.Time {
	for in.Weekday() != wd {
		in = in.Add(1 * 24 * time.Hour)
	}
	return in
}

// If a holiday falls on a Saturday it is celebrated the preceding Friday;
// if a holiday falls on a Sunday it is celebrated the following Monday.
func (h holidaysUSNational) fixWeekend(in time.Time) time.Time {
	switch in.Weekday() {
	case time.Saturday:
		return in.Add(-1 * 24 * time.Hour)
	case time.Sunday:
		return in.Add(1 * 24 * time.Hour)
	default:
		return in
	}
}
