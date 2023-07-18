package never18

import (
	"errors"
	"fmt"
	"math"
	"time"
)

type Age struct {
	Birth time.Time
}

type Report struct {
	Years  int
	Months int
	Days   int
}

var ErrNegativeAge = errors.New("negative age")

func (r Report) String() string {
	return fmt.Sprintf("%d years, %d months, %d days", r.Years, r.Months, r.Days)
}

func (a Age) Nominally(moment time.Time) (Report, error) {
	var (
		years  int
		months int
		days   int
	)

	yearsSub := moment.Year() - a.Birth.Year()
	monthsSub := int(moment.Month()) - int(a.Birth.Month())
	daysSub := moment.Day() - a.Birth.Day()

	if yearsSub < 0 {
		return Report{}, ErrNegativeAge
	}

	if monthsSub >= 0 {
		years = yearsSub
		months = monthsSub
	} else {
		years = yearsSub - 1
		months = 12 + monthsSub
	}

	if daysSub >= 0 {
		days = daysSub
	} else {
		months -= 1
		lastMonth := int(a.Birth.Month()) + months
		if lastMonth > int(time.December) {
			return Report{}, errors.New("last month is greater than December")
		}
		// day may be 28~31, so using duration from the birthday in last month
		dayBegin := time.Date(a.Birth.Year()+years, time.Month(lastMonth), a.Birth.Day(), 0, 0, 0, 0, time.UTC)
		days = int(math.Round(moment.Sub(dayBegin).Hours())) / 24
	}

	return Report{
		Years:  years,
		Months: months,
		Days:   days,
	}, nil
}

func (a Age) Truth(moment time.Time, limitYears int) (Report, error) {
	var (
		years  int
		months int
	)

	nominally, err := a.Nominally(moment)
	if err != nil {
		return Report{}, err
	}

	if nominally.Years <= limitYears {
		years = nominally.Years
		months = nominally.Months
	} else {
		years = limitYears
		months = nominally.Months + ((nominally.Years - limitYears) * 12)
	}

	return Report{
		Years:  years,
		Months: months,
		Days:   nominally.Days,
	}, nil
}
