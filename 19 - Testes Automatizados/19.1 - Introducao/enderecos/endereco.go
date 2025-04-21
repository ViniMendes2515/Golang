package enderecos

import (
	"slices"
	"strings"
)

// TipoDeEndereco recebe um endereço e retorna o tipo de endereço (rua, avenida, estrada ou rodovia) se for válido, ou uma mensagem de erro se não for.
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoLower := strings.ToLower(endereco)
	firstWord := strings.Split(enderecoLower, " ")[0]

	enderecoIsValid := false

	// Verifica se o primeiro elemento do slice tiposValidos é igual ao primeiro elemento do slice enderecoLower utilizando um loop
	// for _, tipo := range tiposValidos {
	// 	if tipo == firstWord {
	// 		enderecoIsValid = true
	// 		break
	// 	}
	// }

	// Faz a mesma coisa utilizando porem usando a função Contains do pacote slices
	enderecoIsValid = slices.Contains(tiposValidos, firstWord)

	if enderecoIsValid {
		return strings.Title(firstWord)
	}

	return "Tipo de endereço inválido"
}
