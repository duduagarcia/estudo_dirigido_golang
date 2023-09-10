// Eduardo Garcia, Eduardo Riboli, Matheus Fernandes e Jocemar Nicolodi
// por Fernando Dotti - fldotti.github.io - PUCRS - Escola Politécnica
// EXERCÍCIO:  dado o programa abaixo
//    1) quantos processos concorrentes são gerados ?
//    2) execute e observe: que se pode supor sobre a velocidade relativa dos mesmos ?
// OBSERVACAO:o sleep no método main serve para este nao acabar, o que acabaria todos processos em execucao.
//     mais adiante veremos outras formas de sincronizar isto

// RESPOSTAS
// 1) 40 processos funcaoA + 1 processo main = total de 41 processos concorentes
// 2) Já que estamos considerando interleaving arbitrário, não podemos tirar
//     qualquer suposição temporal sobre a velocidade relativa dos processos.

package main

import (
	"fmt"
	"time"
)

var N int = 40

func funcaoA(id int, s string) {
	for {
		fmt.Println(s, id)
	}
}

func geraNespacos(n int) string {
	s := "  "
	for j := 0; j < n; j++ {
		s = s + "   "
	}
	return s
}

func main() {
	for i := 0; i < N; i++ {
		go funcaoA(i, geraNespacos(i))
	}
	for true {
		time.Sleep(100 * time.Millisecond)
	}
}
