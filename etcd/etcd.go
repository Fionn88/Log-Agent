package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type collectEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

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

func GetConf(key string) ([]collectEntry, error) {
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
	var collectEntries []collectEntry
	err = json.Unmarshal(ret.Value, &collectEntries)
	if err != nil {
		logrus.Errorf("json unmarshal failed, err:%v", err)
		return nil, err
	}
	return collectEntries, nil
}
