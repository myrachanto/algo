package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"runtime"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

var (
	sema = semaphore.NewWeighted(12)
)

func dosomething(i int) error {
	defer sema.Release(1)
	if i%5 == 0 {
		fmt.Printf("%d is divisble by five\n", i)
	}
	if i%11 == 0 {
		fmt.Printf("%d is divisble by eleven\n", i)
	}
	if i%21 == 0 {
		return errors.New("its divisble by 21")
	}
	return nil
}
func main() {
	errGroup, ctx := errgroup.WithContext(context.Background())
	for i := 0; i < 1000; i++ {
		fmt.Printf("%d ======> is the number of Goroutines running \n", runtime.NumGoroutine())
		if err := sema.Acquire(ctx, 1); err != nil {
			log.Fatal(err)
		}
		i := i
		errGroup.Go(func() error {
			err := dosomething(i)
			return err
		})
	}
	if err := errGroup.Wait(); err != nil {
		log.Fatal(err)
	}

}
