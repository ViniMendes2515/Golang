package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1!"
	variavel2 := "Variavel 2!"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	// Declarar multiplas variaveis tipadas manualmente
	var (
		variavel3 string = "Variavel 3!"
		variavel4 string = "Variavel 4!"
	)

	fmt.Println(variavel3, variavel4)

	// Declarar multiplas variaveis inferidas

	variavel5, variavel6 := "Variavel 5!", "Variavel 6!"

	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1!"
	fmt.Println(constante1)

	// Invertendo valores de variaveis
	variavel5, variavel6 = variavel6, variavel5
}
