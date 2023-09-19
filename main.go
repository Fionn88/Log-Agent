package main

import (
	"fmt"
	"kafka"
	"tailfile"
	"time"

	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
)

type Config struct {
	KafkaConfig   `ini:"kafka"`
	CollectConfig `ini:"collect"`
}

type KafkaConfig struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type CollectConfig struct {
	LogFilePath string `ini:"logfile_path"`
}

func run() (err error) {

	for {
		msg, ok := <-tailfile.TailObj.Lines
		if !ok {
			logrus.Warn("tail file reclose reopen, filename:%s\n", tailfile.TailObj.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("Debug Message")
		fmt.Println("msg: ", msg.Text)
	}

}

func main() {

	// Read The Config File
	var configObj = new(Config)
	err := ini.MapTo(configObj, "config.ini")
	if err != nil {
		logrus.Error("load config.ini failed,err: ", err)
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
	err = kafka.Init([]string{configObj.KafkaConfig.Address})
	if err != nil {
		logrus.Error("init kafka failed,err: ", err)
		return
	}
	logrus.Info("init kafka success")

	// Init TailFil
	err = tailfile.Init(configObj.CollectConfig.LogFilePath)
	if err != nil {
		logrus.Error("init tailfile failed,err: ", err)
	}
	logrus.Info("init tailfile success")

	err = run()
	if err != nil {
		logrus.Error("run failed,err: ", err)
		return
	}

}
