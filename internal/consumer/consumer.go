package consumer

import (
	"sync"

	"github.com/IBM/sarama"
	"github.com/koyo-os/notify-system/internal/config"
	"github.com/koyo-os/notify-system/pkg/logger"
	"go.uber.org/zap/zapcore"
)

type Comsumer struct{
	MessageChan chan []byte
	wg *sync.WaitGroup
	client *sarama.Consumer
	logger *logger.Logger
}

func Init(cfg *config.Config, ch chan []byte) (*Comsumer, error){
	logger := logger.Init()

	config := sarama.NewConfig()

	config.Consumer.Return.Errors = true

	client, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil{
		logger.Error("cant get consumer: ", zapcore.Field{
			Key: "err",
			String: err.Error(),
		})
	}

	var wg sync.WaitGroup

	return &Comsumer{
		client: &client,
		wg: &wg,
		logger: logger,
		MessageChan: ch,
	}, nil
}