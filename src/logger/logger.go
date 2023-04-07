package logger

import (
	"github.com/Sirupsen/logrus"
)

func InitLogger() {
	var textFormatter = new(logrus.TextFormatter)
	textFormatter.TimestampFormat = "2022-09-20 12:49:12.232"
	textFormatter.FullTimestamp = true

	logrus.SetFormatter(textFormatter)
}

func Printf(format string, v ...interface{}) () {
	logrus.Printf(format, v)
}

func Fatalf(format string, v ...interface{}) () {
	logrus.Fatalf(format, v)
}

func Panicf(format string, v ...interface{}) () {
	logrus.Panicf(format, v)
}

func Debugf(format string, v ...interface{}) () {
	logrus.Debugf(format, v)
}

func Warnf(format string, v ...interface{}) () {
	logrus.Warnf(format, v)
}

func Warningf(format string, v ...interface{}) () {
	logrus.Warningf(format, v)
}

func Errorf(format string, v ...interface{}) () {
	logrus.Errorf(format, v)
}

func Print(v ...interface{}) () {
	logrus.Print(format, v)
}

func Fatal(v ...interface{}) () {
	logrus.Fatal(format, v)
}

func Panic(v ...interface{}) () {
	logrus.Panic(format, v)
}

func Debug(v ...interface{}) () {
	logrus.Debug(format, v)
}

func Warn(v ...interface{}) () {
	logrus.Warn(format, v)
}

func Warning(v ...interface{}) () {
	logrus.Warning(format, v)
}

func Error(v ...interface{}) () {
	logrus.Error(format, v)
}
