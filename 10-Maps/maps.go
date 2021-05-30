package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo": "Carlos",
		},
		"curso": {
			"nome": "Engenharia",
			"campus": "Campus ",
		},
		}

		fmt.Println(usuario2["nome"]["primeiro"])
		delete(usuario2, "nome")
		usuario["signo"] // para adicionar uma chave
	}
