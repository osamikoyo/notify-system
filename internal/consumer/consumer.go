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
	client sarama.PartitionConsumer
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

	pititionClient,err := client.ConsumePartition(
		cfg.Topic,
		0,
		sarama.OffsetOldest,
	)

	if err != nil{
		logger.Error("cant get partition consumer: ", zapcore.Field{
			Key: "err",
			String: err.Error(),
		})
		
		return nil, err
	}

	var wg sync.WaitGroup

	return &Comsumer{
		client: pititionClient,
		wg: &wg,
		logger: logger,
		MessageChan: ch,
	}, nil
}

func (c *Comsumer) Listen() {
	c.wg.Add(1)

	c.logger.Info("starting kafka consumer...")

	go func ()  {
		defer c.wg.Done()

		for msg := range c.client.Messages(){
			c.MessageChan <- msg.Value
		}
	}()

	c.wg.Wait()
}