package main

import "fmt"

func main() {

	var variavel1 string = "Variavel 1" //declaração explicida
	variavel2 := "Variavel 2" //declaração implicita, inferencia de tipo
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var(
		variavel3 string = "var 3"
		variavel4 string = "4"
	)
	fmt.Println(variavel3)
	fmt.Println(variavel4)

	const constante1 string = "constante 1"
	fmt.Println(constante1)

	variavel3, variavel4 = variavel4, variavel3 //invertendo o valor das variaveis
	fmt.Println(variavel3)
	fmt.Println(variavel4)


}