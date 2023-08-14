package main

import "fmt"

func main() {

}

// func(param1 tipo,param2 tipo,...)

func sum(a int, b int) int {
	return a + b

}

//parametros com mesmo tipo

func sub(a, b float64) float64 {
	return a - b
}

// multiplos retornos
func trocaValores(a, b float64) (float64, string) {
	return b, "a"

}

//anonima (closure)

func usoAnonima() {
	dobra := func(x int) int {
		return 2 * x
	}

	resultado := dobra(5)
	fmt.Println(resultado)

	calcular := func(f func(int) int, x int) int {
		return f(x)
	}

	resultado = calcular(dobra, 3)
	fmt.Println(resultado)
}
