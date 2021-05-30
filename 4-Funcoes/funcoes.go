package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8){
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da funcao f")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(20, 10)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	// para não precisar retornar as duas, basta colocar um _ no lugar de uma delas (senão o go não permite)
	// resultadoSoma, _ := calculosMatematicos(20, 10)
	// fmt.Println(resultadoSoma)
}