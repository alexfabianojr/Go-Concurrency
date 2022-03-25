package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	go rotina(ch)

	//Vai gerar um deadlock diferente
	//o channel vai ficar esperando ser lido para por o resto dos dados

	fmt.Println(<-ch)
}

func rotina(ch chan int) {
	fmt.Println("A")
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	fmt.Println()
	ch <- 5
	ch <- 6
}
