package main

import "fmt"

var n int

func init() {
	fmt.Println("Função init é executada antes da Main. Cada arquivo detro do pacote pode ter sua func init.")
	n = 20
}

func main () {
	fmt.Println("E a função main é executada depois da init. Só pode ter uma função main por pacote.")
	fmt.Println(n)
}