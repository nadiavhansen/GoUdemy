package main

import "fmt"

func alunoEstaAprovado(n1, n2 float64) bool {
	defer fmt.Println("Média calculada. Resultado será retornado") //defer serve para adiar a execução de um comando. Neste caso, este print será executado logo antes do return da função
	fmt.Println("Entrando na função para calculo da média")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	fmt.Println(alunoEstaAprovado(7, 2))
}