package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

	func diaDaSemana2(numero int)string {
		switch {
		case numero == 1:
			return "Domingo"
			// fallthrough // case numero == 1, ele vai pular o case == 1 e vai para o return do case logo abaixo, neste caso irá retornar "Segunda-feira"
		case numero == 2:
			return "Segunda-feira"
		case numero == 3:
			return "Terça-feira"
		case numero == 4:
			return "Quarta-feira"
		case numero == 5:
			return "Quinta-feira"
		case numero == 6:
			return "Sexta-feira"
		case numero == 7:
			return "Sábado"
		default:
			return "Número inválido"
			// no Go não existe break, ele é automático.
		}
	}

func main() {
	dia := diaDaSemana(1)
	fmt.Println(dia)

	dia2 := diaDaSemana2(4)
	fmt.Println(dia2)
}