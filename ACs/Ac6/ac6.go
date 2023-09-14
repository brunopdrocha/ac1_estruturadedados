package main

import "fmt"

const M = 5

func main() {
	var lista [M]int
	var n = 0

	insereOrd(&lista, &n, 4) // Insere 4
	fmt.Println(lista)

	insereOrd(&lista, &n, 12) // Insere 12
	fmt.Println(lista)

	insereOrd(&lista, &n, 2) // Insere 2 na posição correta
	fmt.Println(lista)

	insereOrd(&lista, &n, 6) // Insere 6 na posição correta
	fmt.Println(lista)

	insereOrd(&lista, &n, 17) // Insere 17
	fmt.Println(lista)

	insereOrd(&lista, &n, 1) // Overflow
	fmt.Println(lista)
}

func insereOrd(v *[M]int, n *int, novoValor int) {
	if *n == M {
		fmt.Println("Overflow")

	}

	// Verifica se o elemento já existe
	for i := 0; i < *n; i++ {
		if v[i] == novoValor {
			fmt.Println("Erro: elemento já existe")

		}
	}

	// Encontra a posição correta para inserção
	pos := 0
	for i := 0; i < *n; i++ {
		if v[i] > novoValor {
			pos = i
			break
		}
		pos++
	}

	// Move os elementos maiores para abrir espaço para o novo valor
	for i := *n; i > pos; i-- {
		v[i] = v[i-1]
	}

	// Insere o novo valor na posição correta
	v[pos] = novoValor
	*n++
}
