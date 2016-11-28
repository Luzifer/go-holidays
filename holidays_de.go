package holidays

func init() {
	registerHolidayDataSource("de", holidaysDENational{})
	registerHolidayDataSource("de-bb", holidaysDEBB{})
	registerHolidayDataSource("de-be", holidaysDEBE{})
	registerHolidayDataSource("de-bw", holidaysDEBW{})
	registerHolidayDataSource("de-by", holidaysDEBY{})
	registerHolidayDataSource("de-hb", holidaysDEHB{})
	registerHolidayDataSource("de-he", holidaysDEHE{})
	registerHolidayDataSource("de-hh", holidaysDEHH{})
	registerHolidayDataSource("de-mv", holidaysDEMV{})
	registerHolidayDataSource("de-ni", holidaysDENI{})
	registerHolidayDataSource("de-nw", holidaysDENW{})
	registerHolidayDataSource("de-rp", holidaysDERP{})
	registerHolidayDataSource("de-sh", holidaysDESH{})
	registerHolidayDataSource("de-sl", holidaysDESL{})
	registerHolidayDataSource("de-sn", holidaysDESN{})
	registerHolidayDataSource("de-st", holidaysDEST{})
	registerHolidayDataSource("de-th", holidaysDETH{})
}

type holidaysDENational struct{}

func (h holidaysDENational) GetIncludes() []string { return []string{} }
func (h holidaysDENational) GetHolidays(year int) []Holiday {
	return []Holiday{
		newHoliday("New Year's Day", map[string]string{"de": "Neujahrstag"}, dateFromNumbers(year, 1, 1)),
		newHoliday("Labor Day", map[string]string{"de": "Tag der Arbeit"}, dateFromNumbers(year, 5, 1)),
		newHoliday("German Unity Day", map[string]string{"de": "Tag der Deutschen Einheit"}, dateFromNumbers(year, 10, 3)),
		newHoliday("Christmas Day", map[string]string{"de": "Weihnachtstag"}, dateFromNumbers(year, 12, 25)),
		newHoliday("Boxing Day", map[string]string{"de": "Zweiter Weihnachtsfeiertag"}, dateFromNumbers(year, 12, 26)),
		newHoliday("Good Friday", map[string]string{"de": "Karfreitag"}, GregorianEasterSunday(year).Add(-2*day)),
		newHoliday("Easter Sunday", map[string]string{"de": "Ostersonntag"}, GregorianEasterSunday(year)),
		newHoliday("Easter Monday", map[string]string{"de": "Ostermontag"}, GregorianEasterSunday(year).Add(1*day)),
		newHoliday("Ascension Day", map[string]string{"de": "Christi Himmelfahrt"}, GregorianEasterSunday(year).Add(39*day)),
		newHoliday("Whit Monday", map[string]string{"de": "Pfingstmontag"}, GregorianEasterSunday(year).Add(50*day)),
	}
}

type holidaysDEBB struct{}

func (h holidaysDEBB) GetIncludes() []string { return []string{"de"} }
func (h holidaysDEBB) GetHolidays(year int) []Holiday {
	return []Holiday{
		newHoliday("Reformation Day", map[string]string{"de": "Reformationstag"}, dateFromNumbers(year, 10, 31)),
	}
}

type holidaysDEBE struct{ holidaysDENational }

type holidaysDEBW struct{}

func (h holidaysDEBW) GetIncludes() []string { return []string{"de"} }
func (h holidaysDEBW) GetHolidays(year int) []Holiday {
	return []Holiday{
		newHoliday("Epiphany", map[string]string{"de": "Heilige Drei Könige"}, dateFromNumbers(year, 1, 6)),
		newHoliday("All Saints", map[string]string{"de": "Allerheiligen"}, dateFromNumbers(year, 11, 1)),
		newHoliday("Corpus Christi", map[string]string{"de": "Fronleichnam"}, GregorianEasterSunday(year).Add(60*day)),
	}
}

type holidaysDEBY struct{ holidaysDEBW }

type holidaysDEHB struct{ holidaysDENational }

type holidaysDEHE struct{}

func (h holidaysDEHE) GetIncludes() []string { return []string{"de"} }
func (h holidaysDEHE) GetHolidays(year int) []Holiday {
	return []Holiday{
		newHoliday("Corpus Christi", map[string]string{"de": "Fronleichnam"}, GregorianEasterSunday(year).Add(60*day)),
	}
}

type holidaysDEHH struct{ holidaysDENational }

type holidaysDEMV struct{ holidaysDEBB }

type holidaysDENI struct{ holidaysDENational }

type holidaysDENW struct{}

func (h holidaysDENW) GetIncludes() []string { return []string{"de"} }
func (h holidaysDENW) GetHolidays(year int) []Holiday {
	return []Holiday{
		newHoliday("All Saints", map[string]string{"de": "Allerheiligen"}, dateFromNumbers(year, 11, 1)),
		newHoliday("Corpus Christi", map[string]string{"de": "Fronleichnam"}, GregorianEasterSunday(year).Add(60*day)),
	}
}

type holidaysDERP struct{ holidaysDENW }

type holidaysDESH struct{ holidaysDENational }

type holidaysDESL struct{}

func (h holidaysDESL) GetIncludes() []string { return []string{"de"} }
func (h holidaysDESL) GetHolidays(year int) []Holiday {
	return []Holiday{
		newHoliday("Assumption Day", map[string]string{"de": "Mariä Himmelfahrt"}, dateFromNumbers(year, 8, 15)),
		newHoliday("All Saints", map[string]string{"de": "Allerheiligen"}, dateFromNumbers(year, 11, 1)),
		newHoliday("Corpus Christi", map[string]string{"de": "Fronleichnam"}, GregorianEasterSunday(year).Add(60*day)),
	}
}

type holidaysDESN struct{ holidaysDEBB }

type holidaysDEST struct{}

func (h holidaysDEST) GetIncludes() []string { return []string{"de"} }
func (h holidaysDEST) GetHolidays(year int) []Holiday {
	return []Holiday{
		newHoliday("Epiphany", map[string]string{"de": "Heilige Drei Könige"}, dateFromNumbers(year, 1, 6)),
		newHoliday("Reformation Day", map[string]string{"de": "Reformationstag"}, dateFromNumbers(year, 10, 31)),
	}
}

type holidaysDETH struct{ holidaysDEBB }
