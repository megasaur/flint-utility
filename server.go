package main

import (
	"net/http"

	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// TODO: need to learn about packages and break main up a bit, a lot of files already

func main() {
	e := echo.New()

	// CORS Middleware
	// TODO: Can probably introduce some security here, only allow the frontend to hit it
	e.Use(middleware.CORS())

	// health check
	e.GET("/", func(c echo.Context) error {
		// TODO: unit test for this endpoint
		return c.String(http.StatusOK, "Backend is alive")
	})

	// season endpoints
	e.GET("/seasons/current", func(c echo.Context) error {
		// TODO: unit test for this endpoint
		return c.JSON(http.StatusOK, getCurrentSeason())
	})

	// coach endpoints
	e.GET("/coaches", func(c echo.Context) error {
		// TODO: unit test for this endpoint
		year, _ := strconv.Atoi(c.QueryParam("year"))
		coaches := getAllCoaches(year)
		return c.JSON(http.StatusOK, coaches)
	})

	// saltyboy endpoint
	e.GET("/saltyboy/round", func(c echo.Context) error {
		// TODO: unit test for this endpoint
		coach := c.QueryParam("coach")
		round, _ := strconv.Atoi(c.QueryParam("round"))
		year, _ := strconv.Atoi(c.QueryParam("year"))
		payload := getSaltyboyRound(coach, round, year)
		// spew.Dump(payload)
		return c.JSON(http.StatusOK, payload)
	})

	// ladder endpoint
	e.GET("/ladder", func(c echo.Context) error {
		// TODO: unit test for this endpoint
		round, _ := strconv.Atoi(c.QueryParam("round"))
		year, _ := strconv.Atoi(c.QueryParam("year"))
		payload := getLadder(round, year)
		return c.JSON(http.StatusOK, payload)
	})

	// TODO: records endpoint with:
	// highest for stats
	// lowest for stats
	// averages

	//LUCKY BOY!
	// when u have a saltyboy that could only beat 1 opponent that round, and that happened to be their opponent!
	e.Logger.Fatal(e.Start(":1323"))
}
