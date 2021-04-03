package main

import (
	"database/sql"
	"encoding/json"
	logger "eqregister/log"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type AppStruct struct {
	DBConn *sql.DB
}

type Question struct {
	ID      string
	Text    string
	Answers []string
	Correct int
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

func (a *AppStruct) InsertQuestion(q *Question) error {
	answers, err := json.Marshal(q.Answers)
	if err != nil {
		return err
	}

	q.ID = uuid.NewString()
	_, err = a.DBConn.Exec(fmt.Sprintf("INSERT into questions (question_uuid,question_text,answer_json,correct_answer) values ( '%s', '%s', '%s', '%d' );", q.ID, q.Text, string(answers), q.Correct))
	if err != nil {
		return err
	}
	return nil
}
