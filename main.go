package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	local_config "github.com/junaozun/game-gate-go/server/config"
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

	cfg := local_config.GatewayConfig{}
	if err := config.LoadConfigFromFile(*cfgPath, &cfg); nil != err {
		panic(err)
	}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, os.Interrupt)
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	// 主循环
QUIT:
	for {
		select {
		case sig := <-sigs:
			log.Printf("Signal: %s", sig.String())
			break QUIT
		case <-ticker.C:
			logrusx.Log.WithFields(logrusx.Fields{
				"goroutine count": runtime.NumGoroutine(),
			}).Info("协程数量")
		}
	}
	logrusx.Log.Info("[main] quiting......")
}
