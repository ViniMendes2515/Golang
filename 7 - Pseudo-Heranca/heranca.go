package main

import "fmt"

type pessoa struct {
	nome      string
	idade     uint8
	altura    uint
	sobrenome string
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
	periodo   uint8
}

func main() {

	pessoa1 := pessoa{"Lucas", 20, 180, "Silva"}

	estudante1 := estudante{pessoa1, "Engenharia", "UFMG", 3}

	fmt.Println(estudante1)
	fmt.Println(estudante1.nome)
}
