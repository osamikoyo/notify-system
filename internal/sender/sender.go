package sender

import (
	"context"
	"sync"

	"github.com/bytedance/sonic"
	"github.com/koyo-os/notify-system/internal/config"
	"github.com/koyo-os/notify-system/internal/models"
	"github.com/koyo-os/notify-system/pkg/logger"
	"go.uber.org/zap/zapcore"
)

type Sender interface{
	Send(models.Notify, context.Context) error
}

type SenderManager struct{
	MessageChan chan []byte
	senders []Sender
	logger *logger.Logger
	cfg *config.Config
	wg *sync.WaitGroup
	ctx context.Context
}

func Init(cfg *config.Config, ch chan []byte, logger *logger.Logger, wg *sync.WaitGroup) *SenderManager {
	var senders []Sender

	if cfg.TgCfg.Use {
		senders = append(senders, InitTgsender(cfg, logger))
	}
	if cfg.EmailCfg.Use {
		senders = append(senders, InitEmailSender(cfg, logger)) 
	}
	if cfg.SmsCfg.Use {
		senders = append(senders, InitSmsSender(cfg, logger))
	}

	return &SenderManager{
		MessageChan: ch,
		cfg: cfg,
		logger: logger,
		senders: senders,
		wg: wg,
	}
}

func (s *SenderManager) Listen() {
	s.logger.Info("sender manager listen...")

	for {
		msg := <- s.MessageChan
		
		var notify models.Notify

		if err := sonic.Unmarshal(msg, &notify);err != nil{
			s.logger.Error("cant unmarshal notify", zapcore.Field{
				Key: "err",
				String: err.Error(),
			})

			return
		}

		for _, sender := range s.senders{
			s.wg.Add(1)
			go func() {
				if err := sender.Send(notify, s.ctx);err != nil{
					s.logger.Error("cant send notify", zapcore.Field{
						Key: "err",
						String: err.Error(),
					})
				}

				s.wg.Done()
			}()
		}

		s.wg.Wait()
	}
}