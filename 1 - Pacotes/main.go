package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	auxiliar.Escrever()
	fmt.Println("Hello, World!")

	erro := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro)
}
