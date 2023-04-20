package list

import "fmt"

type No struct {
    val int
    prox *No

}

 type LinkedList struct{
    cabeca *No
    tamanho int
};

func (lista *LinkedList) ImprimirLista() {
    NoAtual := lista.cabeca
    if NoAtual != nil{
        fmt.Printf("[")
        for NoAtual != nil {
            if NoAtual.prox == nil{
                fmt.Printf("%d", NoAtual.val)
                break
            }
            fmt.Printf("%d, ", NoAtual.val)
            NoAtual = NoAtual.prox
        }
        fmt.Println("]")
    }
}


func (lista *LinkedList) InserirElementoNoFim(valor int){
	auxiliar := &No{val: valor}
	if lista.cabeca == nil{
		lista.cabeca = auxiliar; 
	} else{
        NoAtual := lista.cabeca
        for NoAtual.prox != nil{
            NoAtual = NoAtual.prox
        }
        NoAtual.prox = auxiliar
    }
	lista.tamanho++
}

//Essa função remover o nó que possui o valor informado na chamada
func (l *LinkedList) Remove(val int) bool {
    if l.cabeca == nil {
        return false
    }
    if l.cabeca.val == val {
        l.cabeca = l.cabeca.prox
        l.tamanho--
        return true
    }
    NoAtual := l.cabeca
    for NoAtual.prox != nil {
        if NoAtual.prox.val == val {
            NoAtual.prox = NoAtual.prox.prox
            l.tamanho--
            return true
        }
        NoAtual = NoAtual.prox
    }
    return false
}

func (l *LinkedList) InserirEmPosicao(valor int, posicao int) bool {
    //Verificando se a posição é válida
    if posicao < 0 || posicao > l.tamanho {
        return false
    }
    novoNo := &No{val: valor}
    if posicao == 0 {
        novoNo.prox = l.cabeca
        l.cabeca = novoNo
    } else {
        NoAtual := l.cabeca
        for i := 0; i < posicao-1; i++ {
            NoAtual = NoAtual.prox
        }
        novoNo.prox = NoAtual.prox
        NoAtual.prox = novoNo
    }
    l.tamanho++
    return true
}

func (l *LinkedList) Tamanho () int {
    return l.tamanho
}

