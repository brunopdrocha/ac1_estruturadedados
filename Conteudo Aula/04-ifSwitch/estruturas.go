package main

import "fmt"

func main() {
	var x = 10
	var dia = "Segunda"

	if x > 18 {
		fmt.Println("Voce eh maior de idade")
	} else if x < 0 {
		fmt.Println("Dado invalido")
	} else {
		fmt.Println("Voce eh menor de idade")
	}

	switch dia {
	case "segunda", "terca", "quarta", "quinta", "sexta":
		fmt.Println("Dia util")
	case "sabado", "domingo":
		fmt.Println("Final de Semana")
	default:
		fmt.Println("Dia invalido")

	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("------------------")

	x = 5
	for x > 0 {
		fmt.Println(x)
		x--
	}

	fmt.Println("------------------")

	for x < 10 {
		x++
		if x == 3 {
			continue
		}
		fmt.Println(x)
		if x == 5 {
			break
		}
	}
}
