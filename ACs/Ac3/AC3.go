package main

import (
	"bufio"
	"fmt"
	"os"
	"programa/util"
	"strings"
)

func main() {
	var contatos [5]util.Contato
	var nome, email, opcao string

	leitor := bufio.NewReader(os.Stdin)

	fmt.Println("Lista de contatos!")
	for {
		fmt.Print("Digite (1) para adicionar, (2) para remover,(3) para editar email ou qualquer outra coisa para sair:,(4) para vizualizar e-mails registrados ")
		fmt.Scanln(&opcao)

		switch opcao {
		case "1":
			fmt.Print("Informe o nome: ")
			nome, _ = leitor.ReadString('\n')
			nome = strings.TrimSpace(nome)

			fmt.Print("Informe o email: ")
			email, _ = leitor.ReadString('\n')
			email = strings.TrimSpace(email)

			util.AdicionaContato(util.Contato{Nome: nome, Email: email}, &contatos)
		case "2":
			util.ExcluiContato(&contatos)
		case "3":
			var ind int
			fmt.Print("Informe o índice do contato para editar o email: ")
			fmt.Scanln(&ind)

			fmt.Print("Informe o email: ")
			email, _ = leitor.ReadString('\n')
			email = strings.TrimSpace(email)

			util.EditaEmail(&contatos, ind-1, email)

		case "4":
			util.ExibeContatos(&contatos)
		default:
			return
		}

	}
}
