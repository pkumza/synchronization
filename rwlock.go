// 经典正例
package main

import "time"
import "sync"

var sum int
var rw sync.RWMutex

func read()  {}
func write() {}

func reader() {
	for i := 0; i < 100; i++ {
		rw.RLock()
		read()
		rw.RUnlock()
	}
}
func writer() {
	for i := 0; i < 10; i++ {
		rw.Lock()
		write()
		rw.RLock()
	}
}

func main() {
	sum = 0
	for i := 0; i < 10; i++ {
		go reader()
	}
	go writer()
	time.Sleep(time.Second) // 主进程阻塞一秒
}
