package main

import "fmt"

type Pessoa struct {
	Nome string
}

func main() {
	var p Pessoa
	p = Pessoa{
		Nome: "Victor",
	}

	var pessoas [3]Pessoa

	pessoas[0] = p
	for indice, pessoa := range pessoas {
		if (pessoa == Pessoa{}) {
			fmt.Println("achou espaço vazio no índice", indice)
		}
	}

	pessoas = adiciona(p, pessoas)
	fmt.Println(pessoas)
}

func adiciona(p Pessoa, a [3]Pessoa) [3]Pessoa {
	a[1] = p
	return a
}
