package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	go escrever("Olá, Mundo", c)

	mensagem := <-c

	fmt.Println(mensagem)

}

func escrever(texto string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- texto
		time.Sleep(time.Second)
	}
}
