package sender

import (
	"context"

	"github.com/koyo-os/notify-system/internal/config"
	"github.com/koyo-os/notify-system/internal/models"
	"github.com/koyo-os/notify-system/pkg/logger"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

type SmsSender struct{
	cfg *config.Config
	logger *logger.Logger
	client *twilio.RestClient
	params *api.CreateMessageParams
}

func InitSmsSender(cfg *config.Config, logger *logger.Logger) *SmsSender {
	client := twilio.NewRestClient()
	params := &api.CreateMessageParams{}

	params.SetFrom(cfg.SmsCfg.From)

	return &SmsSender{
		logger: logger,
		cfg: cfg,
		client: client,
		params: params,
	}
}

func (s *SmsSender) Send(msg *models.Notify, _ context.Context) error {
	
}

