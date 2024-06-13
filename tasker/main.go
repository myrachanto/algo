package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	period = time.Second
)

// service struct
type tickerService struct {
}

// NewTickerService returns new service
func NewTickerService() *tickerService {
	return &tickerService{}
}

// Run starts service
func (ds *tickerService) Run(ctx context.Context) {
	ticker := time.NewTicker(period)
	for {
		select {
		case <-ticker.C:
			go func() {
				ds.task()
			}()
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
} 

// periodic task
func (ds *tickerService) task() {
	fmt.Printf("Ticker task, %v\n", time.Now())
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ds := NewTickerService()
	go ds.Run(ctx)
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	<-shutdown
	cancel()
}
