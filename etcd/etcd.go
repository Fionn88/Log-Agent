package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"tailfile"
	"time"
	"workspace/common"

	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	client *clientv3.Client
)

func Init(address []string) (err error) {
	client, err = clientv3.New(clientv3.Config{
		Endpoints: address,
	})
	if err != nil {
		fmt.Println("connect to etcd failed, err:%v", err)
		return

	}
	return

}

func GetConf(key string) (conf []common.CollectEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	resp, err := client.Get(ctx, key)
	if err != nil {
		logrus.Errorf("get conf from etcd by key:%s failed, err:%v", key, err)
		return nil, err
	}
	if len(resp.Kvs) == 0 {
		logrus.Warnf("get conf from etcd by key:%s failed, err:%v", key, err)
		return nil, fmt.Errorf("no data found in etcd")
	}
	ret := resp.Kvs[0]
	// ret.Value
	fmt.Println(string(ret.Value))
	err = json.Unmarshal(ret.Value, &conf)
	if err != nil {
		logrus.Errorf("json unmarshal failed, err:%v", err)
		return nil, err
	}
	return conf, nil
}

func WatchConf(key string) {
	watch := client.Watch(context.Background(), key)
	var newConf []common.CollectEntry
	for wresp := range watch {
		logrus.Infof("watch key:%s", key)
		for _, evt := range wresp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", evt.Type, evt.Kv.Key, evt.Kv.Value)
			err := json.Unmarshal(evt.Kv.Value, &newConf)
			if err != nil {
				logrus.Errorf("json unmarshal new conf failed, err:%v", err)
				continue
			}
			tailfile.SendnewConf(newConf)
		}
	}

}
