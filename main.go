package main

import (
	"eqregister/api"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/question/:id", api.GetQuestionByIDHandler)
	e.Start(":8181")
}