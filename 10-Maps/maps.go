package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Isabelle",
		"sobrenome": "Alencar",
	}

	usuario2 := map[string]map[string]string{
		"nome": {
			"Primeiro": "Isabelle",
			"Segundo":  "Veiga",
		},
	}

	fmt.Println(usuario2["nome"])
	fmt.Println(usuario)
}
