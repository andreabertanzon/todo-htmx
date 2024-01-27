package main

import (
	"context"

	"github.com/andreabertanzon/todo-htmx/templates"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		helloComponent := templates.Index()
		helloComponent.Render(context.Background(), c.Response().Writer)
		return nil
	})
	e.Static("/css", "css")
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":8081"))
}
