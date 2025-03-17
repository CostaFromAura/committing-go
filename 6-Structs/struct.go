package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func main() {
	fmt.Println("Arquivo de structs")

	var u usuario

	u.nome = "Isabelle"
	u.idade = 22

	fmt.Println(u)

	usuario2 := usuario{"Marcelo", 21}

	fmt.Println(usuario2)
}
