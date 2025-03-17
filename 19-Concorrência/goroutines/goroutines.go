package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		c <- "Entrega chegou!"
	}()

	fmt.Println("Aguardando a entrega...")

	mensagem := <-c
	fmt.Println(mensagem)
	fmt.Println("Continuando o programa.")
}
