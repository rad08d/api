package scheduler

import (
	"testing"
	"time"
)

func TestGetRateAvailable(t *testing.T) {
	expectedRate := 500
	startTime := time.Now().UTC()
	endTime := time.Now().UTC().Add(time.Hour)
	rate := Rate{Available, expectedRate}
	avail := Availability{Available, startTime.Hour(), endTime.Hour(), rate}

	computedRate := avail.GetRate(startTime, endTime)
	if computedRate.Price != expectedRate {
		t.Error("Expected rate of 500 but got ", rate)
	}
}

func TestGetRateUnavailable(t *testing.T) {
	expectedRate := 500
	startTime := time.Now().UTC()
	endTime := time.Now().UTC().Add(time.Hour)
	rate := Rate{Available, expectedRate}
	avail := Availability{Available, startTime.Hour(), endTime.Hour(), rate}

	rateEndTimeAfterBounds := avail.GetRate(startTime, endTime.Add(time.Hour))
	if rateEndTimeAfterBounds.Availability != Unavailable && rateEndTimeAfterBounds.Price != 0 {
		t.Error("Start time in availability and End time out of availability price should equal 0 not ", rateEndTimeAfterBounds.Price)
	}

	rateStartAboveEndAfterBounds := avail.GetRate(startTime.Add(2*time.Hour), endTime.Add(3*time.Hour))
	if rateStartAboveEndAfterBounds.Availability != Unavailable && rateStartAboveEndAfterBounds.Price != 0 {
		t.Error("Start time after avaialability and End time after availability price should equal 0 not ", rateStartAboveEndAfterBounds.Price)
	}

	rateStartTimeBeforeBoundsEndTimeAfterBounds := avail.GetRate(startTime.Add(time.Duration(-1)*time.Hour), endTime)
	if rateStartTimeBeforeBoundsEndTimeAfterBounds.Availability != Unavailable && rateStartTimeBeforeBoundsEndTimeAfterBounds.Price != 0 {
		t.Error("Start time below availability and End time in availability price should equal 0 not ", rateStartTimeBeforeBoundsEndTimeAfterBounds.Price)
	}
}

func TestWeekdayAvailability(t *testing.T) {
	expectedRate := 150
	startTime := 10
	endTime := 13
	rate := Rate{Available, expectedRate}
	avail := Availability{Available, startTime, endTime, rate}
	day := Day{time.Monday, []Availability{avail}}

	available := day.GetAvailability(startTime, endTime)
	if available.Status != Available {
		t.Error("Available range is marked as ", available.Status)
	}

	unavailable := day.GetAvailability(12, 14)

	if unavailable.Status != Unavailable {
		t.Error("An unavailable time range is marked as ", unavailable.Status)
	}
}

func TestAddAvailability(t *testing.T) {
	day := Day{time.Sunday, make([]Availability, 24)}
	price := 550
	startHour := time.Now().UTC().Hour()
	endHour := time.Now().Add(time.Hour).UTC().Hour()
	day.AddAvailability(startHour, endHour, price)

	avail := day.TimesAvailable[startHour]
	if avail.Rate.Availability != Available {
		t.Error("An available range is not marked as ", Unavailable)
	}
}

func TestAvailableTime(t *testing.T) {
	expectedRate := 150
	rate := Rate{Available, expectedRate}

	availOverMidnight := Availability{Available, 16, 3, rate}
	if availOverMidnight.AvailableTime() != 11 {
		t.Error("Availability should be 11 but is ", availOverMidnight.AvailableTime())
	}

	availInSameDay := Availability{Available, 5, 23, rate}
	if availInSameDay.AvailableTime() != 18 {
		t.Error("Availability should be 18 but is ", availInSameDay.AvailableTime())
	}

	avail24Hours := Availability{Available, 1, 1, rate}
	if avail24Hours.AvailableTime() != 24 {
		t.Error("Availability should be 24 but is ", avail24Hours.AvailableTime())
	}
}

func TestCheckAvailability(t *testing.T) {
	week := NewWeek()
	wedExpectedRate := 1500
	satExpectedRate := 2000
	startHour := 6
	wedsRateExpected := Rate{Available, wedExpectedRate}
	satRateExpected := Rate{Available, satExpectedRate}
	wed := Day{time.Wednesday, []Availability{Availability{Available, startHour, 18, wedsRateExpected}}}
	sat := Day{time.Saturday, []Availability{Availability{Available, startHour, 20, satRateExpected}}}

	week.Days[time.Wednesday] = wed
	week.Days[time.Saturday] = sat

	wedRateTest, err := week.CheckAvailability("2015-07-01T07:00:00Z", "2015-07-01T12:00:00Z")
	if err != nil {
		t.Error("There was an error checking availability ", err)
	}
	if wedRateTest.Price != wedExpectedRate {
		t.Error("Rate for date range should equal 15000 but equals ", wedRateTest.Price)
	}

	satRateTest, err := week.CheckAvailability("2015-07-04T07:00:00Z", "2015-07-04T12:00:00Z")
	if err != nil {
		t.Error("There was an error checking availability ", err)
	}
	if satRateTest.Price != satExpectedRate {
		t.Error("Rate for date range should equal 2000 but equals ", satRateTest.Price)
	}
	unAvailTest, err := week.CheckAvailability("2015-07-04T03:00:00Z", "2015-07-04T08:00:00Z")
	if unAvailTest.Availability != Unavailable {
		t.Error("An unavailable time range was reported as available")
	}
}
