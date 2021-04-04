package logic

import (
	"eqregister/log"
	"eqregister/model"
	"errors"
	"fmt"
)

var app *model.AppStruct

func init() {

	app = model.GetApp()
}

func CalculateScore(questionIds []string, candidateAnswers []model.Question) (int, error) {
	correctAnswers, err := app.GetAnswers(questionIds)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	trackScore := make(map[string]model.Question)
	for _, q := range correctAnswers {
		trackScore[q.ID] = q
	}
	var answerdCorrectly int
	for _, ans := range candidateAnswers {
		if correct, ok := trackScore[ans.ID]; ok {
			if correct.Correct == ans.Correct {
				answerdCorrectly++
			}

		} else {
			log.Error(errors.New(fmt.Sprintf("%s not found in db", ans.ID)))
			return 0, errors.New("invalid id found" + fmt.Sprint(ans.ID))
		}
	}
	return answerdCorrectly, nil
}

func GetQuestionByID(questionID string) (*model.Question, error) {
	if questionID == "" {
		return nil, errors.New("id is empty")
	}
	return app.GetQuestionByID(questionID)
}
