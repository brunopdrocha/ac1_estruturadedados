package main

import (
	"fmt"
)

type Contato struct {
	Nome  string
	Email string
}

// Adicionar contatos no sistema
func adicionarContato(c Contato, contatos [5]Contato) [5]Contato {
	for i, usuario := range contatos {
		if (usuario == Contato{}) {
			contatos[i] = c
			break
		}

	}
	return contatos
}

// Remover um elemento da Lista
func excluirContato(contatos [5]Contato) [5]Contato {

	for i, usuario := range contatos {
		if (usuario == Contato{}) {
			contatos[i-1] = Contato{}
			break
		}

	}
	return contatos
}

// Copilador codigo
func main() {

	fmt.Println("1-adicionarUsuario")
	fmt.Println("2-removerUsuario")
	fmt.Println("3-Sair")
	var opcao int
	fmt.Scanln(&opcao)

	var dados [5]Contato

	fmt.Println("Informe nome1:")
	var c1 Contato
	c1 = Contato{
		Nome:  "Bruno",
		Email: "Email",
	}
	dados[0] = c1
	var c2 Contato
	c2 = Contato{
		Nome:  "Math",
		Email: "Email2",
	}
	dados[1] = c2

	var c3 Contato
	c3 = Contato{
		Nome:  "Otavio",
		Email: "Email3",
	}
	dados[2] = c3

	var c4 Contato
	c4 = Contato{
		Nome:  "Mario",
		Email: "Email4",
	}
	dados[3] = c4

	var c5 Contato
	c5 = Contato{
		Nome:  "Lucas",
		Email: "Email5",
	}
	dados[4] = c5

	fmt.Println(dados)
	excluirContato(dados)
	fmt.Println(dados)
}
