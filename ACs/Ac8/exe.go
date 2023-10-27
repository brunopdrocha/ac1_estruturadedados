package main

import "fmt"

type No struct {
	chave int
	esq   *No
	dir   *No
}

type Arvore struct {
	raiz *No
}

func buscaArvore(valor int, n *No) int {
	if n == nil {
		return 0
	}
	if valor == n.chave {
		return 1
	}
	if valor < n.chave {
		if n.esq == nil {
			return 2
		}
		n := n.esq
		return buscaArvore(valor, n.esq)
	} else {
		if n.dir == nil {
			return 3
		}
		n := n.dir
		return buscaArvore(valor, n.dir)
	}
}

func insere(valor int, a *Arvore) string {
	no := a.raiz
	var pai *No
	f := buscaArvore(valor, no)

	if f == 1 {
		return "Insercao Invalida"
	}

	novoNo := &No{
		chave: valor,
		esq:   nil,
		dir:   nil,
	}

	for no != nil {
		pai = no
		if valor < no.chave {
			no = no.esq
		} else {
			no = no.dir
		}
	}

	if pai == nil {
		a.raiz = novoNo
	} else if valor < pai.chave {
		pai.esq = novoNo
	} else {
		pai.dir = novoNo
	}

	return "Valor inserido"
}

func imprimeArvore(a *Arvore) {
	imprimePreOrdem(a.raiz)
}

func imprimePreOrdem(n *No) {
	fmt.Println(n.chave)
	if n.esq != nil {
		imprimePreOrdem(n.esq)
	}
	if n.dir != nil {
		imprimePreOrdem(n.dir)
	}
}

func main() {
	arvore := &Arvore{raiz: nil}
	insere(4, arvore)
	insere(2, arvore)
	insere(5, arvore)
	insere(6, arvore)
	insere(1, arvore)

	imprimeArvore(arvore)
	fmt.Println("-------------------------")
	fmt.Println(buscaArvore(0, arvore.raiz))
	fmt.Println(buscaArvore(1, arvore.raiz))
	fmt.Println(buscaArvore(9, arvore.raiz))
}
