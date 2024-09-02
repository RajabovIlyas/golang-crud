package cron

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/robfig/cron/v3"
)

type CronJob struct {
	db *database.Queries
}

func NewCronService(db *database.Queries) *CronJob {
	return &CronJob{db}
}

func (cs *CronJob) DeleteAllToken() {
	c := cron.New()
	_, err := c.AddFunc("0 0 * * *", func() {

		err := cs.db.DeleteOldTokens(context.Background())
		if err != nil {
			return
		}
	})
	if err != nil {
		return
	}
	c.Start()
}
