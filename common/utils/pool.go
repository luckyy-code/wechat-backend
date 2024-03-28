package utils

import (
	"github.com/panjf2000/ants"
	"github.com/zeromicro/go-zero/core/logx"
)

// Init InitPool 初始化routine pool

var Pool *ants.Pool

func init() {
	var err error
	Pool, err = ants.NewPool(12)
	if err != nil {
		logx.Error("ants pool err: ", err)
		return
	}
}
