package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//channel, context
	//troca de mensagem entre goroutines

	ch := converge(trabalho("primeiro"), trabalho("segundo"))
	for x := 0; x < 16; x++ {
		fmt.Println(<-ch)
	}

}

func trabalho(s string) (r chan string) {
	canal := make(chan string)

	go func(s string, c chan string) {
		for i := 1; ; i++ {
			c <- fmt.Sprintf("Função %v diz: %v", s, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e2)))
		}

	}(s, canal)

	return canal
}

func converge(i, p chan string) (r chan string) {
	cnew := make(chan string)

	go func() {
		for {
			cnew <- <-i
		}
	}()

	go func() {
		for {
			cnew <- <-p
		}
	}()

	return cnew
}
