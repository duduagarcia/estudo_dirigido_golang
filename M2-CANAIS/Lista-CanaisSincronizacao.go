package main

import "fmt"

//- A --------------------------

func atividadeAqA(c, fin chan struct{}) {
	fmt.Println(" Atividade A ... comandos")
	c <- struct{}{}
	fin <- struct{}{}
}
func atividadeBqA(c, fin chan struct{}) {
	<-c
	fmt.Println(" Atividade B ... comandos")
	fin <- struct{}{}
}

func questaoA() {
	c := make(chan struct{})   // canal de tamanho zero = sincronizante
	fin := make(chan struct{}) // canal de tamanho zero = para sinalizar fim
	go atividadeAqA(c, fin)    // inicia processo concorrente
	go atividadeBqA(c, fin)    // inicia processo concorrente
	<-fin
	<-fin

	// Q.A Pergunta-se:  os comandos executados nas atividades A e B ?
	// a) serão entrelaçados no tempo
	// b) serão separados no tempo, B antes de A
	// c) serão separados no tempo, A antes de B
}

//- B --------------------------

func atividadeAqB(c chan struct{}) {
	fmt.Println(" Atividade A ... comandos")
	c <- struct{}{}
}
func atividadeBqB(c chan struct{}) {
	fmt.Println(" Atividade B ... comandos")
	c <- struct{}{}
}
func atividadeCqB(c, fin chan struct{}) {
	<-c
	<-c
	fmt.Println(" Atividade C ... comandos")
	fin <- struct{}{}
}
func questaoB() {
	c := make(chan struct{})   // canal de tamanho zero = sincronizante
	fin := make(chan struct{}) // canal de tamanho zero = para sinalizar fim
	go atividadeAqB(c)         // inicia processo concorrente
	go atividadeBqB(c)         // inicia processo concorrente
	go atividadeCqB(c, fin)
	<-fin

	// Q.B Pergunta-se
	// a) atividades A, B e C entrelaçam no tempo ?
	// b) A, B e C serão separadas no tempo, A antes de B, antes de C
	// b) A e B entrelaçam no tempo, ambas acabam antes de B iniciar
}

//- C --------------------------

//  como voce montaria um grafo de dependencias entre atividades  ?
//
// A  ──────► B ─────────► F
// │ │                     ▲
// │ └──────► C ──►E ──────┘
// │               ▲
// └────────► D ───┘

//- D --------------------------

func proc(v int, sync1 chan struct{}, sync2 chan struct{}, fin chan struct{}) {
	fmt.Println(v, " Fase A")
	sync1 <- struct{}{}
	<-sync2
	fmt.Println(v, " Fase B")
	fin <- struct{}{}
}
func questaoC() {
	sync1 := make(chan struct{})
	sync2 := make(chan struct{})
	chfin := make(chan struct{})
	for i := 0; i < 10; i++ {
		go proc(i, sync1, sync2, chfin)
	}
	for i := 0; i < 10; i++ {
		<-sync1
	}
	for i := 0; i < 10; i++ {
		sync2 <- struct{}{}
	}
	for i := 0; i < 10; i++ {
		<-chfin
	}
}

// é possível que processos "proc" diferentes estejam em fases diferentes ?
// como será a saída de questaoC ?

//---------------------------

func main() {
	questaoA()
	questaoB()
	questaoC()
}
