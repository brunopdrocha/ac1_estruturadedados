package main

import "fmt"

const M = 9

func main() {
	lista, tam := [M]int{1, 3, 7, 10}, 4
	lista2, tam2 := [M]int{1, 3, 4, 6, 7, 8, 12}, 7

	fmt.Println(busca1(lista, 4, 7)) //2
	fmt.Println(busca1(lista, 4, 9)) //-1
	fmt.Println()

	fmt.Println(busca2(lista, 4, 7)) //2
	fmt.Println(busca2(lista, 4, 9)) //-1
	fmt.Println()

	fmt.Println(buscaOrd(lista2, 7, 5)) //-1
	fmt.Println(buscaOrd(lista2, 7, 7)) //4
	fmt.Println()

	fmt.Println(buscaBin(lista2, 7, 5)) //-1
	fmt.Println(buscaBin(lista2, 7, 7)) //4
	fmt.Println()

	insere(&lista2, &tam2, 6) //Erro elemento ja existe
	fmt.Println(lista2, tam2)
	insere(&lista2, &tam2, 9) //Insere elemento 9
	fmt.Println(lista2, tam2)
	insere(&lista2, &tam2, 10) //Insere elemnento 10
	fmt.Println(lista2, tam2)
	insere(&lista2, &tam2, 10) //Overflow
	fmt.Println(lista2, tam2)
	fmt.Println()

	remove(&lista, &tam, 6) //Erro elemento n existe
	fmt.Println(lista, tam)
	remove(&lista, &tam, 1) //Remove o elemento 1
	fmt.Println(lista, tam)
	remove(&lista, &tam, 3) //Remove o elemento 3
	fmt.Println(lista, tam)
	remove(&lista, &tam, 7) //Remove o elemento 7
	fmt.Println(lista, tam)
	remove(&lista, &tam, 10) //Remove o elemento 10
	fmt.Println(lista, tam)
	remove(&lista, &tam, 1) //Underflow
	fmt.Println(lista, tam)
}

func insere(v *[M]int, n *int, valor int) {
	if *n == M {
		fmt.Println("Overflow")
	} else {
		if busca1(*v, *n, valor) != -1 {
			fmt.Println("Erro! Elemento ja existe na lista")
		} else {
			v[*n] = valor
			*n++
		}
	}
}

func remove(v *[M]int, n *int, valor int) {
	if *n == 0 {
		fmt.Println("Underflow")
	} else {
		ind := busca1(*v, *n, valor)
		if ind == -1 {
			fmt.Println("Erro! Elemento nao existe na lista")
		} else {
			for i := ind; i <= *n-2; i++ {
				v[i] = v[i+1]
			}
			*n--
			v[*n] = 0
		}
	}
}
func busca1(v [M]int, n, x int) int {
	i := 0

	for i < n {
		if v[i] == x {
			return i
		}
		i++
	}
	return -1
}

//Assumindo que existe pelo menos um espaco vazio no array
func busca2(v [M]int, n, x int) int {
	i := 0
	v[n] = x

	for v[i] != x {
		i++
	}
	if i != n {
		return i
	}
	return -1
}

//Assumindo que existe pelo menos  espaco vazio no array
func buscaOrd(v [M]int, n, x int) int {
	i := 0
	v[n] = x

	for v[i] != x {
		i++
	}
	if i == n || v[i] > x {
		return -1
	}
	return i
}
func buscaBin(v [M]int, n int, x int) int {
	inf := 0
	sup := n - 1

	for inf < sup {
		meio := int(inf+sup) / 2
		if v[meio] == x {
			return meio
		}
		if x < v[meio] {
			sup = meio
		} else {
			inf = meio + 1
		}
	}
	return -1
}
