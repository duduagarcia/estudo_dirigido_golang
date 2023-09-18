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
)

const NJ = 5 // Número de jogadores
const M = 4  // Número de cartas

type carta string // carta é uma string

var ch [NJ]chan carta // NJ canais de itens tipo carta

func jogador(id int, cartasIniciais []carta, ganhador chan int, bateu chan int) {
	mao := cartasIniciais // estado local - as cartas na mão do jogador
	nroDeCartas := M      // quantas cartas ele tem
	nextPlayer := (id + 1) % NJ

	done := make(chan struct{})
	go baterAlternativo(bateu, id, done)

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
			bateu <- id
		} else {
			ch[nextPlayer] <- cartaParaSair // Manda carta escolhida para o próximo jogador.
		}

	}
}

func baterAlternativo(bateu chan int, id int, done chan struct{}) {
	/* fmt.Println("ANTES DO BATEU")
	valido := <-bateu
	fmt.Println("PASSOU DO BATEU, ID:", id)
	if valido == id {
		fmt.Println("JÁ TEM NO BATEU, ID:", id)
		return
	}

	fmt.Println("INSERIU NO BATEU, ID:", id)
	bateu <- id */

	// verificar se alguém ja bateu através do canal ganhador
	// se alguém bateu, verificar se o id do player é igual ao id ganhador
	// se não for, incluir no canal com buffer
	// se for, encerrar o processo para este player

	//valido:= <-bateu
	
	inseriu = false
	for num := range bateu {
		if num == id {
			done <- struct{}{}
			inseriu = true
		} 
	} 
	if !inseriu {
		//fmt.Println("INSERIU NO BATEU, ID:", id)
		bateu <- id
		fmt.Println("INSERIU NO BATEU, ID:", id)
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
	ganhador := make(chan int)
	bateu := make(chan int, 5)

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
		go jogador(i, cartasIniciais, ganhador, bateu)
	}

	// Inicia o jogo
	ch[0] <- cartasDisponiveis[0] // Cada jogador recebe uma carta para iniciar.\

	// Aguarda o término do jogo
	fmt.Println("Aguardando termino")
	// printa os itens do buffer bateu
	
	// Faça a análise desejada aqui, por exemplo, imprimir o número
	/* while bateu != nil {
		fmt.Println("Número recebido:", <-bateu)
	} */
	//fmt.Println("Número recebido:", <-bateu)

    

	// Printa os ganhadores
	for i := 0; i < NJ; i++ {
		fmt.Println("Ganhador", i+1, ":", <-ganhador)
		if i == NJ-1 {
			fmt.Println("DORMINHOCO: ", <-ganhador)
		}
	}

	fmt.Println("Termino")
}
