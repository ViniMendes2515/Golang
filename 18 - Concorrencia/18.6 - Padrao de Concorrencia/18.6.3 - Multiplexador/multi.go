package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	msg1 := escrever("Mensagem 1")
	msg2 := escrever("Mensagem 2")

	canal := multiplexar(msg1, msg2)

	for range 10 {
		fmt.Println(<-canal)
	}
}

func multiplexar(canalEntrada1, canalEntrada2 <-chan string) <-chan string {
	canalSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canalEntrada1:
				canalSaida <- mensagem
			case mensagem := <-canalEntrada2:
				canalSaida <- mensagem
			}
		}
	}()

	return canalSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- texto
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
