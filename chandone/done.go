package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in   chan int
	done func()
}

func doWork(id int, pw *worker) {
	for n := range pw.in {
		fmt.Printf("Worker %d received %c\n", id, n)
		pw.done()
	}
}

func createWorker(id int, wg *sync.WaitGroup) *worker {
	pw := &worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}

	go doWork(id, pw)
	return pw
}

func main() {
	var wg sync.WaitGroup

	//创建10个chan进行读操作
	var workers [10]*worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	//分两次对这10个chan分进行写入数据
	wg.Add(20)
	for i, pworker := range workers {
		pworker.in <- 'a' + i
	}

	for i, pworker := range workers {
		pworker.in <- 'A' + i
	}

	wg.Wait()
}
