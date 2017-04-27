package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"net/http"
	"strconv"
)

type (
	task struct {
		Id    int    `json:"id"`
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

func main() {
	e := echo.New()

	todo[0] = &task{}

	todo[0].Title = "test"

	e.GET("/todo/:id", getTask)

	e.Run(standard.New(":9999"))
}
