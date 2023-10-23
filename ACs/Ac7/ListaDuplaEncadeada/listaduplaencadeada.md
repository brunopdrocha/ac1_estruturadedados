# Lista Duplamente Encadeada

É uma estrutura de dados na qual cada nó contém um ponteiro para o nó anterior e outro para o próximo. Isso permite a navegação em ambas as direções na lista.
A seguir estão as descrições dos nós ordenados em uma lista duplamente encadeadaAntes devemos entendar a base de uma Estrutura da Lista Duplamente Encadeada

# Estrutura Nó 

    Estrutura Nó:
      Dado: valor
      Anterior: ponteiro para o nó anterior na lista duplamente encadeada
      Próximo: ponteiro para o próximo nó na lista duplamente encadeada

    Estrutura Lista Dupla:
        Cab No // Incio da Lista
        Rabo No // Final da Lista

 # Exibição dos Nós em uma Lista Duplamente Encadeada

     Funcao ExibirListaDupla(Lista Encadeada):   
        no = no.cab
        Enquanto no não é Nulo:
          Escrever(no.Dado)
          no= no.Próximo

# Busca de um nó em uma Lista Duplamente Encadeada

        Funcao BuscaListaDupla(Lista Encadeada) (ch int):
            no = ListaEncadeada.cab
            para no não é nulo:
                se no.Dado = ch:
                    retorne no
                no =  no.Proximo
            no = no.Proximo
                
            
# Inserção de um nó em uma Lista Duplamente Encadeada

        Funcao InsereListaDuipla(Lista Encadeada) (ch int):
            Novono = No {Dado: ch}
            se no.cab = nulo:
                ListaEncadeada.cb = Novono
                ListaEncadeada.rabo = Novono
            senao se ch menor ou igual ListaEncadeda.cab.Dado:
                novoNo.Prox = lista.cab
                listaEncadeada.cab.Anterior = novoNo
                listaEncadeada.cab = novoNo
            senao se ch maior ou igual ListaEncadeda.cab.Dado:
                novoNo.Prox = lista.cab
                listaEncadeada.cab.Proximo = novoNo
                listaEncadeada.cab = novoNo
            senao:
                no = no.cab
                para no.Proximo não é nulo e no.Proximo.Dado < ch:
                    no = no.Proximo
                Novono.Proximo = no.Proximo
                Novono.Anterior = no
                no.Proximo.Anterior = Novono
                no.Proximo = NovoNo
                
# Remoção de um nó em uma Lista Duplamente Encadeada

        Funcao RemoveListaDupla(Lista Enceada)(ch int)?
            noRemove = ListaEncadeada.BuscaLista(dado)
            se noRemove.Anterior não é nulo:
                noRemove.Anterior.Proximo = noRemove.Proximo
            senao:
                ListaEncaeada.Cab = noRemove.Proximo
            se noRemove.Proximo não é  nulo:
                noRemove.Proximo.Anterior = noRemove.Anterior
            senao:
                ListaEncadeada = noRemove.Anterior
