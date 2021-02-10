package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type dentista struct {
	pessoa
	dentesarrancados int
	salario          float64
}

type arquiteto struct {
	pessoa
	tipodeconstrucao string
	tamanhodaloucura string
}

type gente interface {
	oibomdia()
}

func main() {

	//go é pass by value
	fmt.Println("com defer, veio primeiro")
	si := []int{1, 2, 30, 2}

	soma := soma1(si...)
	fmt.Println(soma)

	fmt.Println("sem defer, veio depois")

	mauricio := pessoa{"Mauricio", "Matta", 30}
	mrdentista := dentista{pessoa{"Mister Dente", "da Silva", 50}, 876, 9886}
	mrpredio := arquiteto{pessoa{"Mister Predio", "da Silva", 51}, "hospicios", "demais"}

	serhumano(mauricio)
	serhumano(mrpredio)
	serhumano(mrdentista)

	//funcao anonima
	x := 10
	func(x int) {
		fmt.Println(x, "vezes 873648 é:")
		fmt.Println(x * 873648)
	}(x)

	//funcao como expressao
	y := func(x int) {
		fmt.Println(x, "vezes 873648 é:")
		fmt.Println(x * 873648)
	}
	y(x)

	//retornar uma funcao
	k := retornafuncao()
	p := k(3)
	fmt.Println(p)
	fmt.Println(retornafuncao()(3))

	//callback
	t := somenteImpares(soma1, []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}...)
	fmt.Println(t)
}

func soma1(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

func (p pessoa) oibomdia() {
	fmt.Println(p.nome, " diz ", p.idade)
}

func (x dentista) oibomdia() {
	fmt.Println(x.nome, " diz ", x.idade, x.dentesarrancados)

}

func (x arquiteto) oibomdia() {
	fmt.Println(x.nome, " diz ", x.idade, x.tamanhodaloucura)

}

func serhumano(g gente) {
	g.oibomdia()
}

func retornafuncao() func(int) int {

	return func(i int) int {
		return i * 10
	}
}

func somenteImpares(f func(x ...int) int, y ...int) int {
	var slice []int
	for _, v := range y {
		if v%2 == 0 {
			slice = append(slice, v)
		}
	}
	total := f(slice...)
	return total
}
