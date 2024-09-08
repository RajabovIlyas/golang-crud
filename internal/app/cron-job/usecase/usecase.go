package usecase

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/internal/app/cron-job"
	"github.com/RajabovIlyas/golang-crud/internal/app/token"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog"
)

type cronUC struct {
	tokenUC token.UseCase
	logger  zerolog.Logger
}

func NewCronUC(tokenUC token.UseCase, logger zerolog.Logger) cronJob.UseCase {
	return &cronUC{tokenUC, logger}
}

func (c cronUC) DeleteAllToken() {
	newCron := cron.New()
	_, err := newCron.AddFunc("0 0 * * *", func() {

		err := c.tokenUC.DeleteOldTokens(context.Background())
		if err != nil {
			c.logger.Error().Err(err).Msg("cronUC.DeleteAllToken: error deleting old tokens")
			return
		}
	})
	if err != nil {
		c.logger.Error().Err(err).Msg("cronUC.DeleteAllToken: error adding new cron")
		return
	}
	newCron.Start()
}
