package main

import "fmt"

func notasAluno(n1, n2 float32) bool {
	defer fmt.Println("Resultados na tela!")
	fmt.Println("Calculando a média. Aguarde!")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {
	println(notasAluno(5, 5))
}
