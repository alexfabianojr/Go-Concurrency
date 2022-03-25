package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)

	c <- 1 //Operacao bloqueante
	fmt.Println("Só depois que foi lido")
}

//Deadlock é gerado no momento que o canal fica esperando entrar algum dado de uma goroutine para ser lido
func main() {

	//Quando se tem um canal com buffer podemos colocar dentro dele
	//de forma assíncrona até o buffer ficar cheio
	//Quando o buffer encher é ler ele para abrir espaço

	//Quando o canal não tem buffer o primeiro elemento já gera o bloqueio (deadlock)

	c := make(chan int) //canal sem buffer

	go rotina(c)

	fmt.Println(<-c)
	fmt.Println("Foi lido")
	fmt.Println(<-c) //deadlock - não tem mais dados para ser lido
	//Não tem mais nenhuma goroutine para enviar dado ao channel

	fmt.Println("Fim")
}
