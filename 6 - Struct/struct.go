package main

import "fmt"

type Pessoa struct {
	nome     string
	idade    uint8
	Endereco Endereco
}

type Endereco struct {
	logradouro string
	numero     uint8
	cidade     string
}

func main() {
	// Estruturas são tipos de dados compostos que podem conter outros tipos de dados.
	// Elas são úteis para agrupar dados relacionados em uma única unidade.
	// As estruturas são definidas usando a palavra-chave "struct".

	var pessoa1 Pessoa
	pessoa1.nome = "João"
	pessoa1.idade = 30

	fmt.Println(pessoa1)

	enderecoEx := Endereco{"Rua A", 123, "São Paulo"}

	pessoa2 := Pessoa{"Maria", 25, enderecoEx}
	fmt.Println(pessoa2)

	// Caso saiba so o nome por exemplo, deve ser passado com a propriedade para nn dar erro
	pessoa3 := Pessoa{nome: "Pedro"}
	fmt.Println(pessoa3.nome, pessoa3.idade)

	fmt.Println(pessoa2)
}
