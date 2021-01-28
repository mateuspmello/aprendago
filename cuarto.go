package main

import "fmt"

var x bool

//o tipo constante só é atribuida no uso
const (cx = iota *1000
		_ 
		_
		cy
		cz = iota
	)

func main() {

	fmt.Println(x)
	x = true
	fmt.Println(x)

	x = (10 < 100)
	y := (10 == 100)
	z := (10 > 100)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	a := "e"
	b := "é"
	c := "u123"

	fmt.Printf("%v, %v, %v \n", a, b, c)
	d := []byte(a)
	e := []byte(b)
	f := []byte(c)
	fmt.Printf("%v, %v, %v \n", d, e, f)

	g := +0
	fmt.Printf("%v, %T \n", g, g)

	var i uint16
	i = 65535
	fmt.Println(i)
	i += 2
	fmt.Println(i)

	//strings
	s := "Hello, playground"
	fmt.Printf("%v \n %T \n", s, s)

	s = `Hello, 
	  playground`
	fmt.Printf("%v \n %T \n", s, s)

	sb := []byte(s)
	fmt.Printf("%v \n %T \n", sb, sb)

	for _, v := range sb {
		fmt.Printf("%v - %T - %#U - %#x \n", v, v, v, v)
	}

	//sistemas numericos

	k := 100
	fmt.Printf("%d - %b - %#x \n", k, k,k )

	//iota
	fmt.Println(cx, cy, cz)


}
