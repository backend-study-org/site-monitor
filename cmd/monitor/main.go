package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/backend-study-org/site-monitor/internal/config"
	"github.com/backend-study-org/site-monitor/internal/scheduler"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "path to config file")
}

func main() {
	fmt.Println("Site Monitor started")
	flag.Parse()

	if configPath == "" {
		fmt.Println("Must set -config flag")
		return
	}

	conf, err := config.Load(configPath)
	if err != nil {
		panic(err)
	}
	if conf == nil {
		fmt.Println("No config")
		return
	}
	if conf.List == nil {
		fmt.Printf("No sites to check in config file %s\n", configPath)
		return
	}

	sc := scheduler.New(conf)
	sc.Check(time.Now())
	sc.Start()
}
