package logic

import (
	"eqregister/model"
	"fmt"
	"testing"
)

func TestCalculateScore(t *testing.T) {

	score, err := CalculateScore([]string{"42dca56a-68ff-4c2d-9562-b24d701d1e99"}, []model.Question{{ID: "42dca56a-68ff-4c2d-9562-b24d701d1e99", Correct: 2}})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(score)
}

func TestUpdateQuestionByID(t *testing.T) {
	err := UpdateQuestionByID(model.Question{ID: "42dca56a-68ff-4c2d-9562-b24d701d1e99",
		Text:    "new question",
		Answers: []string{"first answer", "second answer", "third answer"},
		Correct: 2})
	if err != nil {
		t.Error(err)
	}
}
