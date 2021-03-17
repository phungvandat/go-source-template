package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/phungvandat/source-template/config/db/pg"
	"github.com/phungvandat/source-template/config/db/redis"
	"github.com/phungvandat/source-template/config/env"
	"github.com/phungvandat/source-template/endpoints"
	"github.com/phungvandat/source-template/pkg/errpkg"
	"github.com/phungvandat/source-template/utils/helper"
	"github.com/phungvandat/source-template/utils/logger"

	httpTransport "github.com/phungvandat/source-template/transports/http"
)

func main() {
	var (
		isProd = env.IsProduction()
		err    error
	)

	if !isProd {
		err = godotenv.Load()
		if err != nil {
			panic(fmt.Sprintf("Failed to load .env file by error: %v", err))
		}
	}

	// Setup locale
	{
		local, err := time.LoadLocation("Asia/Ho_Chi_Minh")
		if err != nil {
			logger.Error("Failed set location time by error:%v", err)
			os.Exit(1)
		}
		time.Local = local
	}

	// Init pg, redis connection
	pg.InitPGConn(env.PGSource())
	defer pg.Close()
	redis.InitRedisConn(env.RedisSource())
	defer redis.Close()

	var (
		httpPort  = env.HTTPPort()
		maxErrChn = 100
		eTracer   = errpkg.NewErrTracer(maxErrChn)
		errChn    = make(chan error)
		eps       = endpoints.MakeServerEndpoints(initService(eTracer))
	)
	defer eTracer.Close()

	// Http handle
	httpHandler := httpTransport.NewHTTPHandler(httpTransport.BuildRouter(eps).Build())
	go helper.Goroutine(func() {
		logger.Info("transport:%v addr:%v", "HTTP", httpPort)
		errChn <- http.ListenAndServe(fmt.Sprintf(":%v", httpPort), httpHandler)
	})

	// Handle function error
	go helper.Goroutine(func() {
		for {
			select {
			case err := <-eTracer.GotErr():
				if err == nil {
					continue
				}
				logger.Trace(err.Error())
			}
		}
	})

	go helper.Goroutine(func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		errChn <- fmt.Errorf("%s", <-ch)
	})

	logger.Error("exit: %v", <-errChn)
}
