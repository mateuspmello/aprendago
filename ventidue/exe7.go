package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	c := make(chan int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go insere(c, 10)
	}

	go remove(c)
	wg.Wait()

}

func insere(c chan int, a int) {

	for i := 0; i < a; i++ {
		c <- i
	}
	wg.Done()

}

func remove(c chan int) {

	for {
		fmt.Println(<-c)
	}

}
