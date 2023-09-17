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
	return
}
