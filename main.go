package main

import (
	"net/http"

	. "github.com/HackPartners/delorean-graph-api/graph"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	InitialiseGraphClient()
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	e.GET("/stations", Stations).Name = "stations"
	e.GET("/stations/:station/up", UpStations).Name = "up-stations"
	e.GET("/stations/:station/down", DownStations).Name = "down-stations"
	e.GET("/paths", Paths).Name = "paths"
	e.GET("/paths/:station/up", UpPaths).Name = "up-paths"
	e.GET("/paths/:station/down", DownPaths).Name = "down-paths"
	e.GET("/paths/:stationA/to/:stationB/up", UpPathsBetweenStations).Name = "up-paths-between-stations"
	e.GET("/paths/:stationA/to/:stationB/down", DownPathsBetweenStations).Name = "down-paths-between-stations"

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
