package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/junaozun/game-gate-go/server/gateway"
	"github.com/junaozun/gogopkg/config"
	"github.com/junaozun/gogopkg/logrusx"
)

var (
	cfgPath = flag.String("config", "gate.yaml", "config file path")
)

func main() {

	go func() {
		logrusx.Log.Info("pprof start.....")
		fmt.Println(http.ListenAndServe(":10083", nil))
	}()

	cfg := &gateway.GateServerConfig{}
	if err := config.LoadConfigFromFile(*cfgPath, cfg); nil != err {
		panic(err)
	}

	// 启动gateway
	g := gateway.NewGateWay(cfg)
	g.Run()
}
