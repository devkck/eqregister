package api

import (
	"eqregister/logic"
	"eqregister/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GetByIDRequest struct {
	ID string `param:"id"validate:"required"`
}

func GetQuestionByIDHandler(c echo.Context) error {
	var req GetByIDRequest
	if err := (&echo.DefaultBinder{}).BindPathParams(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}
	question, err := logic.GetQuestionByID(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}
	c.JSON(http.StatusOK, question)
	return nil
}

func UpdateQuestionByIDHandler(c echo.Context) error {
	var req GetByIDRequest
	if err := (&echo.DefaultBinder{}).BindPathParams(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}
	var modelBody model.Question
	if err := (&echo.DefaultBinder{}).BindBody(c, &modelBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}
	modelBody.ID = req.ID
	q,err := logic.UpdateQuestionByID(modelBody)
    if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}
	c.JSON(http.StatusOK,q)
	return nil
}

func CalculateScoreHandler(c echo.Context) error {
	var req []model.Question
	if err := (&echo.DefaultBinder{}).BindBody(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}
	score, err := logic.CalculateScore(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}
	c.JSON(http.StatusOK, map[string]interface{}{"score": score})
	return nil
}

func InsertQuestionHandler(c echo.Context) error {
	var req GetByIDRequest
	if err := (&echo.DefaultBinder{}).BindPathParams(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}
	var modelBody model.Question
	if err := (&echo.DefaultBinder{}).BindBody(c, &modelBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}
	q,err := logic.InsertQuestion(&modelBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}
	c.JSON(http.StatusOK, q)
	return nil
}
