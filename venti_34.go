package main

import (
	"fmt"
	"runtime"
	"sync"
)

var ctr int = 0

var wgg sync.WaitGroup
var mu sync.Mutex

func counter() {
	mu.Lock()
	fmt.Println(ctr)
	runtime.Gosched()
	ctr++
	mu.Unlock()
	wgg.Done()

}

func main() {

	wgg.Add(10)

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

	wgg.Wait()

}
