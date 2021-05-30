package main

import (
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro) //<nil> nulo pois não há erro no email.

}


/* Comandos no terminal:

go mod init modulo
go build
ls
./modulo - para executar o arquivo pelo binario executavel

*/