package main

import (
	"fmt"
	"kafka"
	"tailfile"
	"workspace/etcd"

	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
)

type Config struct {
	KafkaConfig   `ini:"kafka"`
	CollectConfig `ini:"collect"`
	EtcdConfig    `ini:"etcd"`
}

type EtcdConfig struct {
	Address    string `ini:"address"`
	CollectKey string `ini:"collect_key"`
}

type KafkaConfig struct {
	Address  string `ini:"address"`
	Topic    string `ini:"topic"`
	ChanSize int64  `ini:"chan_size`
}

type CollectConfig struct {
	LogFilePath string `ini:"logfile_path"`
}

func run() {
	// logfile => TailObj => log => client => kafka

	for {

	}

}

func main() {

	// Read The Config File
	var configObj = new(Config)
	err := ini.MapTo(configObj, "config.ini")
	if err != nil {
		logrus.Errorf("load config.ini failed,err:%v", err)
		return
	}
	// cfg, err := ini.Load("config.ini")
	// if err != nil {
	// 	logrus.Error("load config.ini failed,err:%v", err)
	// 	return
	// }
	// kafkaAddress := cfg.Section("kafka").Key("address").String()
	// fmt.Println(kafkaAddress)

	// Init Kafka
	err = kafka.Init([]string{configObj.KafkaConfig.Address}, configObj.KafkaConfig.ChanSize)
	if err != nil {
		logrus.Errorf("init kafka failed,err:%v", err)
		return
	}
	logrus.Info("init kafka success")

	// Init Etcd
	err = etcd.Init([]string{configObj.EtcdConfig.Address})
	if err != nil {
		logrus.Errorf("init etcd failed,err:%v", err)
		return
	}

	allConf, err := etcd.GetConf(configObj.EtcdConfig.CollectKey)
	if err != nil {
		logrus.Errorf("get config from etcd failed,err:%v", err)
		return
	}
	fmt.Println(allConf)

	go etcd.WatchConf(configObj.EtcdConfig.CollectKey)

	// Init TailFil
	err = tailfile.Init(allConf)
	if err != nil {
		logrus.Errorf("init tailfile failed,err:%v", err)
		return
	}
	logrus.Info("init tailfile success")

	run()

}
