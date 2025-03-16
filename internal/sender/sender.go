package sender

import (
	"context"

	"github.com/koyo-os/notify-system/internal/config"
	"github.com/koyo-os/notify-system/internal/models"
	"github.com/koyo-os/notify-system/pkg/logger"
)

type Sender interface{
	Send(models.Notify, context.Context) error
}

type SenderManager struct{
	MessageChan chan []byte
	senders []Sender
	logger *logger.Logger
	cfg *config.Config
}

func Init(cfg *config.Config, ch chan []byte, logger *logger.Logger) (*SenderManager, error) {
	var senders []Sender

	if cfg.TgCfg.Use {
		senders = append(senders, InitTgsender(cfg, logger))
	}
	if cfg.EmailCfg.Use {
		senders = append(senders, InitEmailSender(cfg, logger)) 
	}
	if cfg.SmsCfg.Use {
		
	}
}