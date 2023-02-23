// Code generated by hertz generator.

package main

import (
	"github.com/TremblingV5/DouTok/applications/api/initialize"
	"github.com/TremblingV5/DouTok/pkg/dlog"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/pprof"
)

// 初始化 Hertz API 及 Router
func main() {
	logger := dlog.InitHertzLog(3)
	defer logger.Sync()

	hlog.SetLogger(logger)

	initialize.Init()

	h, shutdown := initialize.InitHertz()
	defer shutdown()

	pprof.Register(h)

	register(h)

	h.Spin()
}
