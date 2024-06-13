package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type ToDo struct {
	ID        int    `json:"id"`
	Body      string `json:"body"`
	Completed bool   `json:"completed"`
}

func main() {
	e := echo.New()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}
	PORT := os.Getenv("port")
	todos := []ToDo{}
	e.GET("/api/todos", func(c echo.Context) error {
		return c.JSON(200, todos)
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
	e.DELETE("/api/todos/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		for i, todo := range todos {
			if todo.ID == id {
				todos = append(todos[:1], todos[i+1:]...)
				return c.String(200, "success")
			}
		}
		return c.String(http.StatusNotFound, "Todo not found")
	})
	e.Logger.Fatal(e.Start(":" + PORT))
}
