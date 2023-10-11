package main

import "fmt"

type No struct {
	chave int
	prox  *No
}

type Lista struct {
	cab *No
}

//Exibe os nos de lista

func (l *Lista) exibe() {

	no := l.cab

	for no != nil {
		fmt.Println(no)
		no = no.prox
	}
}

//busca simples de nos da lista

func (l *Lista) buscaSimples(ch int) *No {

	no := l.cab

	for no != nil {
		if no.chave == ch {
			return no
		}
	}
	return nil
}

//Insere No no final da lista

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

//busca um No em uma lista ordenada
func (l *Lista) busca(ch int) (*No, *No) {
	var ant *No
	no := l.cab

	if no == nil {
		return nil, nil
	}

	for no != nil {
		if no.chave == ch {
			return ant, no
		}
		if no.chave > ch {
			return ant, nil
		}

		ant = no
		no = no.prox
	}
	return ant, nil
}

//insere um No em uma lista ordenada

//remove um No em uma lista ordenada

//Funcao main
func main() {
	var l Lista
	//l insere

	l.insere(2)
	l.insere(4)
	l.insere(6)
	l.insere(8)
	
	l.exibe()

	//l.exibe

	//fmt.Println(l.buscaSimples(0))
	//fmt.Println(l.buscaSimples(7))
	//fmt.Println(l.buscaSimples(3))
	//fmt.Println(l.buscaSimples(4))

}
