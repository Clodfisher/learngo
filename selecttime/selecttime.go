package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generatorSend() chan int {
	out := make(chan int)

	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker received %d\n", n)
	}
}

func createWorker() chan<- int {
	c := make(chan int)
	go worker(c)
	return c
}

func main() {
	//创建两个发送chan
	var c1, c2 = generatorSend(), generatorSend()

	//创建一个接受chan
	var worker = createWorker()

	//创建接受数据缓存器
	var values []int
	tm := time.After(10 * time.Second) //定时器，10秒钟后结束程序
	tick := time.Tick(time.Second)     //定时器，每分钟都打印下缓存切片的长度

	for {
		//	fmt.Println("main func run for ...")
		//当接受到数据时才进行发送chan的创建，否则整个发送数据的case进行阻塞不会进入
		var activeWorker chan<- int
		//取得缓存切片中需要发送到activworker中的数据
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-tick:
			fmt.Println(
				"queue len =", len(values))
		case <-tm:
			fmt.Println("bye...")
			return
			//800毫秒内，select进入不了任意一个case将打印超时
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")

		}
	}
}
