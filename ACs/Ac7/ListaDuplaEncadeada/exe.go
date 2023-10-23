package main

import "fmt"

// Defina a estrutura do nó
type Node2 struct {
	Dado     int
	Anterior *Node2
	Proximo  *Node2
}

// Defina a estrutura da lista duplamente encadeada
type DoubleLinkedList struct {
	Cab   *Node2
	Rabo  *Node2
	Tamanho int
}

// Função para exibir a lista duplamente encadeada
func ExibirListaDupla(lista *DoubleLinkedList) {
	no := lista.Cab
	for no != nil {
		fmt.Println(no.Dado)
		no = no.Proximo
	}
	fmt.Println()
}

// Função para buscar um elemento na lista duplamente encadeada
func BuscaListaDupla(lista *DoubleLinkedList, ch int) *Node2 {
	no := lista.Cab
	for no != nil {
		if no.Dado == ch {
			return no
		}
		no = no.Proximo
	}
	return nil
}

// Função para inserir um elemento na lista duplamente encadeada
func InsereListaDupla(lista *DoubleLinkedList, ch int) {
	novoNo := &Node2{Dado: ch}

	if lista.Cab == nil {
		lista.Cab = novoNo
		lista.Rabo = novoNo
	} else if ch <= lista.Cab.Dado {
		novoNo.Proximo = lista.Cab
		lista.Cab.Anterior = novoNo
		lista.Cab = novoNo
	} else if ch >= lista.Rabo.Dado {
		novoNo.Anterior = lista.Rabo
		lista.Rabo.Proximo = novoNo
		lista.Rabo = novoNo
	} else {
		no := lista.Cab
		for no.Proximo != nil && no.Proximo.Dado < ch {
			no = no.Proximo
		}
		novoNo.Proximo = no.Proximo
		novoNo.Anterior = no
		no.Proximo.Anterior = novoNo
		no.Proximo = novoNo
	}
	lista.Tamanho++
}

// Função para remover um elemento da lista duplamente encadeada
func RemoveListaDupla(lista *DoubleLinkedList, ch int) {
	noRemove := BuscaListaDupla(lista, ch)
	if noRemove != nil {
		if noRemove.Anterior != nil {
			noRemove.Anterior.Proximo = noRemove.Proximo
		} else {
			lista.Cab = noRemove.Proximo
		}
		if noRemove.Proximo != nil {
			noRemove.Proximo.Anterior = noRemove.Anterior
		} else {
			lista.Rabo = noRemove.Anterior
		}
		lista.Tamanho--
	}
}

func main() {
	listaDupla := DoubleLinkedList{}

	// Insira elementos na lista duplamente encadeada
	InsereListaDupla(&listaDupla, 3)
	InsereListaDupla(&listaDupla, 1)
	InsereListaDupla(&listaDupla, 2)
	InsereListaDupla(&listaDupla, 4)

	// Exiba a lista duplamente encadeada
	ExibirListaDupla(&listaDupla)

	// Remova um elemento da lista
	RemoveListaDupla(&listaDupla, 2)

	// Exiba a lista novamente após a remoção
	ExibirListaDupla(&listaDupla)
}