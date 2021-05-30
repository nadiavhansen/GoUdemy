package auxiliar

import "fmt"

//Escrever() registra uma mensagem na tela e também mostra a mensagem da funçao escrever2() pois está no mesmo pacote
func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2()
}