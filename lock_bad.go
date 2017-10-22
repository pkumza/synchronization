// 经典错例
package main

import "fmt"
import "time"

var sum int

func adder() {
	for i := 0; i < 10000; i++ {
		sum++
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
