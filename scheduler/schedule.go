package scheduler

import (
	"errors"
	"time"
)

type Week struct {
	Days [7]Day
}

type Day struct {
	DayOfWeek      time.Weekday
	TimesAvailable []Availability
}

type Availability struct {
	Status    Status
	StartHour int
	EndHour   int
	Rate      Rate
}

type Status int16

const (
	Unavailable Status = 0
	Available   Status = 1
)

type Rate struct {
	Availability Status
	Price        int
}

func NewWeek() Week {
	days := [7]Day{
		Day{time.Sunday, make([]Availability, 24)},
		Day{time.Monday, make([]Availability, 24)},
		Day{time.Tuesday, make([]Availability, 24)},
		Day{time.Wednesday, make([]Availability, 24)},
		Day{time.Thursday, make([]Availability, 24)},
		Day{time.Friday, make([]Availability, 24)},
		Day{time.Saturday, make([]Availability, 24)},
	}
	return Week{days}
}

// CheckAvailability takes a start time string and end time string in the ISO 8601 format
// and checks the availability of that time range.
// This method does not support checking an availability spanning multiple days.
func (w Week) CheckAvailability(startTime string, endTime string) (Rate, error) {
	start, errS := time.Parse(time.RFC3339, startTime)
	if errS != nil {
		return Rate{}, errS
	}
	end, errE := time.Parse(time.RFC3339, endTime)
	if errE != nil {
		return Rate{}, errE
	}
	if start.Day() != end.Day() {
		return Rate{}, errors.New("Start day is not the same as end day.")
	}
	day := w.Days[int(start.Weekday())]
	availability := day.GetAvailability(start.Hour(), end.Hour())

	return availability.GetRate(start, end), nil
}

// AddAvailability adds a new available range to a Day.
// It assumes there will be no overlapping ranges for a single day.
func (d Day) AddAvailability(startHour int, endHour int, price int) {
	d.TimesAvailable[startHour] = Availability{Available, startHour, endHour, Rate{Available, price}}
}

// GetAvailability checks the availability of time slot.
func (d Day) GetAvailability(startHour int, endHour int) Availability {
	for _, v := range d.TimesAvailable {
		if startHour >= v.StartHour && endHour <= v.EndHour {
			return v
		}
	}
	return Availability{Unavailable, 0, 0, Rate{Unavailable, 0}}
}

// GetRate will check the availability hours of a location for a day.
// It operates with the assumption that availability is only on a single day basis.
func (a Availability) GetRate(start time.Time, end time.Time) Rate {
	timeSpan := end.Sub(start).Hours() - 1
	timeAvail := float64(a.AvailableTime())
	if a.StartHour <= start.Hour() {
		if timeSpan <= timeAvail {
			return a.Rate
		}
	}
	return Rate{Unavailable, 0}
}

// AvailableTime computes the time length available for a 24 hour clock
func (a Availability) AvailableTime() int {
	switch end := a.EndHour; {
	case end < a.StartHour:
		return (24 - a.StartHour) + a.EndHour
	case end > a.StartHour:
		return a.EndHour - a.StartHour
	case end == a.StartHour:
		return 24
	default:
		return 0
	}
}
