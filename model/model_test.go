package model

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	app := GetApp()
	err := app.InsertQuestion(&Question{
		Text:    "what is the name of this",
		Answers: []string{"something new", "something differrent", "names that are familiar"},
		Correct: 2,
	})
	if err != nil {
		t.Error(err)
	}
}

func TestGetQuestionByID(t *testing.T) {
	app := GetApp()
	question, err := app.GetQuestionByID("42dca56a-68ff-4c2d-9562-b24d701d1e99")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(question)
}
