package main

import (
	"net/http"

	. "github.com/HackPartners/delorean-graph-api/graph"
	"github.com/labstack/echo/v4"
)

func AllStations(c echo.Context) error {
	return c.JSON(http.StatusOK, GetAllStations())
}
func UpStations(c echo.Context) error {
	station := c.Param("station")
	return c.JSON(http.StatusOK, GetStationsAfter("UP", station))
}
func DownStations(c echo.Context) error {
	station := c.Param("station")
	return c.JSON(http.StatusOK, GetStationsAfter("DOWN", station))
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
