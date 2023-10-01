package tailfile

import (
	"fmt"
	"kafka"
	"time"

	"github.com/Shopify/sarama"
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
)

type tailTask struct {
	path  string
	topic string
	tObj  *tail.Tail
}

func newTailTask(path, topic string) *tailTask {

	tt := &tailTask{
		path:  path,
		topic: topic,
	}
	return tt

}

func (t *tailTask) Init() (err error) {
	cfg := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	t.tObj, err = tail.TailFile(t.path, cfg)
	return

}

func (t *tailTask) run() {
	logrus.Infof("collect for path: %s is running...", t.path)
	for {
		line, ok := <-t.tObj.Lines
		if !ok {
			logrus.Warn("tail file reclose reopen, path:%s\n", t.path)
			time.Sleep(time.Second)
			continue
		}
		if len(line.Text) == 0 {
			continue
		}
		fmt.Println(line.Text)
		msg := &sarama.ProducerMessage{}
		msg.Topic = t.topic
		msg.Value = sarama.StringEncoder(line.Text)

		kafka.ToMsgChan(msg)
	}

}
