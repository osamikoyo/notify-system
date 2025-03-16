package sender

import (
	"context"

	"github.com/bytedance/sonic"
	"github.com/koyo-os/notify-system/internal/config"
	"github.com/koyo-os/notify-system/internal/models"
	"github.com/koyo-os/notify-system/pkg/logger"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
	"go.uber.org/zap/zapcore"
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

func (s *SmsSender) Send(msg models.Notify, _ context.Context) error {
	s.logger.Info("start setup tg message")

	s.params.SetTo(msg.To)
	s.params.SetBody(msg.Message)

	if s.cfg.SmsCfg.PrintResp {
		resp, err := s.client.Api.CreateMessage(s.params)
		if err != nil{
			s.logger.Error("error to create twilio message", zapcore.Field{
				Key: "err",
				String: err.Error(),
			})

			return err
		}

		body, _ := sonic.Marshal(resp)

		s.logger.Info("resp from twilio", zapcore.Field{
			Key: "resp",
			String: string(body),
		})
		
		return nil
	} else {
		_, err := s.client.Api.CreateMessage(s.params)
		if err != nil{
			s.logger.Error("error to create twilio message", zapcore.Field{
				Key: "err",
				String: err.Error(),
			})

			return err
		}
		return nil
	}
}

