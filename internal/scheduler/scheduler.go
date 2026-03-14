package scheduler

import (
	"fmt"
	"strings"
	"time"

	"github.com/backend-study-org/site-monitor/internal/checker"
	"github.com/backend-study-org/site-monitor/internal/config"
)

type Scheduler struct {
	Interval time.Duration
	Sites    []config.Site
}

const DefaultTimeDurationInterval = time.Minute

var stop = make(chan bool)

var ticker *time.Ticker

func New(config *config.Config) *Scheduler {
	var interval time.Duration
	if config.Timeout != 0 {
		interval = time.Duration(config.Timeout)
	} else {
		interval = DefaultTimeDurationInterval
	}
	ticker = time.NewTicker(interval)
	return &Scheduler{
		Interval: interval,
		Sites:    config.List,
	}
}

func (s *Scheduler) Check(t time.Time) {
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
		case <-stop:
			return
		case t := <-ticker.C:
			fmt.Println(strings.Repeat("-", 6) + "TICK" + strings.Repeat("-", 6))
			s.Check(t)
		}
	}
}

func (s *Scheduler) Stop() {
	stop <- true
	if ticker != nil {
		ticker.Stop()
	}
}
