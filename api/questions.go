package api

import (
	"eqregister/logic"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GetByIDRequest struct {
	ID string `param:"id"validate:"required"`
}

func GetQuestionByIDHandler(c echo.Context) error {
	var req GetByIDRequest
	if err := (&echo.DefaultBinder{}).BindPathParams(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	question, err := logic.GetQuestionByID(req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, question)
}
