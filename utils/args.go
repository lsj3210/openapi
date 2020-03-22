package utils

import (
	"flag"
	"fmt"
	"os"
)

var (
	help bool
	conf string
)

//命令行参数说明
func _usage() {
	fmt.Fprintf(os.Stderr, "user center version: v1.0.0\n"+
		"Usage: ucenter [-c filename]\n\n"+
		"Options:\n")
	flag.PrintDefaults()
}

//解析命令行参数
func initArgs() {
	flag.BoolVar(&help, "h", false, "Print help info.")
	flag.StringVar(&conf, "c", "", "Please Input config file.")
	flag.Usage = _usage
	flag.Parse()
	if help {
		flag.Usage()
		os.Exit(-1)
	}
	if conf == "" {
		flag.Usage()
		os.Exit(-1)
	}
}
