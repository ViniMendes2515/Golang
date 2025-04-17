package main

import "time"

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			canal1 <- "Olá canal 1"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			canal2 <- "Olá canal 2"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select {
		case msg1 := <-canal1:
			println(msg1)
		case msg2 := <-canal2:
			println(msg2)
		}
	}

}
