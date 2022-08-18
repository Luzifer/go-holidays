package holidays

import "time"

func init() {
	registerHolidayDataSource("de", holidaysDE{})
	registerHolidayDataSource("de-bb", holidaysDE{state: "bb"})
	registerHolidayDataSource("de-be", holidaysDE{state: "be"})
	registerHolidayDataSource("de-bw", holidaysDE{state: "bw"})
	registerHolidayDataSource("de-by", holidaysDE{state: "by"})
	registerHolidayDataSource("de-hb", holidaysDE{state: "hb"})
	registerHolidayDataSource("de-he", holidaysDE{state: "he"})
	registerHolidayDataSource("de-hh", holidaysDE{state: "hh"})
	registerHolidayDataSource("de-mv", holidaysDE{state: "mv"})
	registerHolidayDataSource("de-ni", holidaysDE{state: "ni"})
	registerHolidayDataSource("de-nw", holidaysDE{state: "nw"})
	registerHolidayDataSource("de-rp", holidaysDE{state: "rp"})
	registerHolidayDataSource("de-sh", holidaysDE{state: "sh"})
	registerHolidayDataSource("de-sl", holidaysDE{state: "sl"})
	registerHolidayDataSource("de-sn", holidaysDE{state: "sn"})
	registerHolidayDataSource("de-st", holidaysDE{state: "st"})
	registerHolidayDataSource("de-th", holidaysDE{state: "th"})
}

type (
	holidaysDE struct{ state string }
)

func (holidaysDE) GetIncludes() []string { return nil }
func (h holidaysDE) GetHolidays(year int) []Holiday {
	var (
		neujahr         = newHoliday("New Year's Day", map[string]string{"de": "Neujahrstag"}, dateFromNumbers(year, 1, 1))
		hlDreiKoenige   = newHoliday("Epiphany", map[string]string{"de": "Heilige Drei Könige"}, dateFromNumbers(year, 1, 6))
		womansDay       = newHoliday("International Women’s Day", map[string]string{"de": "Internationaler Frauentag"}, dateFromNumbers(year, 3, 8))
		tagDerArbeit    = newHoliday("Labor Day", map[string]string{"de": "Tag der Arbeit"}, dateFromNumbers(year, 5, 1))
		marHimmelfahrt  = newHoliday("Assumption of Mary", map[string]string{"de": "Mariä Himmelfahrt"}, dateFromNumbers(year, 8, 15))
		kinderTag       = newHoliday("World Children's Day", map[string]string{"de": "Weltkindertag"}, dateFromNumbers(year, 9, 20))
		tagDerEinheit   = newHoliday("German Unity Day", map[string]string{"de": "Tag der Deutschen Einheit"}, dateFromNumbers(year, 10, 3))
		reformationsTag = newHoliday("Reformation Day", map[string]string{"de": "Reformationstag"}, dateFromNumbers(year, 10, 31))
		allerheiligen   = newHoliday("All Saints", map[string]string{"de": "Allerheiligen"}, dateFromNumbers(year, 11, 1))
		weihnacht1      = newHoliday("Christmas Day", map[string]string{"de": "Weihnachtstag"}, dateFromNumbers(year, 12, 25))
		weihnacht2      = newHoliday("Boxing Day", map[string]string{"de": "Zweiter Weihnachtsfeiertag"}, dateFromNumbers(year, 12, 26))

		karfreitag     = newHoliday("Good Friday", map[string]string{"de": "Karfreitag"}, GregorianEasterSunday(year).Add(-2*day))
		osterSonntag   = newHoliday("Easter Sunday", map[string]string{"de": "Ostersonntag"}, GregorianEasterSunday(year))
		osterMontag    = newHoliday("Easter Monday", map[string]string{"de": "Ostermontag"}, GregorianEasterSunday(year).Add(1*day))
		chrHimmelfahrt = newHoliday("Ascension Day", map[string]string{"de": "Christi Himmelfahrt"}, GregorianEasterSunday(year).Add(39*day))
		pfingstMontag  = newHoliday("Whit Monday", map[string]string{"de": "Pfingstmontag"}, GregorianEasterSunday(year).Add(50*day))
		fronleichnam   = newHoliday("Corpus Christi", map[string]string{"de": "Fronleichnam"}, GregorianEasterSunday(year).Add(60*day))

		national = []Holiday{
			neujahr, tagDerArbeit, tagDerEinheit, weihnacht1, weihnacht2,
			karfreitag, osterSonntag, osterMontag, chrHimmelfahrt, pfingstMontag,
		}
		states = map[string][]Holiday{
			"":   national,
			"bb": append(national, reformationsTag),
			"be": append(national, womansDay),
			"bw": append(national, hlDreiKoenige, fronleichnam, allerheiligen),
			"by": append(national, hlDreiKoenige, fronleichnam, marHimmelfahrt, allerheiligen),
			"hb": append(national, reformationsTag),
			"he": append(national, fronleichnam),
			"hh": append(national, reformationsTag),
			"mv": append(national, reformationsTag),
			"ni": append(national, reformationsTag),
			"nw": append(national, fronleichnam, allerheiligen),
			"rp": append(national, fronleichnam, allerheiligen),
			"sh": append(national, reformationsTag),
			"sl": append(national, fronleichnam, marHimmelfahrt, allerheiligen),
			"sn": append(national, reformationsTag, h.getDRP(year)),
			"st": append(national, hlDreiKoenige, reformationsTag),
			"th": append(national, kinderTag, reformationsTag),
		}
	)

	return states[h.state]
}

func (holidaysDE) getDRP(year int) Holiday {
	var (
		day  time.Time
		dayN = 16
	)

	for {
		day = dateFromNumbers(year, 11, dayN)
		if day.Weekday() == time.Wednesday {
			break
		}

		dayN++
	}

	return newHoliday("Day of Repentance and Prayer", map[string]string{"de": "Buß- und Bettag"}, day)
}
