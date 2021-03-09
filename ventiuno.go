package main

import "fmt"

func main() {

	//channel, comunicacao entre go routines

	canal := make(chan int, 1)

	go func() {
		canal <- 42
	}()
	fmt.Println(<-canal)

	canal2 := make(chan int)

	go send(canal2)

	receive(canal2)

	//range close
	c := make(chan int)

	go meuloop(10, c)

	for v := range c {
		fmt.Println("Recebido do canal: ", v)
	}

	//select
	a := make(chan int)
	b := make(chan int)
	x := 500
	go func(x int) {
		for i := 0; i < x; i++ {
			a <- 1
		}
	}(x / 2)

	go func(x int) {
		for i := 0; i < x; i++ {
			b <- 1
		}
	}(x / 2)

	for i := 0; i < x; i++ {
		select {
		case v := <-a:
			fmt.Println("Canal A: ", v)
		case v := <-b:
			fmt.Println("Canal B: ", v)
		}
	}

	//comma ok
	cc := make(chan int)

	go func() {
		cc <- 0
		close(canal)
	}()

	v, ok := <-cc
	fmt.Println(v, ok)

	v, ok = <-cc
	fmt.Println(v, ok)

}

func send(s chan<- int) {
	s <- 42
}

func receive(r <-chan int) {
	i := <-r
	fmt.Println(i)
}

func meuloop(t int, s chan<- int) {
	for i := 0; i < t; i++ {
		s <- i
	}
	close(s)
}
