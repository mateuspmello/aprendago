package main

import "fmt"

const pi float64 = 3.14

var area float64

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type quadrado struct {
	lado int
}

type circulo struct {
	raio int
}

type figura interface {
	area()
}

func main() {
	fmt.Println("Exercicio 1")
	fmt.Println(retint())
	fmt.Println(retintstr())

	fmt.Println("Exercicio 2")
	fmt.Println(fun(1, 2, 3, 4))
	fmt.Println(fun([]int{4, 3, 2, 1}...))
	fmt.Println(fun2([]int{4, 3, 2, 1}))

	fmt.Println("Exercicio 3")
	dfer()

	fmt.Println("Exercicio 4")
	p := pessoa{
		nome:      "Mateus",
		sobrenome: "Mello",
		idade:     31,
	}
	nome, idade := p.nomeidade()
	fmt.Println(nome, " - ", idade)

	fmt.Println("Exercicio 5")
	q := quadrado{lado: 2}
	c := circulo{raio: 2}
	info(q)
	fmt.Println("Area do quadrado com lado ", q.lado, ": ", area)
	info(c)
	fmt.Println("Area do circulo com raio ", c.raio, ": ", area)

	fmt.Println("Exercicio 6")
	x := func() int {
		return 2
	}()
	fmt.Println("função anonima", x)

	fmt.Println("Exercicio 7")

	fmt.Println("Exercicio 8")
	rf := ff()
	fmt.Println("Retornando função: ", rf())

	fmt.Println("Exercicio 9")
	fmt.Println("Exercicio 10")
	fmt.Println("Exercicio 11")
	fmt.Println("Video gravado")

}

func ff() func() int {
	return func() int {
		return 2
	}
}

func (q quadrado) area() {
	area = float64(q.lado * q.lado)
}

func (c circulo) area() {
	area = 2 * pi * float64(c.raio)
}

func (p pessoa) nomeidade() (string, int) {
	return p.nome, p.idade
}

func info(f figura) {
	f.area()
}

func dfer() {
	defer fmt.Println("primeiro")
	fmt.Println("segundo")
}

func fun(x ...int) int {
	r := 0
	for _, v := range x {
		r += v
	}
	return r
}

func fun2(x []int) int {
	r := 0
	for _, v := range x {
		r += v
	}
	return r
}

func retint() int {
	return 2
}

func retintstr() (int, string) {
	return 2, "dois"
}
