package cron

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	tokensService "github.com/RajabovIlyas/golang-crud/internal/app/services/tokens-service"
	"github.com/robfig/cron/v3"
)

type CronJob struct {
	ts *tokensService.TokensService
}

func NewCronService(p *models.DBConfigParam) *CronJob {
	return &CronJob{ts: tokensService.NewTokensService(p)}
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
