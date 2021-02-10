package main

import "fmt"

type pes struct {
	nome      string
	sobrenome string
	sabor     []sabor
}

type sabor struct {
	nome string
}

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	tracaoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {

	fmt.Println("Exercicio 1")
	pessoa1 := pes{
		nome:      "Mateus",
		sobrenome: "Mello",
		sabor: []sabor{
			sabor{"Creme"},
			sabor{"Flocos"},
		},
	}

	fmt.Println(pessoa1.nome)
	fmt.Println(pessoa1.sobrenome)
	for _, v := range pessoa1.sabor {
		fmt.Println(v.nome)
	}

	pessoa2 := pes{
		nome:      "Aurora",
		sobrenome: "Cabral",
		sabor: []sabor{
			sabor{"Morango"},
		},
	}

	fmt.Println(pessoa2.nome)
	fmt.Println(pessoa2.sobrenome)
	for _, v := range pessoa2.sabor {
		fmt.Println(v.nome)
	}

	fmt.Println("Exercicio 2")
	mapa := map[string]pes{}
	mapa[pessoa1.sobrenome] = pessoa1
	mapa[pessoa2.sobrenome] = pessoa2
	for _, p := range mapa {
		fmt.Println(p.nome)
		fmt.Println(p.sobrenome)
		for _, sabor := range pessoa2.sabor {
			fmt.Println(sabor.nome)
		}
		fmt.Println()

	}

	fmt.Println("Exercicio 3")
	cami := caminhonete{veiculo{4, "preto"}, false}
	sed := sedan{veiculo{4, "azul"}, true}
	fmt.Println(cami)
	fmt.Println(cami.cor)
	fmt.Println(sed)
	fmt.Println(sed.modeloLuxo)

	fmt.Println("Exercicio 4")
	aaa := struct {
		mapa map[int]string
		slc  []string
	}{
		mapa: map[int]string{1: "aaaa"},
		slc:  []string{"aaa", "bbb"},
	}

	fmt.Println(aaa)
}
