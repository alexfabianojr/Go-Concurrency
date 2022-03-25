package main

import (
	"fmt"

	"github.com/alexfabianojr/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		//Enviando dados de um canal para outro
		//Deadlock funciona como feature, trancando for até chegar mais dados
		//Quase como listening de fila
		destino <- <-origem
	}
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)

	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)

	return c
}

func main() {
	c := juntar(
		html.Titulo("https://www.cod3r.com.br", "https://www.google.com.br"),
		html.Titulo("https://www.amazon.com", "https://www.youtube.com"),
	)

	fmt.Println(<-c, " | ", <-c)
	fmt.Println(<-c, " | ", <-c)
}
