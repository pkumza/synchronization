package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	sema "golang.org/x/sync/semaphore"
)

var (
	book  int64
	empty sema.Weighted
	full  sema.Weighted
	wg    sync.WaitGroup
)

func consumer() {
	for i := 0; i < 20; i++ {
		// 随机等 [0,10) 毫秒
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
		empty.Acquire(context.Background(), 1)
		newBook := atomic.AddInt64(&book, -1)
		fmt.Println("Consumer -1, book :", newBook)
		full.Release(1)
	}
	wg.Done()
}

func producer() {
	for i := 0; i < 25; i++ {
		// 随机等 [0,10) 毫秒
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
		full.Acquire(context.Background(), 1)
		newBook := atomic.AddInt64(&book, 1)
		fmt.Println("Producer +1, book :", newBook)
		empty.Release(1)
	}
	wg.Done()
}

func main() {
	ctx := context.Background()
	empty = *sema.NewWeighted(5)
	full = *sema.NewWeighted(5)
	empty.Acquire(ctx, 5)
	wg.Add(2) // wait for producer & consumer
	go consumer()
	go producer()
	wg.Wait()
}
