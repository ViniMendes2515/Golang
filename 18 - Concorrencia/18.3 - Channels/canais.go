package main

import (
	"time"
)

func main() {
	channel := make(chan string)

	go escrever("Olá Mundo!", channel)

	for mensagem := range channel {
		println(mensagem)
	}
}

func escrever(texto string, channel chan string) {
	for range 5 {
		channel <- texto
		time.Sleep(time.Second)
	}

	close(channel)
}
