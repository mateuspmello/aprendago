package main

import (
	"encoding/json"
	"fmt"
	"sort"

	"golang.org/x/crypto/bcrypt"
)

type Celebridade struct {
	Nome          string  `json:"Nome"`
	Sobrenome     string  `json:"Sobrenome"`
	Idade         int     `json:"Idade"`
	Profissao     string  `json:"Profissao"`
	Contabancaria float64 `json:"Contabancaria"`
}

type auto struct {
	nome     string
	potencia int
	consumo  int
}

type ordernarPorPotencia []auto

func (x ordernarPorPotencia) Len() int {
	return len(x)
}
func (x ordernarPorPotencia) Less(i, j int) bool {
	return x[i].potencia < x[j].potencia
}
func (x ordernarPorPotencia) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type ordernarPorConsumo []auto

func (x ordernarPorConsumo) Len() int {
	return len(x)
}
func (x ordernarPorConsumo) Less(i, j int) bool {
	return x[i].consumo > x[j].consumo
}
func (x ordernarPorConsumo) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type ordernarPorLucroProDonoDoPosto []auto

func (x ordernarPorLucroProDonoDoPosto) Len() int {
	return len(x)
}
func (x ordernarPorLucroProDonoDoPosto) Less(i, j int) bool {
	return x[i].consumo < x[j].consumo
}
func (x ordernarPorLucroProDonoDoPosto) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func main() {
	//função Marshall, arranjar coletr
	jamesbond := Celebridade{
		Nome:          "James",
		Sobrenome:     "Bond",
		Idade:         40,
		Profissao:     "Agente Secreto",
		Contabancaria: 100000,
	}

	darthvader := Celebridade{
		Nome:          "Darth",
		Sobrenome:     "Vader",
		Idade:         60,
		Profissao:     "Vilão",
		Contabancaria: 5000000,
	}

	j, err := json.Marshal(jamesbond)
	if err != nil {
		fmt.Println(err)
	}

	d, err := json.Marshal(darthvader)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(j))
	fmt.Println(string(d))

	//Unmarshal JSON
	jp := Celebridade{}
	json.Unmarshal(j, &jp)
	fmt.Println(jp)

	//interface Writer

	//sort
	ss := []string{"atirei", "o", "pau", "no", "gato"}
	sort.Strings(ss)
	fmt.Println(ss)

	ii := []int{2, 9, 1, 6, 5, 4}
	sort.Ints(ii)
	fmt.Println(ii)

	carros := []auto{
		auto{"Chevete", 50, 5},
		auto{"Porsche", 300, 5},
		auto{"Fusca", 20, 30},
	}

	fmt.Println(carros)
	sort.Sort(ordernarPorPotencia(carros))
	fmt.Println(carros)
	sort.Sort(ordernarPorConsumo(carros))
	fmt.Println(carros)
	sort.Sort(ordernarPorLucroProDonoDoPosto(carros))
	fmt.Println(carros)

	//bcrypt
	senha := "20julho1968"

	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if nil != err {
		fmt.Println(err)
	}
	fmt.Println(string(sb))

	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senha)))

}
