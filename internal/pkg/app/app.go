package app

import (
    "github.com/labstack/echo/v4"
    "log"
    "simpleHttp/internal/app/endpoint"
    "simpleHttp/internal/app/service"
    "simpleHttp/internal/pkg/constants"
)

type App struct {
    e    *endpoint.Endpoint
    s    *endpoint.Service
    echo *echo.Echo
}

func New() (*App, error) {
    a := &App{}

    a.s = service.New()

    a.e = endpoint.New(a.s)

    a.echo = echo.New()

    a.echo.Use(endpoint.MW)

    a.echo.GET("/status", a.e.Status)

    return a, nil
}

func (a *App) Run() error {

    log.Println(constants.StartServer, constants.Space, endpoint.NonConst())

    if err := a.echo.Start(":8888"); err != nil {
        log.Fatal(err)
    }

}