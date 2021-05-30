package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) { //os retornos da função são nomeados, e quando são especificados abaixo, precisa somente de =, ao invés de :=
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Println(soma, subtracao)
}