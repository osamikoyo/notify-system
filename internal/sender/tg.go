package sender

import (
	"context"

	"github.com/koyo-os/notify-system/internal/config"
	"github.com/koyo-os/notify-system/internal/models"
	"github.com/koyo-os/notify-system/pkg/logger"
	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/telegram"
	"go.uber.org/zap/zapcore"
)

type TgSender struct{
	client *telegram.Telegram
	logger *logger.Logger
	cfg *config.Config
}

func InitTgsender(cfg *config.Config, logger *logger.Logger) *TgSender {
	service, err := telegram.New(cfg.TgCfg.Token)
	if err != nil{
		logger.Error("cant get telegramm sender", zapcore.Field{
			Key: "err",
			String: err.Error(),
		})

		return nil
	}

	service.AddReceivers(cfg.TgCfg.ChatId)

	notify.UseServices(service)

	return &TgSender{
		client: service,
		logger: logger,
		cfg: cfg,
	}
}

func (t *TgSender) Send(msg models.Notify, ctx context.Context) error {
	return notify.Send(
		ctx,
		msg.Subject,
		msg.Message,
	)
}