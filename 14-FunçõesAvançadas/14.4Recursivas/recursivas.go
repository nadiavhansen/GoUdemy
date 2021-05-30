package main

import "fmt"

func fibonnacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonnacci(posicao-2) + fibonnacci(posicao-1)
}

func main() {
	posicao := uint(12)

	for i := uint(1); i <= posicao; i++ {
		fmt.Println(fibonnacci(i))
	} 
}
