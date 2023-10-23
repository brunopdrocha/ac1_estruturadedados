package main

import (
	"fmt"
)

// Defina a estrutura do nó
type Node struct {
	Dado    int
	Proximo *Node
}

// Defina a estrutura da lista circular
type CircularList struct {
	Inicio *Node
}

// Função para exibir a lista circular
func ExibirListaCircular(lista *CircularList) {
	if lista.Inicio == nil {
		fmt.Println("A lista circular está vazia.")
		return
	}

	noAtual := lista.Inicio
	for {
		fmt.Println(noAtual.Dado)
		noAtual = noAtual.Proximo
		if noAtual == lista.Inicio {
			break
		}
	}
	fmt.Println()
}

// Função para inserir o primeiro nó
func InserirPrimeiroNo(lista *CircularList, valor int) {
	novoNo := &Node{Dado: valor}

	if lista.Inicio == nil {
		novoNo.Proximo = novoNo
	} else {
		atual := lista.Inicio
		for atual.Proximo != lista.Inicio {
			atual = atual.Proximo
		}
		atual.Proximo = novoNo
		novoNo.Proximo = lista.Inicio
	}
	lista.Inicio = novoNo
}

// Função para excluir o primeiro nó
func ExcluirPrimeiroNo(lista *CircularList) {
	if lista.Inicio == nil {
		fmt.Println("A lista circular está vazia.")
		return
	}

	if lista.Inicio.Proximo == lista.Inicio {
		lista.Inicio = nil
	} else {
		atual := lista.Inicio
		for atual.Proximo != lista.Inicio {
			atual = atual.Proximo
		}
		atual.Proximo = lista.Inicio.Proximo
		lista.Inicio = lista.Inicio.Proximo
	}
}

func main() {
	listaCircular := CircularList{}

	// Adicione elementos à lista circular
	InserirPrimeiroNo(&listaCircular, 1)
	InserirPrimeiroNo(&listaCircular, 2)
	InserirPrimeiroNo(&listaCircular, 3)

	// Exiba a lista circular
	ExibirListaCircular(&listaCircular)

	// Exclua o primeiro nó
	ExcluirPrimeiroNo(&listaCircular)

	// Exiba a lista circular novamente
	ExibirListaCircular(&listaCircular)
}
