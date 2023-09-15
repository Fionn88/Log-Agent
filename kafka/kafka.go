package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

var (
	Client sarama.SyncProducer
)

func Init(address []string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	Client, err = sarama.NewSyncProducer(address, config)
	if err != nil {
		logrus.Error("kafka: producer close, err:", err)
		return
	}
	return
}
