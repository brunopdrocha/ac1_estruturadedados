package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x, y, x1, x2, y1, y2, numMeteoritos int
	mensagem := ""
	contFazenda := 0

	for {
		fmt.Scanln(&x1, &y1, &x2, &y2)
		if x1 == 0 && y1 == 0 && x2 == 0 && y2 == 0 {
			break
		}

		contFazenda++
		fmt.Scanln(&numMeteoritos)
		meteoritosNaFazenda := 0

		for meteorito := 1; meteorito <= numMeteoritos; meteorito++ {
			fmt.Scanln(&x, &y)
			if x1 <= x && x <= x2 && y2 <= y && y <= y1 {
				meteoritosNaFazenda++
			}
		}
		mensagem += "Teste" + strconv.Itoa(contFazenda) + "\n" + strconv.Itoa(meteoritosNaFazenda) + "\n\n"

	}
	fmt.Println(mensagem)
}
