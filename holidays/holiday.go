package holidays

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

const day = 24 * time.Hour

var (
	HolidayDataNotFoundError = errors.New("Holiday data not found for requested country-code")
	holidayProviders         = map[string]holidayDataSource{}
)

func registerHolidayDataSource(code string, hds holidayDataSource) {
	if _, ok := holidayProviders[code]; ok {
		panic(fmt.Errorf("Duplicatei definition for country code %q", code))
	}

	holidayProviders[code] = hds
}

// Holiday contains information about an holiday
type Holiday struct {
	// Name contains the english representation of the holidays name
	Name string `json:"name"`
	// LocalizedName contains a map of localizations of the name, key is a ISO 3166-2 country code without subdivision
	LocalizedName map[string]string `json:"localized_name"`
	// Date contains the date in YYYY-MM-DD notation
	Date string `json:"date"`
	// ParsedDate is the Date as a time.Time object
	ParsedDate time.Time `json:"parsed_date"`
}

func newHoliday(name string, localizedName map[string]string, parsedDate time.Time) Holiday {
	return Holiday{
		Name:          name,
		LocalizedName: localizedName,
		Date:          parsedDate.Format("2006-01-02"),
		ParsedDate:    parsedDate,
	}
}

func dateFromNumbers(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}

type holidayDataSource interface {
	GetIncludes() []string
	GetHolidays(year int) []Holiday
}

// GetHolidays returns the holidays for the given ISO 3166-2 countryCode and year
func GetHolidays(countryCode string, year int) ([]Holiday, error) {
	requiredCodes := []string{countryCode}
	result := []Holiday{}

	for len(requiredCodes) > 0 {
		cc := requiredCodes[0]
		hds, ok := holidayProviders[cc]
		if !ok {
			return nil, HolidayDataNotFoundError
		}

		requiredCodes = append(requiredCodes, hds.GetIncludes()...)
		result = append(result, hds.GetHolidays(year)...)

		requiredCodes = requiredCodes[1:]
	}

	sort.Sort(holidays(result))

	return result, nil
}

type holidays []Holiday

func (h holidays) Len() int           { return len(h) }
func (h holidays) Less(i, j int) bool { return h[i].Date < h[j].Date }
func (h holidays) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
