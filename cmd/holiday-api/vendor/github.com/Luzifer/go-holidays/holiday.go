package holidays

//go:generate go-bindata -pkg $GOPACKAGE -o holidays_data.go holidays/...

import (
	"errors"
	"fmt"
	"path"
	"sort"
	"strings"
	"time"

	yaml "gopkg.in/yaml.v2"
)

var HolidayDataNotFoundError = errors.New("Holiday data not found for requested country-code")

// Holiday contains information about an holiday
type Holiday struct {
	// Name contains the english representation of the holidays name
	Name string `json:"name"`
	// LocalizedName contains a map of localizations of the name, key is a ISO 3166-2 country code without subdivision
	LocalizedName map[string]string `json:"localized_name"`
	// Date contains the date in YYYY-MM-DD notation
	Date string `json:"date"`
}

type holidays []Holiday

func (h holidays) Len() int           { return len(h) }
func (h holidays) Less(i, j int) bool { return h[i].Date.Before(h[j].Date) }
func (h holidays) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// GetHolidays returns the holidays for the given ISO 3166-2 countryCode and year
func GetHolidays(countryCode string, year int) ([]Holiday, error) {
	requiredFiles := []string{countryCode}
	result := []Holiday{}

	for len(requiredFiles) > 0 {
		hdf, err := loadDataFile(requiredFiles[0])
		if err != nil {
			return result, err
		}

		requiredFiles = append(requiredFiles, hdf.Includes...)

		easter := GregorianEasterSunday(year)

		for _, h := range hdf.Holidays.Fixed {
			result = append(result, Holiday{
				Name:          h.Name,
				LocalizedName: h.LocalizedName,
				Date:          time.Date(year, time.Month(h.Month), h.Day, 0, 0, 0, 0, time.Local).Format("2006-01-02"),
			})
		}

		for _, h := range hdf.Holidays.EasterBased {
			result = append(result, Holiday{
				Name:          h.Name,
				LocalizedName: h.LocalizedName,
				Date:          easter.Add(time.Duration(h.Difference) * 24 * time.Hour).Format("2006-01-02"),
			})
		}

		requiredFiles = requiredFiles[1:]
	}

	sort.Sort(holidays(result))

	return result, nil
}

type holidaysDataFileEntry struct {
	Name          string            `yaml:"name"`
	LocalizedName map[string]string `yaml:"localized"`
	Month         int               `yaml:"month"`
	Day           int               `yaml:"day"`
	Difference    int               `yaml:"difference"`
}

type holidaysDataFile struct {
	Includes []string `yaml:"includes"`
	Holidays struct {
		Fixed       []holidaysDataFileEntry `yaml:"fixed"`
		EasterBased []holidaysDataFileEntry `yaml:"easter_based"`
	} `yaml:"holidays"`
}

func loadDataFile(countryCode string) (holidaysDataFile, error) {
	r := holidaysDataFile{}

	parts := strings.Split(countryCode, "-")
	if len(parts) == 1 {
		parts = append(parts, "_")
	}

	data, err := Asset(fmt.Sprintf(path.Join("holidays", parts[0], parts[1]) + ".yaml"))
	if err != nil {
		return r, HolidayDataNotFoundError
	}

	return r, yaml.Unmarshal(data, &r)
}
