package main

import "fmt"

//Exercício: Dada uma matriz n x n de valores inteiros, implemente uma função que
//localize um dado valor x. A função deve retornar VERDADEIRO se houver achado, e
//FALSO caso contrário.
func busca(matriz [][]int, n, x int) bool {
	var i, j int
	i = 0

	for i < n {
		j = 0
		for j < n {
			if matriz[i][j] == x {
				return true // retorna valor encontrado
			}
		}
		j++
	}
	return false // n retorna valor encontrado
} //fim func busca
func main() {
	v := []int{3, 4, 5, 6, 7}
	a := []int{4, 8, 12, 2}

	for _, k := range a {
		fmt.Println(k)
	}
	fmt.Println(busca(v, 5, 12))
}
