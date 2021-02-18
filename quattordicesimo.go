package main

import "fmt"

func main() {
	//ponteiros
	x := 10
	y := &x

	*y++

	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(&y)
	fmt.Println(*y)

	fmt.Printf("%T, %T\n", x, y)
	fmt.Printf("%v, %v\n", x, y)

	x = 0
	valor(x)
	ponteiro(&x)
	fmt.Println(x)
}

func valor(x int) {
	x++
	fmt.Println("valor", x)
}

func ponteiro(x *int) {
	*x++
}
