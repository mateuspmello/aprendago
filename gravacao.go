package main

import "fmt"

func main() {
	fmt.Println(funcvariadica(1, 2, 3, 4))
}

func funcvariadica(x ...int) int {
	r := 0
	for _, v := range x {
		r += v
	}
	return r
}
