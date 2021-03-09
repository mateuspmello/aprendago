package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var ctr int32 = 0

var wggg sync.WaitGroup

func counter() {
	fmt.Println(ctr)
	runtime.Gosched()
	atomic.AddInt32(&ctr, 1)
	wggg.Done()

}

func main() {

	wggg.Add(10)

	go counter()
	go counter()
	go counter()
	go counter()
	go counter()
	go counter()
	go counter()
	go counter()
	go counter()
	go counter()

	wggg.Wait()

}
