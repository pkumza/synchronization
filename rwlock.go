// 经典正例
package main

import (
	"fmt"
	"sync"
	"time"
)

var sum int
var rw sync.RWMutex

func read(id int) {
	time.Sleep(time.Millisecond)
	fmt.Printf("%d Read\n", id)
}
func write() {
	time.Sleep(time.Millisecond)
	fmt.Println("Write")
}

func reader(id int) {
	for i := 0; i < 10; i++ {
		rw.RLock()
		fmt.Printf("%d Acquire\n", id)
		read(id)
		rw.RUnlock()
		fmt.Printf("%d Release\n", id)
	}
}
func writer() {
	for i := 0; i < 5; i++ {
		rw.Lock()
		fmt.Println("Writer Acquire")
		write()
		rw.Unlock()
		fmt.Println("Writer Release")
	}
}

func main() {
	sum = 0
	for i := 0; i < 3; i++ {
		go reader(i)
	}
	go writer()
	time.Sleep(time.Second) // 主进程阻塞一秒
}
