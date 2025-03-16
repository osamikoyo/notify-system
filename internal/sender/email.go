package sender

import (
	"context"

	"github.com/koyo-os/notify-system/internal/config"
	"github.com/koyo-os/notify-system/internal/models"
	"github.com/koyo-os/notify-system/pkg/logger"
	"gopkg.in/gomail.v2"
)

type EmailSender struct{
	dialer *gomail.Dialer
	logger *logger.Logger
	cfg *config.Config
}

func InitEmailSender(cfg *config.Config, logger *logger.Logger) *EmailSender {
	dialer := gomail.NewDialer(cfg.EmailCfg.Host, cfg.EmailCfg.Port, cfg.EmailCfg.Username, cfg.EmailCfg.Password)

	return &EmailSender{
		dialer: dialer,
		logger: logger,
		cfg: cfg,
	}
}

func (e *EmailSender) Send(msg models.Notify, _ context.Context) error {
	m := gomail.NewMessage()

	m.SetHeader("From", e.cfg.EmailCfg.From)
	m.SetHeader("To", msg.To)
	m.SetHeader("Subject", msg.Subject)
}