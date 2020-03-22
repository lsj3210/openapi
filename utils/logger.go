package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	// Log 打印系统运行日志
	Log *logrus.Entry
)

//初始化日志
func initLog() {
	c := GetConf()
	logPath := c.Log.File
	level := c.Log.Level
	format := c.Log.Format

	tmp := logrus.New()
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("init logger error!")
	} else {
		tmp.Out = file
	}
	//日志打印级别
	if level == "debug" {
		tmp.SetLevel(logrus.DebugLevel)
	} else if level == "info" {
		tmp.SetLevel(logrus.InfoLevel)
	} else if level == "error" {
		tmp.SetLevel(logrus.ErrorLevel)
	}

	//日志格式
	if format == "text" {
		tmp.Formatter = new(logrus.TextFormatter)
	} else if format == "json" {
		tmp.Formatter = new(logrus.JSONFormatter)
	} else {
		tmp.Formatter = new(logrus.TextFormatter)
	}

	if host, err := os.Hostname(); err != nil {
		fmt.Println("get host name error!")
		os.Exit(-1)
	} else {
		Log = logrus.NewEntry(tmp).WithFields(logrus.Fields{"local": host})
	}
}

//获取日志基本信息
func getLogInfo() (file string, fun string) {
	var ok bool
	var pc uintptr
	var line int
	file, fun = "???", "???"
	pc, file, line, ok = runtime.Caller(2)
	if ok {
		fun = runtime.FuncForPC(pc).Name()
		fun = filepath.Ext(fun)
		fun = strings.TrimPrefix(fun, ".")
		file = filepath.Base(file)
		file = fmt.Sprintf("%s:%d", file, line)
	}
	return
}

func ERR(args ...interface{}) {
	file, fun := getLogInfo()
	Log.WithFields(logrus.Fields{"file": file, "func": fun}).Errorln(args)
}

func ERRF(format string, args ...interface{}) {
	file, fun := getLogInfo()
	Log.WithFields(logrus.Fields{"file": file, "func": fun}).Errorf(format, args)
}

func INFO(args ...interface{}) {
	file, fun := getLogInfo()
	Log.WithFields(logrus.Fields{"file": file, "func": fun}).Infoln(args)
}

func INFOF(format string, args ...interface{}) {
	file, fun := getLogInfo()
	Log.WithFields(logrus.Fields{"file": file, "func": fun}).Infof(format, args)
}

func DEBUG(args ...interface{}) {
	file, fun := getLogInfo()
	Log.WithFields(logrus.Fields{"file": file, "func": fun}).Debugln(args)
}

func DEBUGF(format string, args ...interface{}) {
	file, fun := getLogInfo()
	Log.WithFields(logrus.Fields{"file": file, "func": fun}).Debugf(format, args)
}
