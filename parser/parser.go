package parser

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/rad08d/api/scheduler"
)

type Input struct {
	Rates []RateInput `json:"rates"`
}

type RateInput struct {
	Days  string `json:"days"`
	Times string `json:"times"`
	Price int    `json:"price"`
}

func ParseInput(inputJson []byte) Input {
	var input Input
	err := json.Unmarshal(inputJson, &input)
	if err != nil {
		panic(err)
	}
	return input
}

func InputToWeek(input Input) scheduler.Week {
	week := scheduler.NewWeek()
	for _, r := range input.Rates {
		timeRange := strings.Split(r.Times, "-")
		startHour, err := strconv.ParseInt(timeRange[0], 10, 16)
		if err != nil {
			panic(err)
		}
		endHour, err := strconv.ParseInt(timeRange[1], 10, 16)
		if err != nil {
			panic(err)
		}
		splitDays(week, r.Days, int(startHour)/100, int(endHour)/100, r.Price)
	}
	return week
}

func splitDays(week scheduler.Week, days string, startHour int, endHour int, price int) {
	dayList := strings.Split(days, ",")
	for _, v := range dayList {
		switch v {
		case "sun":
			week.Days[0].AddAvailability(startHour, endHour, price)
		case "mon":
			week.Days[1].AddAvailability(startHour, endHour, price)
		case "tues":
			week.Days[2].AddAvailability(startHour, endHour, price)
		case "wed":
			week.Days[3].AddAvailability(startHour, endHour, price)
		case "thurs":
			week.Days[4].AddAvailability(startHour, endHour, price)
		case "fri":
			week.Days[5].AddAvailability(startHour, endHour, price)
		case "sat":
			week.Days[6].AddAvailability(startHour, endHour, price)
		}
	}
}
