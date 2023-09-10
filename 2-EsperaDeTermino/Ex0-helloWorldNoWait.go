// Eduardo Garcia, Eduardo Riboli, Matheus Fernandes e Jocemar Nicolodi
// Fernando Dotti - fldotti.github.io - PUCRS - Escola Politécnica
// programa da internet
// EXERCICIOS:
//    1) rode o programa abaixo.
//       o que você conclui sobre a execução observada?

// RESPOSTA
// Que a ordem dos processo é aleatória, no meu caso, a saída foi 5x a palavra "hello"
// e depois o programa finalizou, pois o processo concorrente main terminou sua execução;
// Certamente há outros casos diferentes de saída, mas não há como prever seus comportamentos.

package main

import (
	"fmt"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
