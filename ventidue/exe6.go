package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	go insere(c, 100)

	for v := range c {
		fmt.Println(v)
	}

}

func insere(c chan int, a int) {

	for i := 0; i < a; i++ {
		c <- i
	}
	close(c)

}
