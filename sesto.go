package main

import "fmt"

func main() {

	var x interface{}

	for hora := 0; hora <= 12; hora++ {
		fmt.Printf("Hora: %v \n", hora)
		for x := 1; x < 60; x++ {
			fmt.Print(" ", x)
		}
		fmt.Println()
	}

	//for while
	i := 0
	for i < 10 {
		fmt.Printf("I: %v \n", i)
		i++
	}

	i = 0
	for {
		if i < 10 {
			i++
			fmt.Printf("%v é menor que 10 \n", i)
		} else {
			break
		}
	}

	//continue quebra a iteração atual do comando de repetição
	//break    para de iterar e sai do for

	for i = 0; i < 20; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("i não é 3, i é ", i)
	}

	//desafio surpresa
	i = 33
	for {
		if i > 122 {
			break
		}
		fmt.Printf("%d \t %#x \t %v  \n", i, i, string(i))
		i++
	}

	//if
	if y := -1; y >= 0 {
		fmt.Printf("%d \n", y)
	} else {
		fmt.Printf("%d \t %#x \t %v  \n", y, y, y)
	}

	//switch
	k := 10
	switch k {
	case 1:
		fmt.Printf("%d \n", k)
		fallthrough //cai para o proximo
	case 10:
		fmt.Printf("%d \n", 54646)
		fallthrough
	case 11:
		fmt.Printf("%d \n", 465456546)
	default:
		fmt.Printf("%d \n", 999)

	}

	//interagindo com tipos
	x = true
	switch x.(type) {
	case int:
		fmt.Println("é um int")
	case bool:
		fmt.Println("é um bool")
	case string:
		fmt.Println("é um string")
	case float64:
		fmt.Println("é um float")
	}
}
