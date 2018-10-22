package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	"github.com/rad08d/api/parser"
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
		return c.String(http.StatusInternalServerError, "Error processing time rates input")
	}
	schedule := parser.InputToWeek(parser.ParseInput(body))
	rate, err := schedule.CheckAvailability(startTime, endTime)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Uh Oh!")
	}
	return c.JSON(http.StatusOK, rate)
}
