package main

import (
	"manger/api/restful"
	"manger/pkg/misc/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	configPath := os.Getenv("MANGER_CONFIG")
	conf, err := config.NewConfig(configPath)
	if err != nil {
		panic(err)
	}

	// restful
	router, err := restful.NewRouter(conf)

	if err != nil {
		panic(err)
	}

	go router.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			router.Close()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
