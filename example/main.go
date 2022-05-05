package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	live "github.com/tejzpr/echo-live/v4"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Static("assets"))
	e.Use(live.Live())
	e.Logger.Fatal(e.Start(":8080"))
}
