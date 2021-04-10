package main

import (
	"eqregister/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.GET("/question/:id", api.GetQuestionByIDHandler)
	e.PUT("/question/:id", api.UpdateQuestionByIDHandler)
	e.POST("/question", api.InsertQuestionHandler)
	e.GET("/score", api.CalculateScoreHandler)
	e.Start(":8181")
}
