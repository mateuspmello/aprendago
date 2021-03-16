package main

import "fmt"

type ErroEspecial struct {
	err string
}

func (e ErroEspecial) Error() string {
	return e.err
}

func main() {
	v, err := bla(5)
	if err != nil {
		fmt.Println("erro: ", v, err)
	}

}

func bla(i int) (int, error) {
	if i == i {
		return 1, ErroEspecial{"aaaa"}
	}
	return i, nil
}
