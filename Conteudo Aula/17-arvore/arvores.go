package main

import "fmt"

type No struct {
	Chave    int
	Esq, Dir *No
	Altura   int
}

type Arvore struct {
	Raiz *No
}

func main() {
	arv := Arvore{}

	n1 := &No{Chave: 3}
	n2 := &No{Chave: 4}
	n3 := &No{Chave: 8}
	n4 := &No{Chave: 1}
	n5 := &No{Chave: 0}
	n6 := &No{Chave: 7}

	arv.Raiz = n1
	n1.Esq = n2
	n1.Dir = n3
	n2.Esq = n4
	n4.Esq = n5
	n4.Dir = n6

	// imprimeArvore(arv)
	// calculaAltura(arv.Raiz)
	// fmt.Println(arv.Raiz.Altura)
	// fmt.Println(buscaArvore(arv, 0))
	// fmt.Println(buscaArvore(arv, 7))
	// fmt.Println(buscaArvore(arv, 3))
	// fmt.Println(buscaArvore(arv, 10))

	// var noProcurado *No
	// busca2(arv.Raiz, 7, &noProcurado)
	// fmt.Println(noProcurado)
	// TDD -> Test-driven development
	// insere(arv, 4, 8, "e")   // erro, nó já existe na árvore
	// insere(arv, 10, 11, "e") // erro, nó pai não existe
	// insere(arv, 10, 8, "x")  // erro, não existe essa posição
	// insere(arv, 10, 4, "e")  // erro, já tenho um nó nessa posição
	// insere(arv, 10, 8, "e")  // inseriu o nó 10 à esquerda do 8
	// imprimeArvore(arv)

	// fmt.Println(buscaArvorePai(arv, 1))  // 1, 4
	// fmt.Println(buscaArvorePai(arv, 3))  // 3, nil
	// fmt.Println(buscaArvorePai(arv, 10)) // nil, qualquer coisa

	remove(&arv, 1)
	imprimeArvore(arv)
	fmt.Println("---------------")
	remove(&arv, 10)
	imprimeArvore(arv)
	fmt.Println("---------------")
	remove(&arv, 3)
	imprimeArvore(arv)
}

func insere(a Arvore, valorInserir int, valorPai int, posicao string) {
	if a.Raiz == nil {
		fmt.Println("Erro! A árvore está vazia!")
		return
	}

	if posicao != "e" && posicao != "d" {
		fmt.Println("Erro! Não existe essa posição!")
		return
	}

	no := busca(a.Raiz, valorInserir)
	if no != nil {
		fmt.Println("Erro! O nó procurado já existe!")
		return
	}

	noPai := busca(a.Raiz, valorPai)
	if noPai == nil {
		fmt.Println("Erro! O nó pai não existe!")
		return
	}

	if (noPai.Esq != nil && posicao == "e") || (noPai.Dir != nil && posicao == "d") {
		fmt.Println("Erro! Já existe um nó na posição desejada!")
		return
	}

	no = &No{Chave: valorInserir}
	if posicao == "e" {
		noPai.Esq = no
	} else {
		noPai.Dir = no
	}
}

func remove(a *Arvore, valor int) *No {
	if a.Raiz == nil { return nil }
	no, noPai := buscaArvorePai(*a, valor)

	if no == nil { return nil }

	if noPai == nil {
		a.Raiz = nil
	} else {
		if no == noPai.Esq {
			noPai.Esq = nil
		} else {
			noPai.Dir = nil
		}
	}

	return no
}

func buscaArvorePai(a Arvore, valor int) (*No, *No) {
	if a.Raiz == nil { return nil, nil }

	return buscaPai(a.Raiz, nil, valor)
}

func buscaPai(n *No, nPai *No, valor int) (*No, *No) {
	if n.Chave == valor { return n, nPai }
	var noProcurado *No

	if n.Esq != nil { noProcurado, nPai = buscaPai(n.Esq, n, valor) }
	if noProcurado != nil { return noProcurado, nPai }

	if n.Dir != nil { noProcurado, nPai = buscaPai(n.Dir, n, valor) }

	return noProcurado, nPai
}

func buscaArvore(a Arvore, valor int) *No {
	if a.Raiz == nil { return nil }

	return busca(a.Raiz, valor)
}

func busca(n *No, valor int) *No {
	if n.Chave == valor { return n }
	var noProcurado *No

	if n.Esq != nil { noProcurado = busca(n.Esq, valor) }
	if noProcurado != nil { return noProcurado }

	if n.Dir != nil { noProcurado = busca(n.Dir, valor) }

	return noProcurado
}

func busca2(n *No, valor int, noProcurado **No) {
	if n.Chave == valor { *noProcurado = n }
	if n.Esq != nil { busca2(n.Esq, valor, noProcurado) }
	if n.Dir != nil { busca2(n.Dir, valor, noProcurado) }
}

func imprimeArvore(a Arvore) {
	if a.Raiz != nil { imprimePosOrdem(a.Raiz) }
}

func imprimePreOrdem(n *No) {
	fmt.Println(n.Chave)
	if n.Esq != nil { imprimePreOrdem(n.Esq) }
	if n.Dir != nil { imprimePreOrdem(n.Dir) }
}

func imprimeSimetrico(n *No) {
	if n.Esq != nil { imprimeSimetrico(n.Esq) }
	fmt.Println(n.Chave)
	if n.Dir != nil { imprimeSimetrico(n.Dir) }
}

func imprimePosOrdem(n *No) {
	if n.Esq != nil { imprimePosOrdem(n.Esq) }
	if n.Dir != nil { imprimePosOrdem(n.Dir) }
	fmt.Println(n.Chave)
}

func calculaAltura(n *No) {
	if n.Esq != nil { calculaAltura(n.Esq) }
	if n.Dir != nil { calculaAltura(n.Dir) }
	calculaAlturaNo(n) // visita
}

func calculaAlturaNo(n *No) {
	var altE, altD int

	if n.Esq != nil {
		altE = n.Esq.Altura
	} else {
		altE = 0
	}

	if n.Dir != nil {
		altD = n.Dir.Altura
	} else {
		altD = 0
	}

	if altE > altD {
		n.Altura = altE + 1
	} else {
		n.Altura = altD + 1
	}
}