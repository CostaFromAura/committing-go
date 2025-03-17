package enderecos

import "strings"

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{
		"rua", "avenida", "estrada", "rodovia",
	}

	enderecoFormatado := strings.ToLower(endereco)

	primeiraPalavra := strings.Split(enderecoFormatado, " ")[0]

	tipoEnderecoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			tipoEnderecoValido = true
		}
	}

	if tipoEnderecoValido {
		return primeiraPalavra
	}

	return "Tipo Inválido"
}
