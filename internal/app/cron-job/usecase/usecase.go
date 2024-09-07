package usecase

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/internal/app/cron-job"
	"github.com/RajabovIlyas/golang-crud/internal/app/token"
	"github.com/robfig/cron/v3"
)

type cronUC struct {
	tokenUC token.UseCase
}

func NewCronUC(tokenUC token.UseCase) cronJob.UseCase {
	return &cronUC{tokenUC}
}

func (c cronUC) DeleteAllToken() {
	newCron := cron.New()
	_, err := newCron.AddFunc("0 0 * * *", func() {

		err := c.tokenUC.DeleteOldTokens(context.Background())
		if err != nil {
			return
		}
	})
	if err != nil {
		return
	}
	newCron.Start()
}
