package enderecos

import (
	"fmt"
	"testing"
)

type cenariosDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	cenariosDeTeste := []cenariosDeTeste{
		{"Rua ABC", "rua"},
		{"Avenida ABC", "avenida"},
		{"Bosque da Floresta", "Tipo Inválido"},
		{" ", "Tipo Inválido"},
		{"Estrada da Belinha", "estrada"},
		{"Rodovia Sir Marti", "rodovia"},
	}

	for _, cenario := range cenariosDeTeste {
		enderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if enderecoRecebido != cenario.retornoEsperado {
			fmt.Printf("O endereço inserido %s não é o esperado %s. \n", enderecoRecebido, cenario.retornoEsperado)
		}
	}
}
