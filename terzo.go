package main

import "fmt"

type htht int

var x htht
var y int

func main() {

	fmt.Printf("x: %v  %T \n", x, x)

	x = 42

	fmt.Printf("x: %v  %T \n", x, x)

	y = int(x)
	fmt.Printf("x: %v  %T \n ", y, y)
}
