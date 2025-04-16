package main

import "fmt"

func main() {
	channel := make(chan string, 2)

	channel <- "Olá Mundo!"
	channel <- "Olá Mundo 2!"

	mensagem := <-channel
	mensagem2 := <-channel

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
