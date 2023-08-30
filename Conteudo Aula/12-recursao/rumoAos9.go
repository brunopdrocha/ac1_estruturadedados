package main

import (
	"fmt"
	"strconv"
)

func rumoAos9(n string, grau int) string {
	if n == "9" {
		if grau == 0 {
			grau = 1
		}
		return "Eh multiplo de 9 e n eh multiplo de 9 " + strconv.Itoa(grau)
	}
	if len(n) == 1 {
		return "N eh multiplo de 9"
	}
	soma := 0
	for _, carac := range n {
		digit, _ := strconv.Atoi(string(carac))
		soma += digit
	}
	return rumoAos9(strconv.Itoa(soma), grau+1)
}

func main() {
	var numeros []string
	var numero string

	for {
		fmt.Scanln(&numero)
		if numero == "0" {
			break
		}
		numeros = append(numeros, numero)
	}

	for _, num := range numeros {
		resultado := rumoAos9(num, 0)
		fmt.Println(resultado)
	}
}
