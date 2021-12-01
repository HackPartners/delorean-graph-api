package main

import (
	"net/http"
	"strings"

	. "github.com/HackPartners/delorean-graph-api/graph"
	"github.com/labstack/echo/v4"
)

func Stations(c echo.Context) error {
	from := c.QueryParam("from")
	direction := c.QueryParam("direction")
	if (from == "") && (direction == "") {
		return c.JSON(http.StatusOK, GetAllStations())
	} else if (from != "") && (direction != "") {
		direction = strings.ToUpper(direction)
		return c.JSON(http.StatusOK, GetStationsAfter(direction, from))
	} else {
		return c.String(http.StatusBadRequest, "Both 'from' and 'direction' query params are required, or no params to get a full list of stations.")
	}
}
func UpStations(c echo.Context) error {
	station := c.Param("station")
	return c.JSON(http.StatusOK, GetStationsAfter("UP", station))
}
func DownStations(c echo.Context) error {
	station := c.Param("station")
	return c.JSON(http.StatusOK, GetStationsAfter("DOWN", station))
}
func Paths(c echo.Context) error {
	from := c.QueryParam("from")
	to := c.QueryParam("to")
	direction := strings.ToUpper(c.QueryParam("direction"))
	if (from != "") && (to != "") && (direction != "") {
		return c.JSON(http.StatusOK, GetPathsBetween(direction, from, to))
	} else if (from != "") && (direction != "") {
		return c.JSON(http.StatusOK, GetPathsAfter(direction, from))
	} else {
		return c.String(http.StatusBadRequest, "'to', 'from' and 'direction' query params are required, or 'from' and 'direction' query params.")
	}
}
func UpPaths(c echo.Context) error {
	station := c.Param("station")
	return c.JSON(http.StatusOK, GetPathsAfter("UP", station))
}
func DownPaths(c echo.Context) error {
	station := c.Param("station")
	return c.JSON(http.StatusOK, GetPathsAfter("DOWN", station))
}
func UpPathsBetweenStations(c echo.Context) error {
	stationA := c.Param("stationA")
	stationB := c.Param("stationB")
	return c.JSON(http.StatusOK, GetPathsBetween("UP", stationA, stationB))
}
func DownPathsBetweenStations(c echo.Context) error {
	stationA := c.Param("stationA")
	stationB := c.Param("stationB")
	return c.JSON(http.StatusOK, GetPathsBetween("DOWN", stationA, stationB))
}
