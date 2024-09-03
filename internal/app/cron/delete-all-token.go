package cron

import (
	"context"
	tokensService "github.com/RajabovIlyas/golang-crud/internal/app/services/tokens-service"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/robfig/cron/v3"
)

type CronJob struct {
	ts *tokensService.TokensService
}

func NewCronService(db *database.Queries) *CronJob {
	return &CronJob{ts: tokensService.NewTokensService(db)}
}

func (cs *CronJob) DeleteAllToken() {
	c := cron.New()
	_, err := c.AddFunc("0 0 * * *", func() {

		err := cs.ts.DeleteOldTokens(context.Background())
		if err != nil {
			return
		}
	})
	if err != nil {
		return
	}
	c.Start()
}
