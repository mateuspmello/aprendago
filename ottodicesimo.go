package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	fmt.Println(runtime.NumCPU())
	//add (total de funcoes)
	wg.Add(2)

	go func1()
	go func2()
	//espera!
	wg.Wait()

	contador := 0
	totaldegoroutines := 10

	var wgg sync.WaitGroup
	wgg.Add(totaldegoroutines)
	for i := 0; i < totaldegoroutines; i++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wgg.Done()
		}()
	}

	wgg.Wait()

	var mu sync.Mutex
}

func func1() {
	for i := 0; i < 10; i++ {
		fmt.Println("func1: ", i)
		time.Sleep(20)
	}
	wg.Done()
}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println("func2: ", i)
		time.Sleep(20)
	}
	wg.Done()
}
