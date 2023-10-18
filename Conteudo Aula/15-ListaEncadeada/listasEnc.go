package main

import "fmt"

type No struct {
	chave int
	prox  *No
}

type Lista struct {
	cab *No
}

// exibe os nós da lista
func (l *Lista) exibe() {
	no := l.cab

	for no != nil {
		fmt.Println(no)
		no = no.prox
	}
}

// busca simples dos nós da lista
func (l *Lista) buscaSimples(ch int) *No {
	no := l.cab

	for no != nil {
		if no.chave == ch { return no }
		no = no.prox
	}

	return nil
}

// insere nó no final da lista
func (l *Lista) insere(ch int) {
	novoNo := &No{chave: ch}
	no := l.cab

	if no == nil {
		l.cab = novoNo
	} else {
		for no.prox != nil {
			no = no.prox
		}
		no.prox = novoNo
	}
}

// busca um nó em uma lista ordenada
func (l *Lista) busca(ch int) (*No, *No) {
	var ant *No
	no := l.cab

	if no == nil { return nil, nil }

	for no != nil {
		if no.chave == ch { return ant, no }
		if no.chave > ch { return ant, nil }

		ant = no
		no = no.prox
	}

	return ant, nil
}

// insere um nó em uma lista ordenada
func (l *Lista) insereOrd(ch int) {
	ant, no := l.busca(ch)

	if no != nil { return }

	novoNo := &No{chave: ch}
	if ant == nil {
		novoNo.prox = l.cab
		l.cab = novoNo
	} else {
		novoNo.prox = ant.prox
		ant.prox = novoNo
	}
}

// remove um nó em uma lista ordenada
func (l *Lista) remove(ch int) *No {
	ant, no := l.busca(ch)

	if no == nil { return nil }

	if ant == nil {
		l.cab = no.prox
	} else {
		ant.prox = no.prox
	}

	return no
}

func main() {
	var l Lista

	l.insereOrd(2)
	l.insereOrd(1)
	l.insereOrd(9)
	l.insereOrd(3)

	l.remove(0)
	l.remove(1)
	l.remove(9)
	l.exibe()
	// fmt.Println(l.buscaSimples(0))
	// fmt.Println(l.buscaSimples(7))
	// fmt.Println(l.buscaSimples(3))
	// fmt.Println(l.buscaSimples(4))
}