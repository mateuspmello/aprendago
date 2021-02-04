package main

import (
	"fmt"
)

var x [5]int
var y [6]int

func main() {

	//array
	x[0] = 1
	x[1] = 1
	fmt.Println(x[0], x[1])
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Println(len(x))

	//slice
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	slice2 := append(slice, 6)
	fmt.Println(slice2)
	fmt.Println(slice[3])
	slice[3] = 312
	fmt.Println(slice[3])
	//não funciona out of range
	// slice[20] = 1

	sliceStr := []string{"banana", "maça", "jaca"}
	for índice, valor := range sliceStr {
		fmt.Println("No índice", índice, "temos o valor: ", valor)
	}

	//fatiando fatia
	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatro queijos", "margerittha"}

	fatia := sabores[1:2]
	fmt.Println(fatia, "\n", sabores[0:1], "\n", sabores[1:2], "\n", sabores[2:3], "\n", sabores[3:4], "\n", sabores[4:5])

	//é fatiando que se deleta um item de uma slice
	sabores = append(sabores[:3], sabores[4:len(sabores)]...)
	fmt.Println((sabores))

	umaslice := []int{1, 2, 3, 4}
	outraslice := []int{1, 2, 3, 4, 5}

	umaslice = append(umaslice, 5, 6, 7, 8)
	//... serve para pegar o conteudo dentro do slice (unfuri => desenrolar)
	umaslice = append(umaslice, outraslice...)
	fmt.Println(umaslice)

	//make slice
	slicemake := make([]int, 5, 10)
	slicemake[0], slicemake[1], slicemake[2], slicemake[3], slicemake[4] = 1, 2, 3, 4, 5
	slicemake = append(slicemake, 6)
	slicemake = append(slicemake, 7)
	slicemake = append(slicemake, 8)
	slicemake = append(slicemake, 9)
	slicemake = append(slicemake, 10)
	fmt.Println(slicemake, len(slicemake), cap(slicemake))

	//dobra a capacity e cria um novo array
	slicemake = append(slicemake, 11)
	fmt.Println(slicemake, len(slicemake), cap(slicemake))

	//map
	amigos := map[string]int{
		"alfredo": 555432,
		"joana":   9996674,
	}
	fmt.Println(amigos)
	fmt.Println(amigos["joana"])
	amigos["gopher"] = 4444
	fmt.Println(amigos["gopher"])

	//ok para saber se existe o valor
	será, ok := amigos["fantamas"]
	fmt.Println(será, ok)

	//maps nao possuem ordem
	qlqrcoisa := map[int]string{
		123: "muito legal",
		98:  "menos legal um pouquinho",
		983: "é massa",
		18:  "idade da festa",
	}

	fmt.Println(qlqrcoisa)

	total := 0
	for key, _ := range qlqrcoisa {
		total += key
	}

	fmt.Println(total)

}
