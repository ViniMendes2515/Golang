package enderecos

import "testing"

type cenarioTeste struct {
	recebido string
	esperado string
}

func TestTipoDeEndereco(t *testing.T) {

	t.Parallel()

	cenarios := []cenarioTeste{
		{"Rua das Flores", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Estrada do Sol", "Estrada"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"", "Tipo de endereço inválido"},
	}

	for _, cenario := range cenarios {
		recebido := TipoDeEndereco(cenario.recebido)
		if recebido != cenario.esperado {
			t.Errorf("Esperado: %s, Recebido: %s", cenario.esperado, recebido)
		}
	}

}

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Errorf("1 não é igual a 1")
	}
}
