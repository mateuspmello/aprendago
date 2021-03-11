package main

import (
	"fmt"
)

//Usando uma função anônima auto-executável
//Usando buffer
func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	c = make(chan int, 2)

	c <- 42

	fmt.Println(<-c)

}
