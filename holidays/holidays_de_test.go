package holidays

import (
	"testing"
	"time"
)

func TestDRPDate(t *testing.T) {
	for year, exp := range map[int]time.Time{
		2021: dateFromNumbers(2021, 11, 17),
		2022: dateFromNumbers(2022, 11, 16),
		2023: dateFromNumbers(2023, 11, 22),
		2024: dateFromNumbers(2024, 11, 20),
		2025: dateFromNumbers(2025, 11, 19),
		2026: dateFromNumbers(2026, 11, 18),
		2027: dateFromNumbers(2027, 11, 17),
	} {
		h := holidaysDE{}.getDRP(year)
		if !h.ParsedDate.Equal(exp) {
			t.Errorf("DRP %d: Expected %s, got %s", year, exp, h.ParsedDate)
		}
	}
}
