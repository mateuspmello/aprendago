package main

import "fmt"

var y = "strinhs"

type hotdog int

var b hotdog

func main() {
	b = 99
	h := 10
	fmt.Println(b == h)
	/*
		zero (antes da inicialização)
		ints: 0
		float: 0.0
		string: ""
		booleans: false
		pointers, functions, interfaces, slices, channels, maps: nil
	*/

	x := 16. //operador curto de inicialização (marmota)

	z := true

	fmt.Println(x, y, z)
	fmt.Printf("x: %v, %T \n", x, x)
	fmt.Printf("y: %v, %T \n\n", y, y)

	x, k := 20, 10
	fmt.Printf("x new value: %v, %T \n", x, x)
	fmt.Printf("k: %v, %T \n", k, k)

	qualquercoisa(10)

	fmt.Println(fmt.Sprint("aqui", x, y, z))

}

func qualquercoisa(x int) {
	fmt.Println(x)
	fmt.Println(y)
}
