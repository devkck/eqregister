package main

import (
	"eqregister/api"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/question/:id", api.GetQuestionByIDHandler)
	e.PUT("/question/:id", api.UpdateQuestionByIDHandler)
	e.POST("/question", api.InsertQuestionHandler)
	e.GET("/score", api.CalculateScoreHandler)
	e.Start(":8181")
}
