package tailfile

import (
	"fmt"

	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
)

var (
	TailObj *tail.Tail
)

func Init(filename string) (err error) {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	TailObj, err := tail.TailFile(filename, config)
	if err != nil {
		logrus.Error("tailfile: create tailobj for path: %s failed, err:%v\n", filename, err)
		return
	}
	fmt.Println(TailObj)
	if err != nil {
		logrus.Error("run failed,err: ", err)
		return
	}

	// for {
	// 	msg, ok := <-TailObj.Lines
	// 	if !ok {
	// 		logrus.Warn("tail file reclose reopen, filename:%s\n", TailObj.Filename)
	// 		time.Sleep(time.Second)
	// 		continue
	// 	}
	// 	fmt.Println("Debug Message")
	// 	fmt.Println("msg: ", msg.Text)
	// }
	return
}
