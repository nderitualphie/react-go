package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ToDo struct {
	ID        int    `json:"id"`
	Body      string `json:"body"`
	Completed bool   `json:"completed"`
}

func main() {
	e := echo.New()
	todos := []ToDo{}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
	e.POST("/api/todos", func(c echo.Context) error {
		todo := ToDo{}
		if err := c.Bind(&todo); err != nil {
			return err
		}
		if todo.Body == "" {
			return c.String(http.StatusBadRequest, "body cannot be empty")
		}
		todo.ID = len(todos) + 1
		todos = append(todos, todo)
		return c.JSON(201, todo)
	})
	//UPDATE TODO
	e.PATCH("/api/todos/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		for i, todo := range todos {
			if todo.ID == id {
				todos[i].Completed = true
				return c.JSON(200, todos[i])
			}
		}
		return c.String(http.StatusNotFound, "Todo not found")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
