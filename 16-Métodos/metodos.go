package main

import "fmt"

type usuario struct {
	nome  string
	idade uint
}

func (u usuario) salvar() {
	fmt.Printf("salvando os dados do Usuário %s no banco de dados", u.nome)
}

func main() {
	usuario1 := usuario{"Isabelle", 22}
	fmt.Println(usuario1)
	usuario1.salvar()
}
