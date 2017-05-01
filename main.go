package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"net/http"
	"strconv"
)

type (
	task struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	}
)

var (
	todo = map[int]*task{}
	seq  = 1
)

func getTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, todo[id])
}

func updateTask(c echo.Context) error {
	t := new(task)
	err := c.Bind(t)

	if err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	todo[id].Title = t.Title
	return c.JSON(http.StatusOK, todo[id])
}

func createTask(c echo.Context) error {
	t := &task{
		ID: seq,
	}

	err := c.Bind(t)

	if err != nil {
		return err
	}
	todo[t.ID] = t
	seq++
	return c.JSON(http.StatusCreated, t)
}

func deleteTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(todo, id)
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()

	todo[0] = &task{}

	todo[0].Title = "test"

	e.GET("/todo/:id", getTask)
	e.PUT("/todo/:id", updateTask)
	e.POST("/todo/", createTask)
	e.DELETE("todo/:id", deleteTask)

	e.Run(standard.New(":9999"))
}
