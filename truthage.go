package never18

import (
	"math"
	"time"
)

type FalsehoodAge struct {
	Years  int
	Months int
	Days   int
}

type TruthAge struct {
	Years  int
	Months int
	Days   int
}

func GetTruthAge(duration time.Duration, limit int) TruthAge {
	return GetFalsehoodAgeFromDuration(duration).TruthAge(limit)
}

func (f FalsehoodAge) excessYears(limit int) int {
	return f.Years - limit
}

func (f FalsehoodAge) TruthAge(limit int) TruthAge {
	if f.Years <= limit {
		return TruthAge(f)
	}

	return TruthAge{
		Years:  limit,
		Months: f.Months + (f.excessYears(limit) * 12),
		Days:   f.Days,
	}
}

func divmod(numerator, denominator int) (quotient int, remainder int) {
	quotient = numerator / denominator
	remainder = numerator % denominator
	return quotient, remainder
}

func GetFalsehoodAgeFromDuration(duration time.Duration) FalsehoodAge {
	allhours := int(math.Round(duration.Hours()))
	allDays := allhours / 24
	allMonths, days := divmod(allDays, 30) // Not accurate, days in a month may be 28~31
	years, months := divmod(allMonths, 12)

	return FalsehoodAge{
		Years:  years,
		Months: months,
		Days:   days,
	}
}
