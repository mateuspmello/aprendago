package main

import "fmt"

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	titulo  string
	salario int
}

func main() {
	//struct
	cliente1 := cliente{
		nome:      "Mateus",
		sobrenome: "Mello",
		fumante:   false,
	}

	cliente2 := cliente{"Joana", "Pereira", true}

	fmt.Println(cliente1)
	fmt.Println(cliente2)

	//struct embutido
	pessoa1 := pessoa{
		nome:  "Alfredo",
		idade: 30,
	}

	pessoa2 := profissional{
		pessoa:  pessoa{"Joana", 29},
		titulo:  "analista",
		salario: 5000,
	}

	pessoa3 := pessoa{"Mauricio", 40}
	pessoa4 := profissional{pessoa{"Joao", 41}, "Dev", 5000}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
	fmt.Println(pessoa3)
	fmt.Println(pessoa4)

	//struct anônimo
	sa := struct {
		nome  string
		idade int
	}{"Mateus", 50}

	fmt.Println(sa)

}
