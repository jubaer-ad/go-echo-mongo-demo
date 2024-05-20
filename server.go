package main

import (
	"myapp/configs"
	"myapp/routes"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	client := configs.ConnectDB()

	routes.UserRoute(e)

	configs.GetCollection(client, "testDB")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/show", show)
	e.Logger.Fatal(e.Start(":5555"))
}

//e.GET("/show", show)
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	other := c.QueryParam("other")
	return c.JSON(http.StatusOK, &echo.Map{"data":"team:" + team + ", member:" + member + ", other:" + other})
}

