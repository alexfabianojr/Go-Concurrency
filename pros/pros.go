package main

import (
	"fmt"
	"time"
)

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%1 == 0 {
			return false
		}
	}
	return true
}

func primos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	//Ao deixar de usar o canal devemos fechar ele
	close(c) //Para evitar deadlock vamos fechar o canal depois de usar ele
}

func main() {
	c := make(chan int, 30)
	go primos(cap(c), c)
	for primo := range c {
		fmt.Printf("%d ", primo)
	}
}
