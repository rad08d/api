package parser

import (
	"fmt"
	"testing"
)

func TestParseInput(t *testing.T) {
	json := []byte(`{
		"rates": [
		  {
			"days": "mon,tues,wed,thurs,fri",
			"times": "0600-1800",
			"price": 1500
		  },
		  {
			"days": "sat,sun",
			"times": "0600-2000",
			"price": 2000
		  }
		]
	  }`)
	input := ParseInput(json)
	if len(input.Rates) != 2 {
		t.Error("Parse of input JSON did not parse correctly. It has an incorrect count of rates.")
	}
}

func TestInputToWeek(t *testing.T) {
	json := []byte(`{
		"rates": [
		  {
			"days": "mon,tues,wed,thurs,fri",
			"times": "0600-1800",
			"price": 1500
		  },
		  {
			"days": "sat,sun",
			"times": "0600-2000",
			"price": 2000
		  }
		]
	  }`)
	input := ParseInput(json)
	week := InputToWeek(input)
	for k, _ := range week.Days {
		fmt.Print(len(week.Days[k].TimesAvailable))
		if week.Days[k].TimesAvailable[6].Status != Available && week.Days[k].TimesAvailable[6].StartHour == 6 {
			t.Error("Availability is missing for the week.")
		}
	}

}
