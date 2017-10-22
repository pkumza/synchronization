// 经典正例
package main

import "fmt"
import "time"
import "sync"

var sum int
var l sync.Mutex

func adder() {
	for i := 0; i < 10000; i++ {
		l.Lock()
		sum++
		l.Unlock()
	}
}

func main() {
	sum = 0
	for i := 0; i < 10; i++ {
		go adder()
	}
	time.Sleep(time.Second * 1)
	fmt.Println(sum)
}
