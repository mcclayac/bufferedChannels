package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

var (
	running int64 = 0
)

func work() {
	atomic.AddInt64(&running, 1)
	fmt.Printf("[%d ", running)
	time.Sleep(time.Duration(rand.Intn(4)) * time.Second)
	atomic.AddInt64(&running, -1)
	fmt.Printf("]")
}

func worker(sema chan bool) {
	//fmt.Printf("<-sema")
	<-sema
	//fmt.Printf("work()")
	work()
	sema <- true
	//fmt.Printf("sema <- true")

}

func main() {
	fmt.Printf("Starting buffered Channel\n\n")

	sema := make(chan bool, 20)

	for i := 0; i < 1000; i++ {
		fmt.Printf("Starting worker :%d\t", i)
		go worker(sema)
		//fmt.Printf("stopping worker :%d", i)
	}
	fmt.Printf("\n\n")

	for i := 0; i < cap(sema); i++ {
		sema <- true
	}

	time.Sleep((30 * time.Second))
}
