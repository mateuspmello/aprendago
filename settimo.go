package main

import "fmt"

func main() {

	//exercicio 1
	i := 0
	for {
		if i > 10000 {
			break
		}
		fmt.Printf("%v \t", i)
		i++
	}

	//exercicio 2
	for i = 64; i <= 90; i++ {
		fmt.Printf("%v \n", i)
		j := 1
		for j <= 3 {
			fmt.Printf("\t%#U \n", i)
			j++
		}
	}

	//exercicio 3
	i = 1989
	for i <= 2021 {
		fmt.Printf("%v, ", i)
		i++
	}

	fmt.Println()
	//exercicio 4
	i = 1989
	for {
		if i > 2021 {
			break
		}
		fmt.Printf("%v, ", i)
		i++
	}

	fmt.Println()
	//exercicio 5
	for i = 10; i <= 100; i++ {
		fmt.Println(i, "%4 = ", i%4)
	}

	//exercicio 6
	i = 10
	if i < 10 {
		fmt.Println("x é menor que 10")
	} else {
		fmt.Println("x é maior ou igual a 10")
	}

	//exercicio 7
	i = 9
	if i < 10 {
		fmt.Println("x é menor que 10")
	} else if i == 10 {
		fmt.Println("x é igual a 10")
	} else {
		fmt.Println("x é maior que 10")
	}

	//exercicio 8
	switch {
	case i == 9:
		fmt.Println("x é 9")
	case i > 10:
		fmt.Println("x é  maior que 9")
	}

	//exercicio 9
	esportefavorito := "tenis"
	switch esportefavorito {
	case "tenis":
		fmt.Println("x é ", esportefavorito)
	case "futebol":
		fmt.Println("x é ", esportefavorito)
	}

}
