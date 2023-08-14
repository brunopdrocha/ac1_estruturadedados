package main

import "fmt"

func main() {
	// Array é uma coleção de dados do mesmo tipo
	// Arrays possuem tamanho fixo, definido em tempo de compilação

	var filme [5]string

	filme[0] = "Senhor dos Anéis"
	filme[1] = "Barney"
	filme[2] = "Tomás e seus Amigos"
	filme[3] = "Backyardigans"
	filme[4] = "Carros"

	// Declaração curta
	numeros := [4]int{0, 2, 4, 6}
	fmt.Println(numeros)

	// slices --> Estruturas flexíveis, de tamanho dinâmico
	var novosNumeros []int

	novosNumeros = numeros[1:]  // vai até o final do array
	novosNumeros = numeros[1:3] // vai até o elemento de índice 2
	fmt.Println(novosNumeros)

	numeros[2] = 7
	fmt.Println(novosNumeros)

	nomes := []string{"Bruno", "Matheus"}
	fmt.Println(nomes)

	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=")

	for i := 0; i < len(numeros); i++ {
		fmt.Println(numeros[i])
	}

	for indice, num := range numeros {
		fmt.Println("Índice:", indice, "Valor:", num)
	}
	fmt.Println("-=-=-=-=-=-=-=-=-=-=")
	fmt.Println(numeros)
	modificarArray(numeros) // não altera array original
	fmt.Println(numeros)

	fmt.Println(novosNumeros)
	modificarSlice(novosNumeros) // altera o slice original
	fmt.Println(novosNumeros)
	fmt.Println(numeros)

	novoSlice := criarSlice()
	fmt.Println(novoSlice)
}

func modificarArray(a [4]int) {
	a[0] = 999
}

func modificarSlice(s []int) {
	s[0] = 999
}

func criarSlice() []int {
	novoSlice := []int{10, 20, 30}
	return novoSlice
}
