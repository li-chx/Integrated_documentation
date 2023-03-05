package common

import (
	"log"
	"time"
)

//  一些必要的常量

const (
	CodeExpires = -2
	CodeError   = -1
	CodeSuccess = 0
)

var ChinaTime *time.Location

type Gender int

const (
	Male Gender = iota + 1
	FeMale
)

func init() {
	var err error
	ChinaTime, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Panicln(err)
	}
}
