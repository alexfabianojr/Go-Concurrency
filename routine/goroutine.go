package main

import (
	"fmt"
	"time"
)

func main() {

	//Bloco sem concorrencia
	fale("Maria", "Pq vc nao fala comigo?", 3)
	fale("Joao", "So posso falar depois de voce", 1)

	//Bloco com concorrencia
	go fale("Pedro", "Ei....", 500)
	go fale("Ana", "Opa...", 500)

	//Palavra reservada go funciona "como" abrir uma thread, no caso goroutines

	//A go routine so vai continuar funcionando se a thread principal do programa continuar viva (main)

	fmt.Println("Fim!") //Para manter a thread viva
}

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteracao %d)\n", pessoa, texto, i+1)
	}
}
