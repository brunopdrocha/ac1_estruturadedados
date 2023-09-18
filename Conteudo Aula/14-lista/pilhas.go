package main

import "fmt"

const M = 5

func main() {
	var pilha [M]int
	topo := -1

	push(&pilha, &topo, 2)
	fmt.Println(pilha)
	push(&pilha, &topo, 4)
	fmt.Println(pilha)
	push(&pilha, &topo, 1)
	fmt.Println(pilha)
	push(&pilha, &topo, 5)
	fmt.Println(pilha)
	push(&pilha, &topo, 7)
	fmt.Println(pilha)
	push(&pilha, &topo, 3) //Overflow
	fmt.Println(pilha)

	pop(&pilha, &topo)
	fmt.Println(pilha)
	pop(&pilha, &topo)
	fmt.Println(pilha)
	pop(&pilha, &topo)
	fmt.Println(pilha)
	pop(&pilha, &topo)
	fmt.Println(pilha)
	pop(&pilha, &topo)
	fmt.Println(pilha)
	pop(&pilha, &topo) //Underflow
	fmt.Println(pilha)

}

func push(p *[M]int, topo *int, valor int) {
	if *topo+1 == M {
		fmt.Println("Overflow")
		return
	}
	*topo++
	p[*topo] = valor
}

func pop(p *[M]int, topo *int) int {

	if *topo == -1 {
		fmt.Println("Underflow")
		return -1
	}

	valorRemovido := p[*topo]
	p[*topo] = 0
	fmt.Println(valorRemovido)
	*topo--
	return valorRemovido

}
