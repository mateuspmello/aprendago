package main

import "fmt"

type pessoa struct {
	nome string
}

func main() {

	v := 1
	fmt.Println(&v)
	pessoa := pessoa{
		nome: "mello",
	}
	fmt.Println(pessoa.nome)
	mudeMe(&pessoa)
	fmt.Println(pessoa.nome)

	fmt.Println(pessoa.nome)
	mudeMeP(pessoa)
	fmt.Println(pessoa.nome)
}

func mudeMe(p *pessoa) {
	(*p).nome = "Mateus"
}

func mudeMeP(p pessoa) {
	p.nome = "Jj"
}
