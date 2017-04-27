package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"net/http"
	"strconv"
)

type (
	task struct {
		id    int    `json:"id"`
		title string `json:"title"`
	}
)

var (
	todo = map[int]*task{}
	seq  = 1
)

func getTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, *todo[id])
}

func main() {
	e := echo.New()

	todo[0] = &task{
		id:    0,
		title: "ichigo",
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/todo/:id", getTask)

	e.Run(standard.New(":9999"))
}
