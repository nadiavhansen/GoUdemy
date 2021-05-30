package main

import (
	"fmt"
	"time"
	// "time"
)


func main() {

	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando o i")
		time.Sleep(time.Second)
	}
	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando o j", j)
		time.Sleep(time.Second)
	}
	fmt.Println(j) não é possível printar o j se o mesmo foi criado dentro do for

	nomes := [3]string{"João", "Pedro", "Lucas"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	nomess := [3]string{"João", "Pedro", "Lucas"}
	for _, nome := range nomess {
		fmt.Println(nome)
	}
	
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra)) //deve ser especificado string pois sem ele retornam os valores da tabela ascii de cada letra
	}
	
	usuario := map[string]string {
		"nome": "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario{
		fmt.Println(chave, valor)
	}

	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}

}