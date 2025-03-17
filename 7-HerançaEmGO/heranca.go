package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    float32
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Isabelle", "Alencar", 22, 1.73}

	fmt.Println(p1)

	e1 := estudante{p1, "Psicologia", "Uninassau"}
	fmt.Println(e1)
}
