package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andreabertanzon/todo-htmx/models"
	"github.com/andreabertanzon/todo-htmx/templates"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost}, // And other methods as needed
	}))

	e.GET("/", func(c echo.Context) error {
		helloComponent := templates.Index()
		helloComponent.Render(context.Background(), c.Response().Writer)
		return nil
	})

	e.GET("/todo/add", func(c echo.Context) error {
		fmt.Println("GET /todo/add")
		helloComponent := templates.TodoForm()
		helloComponent.Render(context.Background(), c.Response().Writer)
		return nil
	})

	e.POST("/todo/create", func(c echo.Context) error {
		title := c.FormValue("title")
		description := c.FormValue("description")
		fmt.Println(title, description)
		todo := models.Todo{Title: title, Description: description}
		todoList := []models.Todo{todo}
		helloComponent := templates.TodoList(todoList)
		helloComponent.Render(context.Background(), c.Response().Writer)
		return nil
	})

	e.Static("/css", "css")
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":8083"))
}
