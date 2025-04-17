package main

func main() {
	tarefas := make(chan int, 40)
	resultado := make(chan int, 40)

	go worker(tarefas, resultado)
	go worker(tarefas, resultado)
	go worker(tarefas, resultado)

	for i := range 40 {
		tarefas <- i
	}

	for range 40 {
		numero := <-resultado
		println(numero)
	}

	close(resultado)
}

func worker(tarefas <-chan int, resultado chan<- int) {
	for numero := range tarefas {
		resultado <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-1) + fibonacci(posicao-2)
}
