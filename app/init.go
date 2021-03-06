package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/DSMdongly/pnf/config"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	Echo *echo.Echo
)

func Init() {
	Echo = echo.New()

	Echo.Static("/static", "app/static")

	Echo.Use(middleware.Recover())
	Echo.Use(middleware.RequestID())
	Echo.Use(middleware.CORS())
}

func Awake() {
	go func() {
		time.Sleep(time.Minute * 1)

		for {
			http.Get("https://localhost/")
			time.Sleep(time.Minute * 3)
		}
	}()
}

func Start() {
	Echo.Logger.Fatal(Echo.Start(fmt.Sprintf(":%s", config.HTTP["PORT"])))
}
