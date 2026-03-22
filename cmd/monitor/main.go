package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/backend-study-org/site-monitor/internal/config"
	"github.com/backend-study-org/site-monitor/internal/scheduler"
)

var configPath string

func main() {
	fmt.Println("Site Monitor started. Press Ctrl+C to stop.")
	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()

	if configPath == "" {
		fmt.Println("Must set -config flag")
		return
	}

	conf, err := config.Load(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail to load config %v\n", err)
		os.Exit(1)
	}
	if conf.List == nil {
		fmt.Printf("No sites to check in config file %s\n", configPath)
		return
	}

	sc, stop := scheduler.New(conf)

	go func() {
		sc.Check(time.Now())
		sc.Start()
	}()

	<-stop
	fmt.Println("Shutting down...")
	sc.Stop()
	fmt.Println("Site Monitor stopped.")
}
