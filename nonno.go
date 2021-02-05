package main

import "fmt"

var arr [5]int

func main() {

	//exercicio 1
	for i := range arr {
		arr[i] = i * 2
	}

	fmt.Println(arr)
	fmt.Printf("tipo do array: %T.\n", arr)

	//exercicio 2
	slice := []int{}
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4)
	slice = append(slice, 5)
	slice = append(slice, 6)
	slice = append(slice, 7)
	slice = append(slice, 8)
	slice = append(slice, 9)
	slice = append(slice, 10)

	for id, val := range slice {
		fmt.Printf("valor do slice na posicao %v:  %v\n", id, val)
	}

	//exercicio 3
	fmt.Println("Do primeiro ao terceiro", slice[:3])
	fmt.Println("Do quinto ao último item", slice[4:])
	fmt.Println("Do segundo ao sétimo item", slice[1:7])
	fmt.Println("Do terceuro ao penúltimo item", slice[2:len(slice)-1])

	//exercicio 4
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)

	//exercicio 5
	x = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y = x[:3]

	y = append(y, x[6:]...)
	fmt.Println("exercicio 5")
	fmt.Println(y)

	//exercicio 6
	state := make([]string, 2, 26)
	state[0], state[1] = "Acre", "Alagoas"

	fmt.Println("exercicio 6")
	fmt.Println(len(state), cap(state))
	fmt.Println(state)

	//exercicio 7
	str := [][]string{
		[]string{"Mateus", "Mello", "Programar"},
		[]string{"Sofia", "Mello", "Rir"},
		[]string{"Aurora", "Mello", "Brincar"},
	}
	fmt.Println("exercicio 7")
	fmt.Println("Nome\tSobrenome\tHobby")
	for _, v := range str {
		fmt.Printf("%v\t%v\t%v\n", v[0], v[1], v[2])
	}

	//exercicio 8
	mapa := map[string][]string{
		"mello_mateus": []string{"tenis", "pão", "programar"},
		"mello_sofia":  []string{"trekking", "praia", "comer"},
		"mello_aurora": []string{"brincar", "brincar", "brincar"},
	}

	fmt.Println("exercicio 8")
	for key, value := range mapa {
		fmt.Println(key, " - ", value)
	}

	//exercicio 9
	mapa["mello_francisco"] = []string{"comer", "chorar", "cagar"}
	fmt.Println("exercicio 9")
	for key, value := range mapa {
		fmt.Println(key, " - ", value)
	}

	//exercico 10
	fmt.Println("exercicio 10")
	delete(mapa, "mello_francisco")
	for key, value := range mapa {
		fmt.Println(key, " - ", value)
	}
}
