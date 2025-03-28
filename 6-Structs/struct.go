package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	peso  uint64
}

func main() {
	fmt.Println("Arquivo de structs")

	var u usuario

	u.nome = "Teste"
	u.idade = 22
	u.peso = 90

	fmt.Println(u)

	usuario2 := usuario{"Marcelo", 21, 90}

	fmt.Println(usuario2)
}
