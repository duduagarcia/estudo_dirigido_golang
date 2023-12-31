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
	"strconv"
)

const NJ = 5 // Número de jogadores
const M = 4  // Número de cartas
var ganhadores = []string{}

type carta string // carta é uma string

var ch [NJ]chan carta // NJ canais de itens tipo carta

func jogador(id int, cartasIniciais []carta, ganhador chan int, done chan struct{}) {
	mao := cartasIniciais // estado local - as cartas na mão do jogador
	nroDeCartas := M      // quantas cartas ele tem
	nextPlayer := (id + 1) % NJ

	go baterAlternativo(id, ganhador, done)

	for {
		// Jogador tem o número "normal" de cartas, espera uma carta na entrada.
		fmt.Println(id, " espera carta.")
		cartaRecebida := <-ch[id] // Recebe carta na entrada.
		mao = append(mao, cartaRecebida)
		nroDeCartas++
		fmt.Println(id, " recebeu carta:", cartaRecebida)

		// Jogador tem uma carta a mais, escolhe uma e escreve na saída.
		fmt.Println(id, " joga")
		cartaParaSair := mao[0] // Escolha uma carta para passar adiante.
		mao = mao[1:]           // Remove a carta da mão.
		nroDeCartas--
		// Lógica para verificar se o jogador pode bater.
		if podeBater(mao) {
			fmt.Println(id, " bateu!")
			ganhador <- id
			ganhadores = append(ganhadores, strconv.Itoa(id))
			return
		} else {
			ch[nextPlayer] <- cartaParaSair // Manda carta escolhida para o próximo jogador.
		}

	}
}

func baterAlternativo(id int, ganhador chan int, done chan struct{}) {
	<-done

	var contains = false
	for _, winner := range ganhadores {
		if winner == strconv.Itoa(id) {
			contains = true
			done <- struct{}{}
		}
	}
	if !contains {
		ganhadores = append(ganhadores, strconv.Itoa(id))
		ganhador <- id
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
		if carta != primeiraCarta && carta != "@" {
			return false
		}
	}
	return true
}

func main() {
	ganhador := make(chan int, 1)
	done := make(chan struct{})

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
		go jogador(i, cartasIniciais, ganhador, done)
	}

	// Inicia o jogo
	ch[0] <- cartasDisponiveis[0] // Cada jogador recebe uma carta para iniciar.\

	// Aguarda o término do jogo
	fmt.Println("Aguardando termino")

	// Printa os ganhadores
	for i := 0; i < NJ; i++ {
		if i == NJ-1 {
			fmt.Println("DORMINHOCO: ", <-ganhador)
			break
		} else {
			fmt.Println("Ganhador", i+1, ":", <-ganhador)
		}
		done <- struct{}{}
	}

	fmt.Println("Termino")
}
