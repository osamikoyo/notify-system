package app

import (
	"fmt"
	"sync"
	"time"

	"github.com/koyo-os/notify-system/internal/config"
	"github.com/koyo-os/notify-system/internal/consumer"
	"github.com/koyo-os/notify-system/internal/sender"
	"github.com/koyo-os/notify-system/pkg/logger"
	"go.uber.org/zap/zapcore"
)

func App() error {
	counter := 0

	var (
		messageChan chan []byte
		wg sync.WaitGroup
	)
	go func ()  {
		counter++

		if counter == -1{
			return
		}

		time.Sleep(time.Second)
	}()

	fmt.Print(`
 __    _  _______  _______  ___   _______  __   __ 
|  |  | ||       ||       ||   | |       ||  | |  |
|   |_| ||   _   ||_     _||   | |    ___||  |_|  |
|       ||  | |  |  |   |  |   | |   |___ |       |
|  _    ||  |_|  |  |   |  |   | |    ___||_     _|
| | |   ||       |  |   |  |   | |   |      |   |  
|_|  |__||_______|  |___|  |___| |___|      |___|  
`)

	logger := logger.Init()

	logger.Info("start notify...")

	logger.Info("getting config.yaml...")

	cfg, err := config.Init()
	if err != nil{
		logger.Error("cant get config, stopped...", zapcore.Field{
			Key: "err",
			String: err.Error(),
		})
		return err
	}

	logger.Info("get consumer...")

	consumer, err := consumer.Init(cfg, messageChan)
	if err != nil{
		logger.Error("error get consumer, stopped", zapcore.Field{
			Key: "err",
			String: err.Error(),
		}, zapcore.Field{
			Key: "url",
			String: cfg.KafkaUrl,
		})
	}

	logger.Info("get masanger..,")

	manager := sender.Init(
		cfg,
		messageChan,
		logger,
		&wg,
	)

	logger.Info("start listeners...")

	manager.Listen()
	go consumer.Listen()

	return nil
}