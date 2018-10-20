package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/rad08d/api/parser"
	"github.com/rad08d/api/scheduler"
)

func Health(c echo.Context) error {
	return c.String(http.StatusOK, "healthy....maybe")
}

func CheckAvailability(c echo.Context) error {
	startTime := c.QueryParam("startTime")
	endTime := c.QueryParam("endTime")
	fmt.Println("Start Time: ", startTime)
	fmt.Println("End time: ", endTime)
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, "whooppse")
	}
	schedule := parser.InputToWeek(parser.ParseInput(body))
	rate, err := schedule.CheckAvailability(startTime, endTime)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Uh Oh!")
	}
	if rate.Availability == scheduler.Available {
		return c.String(http.StatusOK, strconv.Itoa(rate.Price))
	}
	// Not available
	return c.String(http.StatusOK, "Unavailable")
}
