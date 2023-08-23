package util

import "fmt"

func (c *Contato) AlterarEmail(novoEmail string) {
	c.Email = novoEmail
}
func AdicionaContato(c Contato, a *[5]Contato) {
	for ind, contato := range a {
		if (contato == Contato{}) {
			a[ind] = c
			break
		}
	}

}

func ExcluiContato(a *[5]Contato) {
	if (a[0] == Contato{}) {
		fmt.Println("Lista de contatos vazia!")

	}

	for ind, contato := range a {
		if (contato == Contato{}) {
			a[ind-1] = Contato{}
		}
	}

}
func EditaEmail(listaCont *[5]Contato, ind int, novoEmail string) {
	listaCont[ind].AlterarEmail(novoEmail)
}

func ExibeContatos(listaCont *[5]Contato) {
	
	fmt.Println("Lista de Contatos:")
	for ind, contato := range listaCont {
		if contato.Nome != "" || contato.Email != "" {
			fmt.Printf("[%d] Nome: %s | Email: %s\n", ind+1, contato.Nome, contato.Email)
		}
	}
}