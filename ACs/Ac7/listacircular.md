# Lista Circular
  
  Uma lista circular é uma estrutura de dados onde o último nó está conectado ao primeiro, criando um ciclo. 
  Nela, podemos realizar três operações principais: exibição dos nós, inserção de um nó no início e exclusão do primeiro nó.
  Antes devemos entendar a base de uma Estrutura do Nó


# Estrutura do Nó:

    Estrutura Nó:
        Dado: valor
        Próximo: ponteiro para o próximo nó na lista circular

# Exibição dos Nós em uma Lista Circular

    funcao ExibirListaCircular(Lista Circular):
        no = Lista.cab
        para no não nulo:
          escreve(no)
          no = no.proximo
# Inserção de um nó no início da Lista

    funcao InserirPrimeiroNo(Lista Circular)(valor int):
      novoNo = No{Dado = Valor, Prox = ListaCircular.Inicio}
      se ListaCircular.Incio == nulo:
        newNovo.proximo = novoNo
        senao 
          atual = ListaCircular.proximo
          enquanto atual.proximo não é ListaCircular.proximo
            atual = atual.proximo
          atual.proximo = novoNo
      ListaCircular.inicio = novoNo
      
  

#  Exclusão do Primeiro Nó

    funcao ExcluirPrimeiroNo(Lista Circular)

     se ListaCircular.Inicio.Proximo == ListaCircular.Incio:
       ListaCircular.Inicio = nil
       senao
       atual = ListaCircular= inicio
       enquanto atual.proximo não é ListaCircular.Proximo:
         atual = atual.Proximo

        atual.Proximo = ListaCircular.Proximo
        ListaCircular.Inicio = ListaCircular.Inicia.Proximo

      
      
    

