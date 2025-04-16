package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escreverFuncaoAnonima("Olá Mundo!")
		waitGroup.Done()
	}()

	go escrever("Olá Devs!", &waitGroup)

	waitGroup.Wait()

	fmt.Println("Fim do programa!")

}

func escreverFuncaoAnonima(texto string) {
	for range 5 {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func escrever(texto string, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 5 {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
