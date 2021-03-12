package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/phungvandat/source-template/endpoints"
	"github.com/phungvandat/source-template/utils/config/env"
	"github.com/phungvandat/source-template/utils/errs"
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

	var (
		httpAddr  = ":4000"
		maxErrChn = 100
		eTracer   = errs.NewErrTracer(maxErrChn)
		errChn    = make(chan error)
		svc       = initService(eTracer)
		eps       = endpoints.MakeServerEndpoints(svc)
	)
	defer eTracer.Close()

	// Http handle
	var (
		httpHandler = httpTransport.NewHTTPHandler(httpTransport.BuildRouter(eps).Build())
	)
	go func() {
		logger.Info("transport:%v addr:%v", "HTTP", httpAddr)
		errChn <- http.ListenAndServe(httpAddr, httpHandler)
	}()

	// Handle function error
	go func() {
		for {
			select {
			case err := <-eTracer.GotErr():
				logger.Error(err.Error())
			}
		}
	}()

	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		errChn <- fmt.Errorf("%s", <-ch)
	}()

	logger.Error("exit: %v", <-errChn)
}
