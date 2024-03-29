// Path: otpclient
// FileName: errors.go
// Created by dkedTeam
// Author: GJing
// Date: 2023/8/2$ 15:18$

package ks_error

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
)

type OtpError int

func (e OtpError) Error() string {
	return fmt.Sprintf("%d", e)
}

const (
	OtpConfigErr OtpError = 1001 + iota //config配置错误
	OtpAuthTokenErr
	OtpServerCodeErr
)

var Debug = false

func DebugLog(v ...interface{}) {
	func() {
		_, file, lineNo, _ := runtime.Caller(2)
		LogFile := file + ":" + strconv.Itoa(lineNo)
		if Debug {
			fmt.Println(LogFile, v)
			log.Println(LogFile, v)
		}
	}()

}
