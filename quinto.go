package main

import "fmt"

const z = 10
const k int = 10

func main(){

	x := 1000

	fmt.Printf("%d, %b, %x", x, x, x)


	fmt.Println()
	y := (10 == 100)
	fmt.Println(y)
	y = (10 <= 100)
	fmt.Println(y)
	y = (10 > 100)
	fmt.Println(y)
	y = (10 < 100)
	fmt.Println(y)

	var aa int = 20
	aa = z
	fmt.Println(aa)
	var ab float64 = 20
	ab = z
	fmt.Println(ab)

	v := 10
	fmt.Printf("%d, %b, %x\n", v, v, v)
	v = v << 1
	fmt.Printf("%d, %b, %x\n", v, v, v)
	
}