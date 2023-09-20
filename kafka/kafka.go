package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

var (
	client  sarama.SyncProducer
	MsgChan chan *sarama.ProducerMessage
)

func Init(address []string, chansize int64) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer(address, config)
	if err != nil {
		logrus.Error("kafka: producer close, err:", err)
		return
	}
	MsgChan = make(chan *sarama.ProducerMessage, chansize)
	go SendMsg()
	return
}

func SendMsg() {
	for {
		select {
		case msg := <-MsgChan:
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				logrus.Warning("send message failed,", err)
				return
			}
			logrus.Info("send message to kafka success, pid:", pid, ", offset:", offset)
		}
	}

}
