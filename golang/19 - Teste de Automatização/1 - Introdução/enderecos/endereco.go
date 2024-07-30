package enderecos

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// TipoDeEndereco mostra que tipo de endereço o endereço é.
func TipoDeEnderecos(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "viela", "rodovia"}

	enderecoLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoLetraMinuscula, " ")[0]

	enderecoTemTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemTipoValido = true
		}
	}

	if enderecoTemTipoValido {
		retorno := cases.Title(language.BrazilianPortuguese)
		return retorno.String(primeiraPalavraDoEndereco)
	}

	return "Tipo não definido"

}
