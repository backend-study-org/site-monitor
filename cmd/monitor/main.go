package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/backend-study-org/site-monitor/internal/checker"
	"github.com/backend-study-org/site-monitor/internal/config"
)

var configPath string

func main() {
	fmt.Println("Site Monitor started")
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
	sites := conf.List
	for _, site := range sites {
		resp := checker.CheckSite(site.URL)

		if resp.Error == nil && resp.Code == 200 {
			fmt.Printf("Site %s ok\n", resp.URL)
		} else {
			fmt.Printf("Site %s NOT ok\n", resp.URL)
		}
	}
}
