package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(channel chan<- int) {
	for i := 1; i <= 25; i++ {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
		channel <- i
		fmt.Printf("Produced %d\n", i)
	}
}
func consumer(channel <-chan int) {
	for j := 1; j <= 20; j++ {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
		i := <-channel // 此处会阻塞, 如果信道中没有数据的话
		fmt.Printf("Received %d\n", i)
	}
}
func main() {
	channel := make(chan int, 5) // 定义带有5个缓冲区的信道
	go producer(channel)         // producer产出的结果放入管道
	consumer(channel)            // 主线程从信道中取数据
	close(channel)
}
