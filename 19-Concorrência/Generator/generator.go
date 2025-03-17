package main

import (
	"fmt"
	"time"
)

func main() {
	canaldoprint := escrever("Olá Mundo!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canaldoprint)
	}

}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
