package enderecos

import (
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

// Teste Unitario

func TestTipoDeEndereco(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{

		{"Rua ABC", "Rua"},
		{"Avenida ABC", "Avenida"},
		{"Rodovia ABC", "Rodovia"},
		{"Viela ABC", "Viela"},
		{"Travessa ABC", "Tipo não definido"},
	}

	for _, teste := range cenariosDeTeste {
		TipoDeEnderecoRecebido := TipoDeEnderecos(teste.enderecoInserido)
		if TipoDeEnderecoRecebido != teste.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", TipoDeEnderecoRecebido, teste.retornoEsperado)
		}
	}
}
