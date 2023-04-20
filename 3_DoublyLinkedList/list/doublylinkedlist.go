package list

import "fmt"

type No struct {
    val int
    ant  *No
    prox  *No
}

type DoublyLinkedList struct {
    cabeca *No
    cauda *No
    tam int
}

func (lista *DoublyLinkedList) ImprimirLista() {
    // usamos o aux para percorrer a lista
    if lista.cabeca != nil {
        aux := lista.cabeca;
        // navega partindo da cabeça até chegar um valor inválido
        fmt.Printf("[");
        for {
            fmt.Printf("%d", aux.val);
            aux = aux.prox;
            if (aux != nil) {
                fmt.Printf(", ");
            }else{
                break;
            }
        } 
        fmt.Printf("]")
    }else {
        fmt.Println("A lista está vazia!");
    }
    fmt.Println("")
}

func (lista *DoublyLinkedList) InserirElementoNoFim (valor int){
    NovoNo := &No{val: valor}
    if lista.cauda == nil {
        lista.cauda = NovoNo
        lista.cabeca = NovoNo
    } else{
        NovoNo.ant = lista.cauda
        lista.cauda.prox = NovoNo
        lista.cauda = NovoNo
    }
    lista.tam++
}

func (lista *DoublyLinkedList) InserirElementoNoInicio (valor int){
    NovoNo := &No{val: valor}
    if lista.cabeca == nil {
        lista.cabeca = NovoNo
        lista.cauda = NovoNo
    } else{
        NovoNo.prox = lista.cabeca
        lista.cabeca.ant = NovoNo
        lista.cabeca = NovoNo
    }
    lista.tam++
}

func (lista *DoublyLinkedList) InserirElementoEmPosicao (valor, pos int) {
    if pos < 0 || pos > lista.tam {
        fmt.Print("indice fora do permitido")
        return
    }

    NovoNo := &No{val: valor}
    if lista.cabeca == nil {
        lista.cabeca = NovoNo
        lista.cauda = NovoNo
    } else if pos == 0 { 
        NovoNo.prox = lista.cabeca
        lista.cabeca.ant = NovoNo
        lista.cabeca = NovoNo
    } else if lista.tam == pos {
        lista.cauda.prox = NovoNo
        NovoNo.ant = lista.cauda
        lista.cauda = NovoNo
    } else{
        aux := lista.cabeca
        for i:= 0; i <pos-1; i++ {
            aux = aux.prox
        }
        NovoNo.prox = aux.prox
        NovoNo.ant = aux
        aux.prox.ant = NovoNo
        aux.prox = NovoNo

    }
    lista.tam++
    return 
}

func (l *DoublyLinkedList) Remover(pos int) {
    if pos < 0 || pos > l.tam-1 {
        fmt.Println("indice fora do permitido")
        return
    }
    NoAtual := l.cabeca

    for i:=0; i<pos; i++ {
        NoAtual = NoAtual.prox //percorrendo a lista
    }
    // ajustando os valores dos ponteiros ant e prox
    if NoAtual.ant != nil {
        NoAtual.ant.prox = NoAtual.prox 
    } else {
        l.cabeca = NoAtual.prox //caso em que o nó removido era a cabeça
    }
    if NoAtual.prox != nil {
        NoAtual.prox.ant = NoAtual.ant
    } else {
        // se o nó a ser removido for a cauda da lista
        l.cauda = NoAtual.ant
    }

    l.tam--
    
}