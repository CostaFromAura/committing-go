package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raça  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Martin Scorsese", "Cocker Spaniel Inglês", 3}
	fmt.Println(c)

	cachorroEmJson, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJson)

	BufferedCachorroJSON := bytes.NewBuffer(cachorroEmJson)

	fmt.Println(BufferedCachorroJSON)
}
