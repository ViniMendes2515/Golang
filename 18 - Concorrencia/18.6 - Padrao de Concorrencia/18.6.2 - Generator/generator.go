package main

import "time"

func main() {
	canal := escrever("Olá Mundo!")

	for range 10 {
		println(<-canal)
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- texto
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
