package main

import (
	"fmt"
	"strings"
)

type Pessoa struct {
	Nome  string
	Email string
}

// Use pointer receiver to modify the instance
func (p *Pessoa) alterarEmail(novoEmail string) {
	p.Email = novoEmail
}

// Return updated array and fixed the function signature
func adicionarPessoa(p Pessoa, a *[5]Pessoa) {
	for ind, pessoa := range a {
		if (pessoa == pessoa) {
			a[ind] = p
			break
		}
	}
}

func main() {
	var pessoas [5]Pessoa
	p1 := Pessoa{
		Nome:  "AAA",
		Email: "BBB",
	}

	fmt.Println(pessoas)
	adicionarPessoa(p1, &pessoas)
	fmt.Println(pessoas)

	x := 5
	y := x
	fmt.Println(x, y) // 5 e 5

	x = 6
	fmt.Println(x, y) // 6 e 5

	z := &x
	fmt.Println(z)  // Endereço de memória em que x está armazenado
	fmt.Println(*z) // Desreferenciação -> retorna o valor que está sendo apontado por z

	x = 7
	fmt.Println(x, *z) // 7 e 7

	var w *int
	fmt.Println(w) // nil
	// fmt.Println(*w) --> Isso causará um erro!

	mensagem := "Ola, mundo"
	alteraMensagem(&mensagem)
	fmt.Println(mensagem)

	a := &z
	fmt.Println(z)   // Endereço de x
	fmt.Println(a)   // Endereço de z
	fmt.Println(*a)  // Endereço de x
	fmt.Println(**a) // Valor de x
}

func alteraMensagem(msg *string) {
	*msg = strings.ReplaceAll(*msg, "mundo", "turma")
}
