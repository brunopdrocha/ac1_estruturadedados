package main

import (
	"fmt"
)

var contadorMovimientos int

func Hanoi(n int, origem string, destino string, trabalho string) {
	if n > 0 {
		Hanoi(n-1, origem, trabalho, destino)
		contadorMovimientos++
		fmt.Printf("Movimiento %d: %s --> %s\n", contadorMovimientos, origem, destino)
		Hanoi(n-1, trabalho, destino, origem)
	}
}

func main() {
	var n int
	var A, B, C string
	A = "A"
	B = "B"
	C = "C"
	contadorMovimientos = 0

	fmt.Println("Informe a quantidade de discos:")
	fmt.Scanln(&n)
	Hanoi(n, A, C, B)
	fmt.Printf("Quantidade de movimentos total de %d.\n", contadorMovimientos)
}
