package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil { //recover é escrito dessa forma
		fmt.Println("Execução recuperada com sucesso")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA FOI EXATAMENTE 6") //o panic não é um erro, mas é como se fosse uma exceção,
	//onde é executado e para o programa em seguida, para o programa não pagar necessida do recover

}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
}