// multiplicacao de matrizes
// por Fernando Dotti - PUCRS

// Os processos "calculateCell" são sincronizados com canais, e acessam a mesma matriz de dados.
// Este tipo de modelagem pode ser utilizado para processamento de imagens.
// Cada processo fica responsável por uma parte da imagem - aqui correspondendo a uma célula da matriz.

package main

import (
	"fmt"
)

const (
	N = 10
)

func generateMatrix(n int) [][]int {

	out := make([][]int, n) // gera matriz
	for i := 0; i < len(out); i++ {
		out[i] = make([]int, n)
	}

	for i := 0; i < n; i++ { // preenche com valores crescentes
		for j := 0; j < n; j++ {
			out[i][j] = i*10 + j
		}
	}
	return out
}

func processMatrix(m [][]int, iterations int) {
	// ESTE PROCESSO CRIA UM PROCESSO PARA CADA CELULA E FICA EM LOOP IMPRIMINDO OS VALORES DA MATRIZ

	arrive := make(chan struct{})
	wait := make(chan struct{})

	for i := 0; i < N; i++ { // gera N^2 processos para calcular celulas
		for j := 0; j < N; j++ {
			go calculateCell(i, j, m, arrive, wait, iterations)
		}
	}

	for i := 1; i <= 2*iterations; i++ { // a cada iteracao, duas vezes por iteracao

		for k := 0; k < N*N; k++ { // espera N^2 processos chegaram no ponto de sincronizacao ...
			<-arrive
		}

		if (i % 2) == 0 {
			fmt.Println(m)
		}
		for k := 0; k < N*N; k++ { // e entao libera-os ...
			wait <- struct{}{}
		}

	}

}

func calculateCell(i int, j int, m [][]int, arrive, wait chan struct{}, iter int) {
	// ESTE PROCESSO CALCULA EM LOOP OS NOVOS VALORES DA POSICAO I J DA MATRIZ, PASSADA POR PARAMETRO

	for k := 1; k <= iter; k++ {

		// le valores e calcula
		sum := m[i][j] + m[i][(j+1)%N] + m[i][(N+j-1)%N] + m[(i+1)%N][j] + m[(N+i-1)%N][j]

		arrive <- struct{}{} // chega em ponto de sincronizacao
		<-wait               // sai do ponto de sincronizacao

		m[i][j] = sum / 5 // escreve media

		arrive <- struct{}{} // chega em ponto de sincronizacao
		<-wait               // sai do ponto de sincronizacao

	}
	// fmt.Println("cell ", i, ",", j, " = ", m[i][j])
}

func main() {

	A := generateMatrix(N)
	fmt.Println(A)

	processMatrix(A, 10)
	fmt.Println(A)

}
