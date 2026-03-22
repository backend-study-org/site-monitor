package scheduler

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/backend-study-org/site-monitor/internal/checker"
	"github.com/backend-study-org/site-monitor/internal/config"
)

type Scheduler struct {
	Interval time.Duration
	Sites    []config.Site
	wg       sync.WaitGroup
}

const DefaultTimeDurationInterval = time.Minute

var stopSignal chan os.Signal
var wg sync.WaitGroup

var ticker *time.Ticker

func New(config *config.Config) (*Scheduler, chan os.Signal) {
	var interval time.Duration
	if config.Timeout != 0 {
		interval = time.Duration(config.Timeout)
	} else {
		interval = DefaultTimeDurationInterval
	}
	ticker = time.NewTicker(interval)

	stopSignal = make(chan os.Signal)
	signal.Notify(stopSignal, syscall.SIGTERM, syscall.SIGINT)

	return &Scheduler{
		Interval: interval,
		Sites:    config.List,
	}, stopSignal
}

func (s *Scheduler) Check(t time.Time) {
	wg.Add(1)
	defer wg.Done()
	
	for _, site := range s.Sites {
		tf := fmt.Sprintf("[%d-%02d-%02d %02d:%02d:%02d]", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
		resp := checker.CheckSite(site.URL)

		if resp.Error == nil && resp.Code == 200 {
			fmt.Printf("%s Site %s ok\n", tf, resp.URL)
		} else {
			fmt.Printf("%s Site %s NOT ok\n", tf, resp.URL)
		}
	}
	fmt.Println("All sites checked")
}

func (s *Scheduler) Start() {
	for {
		select {
		case <-stopSignal:
			return
		case t := <-ticker.C:
			fmt.Println(strings.Repeat("-", 6) + "TICK" + strings.Repeat("-", 6))
			s.Check(t)
		}
	}
}

func (s *Scheduler) Stop() {
	if ticker != nil {
		ticker.Stop()
	}
	wg.Wait()
}
