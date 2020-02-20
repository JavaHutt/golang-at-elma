package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx := context.Background()
	canceledContext, cancel := context.WithCancel(ctx)
	go func() {
		time.Sleep(7 * time.Second)
		cancel()
	}()
	interrupts := make(chan os.Signal)
	signal.Notify(interrupts, os.Interrupt)

	select {
	case <-canceledContext.Done():
		fmt.Println("context canceled", ctx.Err())
	case <-interrupts:
		fmt.Println("got signal")
	case <-time.After(10 * time.Second):
		fmt.Println("10 second expired")
	}
}
