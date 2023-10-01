package tailfile

import (
	"workspace/common"

	"github.com/sirupsen/logrus"
)

// tail file manager

var (
	ttMgr *tailTaskMgr
)

type tailTaskMgr struct {
	tailTaskMap      map[string]*tailTask
	collectEntryList []common.CollectEntry
	confChan         chan []common.CollectEntry
}

func Init(allConf []common.CollectEntry) (err error) {

	ttMgr = &tailTaskMgr{
		tailTaskMap:      make(map[string]*tailTask, 20),
		collectEntryList: allConf,
		confChan:         make(chan []common.CollectEntry),
	}

	for _, conf := range allConf {

		tt := newTailTask(conf.Path, conf.Topic)
		err = tt.Init()
		if err != nil {
			logrus.Errorf("create tailObj for path: %s failed, err:%v\n", conf.Path, err)
			continue
		}
		logrus.Infof("create a tail task for path: %s success\n", conf.Path)
		ttMgr.tailTaskMap[conf.Path] = tt
		go tt.run()
	}

	go ttMgr.watch()

	return
}

func (t *tailTaskMgr) isExited(conf common.CollectEntry) bool {
	_, ok := t.tailTaskMap[conf.Path]
	return ok

}

func (t *tailTaskMgr) watch() {
	for {
		newConf := <-t.confChan
		logrus.Infof("receive new conf: %v\n", newConf)
		for _, conf := range newConf {
			if t.isExited(conf) {
				continue
			}
			tt := newTailTask(conf.Path, conf.Topic)
			err := tt.Init()
			if err != nil {
				logrus.Errorf("create tailObj for path: %s failed, err:%v\n", conf.Path, err)
			}
			ttMgr.tailTaskMap[conf.Path] = tt
			go tt.run()

		}
	}
}

func SendnewConf(newConf []common.CollectEntry) {
	ttMgr.confChan <- newConf

}
