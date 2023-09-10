package main

import "fmt"

func main() {

	A := []int{1, 4, 5, 7, 9, 11}
	alvo := 11
	alvo2 := 30

	fmt.Println(ParNumero(A, alvo))

	fmt.Println(ParNumero(A, alvo2))
}

// Dado um array de números inteiros positivos,
// considerado ordenado crescentemente, e um valor alvo,
// encontre um par de números no array cuja soma seja igual
// ao valor alvo. Se nenhum par for encontrado, retorne um
// valor (-1, -1) indicando que nenhum par foi encontrado.
// Resolva esse problema com um algoritmo cuja complexidade é O(n).
func ParNumero(v []int, alvo int) (int, int) {

	i := 0          //INDICE(SETA 1)
	j := len(v) - 1 //INDICE(SETA Println(a)

	for i < j {
		soma := v[i] + v[j] //SOMA DOS INDICES DO VETOR

		if soma == alvo {
			return v[i], v[j] //Retorno dos Pares
		} else if soma < alvo {
			i++
		} else {
			j--
		}

	} //fim do for
	return -1, -1 //Caso nao retorne nenhum valor
}
