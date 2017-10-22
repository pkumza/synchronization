package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg     sync.WaitGroup // 等子任务
	locker *sync.Mutex
	cond   *sync.Cond
	iPhone int // 苹果专卖店iPhone库存
)

func iFans(ID int) { // 果粉消费者
	cond.L.Lock()
	for iPhone <= 0 {
		cond.Wait()
	}
	iPhone--
	fmt.Printf("iFans %d get iPhone X\n", ID)
	cond.L.Unlock()
	wg.Done()
}

func appleStore() { // 苹果专卖店
	time.Sleep(time.Second)
	fmt.Println("Receive an iPhone from factory")
	cond.L.Lock()
	iPhone++
	cond.L.Unlock()
	cond.Signal()

	time.Sleep(time.Second)
	fmt.Println("Receive an iPhone from factory")
	cond.L.Lock()
	iPhone++
	cond.L.Unlock()
	cond.Signal()

	time.Sleep(time.Second * 4)
	fmt.Println("Receive 100 iPhones from factory")
	cond.L.Lock()
	iPhone += 100
	cond.L.Unlock()
	cond.Broadcast()

	wg.Done()
}

func main() {
	wg = sync.WaitGroup{}
	locker = new(sync.Mutex)
	cond = sync.NewCond(locker)
	iPhone = 4 // 库存4台iPhone
	wg.Add(1)
	go appleStore() // 店铺开门
	for i := 1; i <= 8; i++ {
		wg.Add(1)
		go iFans(i) // 果粉排队
	}
	wg.Wait()
}
