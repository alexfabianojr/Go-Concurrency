package main

import (
	"fmt"
	"time"
)

func main() {

	//Channel (canal) é a forma de comunicação entre goroutines
	//Channel é um tipo

	ch := make(chan int, 1)

	ch <- 1 //Enviando dados para o canal (escrita) - ENTRANDO

	<-ch //Recebendo dados do canal (leitura) - SAIDA

	ch <- 2

	fmt.Println(<-ch)

	channel := make(chan int)

	go doisTresQuatroVezes(2, channel)

	fmt.Println("A")

	//Isso aqui faz a goroutine principal esperar as outras
	a, b := <-channel, <-channel //Recebendo dados do canal

	fmt.Println("B")

	fmt.Println(a, b)

	fmt.Println(<-channel)

}

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)

	c <- 2 * base //enviando dados para o canal

	time.Sleep(time.Second)

	c <- 3 * base

	time.Sleep(time.Second)

	c <- 4 * base
}
