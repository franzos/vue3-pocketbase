package main

import (
	"log"

	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*Log)(nil)

type Log struct {
	models.BaseModel

	Message    string `db:"message" json:"message"`
	Topic      string `db:"topic" json:"topic"`
	Data       string `db:"data" json:"data"`
	Level      int    `db:"level" json:"level"`
	Forwarder  string `db:"forwarder" json:"forwarder"`
	Submission string `db:"submission" json:"submission"`
}

type LogsApi struct {
	dao  *daos.Dao
	data Log
}

func (m *Log) TableName() string {
	return "logs"
}

func createLog(api LogsApi) error {
	newLog := &Log{
		Message:    api.data.Message,
		Topic:      api.data.Topic,
		Data:       api.data.Data,
		Level:      api.data.Level,
		Forwarder:  api.data.Forwarder,
		Submission: api.data.Submission,
	}

	if err := api.dao.Save(newLog); err != nil {
		log.Printf("Error saving log: %v\n", err)
		return err
	}

	return nil
}
