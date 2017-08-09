package log

import (
	"log"
	"net"

	"github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func InitLog() {
	logger.Formatter = &logrus.JSONFormatter{}
}

func setLogStash() {
	conn, err := net.Dial("tcp", "logstash.mycompany.net:8911")
	if err != nil {
		log.Fatal(err)
	}
	hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{"type": "myappName"}))

	if err != nil {
		log.Fatal(err)
	}
	logger.Hooks.Add(hook)
	ctx := logger.WithFields(logrus.Fields{
		"method": "main",
	})
	ctx.Info("Hello World!")
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

func Print(args ...interface{}) {
	logger.Println(args...)
}

func Println(args ...interface{}) {
	logger.Println(args...)
}
