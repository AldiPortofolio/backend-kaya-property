package main

import (
	"fmt"
	"kaya-backend/routers"
	"kaya-backend/scheduler"
	"kaya-backend/utils"
	"kaya-backend/utils/redis"
	"os"
	"os/signal"
	"time"

	"syscall"

	"github.com/astaxie/beego/logs"
	"github.com/go-co-op/gocron"
)

// main ..
func main() {

	var errChan = make(chan error, 1)
	go func() {
		listenAddress := utils.GetEnv("KAYA_API_SERVICE_PORT", "0.0.0.0:4000")
		Scheduler()
		redisss, err := redis.RedisStore().Ping().Result()
		if err != nil {
			panic(err)
		}

		fmt.Println("Starting listen address: ", listenAddress)
		fmt.Println("Redis status: ", redisss)
		errChan <- routers.Server(listenAddress)
	}()

	var signalChan = make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	select {
	case <-signalChan:
		fmt.Println("Got an interrupt, exiting...")
	case err := <-errChan:
		if err != nil {
			fmt.Println("Error while running api, exiting...: ", err)
		}
	}
}


//CheckMutation ..
func Scheduler() {
	// defines a new scheduler that schedules and runs jobs
	s1 := gocron.NewScheduler(time.UTC)
	_, wErr := s1.Every(30).Minute().Do(CheckMutation)

	fmt.Println("gocron -------")
	logs.Info("Start cron ----- ")
	if wErr != nil {
		logs.Error("error creating job: %v", wErr)
		// return wErr
	}
	s1.StartAsync()
}


// CheckMutation ..
func CheckMutation() {
	err := scheduler.CheckMutation
	if err != nil {
		logs.Error("error Checkmutation : %v", err)
	}
}

