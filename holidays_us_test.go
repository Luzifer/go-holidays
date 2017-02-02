package holidays

import (
	"testing"
	"time"
)

func TestHolidaysUS(t *testing.T) {
	// List of known dates according to US Office of Personnel Management
	// https://www.opm.gov/policy-data-oversight/snow-dismissal-procedures/federal-holidays/

	knownDates := map[int]map[string]time.Time{
		2011: map[string]time.Time{
			"New Year's Day":                      dateFromNumbers(2010, 12, 31), // 1 = Saturday... What the hell are you doing?!?
			"Birthday of Martin Luther King, Jr.": dateFromNumbers(2011, 1, 17),
			"Washington's Birthday":               dateFromNumbers(2011, 2, 21),
			"Memorial Day":                        dateFromNumbers(2011, 5, 30),
			"Independence Day":                    dateFromNumbers(2011, 7, 4),
			"Labor Day":                           dateFromNumbers(2011, 9, 5),
			"Columbus Day":                        dateFromNumbers(2011, 10, 10),
			"Veterans Day":                        dateFromNumbers(2011, 11, 11),
			"Thanksgiving Day":                    dateFromNumbers(2011, 11, 24),
			"Christmas Day":                       dateFromNumbers(2011, 12, 26), // 25 = Sunday
		},
		2014: map[string]time.Time{
			"New Year's Day":                      dateFromNumbers(2014, 1, 1),
			"Birthday of Martin Luther King, Jr.": dateFromNumbers(2014, 1, 20),
			"Washington's Birthday":               dateFromNumbers(2014, 2, 17),
			"Memorial Day":                        dateFromNumbers(2014, 5, 26),
			"Independence Day":                    dateFromNumbers(2014, 7, 4),
			"Labor Day":                           dateFromNumbers(2014, 9, 1),
			"Columbus Day":                        dateFromNumbers(2014, 10, 13),
			"Veterans Day":                        dateFromNumbers(2014, 11, 11),
			"Thanksgiving Day":                    dateFromNumbers(2014, 11, 27),
			"Christmas Day":                       dateFromNumbers(2014, 12, 25),
		},
		2017: map[string]time.Time{
			"New Year's Day":                      dateFromNumbers(2017, 1, 2), // 1 = Sunday
			"Birthday of Martin Luther King, Jr.": dateFromNumbers(2017, 1, 16),
			"Washington's Birthday":               dateFromNumbers(2017, 2, 20),
			"Memorial Day":                        dateFromNumbers(2017, 5, 29),
			"Independence Day":                    dateFromNumbers(2017, 7, 4),
			"Labor Day":                           dateFromNumbers(2017, 9, 4),
			"Columbus Day":                        dateFromNumbers(2017, 10, 9),
			"Veterans Day":                        dateFromNumbers(2017, 11, 10), // 11 = Saturday
			"Thanksgiving Day":                    dateFromNumbers(2017, 11, 23),
			"Christmas Day":                       dateFromNumbers(2017, 12, 25),
		},
	}

	for year, knownHolidays := range knownDates {
		generatedHolidays, err := GetHolidays("us", year)
		if err != nil {
			t.Fatalf("[US] Could not load holidays for %d: %s", year, err)
		}

		for name, date := range knownHolidays {

			found := false
			for _, hd := range generatedHolidays {
				if hd.Name == name && hd.ParsedDate == date {
					found = true
				}
			}
			if !found {
				t.Errorf("[US] Did not find %q on %s.", name, date.Format("2006-01-02"))
			}

		}
	}
}
