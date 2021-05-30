package main

import (
	"errors"
	"fmt"
)

func main() {
	//somente int ele vai reconhecer se o pc é 32 ou 64
	var numero1 int64 = 10000000000000
	fmt.Println(numero1)

	//alias
	//int32 = rune
	//uint8 = byte

	//tipo error

	var erro error = errors.New("Erro interno aquiii")
	fmt.Println(erro)
}
