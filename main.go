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

var todoList []models.Todo = []models.Todo{}

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost}, // And other methods as needed
	}))

	e.GET("/", func(c echo.Context) error {
		helloComponent := templates.Index(todoList)
		helloComponent.Render(context.Background(), c.Response().Writer)
		return nil
	})

	e.GET("/todo/form", func(c echo.Context) error {
		fmt.Println("GET /todo/add")
		helloComponent := templates.TodoForm()
		helloComponent.Render(context.Background(), c.Response().Writer)
		return nil
	})

	e.POST("/todo/create", func(c echo.Context) error {
		title := c.FormValue("title")
		description := c.FormValue("description")
		fmt.Println(title, description)
		todo := models.NewTodo(title, description, false)
		todoList = append(todoList, *todo)
		helloComponent := templates.TodoList(todoList)
		helloComponent.Render(context.Background(), c.Response().Writer)
		return nil
	})

	e.POST("/todo/status", func(c echo.Context) error {
		guid := c.FormValue("guid")
		completed := c.FormValue("completed")

		fmt.Printf("GUID: %s, STATUS: %v \n", guid, completed)
		for i, todo := range todoList {
			if todo.Guid == guid {
				todoList[i].Completed = completed == "true"
			}
		}
		return nil
	})

	e.Static("/css", "css")
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":8083"))
}
