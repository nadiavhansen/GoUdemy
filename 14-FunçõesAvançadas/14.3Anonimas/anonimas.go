package main

import "fmt"

func main() {
	
	 func() {
		fmt.Println("Olá Mundo")
	}()

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando parâmetro")

	fmt.Println(retorno)

} //funções anônimas são as que não recebem nome, mas deve ser passado um parenteses logo após o fechamento das chaves