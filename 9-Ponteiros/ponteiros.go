package main

import "fmt"

func main() {
	fmt.Println("Revisando os ponteiros")

	numero := 100

	fmt.Println(numero)

	var ponteiro *int

	numero = 200

	ponteiro = &numero

	fmt.Println(numero, *ponteiro)

}
