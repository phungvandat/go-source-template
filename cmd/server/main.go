package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/phungvandat/source-template/utils/config/env"
	"github.com/phungvandat/source-template/utils/errs"
	"github.com/phungvandat/source-template/utils/logger"
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
		maxErrChn = 100
		tErr      = errs.NewErrTracer(maxErrChn)
		errChn    = make(chan error)
	)
	defer tErr.Close()

	// Handle function error
	go func() {
		for {
			select {
			case err := <-tErr.GotErr():
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
