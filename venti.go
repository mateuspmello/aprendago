package main

import (
	"fmt"
	"sync"
)

type pessoa struct {
	nome  string
	idade int
}

type mulher struct {
	nome string
}

type humano interface {
	falar()
}

func (p *pessoa) falar() {
	fmt.Printf("%v : Falei\n", p.nome)
}

func (m mulher) falar() {
	fmt.Println("Mulher falando")
}

func dizerAlgumaCoisa(h humano) {
	h.falar()
}

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	go func() {
		p := pessoa{
			nome:  "Mateus",
			idade: 32,
		}
		fmt.Println(p.nome)
		dizerAlgumaCoisa(&p)
		fmt.Println(p.nome)
		wg.Done()
	}()

	go func() {
		fmt.Println("Segunda gorountine")
		wg.Done()
	}()

	wg.Wait()

}
