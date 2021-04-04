package api

import (
	"github.com/labstack/echo/v4"
	"net/http/httptest"
	"testing"
)

func TestGetQuestionByIDHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest("GET", "/question/6fcb4995-ec59-48bc-923d-1c43a540addb", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if err := GetQuestionByIDHandler(c); err != nil {
		t.Error(err)
	}

}
