package tailfile

import (
	"context"
	"fmt"
	"kafka"
	"time"

	"github.com/Shopify/sarama"
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
)

type tailTask struct {
	path   string
	topic  string
	tObj   *tail.Tail
	ctx    context.Context
	cancel context.CancelFunc
}

func newTailTask(path, topic string) *tailTask {
	ctx, cancel := context.WithCancel(context.Background())

	tt := &tailTask{
		path:   path,
		topic:  topic,
		ctx:    ctx,
		cancel: cancel,
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
		select {
		case <-t.ctx.Done():
			logrus.Infof("%s is topping...", t.path)
			return
		case line, ok := <-t.tObj.Lines:
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

}
