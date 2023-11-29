package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	benchmark(10)
	benchmark(100)
	benchmark(1000)
	benchmark(10000)
	benchmark(100000)
	benchmark(1000000)
}

func benchmark(tamanho int) {
	numerosQuick := make([]int, tamanho)
	var numerosBubble, numerosMerge []int

	for i := 0; i < tamanho; i++ {
		numerosQuick[i] = rand.Intn(tamanho * 2)
	}

	numerosBubble = append(numerosBubble, numerosQuick...)
	numerosMerge = append(numerosMerge, numerosQuick...)

	chamaQuicksort(numerosQuick)
	chamaMergesort(numerosMerge)
	chamaBubblesort(numerosBubble)
	fmt.Println("----------------------------------------")
}

func chamaBubblesort(numeros []int) {
	t1 := time.Now()
	bubblesort(&numeros)
	t2 := time.Now()
	fmt.Println("Bubblesort /", len(numeros), "elementos / tempo:", t2.Sub(t1))
}

func chamaQuicksort(numeros []int) {
	t1 := time.Now()
	quicksort(&numeros, 0, len(numeros) - 1)
	t2 := time.Now()
	fmt.Println("Quicksort /", len(numeros), "elementos / tempo:", t2.Sub(t1))
}

func chamaMergesort(numeros []int) {
	aux := make([]int, len(numeros))
	t1 := time.Now()
	mergesort(&numeros, &aux, 0, len(numeros) - 1)
	t2 := time.Now()
	fmt.Println("Mergesort /", len(numeros), "elementos / tempo:", t2.Sub(t1))
}

func mergesort(lista, aux *[]int, inicio, fim int) {
	if inicio < fim {
		meio := int((inicio + fim) / 2)
		mergesort(lista, aux, inicio, meio)
		mergesort(lista, aux, meio + 1, fim)
		intercala(lista, aux, inicio, meio, fim)
	}
}

func intercala(lista, aux *[]int, inicio, meio, fim int) {
	i, j := inicio, meio + 1
	for k := inicio; k <= fim; k++ {
		(*aux)[k] = (*lista)[k]
	}

	k := i
	for i <= meio && j <= fim {
		if (*aux)[i] <= (*aux)[j] {
			(*lista)[k] = (*aux)[i]
			i++
		} else {
			(*lista)[k] = (*aux)[j]
			j++
		}
		k++
	}

	for i <= meio {
		(*lista)[k] = (*aux)[i]
		i++
		k++
	}

	for j <= fim {
		(*lista)[k] = (*aux)[j]
		j++
		k++
	}
}

func quicksort(lista *[]int, inicio, fim int) {
	i, j := inicio, fim
	pivo := (*lista)[i]

	for i <= j {
		for (*lista)[i] < pivo {
			i++
		}
		for (*lista)[j] > pivo {
			j--
		}

		if i <= j {
			(*lista)[i], (*lista)[j] = (*lista)[j], (*lista)[i]
			i++
			j--
		}
	}

	if j > inicio { quicksort(lista, inicio, j) }
	if i < fim { quicksort(lista, i, fim) }
}

func bubblesort(lista *[]int) {
	trocou := true
	limite := len(*lista)

	for trocou {
		trocou = false
		limite--
		for i := 0; i < limite; i++ {
			if (*lista)[i] > (*lista)[i + 1] {
				(*lista)[i], (*lista)[i + 1] = (*lista)[i + 1], (*lista)[i]
				trocou = true
			}
		}
	}
}