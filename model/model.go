package model

import (
	"database/sql"
	"encoding/json"
	logger "eqregister/log"
	"errors"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"strings"
)

type AppStruct struct {
	DBConn *sql.DB
}

type Question struct {
	ID      string        `json:"id"`
	Text    string        `json:"text"`
	Answers []string      `json:"answers"`
	Correct int           `json:"correct"`
}

var db *sql.DB

func init() {
	conn, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", "localhost", "5432", "eqregister", "eqregister", "questions"))
	db = conn
	if err != nil {
		logger.Error(err)
	}
}

func GetApp() *AppStruct {
	conn, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", "localhost", "5432", "eqregister", "eqregister", "questions"))
	if err != nil {
		logger.Error(err)
	}

	db = conn
	return &AppStruct{DBConn: db}
}

func (a *AppStruct) GetAnswers(questionIds []string) ([]Question, error) {
	if len(questionIds) == 0 {
		return nil, nil
	}
	questionString := strings.Join(questionIds, "','")
	result, err := a.DBConn.Query(fmt.Sprintf("select question_uuid,question_text,answer_json,correct_answer from questions where question_uuid in(%s)", "'"+questionString+"'"))
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	var answers []Question
	for result.Next() {
		var question_uuid sql.NullString
		var question_text sql.NullString
		var answer_json sql.NullString
		var correct_answer sql.NullInt64
		var q Question
		if err := result.Scan(&question_uuid, &question_text, &answer_json, &correct_answer); err != nil {
			logger.Error(err)
			return nil, err
		}

		if question_uuid.Valid {
			q.ID = question_uuid.String
		}

		if question_text.Valid {
			q.Text = question_text.String
		}

		if answer_json.Valid {
			var answerSlice []string

			if err := json.Unmarshal([]byte(answer_json.String), &answerSlice); err != nil {
				logger.Error(err)
			}
			q.Answers = answerSlice
		}

		if correct_answer.Valid {
			q.Correct = int(correct_answer.Int64)
		}
		answers = append(answers, q)

	}
	return answers, nil
}

func (a *AppStruct) GetQuestionByID(ID string) (*Question, error) {
	result := a.DBConn.QueryRow(fmt.Sprintf("select question_uuid,question_text,answer_json,correct_answer from questions where question_uuid='%s'", ID))
	var question_uuid sql.NullString
	var question_text sql.NullString
	var answer_json sql.NullString
	var correct_answer sql.NullInt64

	if err := result.Scan(&question_uuid, &question_text, &answer_json, &correct_answer); err != nil {
		logger.Error(err)
	}

	var q Question
	if question_uuid.Valid {
		q.ID = question_uuid.String
	}

	if question_text.Valid {
		q.Text = question_text.String
	}

	if answer_json.Valid {
		var answerSlice []string

		if err := json.Unmarshal([]byte(answer_json.String), &answerSlice); err != nil {
			logger.Error(err)
		}
		q.Answers = answerSlice
	}

	if correct_answer.Valid {
		q.Correct = int(correct_answer.Int64)
	}

	return &q, nil
}

func (a *AppStruct) IsValidID(ID string) (bool, error) {

	question, err := a.GetQuestionByID(ID)
	if err != nil {
		return false, err
	}
	if question.ID == "" {
		return false, nil
	}
	return true, nil
}

func (a *AppStruct) UpdateQuestion(question Question) (*Question,error) {
	if question.ID == "" {
		return nil,errors.New("invalid id")
	}

	if len(question.Answers) < question.Correct || question.Correct < 0 {
		return nil,errors.New(" invalid correct answer index")
	}

	answerJson, err := json.Marshal(question.Answers)
	if err != nil {
		return nil,err
	}
	query := fmt.Sprintf("update questions set question_text='%s', answer_json='%s', correct_answer='%d' where question_uuid='%s'", question.Text, string(answerJson), question.Correct, question.ID)

	_, err = a.DBConn.Exec(query)
	if err != nil {
		return nil,err
	}
	return &question,nil
}

func (a *AppStruct) InsertQuestion(q *Question) (*Question,error) {
	
    if len(q.Answers) < q.Correct || q.Correct < 0 {
		return nil,errors.New(" invalid correct answer index")
	}
	answers, err := json.Marshal(q.Answers)
	if err != nil {
		return nil,err
	}
	q.ID = uuid.NewString()
	_, err = a.DBConn.Exec(fmt.Sprintf("INSERT into questions (question_uuid,question_text,answer_json,correct_answer) values ( '%s', '%s', '%s', '%d' );", q.ID, q.Text, string(answers), q.Correct))
	if err != nil {
		return nil,err
	}
	return q,nil
}
