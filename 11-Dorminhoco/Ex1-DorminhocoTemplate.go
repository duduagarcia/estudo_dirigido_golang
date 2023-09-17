// Eduardo Garcia, Eduardo Riboli, Matheus Fernandes e Jocemar Nicolodi
// por Fernando Dotti - fldotti.github.io - PUCRS - Escola Politécnica
// PROBLEMA:
//   o dorminhoco especificado no arquivo Ex1-ExplanacaoDoDorminhoco.pdf nesta pasta
// ESTE ARQUIVO
//   Um template para criar um anel generico.
//   Adapte para o problema do dorminhoco.
//   Nada está dito sobre como funciona a ordem de processos que batem.
//   O ultimo leva a rolhada ...
//   ESTE  PROGRAMA NAO FUNCIONA.    É UM RASCUNHO COM DICAS.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const NJ = 5 // Número de jogadores
const M = 4  // Número de cartas

type carta string // carta é uma string

var ch [NJ]chan carta // NJ canais de itens tipo carta
var wg sync.WaitGroup // WaitGroup para esperar que todos os jogadores terminem

func jogador(id int, cartasIniciais []carta) {
	mao := cartasIniciais // estado local - as cartas na mão do jogador
	nroDeCartas := M      // quantas cartas ele tem
	nextPlayer := (id + 1) % NJ

	for {
		if nroDeCartas == M {
			// Jogador tem o número "normal" de cartas, espera uma carta na entrada.
			fmt.Println(id, " espera carta.")
			cartaRecebida, ok := <-ch[id] // Recebe carta na entrada.
			if !ok {
				// Canal fechado, o jogo terminou.
				wg.Done()
				return
			}
			mao = append(mao, cartaRecebida)
			nroDeCartas++
			fmt.Println(id, " recebeu carta:", cartaRecebida)

			// Lógica para verificar se o jogador pode bater.
			if podeBater(mao) {
				fmt.Println(id, " bateu!")
				ch[nextPlayer] <- mao[0] // Passa a carta para o próximo jogador.
				return
			}
		} else {
			// Jogador tem uma carta a mais, escolhe uma e escreve na saída.
			fmt.Println(id, " joga")
			cartaParaSair := mao[0] // Escolha uma carta para passar adiante.
			// fmt.Println("jogador: " + "Carta para sair: ", id, cartaParaSair)
			// fmt.Println("Mao pré jogador: ", id, mao)
			mao = mao[1:]             // Remove a carta da mão.
			// fmt.Println("Mao pós jogador: ", id, mao)
			fmt.Println("Carta para Sair:", cartaParaSair, " Proximo jogador:", nextPlayer)
			fmt.Println(ch)
			ch[nextPlayer] <- cartaParaSair // Manda carta escolhida para o próximo jogador.
			//deadlock

			nroDeCartas--

			// Recebe carta na entrada.
			cartaRecebida, ok := <-ch[id]
			if !ok {
				// Canal fechado, o jogo terminou.
				wg.Done()
				return
			}
			mao = append(mao, cartaRecebida)
			fmt.Println(id, " recebeu carta:", cartaRecebida)

			// Lógica para verificar se o jogador pode bater.
			if podeBater(mao) {
				fmt.Println(id, " bateu!")
				ch[nextPlayer] <- mao[0] // Passa a carta para o próximo jogador.
				return
			}
		}
	}
}

func podeBater(mao []carta) bool {
	// Implemente a lógica para verificar se o jogador pode bater.
	// Neste ponto, você deve verificar se o jogador tem M cartas iguais
	// e retornar true se puder bater, caso contrário, retorne false.
	// Exemplo: Verifique se todas as cartas na mão são iguais.
	if len(mao) != M {
		return false
	}
	primeiraCarta := mao[0]
	for _, carta := range mao {
		if carta != primeiraCarta && carta != "@"{
			return false
		}
	}
	return true
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Inicializa canais de comunicação entre processos
	for i := 0; i < NJ; i++ {
		ch[i] = make(chan carta)
	}

	// Cria um baralho com NJ*M cartas
	baralho := make([]carta, NJ*M+1)
	cartasDisponiveis := make([]carta, NJ*M+1)

	for i := 0; i < NJ; i++ {
		for j := 0; j < M; j++ {
			carta := carta(fmt.Sprintf("%d", i+1))
			baralho[i*M+j] = carta
			cartasDisponiveis[i*M+j] = carta
		}
	}
	baralho[NJ*M] = "@"

	// Embaralha as cartas
	rand.Shuffle(len(baralho), func(i, j int) {
		baralho[i], baralho[j] = baralho[j], baralho[i]
	})

	// Distribui cartas iniciais para os jogadores
	for i := 0; i < NJ; i++ {
		cartasIniciais := baralho[i*M : (i+1)*M]
		go jogador(i, cartasIniciais)
	}

	// Inicia o jogo
	for i := 0; i < NJ; i++ {
		wg.Add(1)
		ch[i] <- cartasDisponiveis[i*M] // Cada jogador recebe uma carta para iniciar.\
	}

	fmt.Println("Aguardando termino")
	// Aguarda o término do jogo
	wg.Wait()

	fmt.Println("Termino")

	// Fecha os canais para indicar que o jogo terminou
	for i := 0; i < NJ; i++ {
		close(ch[i])
		fmt.Println("fechou\n")
	}
}
