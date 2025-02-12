package main

import (
	"encoding/json"
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
	mlog "mall/log"
	"os"
)

var mp map[int]string
var logger *logrus.Logger

type MessageHandler struct {
}

func (h *MessageHandler) HandleMessage(message *nsq.Message) error {
	var msg mlog.LogBody

	err := json.Unmarshal(message.Body, &msg)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to unmarshal message")
		return nil
	}
	output(msg)

	return nil

}
func output(msg mlog.LogBody) {
	//fmt.Println(msg.Name, time.Now().Format("2006-01-02 15:04:05"), mp[msg.Level]+":", msg.Message)
	logger.WithFields(logrus.Fields{
		"name":    msg.Name,
		"level":   mp[msg.Level],
		"message": msg.Message,
	}).Info("Log received")
}
func main() {
	logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{}) // 设置为 JSON 格式
	logger.SetLevel(logrus.InfoLevel)

	file, err := os.OpenFile("log/runtime_log/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.SetOutput(file)
	} else {
		logger.Fatal("Failed to open log file")
	}

	consumer, err := nsq.NewConsumer("log", "output", nsq.NewConfig())
	if err != nil {
		panic(err.Error())
	}
	consumer.AddHandler(&MessageHandler{})
	if err = consumer.ConnectToNSQD("127.0.0.1:4150"); err != nil {
		panic(err.Error())
	}
	mp = map[int]string{
		mlog.Debug: "DEBUG",
		mlog.Info:  "INFO",
		mlog.Warn:  "WARN",
		mlog.Error: "ERROR",
	}
	logger.Info("log start.....")
	select {}
}
