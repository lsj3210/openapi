package utils

import "gopkg.in/resty.v1"

func DoHttp() *resty.Request {
	runMode := Conf.Http.RunMode
	if runMode == "pro" {
		return resty.SetDebug(false).R()
	} else {
		return resty.SetDebug(true).R()
	}
}
